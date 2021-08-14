package graph_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	_ "github.com/mattn/go-sqlite3"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/ent/enttest"
	"github.com/msal4/hassah_school_server/ent/school"
	"github.com/msal4/hassah_school_server/graph"
	"github.com/msal4/hassah_school_server/graph/model"
	"github.com/msal4/hassah_school_server/service"
	"github.com/msal4/hassah_school_server/testutil"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type schoolTestSuite struct {
	suite.Suite
	mc *minio.Client
}

func (s *schoolTestSuite) newService(db string) *service.Service {
	ec := enttest.Open(s.T(), dialect.SQLite, fmt.Sprintf("file:%s?mode=memory&cache=shared&_fk=1", db), enttest.WithOptions(ent.Log(s.T().Log)))
	srv, err := service.New(ec, s.mc, nil)
	s.Require().NoError(err)
	return srv
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
	srv := s.newService("123rand")
	gc := client.New(handler.NewDefaultServer(graph.NewSchema(srv)))
	ec := srv.EC

	type response struct {
		Schools struct {
			TotalCount int
			Edges      []struct {
				Node struct {
					ID        string
					Name      string
					Image     string
					Active    bool
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
        active
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
		defer srv.EC.School.Delete().ExecX(ctx)

		const expectedLen = 3
		schools := make([]*ent.School, expectedLen)
		schools[0] = ec.School.Create().SetName("school 1").SetImage("image/1").SetDirectory("test_dir").SaveX(ctx)
		schools[1] = ec.School.Create().SetName("school 2").SetImage("image/2").SetDirectory("test_dir").SaveX(ctx)
		schools[2] = ec.School.Create().SetName("school 3").SetImage("image/3").SetDirectory("test_dir").SetActive(false).SaveX(ctx)

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
        active
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
			require.Equal(t, currentSchool.Active, edge.Node.Active)
			require.Equal(t, currentSchool.Image, edge.Node.Image)
		}
	})

	s.T().Run("order", func(t *testing.T) {
		defer ec.School.Delete().ExecX(ctx)

		ec.School.Create().SetName("school 1").SetImage("image/1").SetDirectory("test_dir").SetCreatedAt(time.Now().Add(time.Minute)).SaveX(ctx)
		ec.School.Create().SetName("school 2").SetImage("image/2").SetDirectory("test_dir").SetCreatedAt(time.Now().Add(time.Hour)).SaveX(ctx)
		ec.School.Create().SetName("school 3").SetImage("image/3").SetDirectory("test_dir").SetActive(false).SaveX(ctx)

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
        active
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
			require.Equal(t, want[i].Active, edge.Node.Active)
			require.Equal(t, want[i].Image, edge.Node.Image)
		}
	})

	s.T().Run("order & filter", func(t *testing.T) {
		defer ec.School.Delete().ExecX(ctx)

		ec.School.Create().SetName("school 1").SetDirectory("test_dir").SetImage("image/1").SetCreatedAt(time.Now().Add(time.Minute)).SaveX(ctx)
		ec.School.Create().SetName("school 2").SetDirectory("test_dir").SetImage("image/2").SetCreatedAt(time.Now().Add(time.Hour)).SaveX(ctx)
		ec.School.Create().SetName("school 3").SetDirectory("test_dir").SetImage("image/3").SetActive(false).SaveX(ctx)

		var resp response

		const query = `query {
schools(orderBy: {field: CREATED_AT, direction: ASC}, where: {active: false}) {
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
        active
        createdAt
        updatedAt
      }
      cursor
    }
  }
}`

		want := ec.School.Query().Order(ent.Asc(school.FieldCreatedAt)).Where(school.Active(false)).AllX(ctx)

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
			require.Equal(t, false, edge.Node.Active)
			require.Equal(t, want[i].Image, edge.Node.Image)
		}
	})
}

func (s *schoolTestSuite) TestAddSchool() {
	srv := s.newService("tesw2")
	server := handler.NewDefaultServer(graph.NewSchema(srv))
	gc := client.New(server)
	ec := srv.EC
	ctx := context.Background()

	type response struct {
		Data *struct {
			AddSchool *struct {
				ID        string `json:"id"`
				Name      string `json:"name"`
				Image     string `json:"image"`
				Active    bool   `json:"active"`
				CreatedAt string `json:"created_at"`
				UpdatedAt string `json:"updated_at"`
			} `json:"addSchool"`
		} `json:"data"`
		Errors []struct {
			Message string   `json:"message"`
			Path    []string `json:"path"`
		} `json:"errors,omitempty"`
	}

	s.T().Run("missing image", func(t *testing.T) {
		var resp response
		err := gc.Post("mutation { addSchool(input: {name: \"a school without an image\"}) { id name image active createdAt updatedAt }}", &resp.Data)
		require.Error(t, err)
	})

	s.T().Run("with image", func(t *testing.T) {
		defer ec.School.Delete().ExecX(ctx)

		w := httptest.NewRecorder()

		imgFile, err := os.Open("../testfiles/stanford.png")
		require.NoError(t, err)

		operations := `{
			"query": "mutation ($image: Upload!) { addSchool(input: {name: \"a school with an image\", image: $image}) { id name image active createdAt updatedAt }}", 
			"variables": {"image": null}
		}`

		mapData := `{"0": ["variables.image"]}`

		r := createMultipartRequest(t, operations, mapData, file{
			mapKey: "0",
			File:   imgFile,
		})

		server.ServeHTTP(w, r)

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

const testID = "2710c203-7842-4356-8d9f-12f9da4722a2"

func (s *schoolTestSuite) TestUpdateSchool() {
	srv := s.newService("sd3tesw2")
	defer srv.EC.Close()
	server := handler.NewDefaultServer(graph.NewSchema(srv))
	gc := client.New(server)
	ec := srv.EC
	ctx := context.Background()

	type response struct {
		Data *struct {
			UpdateSchool *struct {
				ID        string `json:"id"`
				Name      string `json:"name"`
				Image     string `json:"image"`
				Active    bool   `json:"active"`
				CreatedAt string `json:"created_at"`
				UpdatedAt string `json:"updated_at"`
			} `json:"updateSchool"`
		} `json:"data"`
		Errors []struct {
			Message string   `json:"message"`
			Path    []string `json:"path"`
		} `json:"errors,omitempty"`
	}

	s.T().Run("invalid", func(t *testing.T) {
		var resp response
		err := gc.Post(fmt.Sprintf("mutation { updateSchool(id: %q, input: {name: \"a school without an image\"}) { id name image active createdAt updatedAt }}", testID), &resp.Data)
		fmt.Println(err)
		require.Error(t, err)
	})

	f := testutil.OpenFile(s.T(), "../testfiles/harvard.jpg")
	defer f.Close()
	sch, err := srv.AddSchool(ctx,
		model.AddSchoolInput{
			Name: "test schoo",
			Image: graphql.Upload{
				File:     f,
				Filename: f.File.Name(),
				Size:     f.Size(),
			},
			Active: true,
		},
	)
	require.NoError(s.T(), err)

	s.T().Run("valid", func(t *testing.T) {
		defer ec.School.Delete().ExecX(ctx)

		w := httptest.NewRecorder()

		imgFile, err := os.Open("../testfiles/stanford.png")
		defer imgFile.Close()
		require.NoError(t, err)

		operations := fmt.Sprintf(`{
"query": "mutation ($image: Upload!) { updateSchool(id: \"%s\", input: {name: \"a school with an image\", image: $image}) { id name image active createdAt updatedAt }}", 
			"variables": {"image": null}
		}`, sch.ID)

		mapData := `{"0": ["variables.image"]}`

		r := createMultipartRequest(t, operations, mapData, file{
			mapKey: "0",
			File:   imgFile,
		})

		server.ServeHTTP(w, r)

		var resp response
		require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
		require.Nil(t, resp.Errors)
		require.NotNil(t, resp.Data, "data is nil")
		require.NotNil(t, resp.Data.UpdateSchool, "data.updateSchool is nil")
		require.NotEmpty(t, resp.Data.UpdateSchool.ID)
		require.Equal(t, "a school with an image", resp.Data.UpdateSchool.Name)
	})
}

func (s *schoolTestSuite) TestDeleteSchool() {
	srv := s.newService("rsdj2")
	ec := srv.EC
	server := handler.NewDefaultServer(graph.NewSchema(srv))
	gc := client.New(server)
	ctx := context.Background()

	type response struct {
		DeleteSchool bool
	}

	s.T().Run("exists", func(t *testing.T) {
		defer ec.School.Delete().ExecX(ctx)

		sch := ec.School.Create().SetName("test school").SetDirectory("test_dir").SetImage("test/image").SaveX(ctx)

		var resp response
		gc.MustPost(fmt.Sprintf(`mutation { deleteSchool(id:"%s") } `, sch.ID.String()), &resp)
		require.True(t, resp.DeleteSchool)

		schools := ec.School.Query().AllX(ctx)
		require.Empty(t, schools)
	})

	s.T().Run("does not exist", func(t *testing.T) {
		defer ec.School.Delete().ExecX(ctx)

		ec.School.Create().SetName("test school").SetDirectory("test_dir").SetImage("test/image").SaveX(ctx)

		var resp response
		err := gc.Post(`mutation { deleteSchool(id:"2710c203-7842-4356-8d9f-12f9da4722a2") } `, &resp)
		require.Error(t, err)

		schools := ec.School.Query().AllX(ctx)
		require.Len(t, schools, 1)
	})
}
