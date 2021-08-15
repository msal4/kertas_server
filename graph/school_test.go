package graph_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/99designs/gqlgen/graphql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/msal4/hassah_school_server/auth"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/ent/user"
	"github.com/msal4/hassah_school_server/graph"
	"github.com/msal4/hassah_school_server/graph/model"
	"github.com/msal4/hassah_school_server/service"
	"github.com/msal4/hassah_school_server/testutil"
	"github.com/stretchr/testify/require"
)

func TestSchools(t *testing.T) {
	s := newService(t)
	defer s.EC.Close()

	srv := graph.NewServer(s, false)
	ctx := context.Background()

	type response struct {
		Data *struct {
			Schools *struct {
				TotalCount int
				Edges      []struct {
					Node struct {
						ID        string `json:"id"`
						Name      string `json:"name"`
						Image     string `json:"image"`
						Active    bool   `json:"active"`
						CreatedAt string `json:"created_at"`
						UpdatedAt string `json:"udpated_at"`
					} `json:"edges"`
					Cursor *string
				} `json:"edges"`
				PageInfo struct {
					HasNextPage     bool
					HasPreviousPage bool
					StartCursor     *string
					EndCursor       *string
				} `json:"pageInfo"`
			} `json:"schools"`
		} `json:"data"`

		Errors []struct {
			Message string   `json:"message"`
			Path    []string `json:"path"`
		} `json:"errors,omitempty"`
	}

	operations := []byte(`{
		"query": "{ schools { totalCount pageInfo { hasNextPage hasPreviousPage startCursor endCursor } edges { node { id name image active createdAt updatedAt } cursor } } }"
		}`)

	t.Run("unauthorized", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodPost, "/graphql", bytes.NewBuffer(operations))
		w := httptest.NewRecorder()

		srv.ServeHTTP(w, r)

		require.Equal(t, http.StatusUnauthorized, w.Result().StatusCode)
	})

	t.Run("super admin", func(t *testing.T) {
		var resp response

		r := httptest.NewRequest(http.MethodPost, "/graphql", bytes.NewBuffer(operations))
		w := httptest.NewRecorder()
		r.Header.Set("content-type", "application/json")

		u := createSuperAdmin(ctx, s, "testuser223")
		data, err := auth.GenerateTokens(*u, s.Config.AuthConfig)
		require.NoError(t, err)
		require.NotNil(t, data)

		r.Header.Set("authorization", fmt.Sprintf("Bearer %s", data.AccessToken))

		srv.ServeHTTP(w, r)

		require.Equal(t, http.StatusOK, w.Result().StatusCode)

		err = json.NewDecoder(w.Body).Decode(&resp)
		require.NoError(t, err)

		require.NotNil(t, resp.Data)
		require.NotNil(t, resp.Data.Schools)
		require.Empty(t, resp.Data.Schools.Edges)
		require.Nil(t, resp.Errors)
	})

	t.Run("student", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodPost, "/graphql", bytes.NewBuffer(operations))
		w := httptest.NewRecorder()
		r.Header.Set("content-type", "application/json")

		u := createStudent(ctx, s, "22testuser223")
		data, err := auth.GenerateTokens(*u, s.Config.AuthConfig)
		require.NoError(t, err)
		require.NotNil(t, data)

		r.Header.Set("authorization", fmt.Sprintf("Bearer %s", data.AccessToken))

		srv.ServeHTTP(w, r)

		var resp response
		err = json.NewDecoder(w.Body).Decode(&resp)
		require.NoError(t, err)

		require.Nil(t, resp.Data.Schools)
		require.NotEmpty(t, resp.Errors)
		require.Equal(t, auth.UnauthorizedErr.Error(), resp.Errors[0].Message)
	})
}

func createSuperAdmin(ctx context.Context, s *service.Service, username string) *ent.User {
	return s.EC.User.Create().SetName("test userd" + username).SetUsername(username).
		SetPhone("077059333812").SetDirectory("diresss22").SetRole(user.RoleSuperAdmin).SetImage("sss").SetPassword("mipassword22@@@@5").SaveX(ctx)
}

func createStudent(ctx context.Context, s *service.Service, username string) *ent.User {
	sch := s.EC.School.Create().SetName("schooltest").SetDirectory("fsss").SetImage("fss").SaveX(ctx)
	stage := s.EC.Stage.Create().SetName("2nd").SetTuitionAmount(122).SetSchool(sch).SaveX(ctx)
	return s.EC.User.Create().SetName("test userd" + username).SetUsername(username).
		SetPhone("077059333812").SetDirectory("diresss22").SetPassword("mipassword22@@@@5").SetSchool(sch).SetStage(stage).SaveX(ctx)
}

func TestAddSchool(t *testing.T) {
	s := newService(t)
	srv := graph.NewServer(s, false)
	ec := s.EC
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

	u := createSuperAdmin(ctx, s, "whatev22223")
	data, err := auth.GenerateTokens(*u, s.Config.AuthConfig)
	require.NoError(t, err)

	t.Run("missing image", func(t *testing.T) {
		var resp response

		w := httptest.NewRecorder()
		r := createRequest(t, "mutation { addSchool(input: {name: \"a school without an image\"}) { id name image active createdAt updatedAt }}", "{}")
		r.Header.Set("authorization", fmt.Sprintf("Bearer %s", data.AccessToken))

		srv.ServeHTTP(w, r)

		err = json.NewDecoder(w.Body).Decode(&resp)
		require.NoError(t, err)

		require.NotEmpty(t, resp.Errors)
	})

	t.Run("super admin", func(t *testing.T) {
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

		r.Header.Set("authorization", fmt.Sprintf("Bearer %s", data.AccessToken))

		srv.ServeHTTP(w, r)

		var resp response
		require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
		require.NotNil(t, resp.Data, "data is nil")
		require.NotNil(t, resp.Data.AddSchool, "data.addSchool is nil")
		require.NotEmpty(t, resp.Data.AddSchool.ID)
		require.Equal(t, "a school with an image", resp.Data.AddSchool.Name)
	})

	t.Run("student", func(t *testing.T) {
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

		u.Role = user.RoleStudent
		data, err = auth.GenerateTokens(*u, s.Config.AuthConfig)
		require.NoError(t, err)

		r.Header.Set("authorization", fmt.Sprintf("Bearer %s", data.AccessToken))

		srv.ServeHTTP(w, r)

		var resp response
		require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
		require.NotEmpty(t, resp.Errors)
		require.Equal(t, auth.UnauthorizedErr.Error(), resp.Errors[0].Message)
	})
}

const testID = "2710c203-7842-4356-8d9f-12f9da4722a2"

func TestUpdateSchool(t *testing.T) {
	s := newService(t)
	defer s.EC.Close()

	srv := graph.NewServer(s, false)
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

	t.Run("unauthenticated", func(t *testing.T) {
		r := createRequest(t, fmt.Sprintf("mutation { updateSchool(id: %q, input: {name: \"a school without an image\"}) { id name image active createdAt updatedAt }}", testID), "{}")

		w := httptest.NewRecorder()

		srv.ServeHTTP(w, r)

		require.Equal(t, http.StatusUnauthorized, w.Code)
	})

	f := testutil.OpenFile(t, "../testfiles/harvard.jpg")
	defer f.Close()
	sch, err := s.AddSchool(ctx,
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
	require.NoError(t, err)

	u := createSuperAdmin(ctx, s, "superuser224")
	data, err := auth.GenerateTokens(*u, s.Config.AuthConfig)
	require.NoError(t, err)

	imgFile, err := os.Open("../testfiles/stanford.png")
	defer imgFile.Close()
	require.NoError(t, err)

	operations := fmt.Sprintf(`{
"query": "mutation ($image: Upload!) { updateSchool(id: \"%s\", input: {name: \"a school with an image\", image: $image}) { id name image active createdAt updatedAt }}",
			"variables": {"image": null}
		}`, sch.ID)

	mapData := `{"0": ["variables.image"]}`

	t.Run("super admin", func(t *testing.T) {
		w := httptest.NewRecorder()

		r := createMultipartRequest(t, operations, mapData, file{
			mapKey: "0",
			File:   imgFile,
		})

		r.Header.Set("authorization", "Bearer "+data.AccessToken)

		srv.ServeHTTP(w, r)

		var resp response
		require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
		require.Empty(t, resp.Errors)
		require.NotNil(t, resp.Data, "data must not be nil")
		require.NotNil(t, resp.Data.UpdateSchool, "data.updateSchool must not be nil")
		require.NotEmpty(t, resp.Data.UpdateSchool.ID)
		require.Equal(t, "a school with an image", resp.Data.UpdateSchool.Name)
	})

	t.Run("school admin", func(t *testing.T) {
		w := httptest.NewRecorder()

		imgFile.Seek(0, 0)

		r := createMultipartRequest(t, operations, mapData, file{
			mapKey: "0",
			File:   imgFile,
		})

		u.Role = user.RoleSchoolAdmin

		data, err = auth.GenerateTokens(*u, s.Config.AuthConfig)
		require.NoError(t, err)

		r.Header.Set("authorization", "Bearer "+data.AccessToken)

		srv.ServeHTTP(w, r)

		var resp response
		require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
		require.NotEmpty(t, resp.Errors)
		require.Equal(t, auth.UnauthorizedErr.Error(), resp.Errors[0].Message)
	})
}

func TestDeleteSchoolPermanently(t *testing.T) {
	s := newService(t)
	defer s.EC.Close()

	ec := s.EC
	srv := graph.NewServer(s, false)
	ctx := context.Background()

	type response struct {
		Data struct {
			DeleteSchoolPermanently bool `json:"deleteSchoolPermanently"`
		} `json:"data"`

		Errors []struct {
			Message string   `json:"message"`
			Path    []string `json:"path"`
		} `json:"errors,omitempty"`
	}

	u := createSuperAdmin(ctx, s, "superuser224")
	data, err := auth.GenerateTokens(*u, s.Config.AuthConfig)
	require.NoError(t, err)

	t.Run("super admin", func(t *testing.T) {
		var resp response
		sch := ec.School.Create().SetName("test school").SetDirectory("test_dir").SetImage("test/image").SaveX(ctx)

		r := createRequest(t, fmt.Sprintf(`mutation { deleteSchoolPermanently(id:"%s") }`, sch.ID.String()), "{}")
		r.Header.Set("authorization", "Bearer "+data.AccessToken)

		w := httptest.NewRecorder()

		srv.ServeHTTP(w, r)

		require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))

		require.True(t, resp.Data.DeleteSchoolPermanently)
	})

	t.Run("school admin", func(t *testing.T) {
		defer ec.School.Delete().ExecX(ctx)

		sch := ec.School.Create().SetName("test school").SetDirectory("test_dir").SetImage("test/image").SaveX(ctx)

		r := createRequest(t, fmt.Sprintf(`mutation { deleteSchoolPermanently(id:"%s") }`, sch.ID.String()), "{}")

		u.Role = user.RoleSchoolAdmin
		data, err := auth.GenerateTokens(*u, s.Config.AuthConfig)
		require.NoError(t, err)

		r.Header.Set("authorization", "Bearer "+data.AccessToken)

		w := httptest.NewRecorder()

		srv.ServeHTTP(w, r)

		var resp response
		require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))

		require.NotEmpty(t, resp.Errors)
		require.Equal(t, auth.UnauthorizedErr.Error(), resp.Errors[0].Message)
	})
}

func TestDeleteSchool(t *testing.T) {
	s := newService(t)
	defer s.EC.Close()

	ec := s.EC
	srv := graph.NewServer(s, false)
	ctx := context.Background()

	type response struct {
		Data struct {
			DeleteSchool bool `json:"deleteSchool"`
		} `json:"data"`

		Errors []struct {
			Message string   `json:"message"`
			Path    []string `json:"path"`
		} `json:"errors,omitempty"`
	}

	u := createSuperAdmin(ctx, s, "superuser224")
	data, err := auth.GenerateTokens(*u, s.Config.AuthConfig)
	require.NoError(t, err)

	t.Run("super admin authorized", func(t *testing.T) {
		var resp response
		sch := ec.School.Create().SetName("test school").SetDirectory("test_dir").SetImage("test/image").SaveX(ctx)

		r := createRequest(t, fmt.Sprintf(`mutation { deleteSchool(id:"%s") }`, sch.ID.String()), "{}")
		r.Header.Set("authorization", "Bearer "+data.AccessToken)

		w := httptest.NewRecorder()

		srv.ServeHTTP(w, r)

		require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))

		require.True(t, resp.Data.DeleteSchool)
	})

	t.Run("school admin not authorized", func(t *testing.T) {
		defer ec.School.Delete().ExecX(ctx)

		sch := ec.School.Create().SetName("test school").SetDirectory("test_dir").SetImage("test/image").SaveX(ctx)

		r := createRequest(t, fmt.Sprintf(`mutation { deleteSchool(id:"%s") }`, sch.ID.String()), "{}")

		u.Role = user.RoleSchoolAdmin
		data, err := auth.GenerateTokens(*u, s.Config.AuthConfig)
		require.NoError(t, err)

		r.Header.Set("authorization", "Bearer "+data.AccessToken)

		w := httptest.NewRecorder()

		srv.ServeHTTP(w, r)

		var resp response
		require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))

		require.NotEmpty(t, resp.Errors)
		require.Equal(t, auth.UnauthorizedErr.Error(), resp.Errors[0].Message)
	})
}
