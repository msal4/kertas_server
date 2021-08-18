package server_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/msal4/hassah_school_server/auth"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/ent/user"
	"github.com/msal4/hassah_school_server/server"
	"github.com/msal4/hassah_school_server/util/ptr"
	"github.com/stretchr/testify/require"
)

func TestStages(t *testing.T) {
	s := newService(t)
	defer s.EC.Close()

	srv := server.NewServer(s, false)
	ctx := context.Background()

	query := "{ stages { totalCount } } "

	t.Run("super admin", func(t *testing.T) {
		var resp errsResponse

		r := createRequest(t, query, "{}")

		w := httptest.NewRecorder()

		u := createSuperAdmin(ctx, s, "testuser223")
		data, err := auth.GenerateTokens(*u, s.Config.AuthConfig)
		require.NoError(t, err)
		require.NotNil(t, data)

		r.Header.Set("authorization", fmt.Sprintf("Bearer %s", data.AccessToken))

		srv.ServeHTTP(w, r)

		require.Equal(t, http.StatusOK, w.Result().StatusCode)

		err = json.NewDecoder(w.Body).Decode(&resp)
		require.NoError(t, err)
		require.Nil(t, resp.Errors)
	})

	t.Run("student", func(t *testing.T) {
		r := createRequest(t, query, "{}")
		w := httptest.NewRecorder()

		u := createStudent(ctx, s, "22testuser223")
		data, err := auth.GenerateTokens(*u, s.Config.AuthConfig)
		require.NoError(t, err)
		require.NotNil(t, data)

		r.Header.Set("authorization", fmt.Sprintf("Bearer %s", data.AccessToken))

		srv.ServeHTTP(w, r)

		var resp errsResponse
		err = json.NewDecoder(w.Body).Decode(&resp)
		require.NoError(t, err)

		require.Nil(t, resp.Errors)
	})

	t.Run("not authenticated", func(t *testing.T) {
		r := createRequest(t, query, "{}")
		w := httptest.NewRecorder()

		srv.ServeHTTP(w, r)

		var resp errsResponse
		err := json.NewDecoder(w.Body).Decode(&resp)
		require.NoError(t, err)

		require.Nil(t, resp.Errors)
	})
}

func TestAddStage(t *testing.T) {
	s := newService(t)
	defer s.EC.Close()

	srv := server.NewServer(s, false)
	ctx := context.Background()

	suAdmin := createSuperAdmin(ctx, s, "hello23super")
	schAdmin := *suAdmin
	schAdmin.Role = user.RoleSchoolAdmin
	teacher := schAdmin
	teacher.Role = user.RoleTeacher
	student := schAdmin
	student.Role = user.RoleStudent

	sch := createSchool(ctx, s, "test school", "ch/image.pgg")

	query := `mutation { addStage(input: {name: "a stage", tuition_amount: 20, school_id: %q}) { id }}`

	cases := []struct {
		desc string
		user *ent.User
		want *string
	}{
		{"super admin is authorized", suAdmin, nil},
		{"school admin is authorized", &schAdmin, nil},
		{"teacher is not authorized", &teacher, ptr.Str(auth.UnauthorizedErr.Error())},
		{"student is not authorized", &student, ptr.Str(auth.UnauthorizedErr.Error())},
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			var resp errsResponse

			r := createRequest(t, fmt.Sprintf(query, sch.ID), `{}`)
			w := httptest.NewRecorder()

			data := genTokens(t, c.user, s)

			setAuth(r, data.AccessToken)

			srv.ServeHTTP(w, r)

			parseBody(t, w, &resp)

			if c.want == nil {
				require.Nil(t, resp.Errors)
				return
			}

			require.NotEmpty(t, resp.Errors)
			require.Equal(t, *c.want, resp.Errors[0].Message)
		})
	}
}

func TestUpdateStage(t *testing.T) {
	s := newService(t)
	defer s.EC.Close()

	srv := server.NewServer(s, false)
	ctx := context.Background()

	suAdmin := createSuperAdmin(ctx, s, "hello23super")
	schAdmin := *suAdmin
	schAdmin.Role = user.RoleSchoolAdmin
	teacher := schAdmin
	teacher.Role = user.RoleTeacher
	student := schAdmin
	student.Role = user.RoleStudent

	st := createStage(ctx, s, "stage", 3523)

	query := fmt.Sprintf(`mutation { updateStage(id: %q, input: {name: "a stage with a new name", tuition_amount: 20}) { id }}`, st.ID)

	cases := []struct {
		desc string
		user *ent.User
		want *string
	}{
		{"super admin is authorized", suAdmin, nil},
		{"school admin is authorized", &schAdmin, nil},
		{"teacher is not authorized", &teacher, ptr.Str(auth.UnauthorizedErr.Error())},
		{"student is not authorized", &student, ptr.Str(auth.UnauthorizedErr.Error())},
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			var resp errsResponse

			r := createRequest(t, query, `{}`)
			w := httptest.NewRecorder()

			data := genTokens(t, c.user, s)

			setAuth(r, data.AccessToken)

			srv.ServeHTTP(w, r)

			parseBody(t, w, &resp)

			if c.want == nil {
				require.Nil(t, resp.Errors)
				return
			}

			require.NotEmpty(t, resp.Errors)
			require.Equal(t, *c.want, resp.Errors[0].Message)
		})
	}
}

func TestDeleteStage(t *testing.T) {
	s := newService(t)
	defer s.EC.Close()

	srv := server.NewServer(s, false)
	ctx := context.Background()

	suAdmin := createSuperAdmin(ctx, s, "hello23super")
	schAdmin := *suAdmin
	schAdmin.Role = user.RoleSchoolAdmin
	teacher := schAdmin
	teacher.Role = user.RoleTeacher
	student := schAdmin
	student.Role = user.RoleStudent

	st := createStage(ctx, s, "stage", 3523)

	query := fmt.Sprintf(`mutation { deleteStage(id: %q) }`, st.ID)

	cases := []struct {
		desc string
		user *ent.User
		want *string
	}{
		{"super admin is authorized", suAdmin, nil},
		{"school admin is authorized", &schAdmin, nil},
		{"teacher is not authorized", &teacher, ptr.Str(auth.UnauthorizedErr.Error())},
		{"student is not authorized", &student, ptr.Str(auth.UnauthorizedErr.Error())},
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			var resp errsResponse

			r := createRequest(t, query, `{}`)
			w := httptest.NewRecorder()

			data := genTokens(t, c.user, s)

			setAuth(r, data.AccessToken)

			srv.ServeHTTP(w, r)

			parseBody(t, w, &resp)

			if c.want == nil {
				require.Nil(t, resp.Errors)
				return
			}

			require.NotEmpty(t, resp.Errors)
			require.Equal(t, *c.want, resp.Errors[0].Message)
		})
	}
}

func TestDeleteStagePermanently(t *testing.T) {
	s := newService(t)
	defer s.EC.Close()

	srv := server.NewServer(s, false)
	ctx := context.Background()

	suAdmin := createSuperAdmin(ctx, s, "hello23super")
	schAdmin := *suAdmin
	schAdmin.Role = user.RoleSchoolAdmin
	teacher := schAdmin
	teacher.Role = user.RoleTeacher
	student := schAdmin
	student.Role = user.RoleStudent

	query := `mutation { deleteStagePermanently(id: %q) }`

	cases := []struct {
		desc string
		user *ent.User
		want *string
	}{
		{"super admin is authorized", suAdmin, nil},
		{"school admin is authorized", &schAdmin, nil},
		{"teacher is not authorized", &teacher, ptr.Str(auth.UnauthorizedErr.Error())},
		{"student is not authorized", &student, ptr.Str(auth.UnauthorizedErr.Error())},
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			var resp errsResponse

			st := createStage(ctx, s, "stage", 3523)

			r := createRequest(t, fmt.Sprintf(query, st.ID), `{}`)
			w := httptest.NewRecorder()

			data := genTokens(t, c.user, s)

			setAuth(r, data.AccessToken)

			srv.ServeHTTP(w, r)

			parseBody(t, w, &resp)

			if c.want == nil {
				require.Nil(t, resp.Errors)
				return
			}

			require.NotEmpty(t, resp.Errors)
			require.Equal(t, *c.want, resp.Errors[0].Message)
		})
	}
}
