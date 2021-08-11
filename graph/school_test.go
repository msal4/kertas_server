package graph_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	_ "github.com/mattn/go-sqlite3"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/ent/enttest"
	"github.com/msal4/hassah_school_server/ent/schema"
	"github.com/msal4/hassah_school_server/ent/school"
	"github.com/msal4/hassah_school_server/graph"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type schoolTestSuite struct {
	suite.Suite
	mc *minio.Client
}

func TestSchool(t *testing.T) {
	suite.Run(t, &schoolTestSuite{})
}

func (s *schoolTestSuite) SetupTest() {
	var err error
	s.mc, err = minio.New("localhost:9000", &minio.Options{
		Creds: credentials.NewStaticV4("minioadmin", "minioadmin", ""),
	})
	s.Require().Nilf(err, "instantiating minio client: %v", err)
	_, err = s.mc.ListBuckets(context.Background())
	s.Require().Nilf(err, "connecting to minio: %v", err)
}

func (s *schoolTestSuite) TestSchools() {
	ec := enttest.Open(s.T(), dialect.SQLite, "file:ent2?mode=memory&cache=shared&_fk=1", enttest.WithOptions(ent.Log(s.T().Log)))
	gc := client.New(handler.NewDefaultServer(graph.NewSchema(ec, s.mc, rand.NewSource(0))))

	type response struct {
		Schools struct {
			TotalCount int
			Edges      []struct {
				Node struct {
					ID        string
					Name      string
					Image     string
					Status    schema.Status
					CreatedAt string
					UpdatedAt string
				}
				Cursor string
			}
			PageInfo struct {
				HasNextPage     bool
				HasPreviousPage bool
				StartCursor     *string
				EndCursor       *string
			}
		}
	}
	ctx := context.Background()

	s.T().Run("empty", func(t *testing.T) {
		var resp response

		const query = `query {
  schools {
    totalCount
    pageInfo {
      hasNextPage
      hasPreviousPage
      startCursor
      endCursor
    }
    edges {
      node {
        id
        name
        image
        status
        createdAt
        updatedAt
      }
      cursor
    }
  }
}`

		gc.MustPost(query, &resp)

		require.Empty(t, resp.Schools.Edges)
		require.Zero(t, resp.Schools.TotalCount)
		require.False(t, resp.Schools.PageInfo.HasNextPage)
		require.False(t, resp.Schools.PageInfo.HasPreviousPage)
		require.Nil(t, resp.Schools.PageInfo.EndCursor)
		require.Nil(t, resp.Schools.PageInfo.StartCursor)
	})

	s.T().Run("unordered", func(t *testing.T) {
		defer ec.School.Delete().ExecX(ctx)

		const expectedLen = 3
		schools := make([]*ent.School, expectedLen)
		schools[0] = ec.School.Create().SetName("school 1").SetImage("image/1").SaveX(ctx)
		schools[1] = ec.School.Create().SetName("school 2").SetImage("image/2").SetStatus(schema.StatusDeleted).SaveX(ctx)
		schools[2] = ec.School.Create().SetName("school 3").SetImage("image/3").SetStatus(schema.StatusDisabled).SaveX(ctx)

		var resp response

		const query = `query {
  schools {
    totalCount
    pageInfo {
      hasNextPage
      hasPreviousPage
      startCursor
      endCursor
    }
    edges {
      node {
        id
        name
        image
        status
        createdAt
        updatedAt
      }
      cursor
    }
  }
}`

		gc.MustPost(query, &resp)

		require.Len(t, resp.Schools.Edges, expectedLen)
		require.Equal(t, resp.Schools.TotalCount, expectedLen)
		require.False(t, resp.Schools.PageInfo.HasNextPage)
		require.False(t, resp.Schools.PageInfo.HasPreviousPage)
		require.NotNil(t, resp.Schools.PageInfo.EndCursor)
		require.NotNil(t, resp.Schools.PageInfo.StartCursor)

		for _, edge := range resp.Schools.Edges {
			var currentSchool *ent.School
			for _, school := range schools {
				if school.ID.String() == edge.Node.ID {
					currentSchool = school
					break
				}
			}
			require.NotNil(t, currentSchool)
			require.NotNil(t, edge.Cursor)
			require.Equal(t, currentSchool.Name, edge.Node.Name)
			require.Equal(t, currentSchool.Status, edge.Node.Status)
			require.Equal(t, currentSchool.Image, edge.Node.Image)
		}
	})

	s.T().Run("order", func(t *testing.T) {
		defer ec.School.Delete().ExecX(ctx)

		ec.School.Create().SetName("school 1").SetImage("image/1").SetCreatedAt(time.Now().Add(time.Minute)).SaveX(ctx)
		ec.School.Create().SetName("school 2").SetImage("image/2").SetStatus(schema.StatusDeleted).SetCreatedAt(time.Now().Add(time.Hour)).SaveX(ctx)
		ec.School.Create().SetName("school 3").SetImage("image/3").SetStatus(schema.StatusDisabled).SaveX(ctx)

		var resp response

		const query = `query {
schools(orderBy: {field: CREATED_AT, direction: ASC}) {
    totalCount
    pageInfo {
      hasNextPage
      hasPreviousPage
      startCursor
      endCursor
    }
    edges {
      node {
        id
        name
        image
        status
        createdAt
        updatedAt
      }
      cursor
    }
  }
}`

		want := ec.School.Query().Order(ent.Asc(school.FieldCreatedAt)).AllX(ctx)

		gc.MustPost(query, &resp)

		require.Len(t, resp.Schools.Edges, len(want))
		require.Equal(t, resp.Schools.TotalCount, len(want))
		require.False(t, resp.Schools.PageInfo.HasNextPage)
		require.False(t, resp.Schools.PageInfo.HasPreviousPage)
		require.NotNil(t, resp.Schools.PageInfo.EndCursor)
		require.NotNil(t, resp.Schools.PageInfo.StartCursor)

		for i, edge := range resp.Schools.Edges {
			require.NotNil(t, edge.Cursor)
			require.Equal(t, want[i].ID.String(), edge.Node.ID)
			require.Equal(t, want[i].Name, edge.Node.Name)
			require.Equal(t, want[i].Status, edge.Node.Status)
			require.Equal(t, want[i].Image, edge.Node.Image)
		}
	})

	s.T().Run("order & filter", func(t *testing.T) {
		defer ec.School.Delete().ExecX(ctx)

		ec.School.Create().SetName("school 1").SetImage("image/1").SetCreatedAt(time.Now().Add(time.Minute)).SaveX(ctx)
		ec.School.Create().SetName("school 2").SetImage("image/2").SetStatus(schema.StatusDeleted).SetCreatedAt(time.Now().Add(time.Hour)).SaveX(ctx)
		ec.School.Create().SetName("school 3").SetImage("image/3").SetStatus(schema.StatusDisabled).SaveX(ctx)

		var resp response

		const query = `query {
schools(orderBy: {field: CREATED_AT, direction: ASC}, where: {status: DISABLED}) {
    totalCount
    pageInfo {
      hasNextPage
      hasPreviousPage
      startCursor
      endCursor
    }
    edges {
      node {
        id
        name
        image
        status
        createdAt
        updatedAt
      }
      cursor
    }
  }
}`

		want := ec.School.Query().Order(ent.Asc(school.FieldCreatedAt)).Where(school.StatusEQ(schema.StatusDisabled)).AllX(ctx)

		gc.MustPost(query, &resp)

		require.Len(t, resp.Schools.Edges, len(want))
		require.Equal(t, resp.Schools.TotalCount, len(want))
		require.False(t, resp.Schools.PageInfo.HasNextPage)
		require.False(t, resp.Schools.PageInfo.HasPreviousPage)
		require.NotNil(t, resp.Schools.PageInfo.EndCursor)
		require.NotNil(t, resp.Schools.PageInfo.StartCursor)

		for i, edge := range resp.Schools.Edges {
			require.NotNil(t, edge.Cursor)
			require.Equal(t, want[i].ID.String(), edge.Node.ID)
			require.Equal(t, want[i].Name, edge.Node.Name)
			require.Equal(t, schema.StatusDisabled, edge.Node.Status)
			require.Equal(t, want[i].Image, edge.Node.Image)
		}
	})
}

func (s *schoolTestSuite) TestAddSchool() {
	ec := enttest.Open(s.T(), dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1", enttest.WithOptions(ent.Log(s.T().Log)))
	srv := handler.NewDefaultServer(graph.NewSchema(ec, s.mc, rand.NewSource(0)))
	gc := client.New(srv)

	type response struct {
		Data *struct {
			AddSchool *struct {
				ID        string        `json:"id"`
				Name      string        `json:"name"`
				Image     string        `json:"image"`
				Status    schema.Status `json:"status"`
				CreatedAt string        `json:"created_at"`
				UpdatedAt string        `json:"updated_at"`
			} `json:"addSchool"`
		} `json:"data"`
		Errors []struct {
			Message string   `json:"message"`
			Path    []string `json:"path"`
		} `json:"errors,omitempty"`
	}

	s.T().Run("missing image", func(t *testing.T) {
		var resp response
		err := gc.Post("mutation { addSchool(input: {name: \"a school without an image\"}) { id name image status createdAt updatedAt }}", &resp.Data)
		require.Error(t, err)
	})

	s.T().Run("with image", func(t *testing.T) {
		w := httptest.NewRecorder()

		imgFile, err := os.Open("../testfiles/stanford.png")
		require.NoError(t, err)

		operations := `{
			"query": "mutation ($image: Upload!) { addSchool(input: {name: \"a school with an image\", image: $image}) { id name image status createdAt updatedAt }}", 
			"variables": {"image": null}
		}`

		mapData := `{"0": ["variables.image"]}`

		r := createMultipartRequest(t, operations, mapData, file{
			mapKey: "0",
			File:   imgFile,
		})

		srv.ServeHTTP(w, r)

		var resp response
		require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
		require.NotNil(t, resp.Data, "data is nil")
		require.NotNil(t, resp.Data.AddSchool, "data.addSchool is nil")
		require.NotEmpty(t, resp.Data.AddSchool.ID)
		require.Equal(t, "a school with an image", resp.Data.AddSchool.Name)
	})
}

type file struct {
	mapKey string
	*os.File
}

func createMultipartRequest(t *testing.T, operations, mapData string, f file) *http.Request {
	b := new(bytes.Buffer)
	w := multipart.NewWriter(b)
	require.NoError(t, w.WriteField("operations", operations))
	require.NoError(t, w.WriteField("map", mapData))

	ff, err := w.CreateFormFile(f.mapKey, f.Name())
	require.NoError(t, err)
	_, err = io.Copy(ff, f)
	require.NoError(t, err)

	require.NoError(t, w.Close())

	r := httptest.NewRequest(http.MethodPost, "/graphql", b)

	r.Header.Set("content-type", w.FormDataContentType())

	return r
}
