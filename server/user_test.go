package server_test

import (
	"context"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/msal4/hassah_school_server/auth"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/ent/user"
	"github.com/msal4/hassah_school_server/server"
	"github.com/msal4/hassah_school_server/service"
	"github.com/msal4/hassah_school_server/testutil"
	"github.com/msal4/hassah_school_server/util/ptr"
	"github.com/stretchr/testify/require"
)

type errsResponse struct {
	Errors []struct {
		Message string   `json:"message"`
		Path    []string `json:"path"`
	} `json:"errors,omitempty"`
}

func TestUsers(t *testing.T) {
	s := newService(t)
	defer s.EC.Close()

	srv := server.NewServer(s, false)
	ctx := context.Background()

	suAdmin := createSuperAdmin(ctx, s, "hello23super")
	schAdmin := createSchoolAdmin(ctx, s, "hellosdjflksdjflsdj")
	teacher := createTeacher(ctx, s, "jsjdflks4444")
	student := createStudent(ctx, s, "jflksjdflksdjfklsjd")

	operations := `{ users { totalCount pageInfo { hasNextPage hasPreviousPage startCursor endCursor } edges { node { id } cursor } } }`

	cases := []struct {
		desc string
		user *ent.User
		want *string
	}{
		{"super admin is authorized", suAdmin, nil},
		{"school admin is not authorized", schAdmin, nil},
		{"teacher is not authorized", teacher, ptr.Str(auth.UnauthorizedErr.Error())},
		{"student is not authorized", student, ptr.Str(auth.UnauthorizedErr.Error())},
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			var resp errsResponse

			r := createRequest(t, operations, "{}")
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

func TestAddUser(t *testing.T) {
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

	sch := s.EC.School.Create().SetName("test school").SetDirectory("test_dir").SetImage("test/image").SaveX(ctx)
	stage := s.EC.Stage.Create().SetName("2nd").SetTuitionAmount(122).SetSchool(sch).SetDirectory("fsdjfld").SaveX(ctx)

	operations := `{
"query": "mutation ($image: Upload!) { addUser(input: {stageID: \"%s\", name: \"a test user\", phone: \"077059333812\", username: \"minamo123%d\", password: \"helo234444488@@@@8\" image: $image}) { id name updatedAt }}",
			"variables": {"image": null}
		}`

	cases := []struct {
		desc string
		user ent.User
		want *string
	}{
		{"super admin is authorized", *suAdmin, nil},
		{"school admin is not authorized", schAdmin, nil},
		{"teacher is not authorized", teacher, ptr.Str(auth.UnauthorizedErr.Error())},
		{"student is not authorized", student, ptr.Str(auth.UnauthorizedErr.Error())},
	}

	f := testutil.OpenFile(t, "../testfiles/harvard.jpg")

	for i, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			var resp errsResponse

			f.Seek(0, 0)

			r := createMultipartRequest(t, fmt.Sprintf(operations, stage.ID.String(), i), `{"0": ["variables.image"]}`, file{File: f.File, mapKey: "0"})
			w := httptest.NewRecorder()

			data := genTokens(t, &c.user, s)

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

func TestUpdateUser(t *testing.T) {
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

	u := createStudent(ctx, s, "random234user")

	operations := `mutation { updateUser(id: %q, input: {name: "a test user"}) { id name updatedAt }}`

	cases := []struct {
		desc string
		user ent.User
		want *string
	}{
		{"super admin is authorized", *suAdmin, nil},
		{"school admin is not authorized", schAdmin, nil},
		{"teacher is not authorized", teacher, ptr.Str(auth.UnauthorizedErr.Error())},
		{"student is not authorized", student, ptr.Str(auth.UnauthorizedErr.Error())},
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			var resp errsResponse

			r := createRequest(t, fmt.Sprintf(operations, u.ID.String()), "{}")
			w := httptest.NewRecorder()

			data := genTokens(t, &c.user, s)

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

func TestDeleteUser(t *testing.T) {
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

	operations := `mutation { deleteUser(id: %q) }`

	cases := []struct {
		desc string
		user ent.User
		want *string
	}{
		{"super admin is authorized", *suAdmin, nil},
		{"school admin is not authorized", schAdmin, nil},
		{"teacher is not authorized", teacher, ptr.Str(auth.UnauthorizedErr.Error())},
		{"student is not authorized", student, ptr.Str(auth.UnauthorizedErr.Error())},
	}

	for i, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			var resp errsResponse

			u := createStudent(ctx, s, fmt.Sprintf("random234u%dser", i))

			r := createRequest(t, fmt.Sprintf(operations, u.ID.String()), "{}")
			w := httptest.NewRecorder()

			data := genTokens(t, &c.user, s)

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

func TestDeleteUserPermanently(t *testing.T) {
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

	operations := `mutation { deleteUserPermanently(id: %q) }`

	cases := []struct {
		desc string
		user ent.User
		want *string
	}{
		{"super admin is authorized", *suAdmin, nil},
		{"school admin is not authorized", schAdmin, nil},
		{"teacher is not authorized", teacher, ptr.Str(auth.UnauthorizedErr.Error())},
		{"student is not authorized", student, ptr.Str(auth.UnauthorizedErr.Error())},
	}

	for i, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			var resp errsResponse

			u := createStudent(ctx, s, fmt.Sprintf("random234u%dser", i))

			r := createRequest(t, fmt.Sprintf(operations, u.ID.String()), "{}")
			w := httptest.NewRecorder()

			data := genTokens(t, &c.user, s)

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

func TestLoginAdmin(t *testing.T) {
	type response struct {
		Data *struct {
			LoginAdmin *struct {
				AccessToken  string `json:"accessToken"`
				RefreshToken string `json:"refreshToken"`
			} `json:"loginAdmin"`
		} `json:"data"`
		Errors []struct {
			Message string   `json:"message"`
			Path    []string `json:"path"`
		} `json:"errors,omitempty"`
	}
	s := newService(t)
	defer s.EC.Close()

	srv := server.NewServer(s, false)
	ctx := context.Background()

	suAdmin := createSuperAdmin(ctx, s, "hello23super")
	schAdmin := createSchoolAdmin(ctx, s, "hello23sup24")
	teacher := createTeacher(ctx, s, "hellostuteachert22")
	student := createStudent(ctx, s, "heldesnt22")

	operations := `mutation { loginAdmin(input: { username: %q, password: %q }) {accessToken refreshToken} }`

	cases := []struct {
		desc string
		user *ent.User
		want *string
	}{
		{"super admin is allowed", suAdmin, nil},
		{"school admin is allowed", schAdmin, nil},
		{"teacher is not allowed", teacher, ptr.Str(service.NotAllowedErr.Error())},
		{"student is not allowed", student, ptr.Str(service.NotAllowedErr.Error())},
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			var resp response

			r := createRequest(t, fmt.Sprintf(operations, c.user.Username, c.user.Password), "{}")
			w := httptest.NewRecorder()

			srv.ServeHTTP(w, r)

			parseBody(t, w, &resp)

			if c.want == nil {
				require.Nil(t, resp.Errors)
				require.NotNil(t, resp.Data)
				require.NotNil(t, resp.Data.LoginAdmin)
				return
			}

			require.NotEmpty(t, resp.Errors)
			require.Equal(t, *c.want, resp.Errors[0].Message)
		})
	}
}

func TestLoginUser(t *testing.T) {
	type response struct {
		Data *struct {
			LoginUser *struct {
				AccessToken  string `json:"accessToken"`
				RefreshToken string `json:"refreshToken"`
			} `json:"loginUser"`
		} `json:"data"`
		Errors []struct {
			Message string   `json:"message"`
			Path    []string `json:"path"`
		} `json:"errors,omitempty"`
	}
	s := newService(t)
	defer s.EC.Close()

	srv := server.NewServer(s, false)
	ctx := context.Background()

	suAdmin := createSuperAdmin(ctx, s, "hello23super")
	schAdmin := createSchoolAdmin(ctx, s, "hello23sup24")
	teacher := createTeacher(ctx, s, "hellostuteachert22")
	student := createStudent(ctx, s, "heldesnt22")

	operations := `mutation { loginUser(input: { username: %q, password: %q }) {accessToken refreshToken} }`

	cases := []struct {
		desc string
		user *ent.User
		want *string
	}{
		{"teacher is allowed", teacher, nil},
		{"student is allowed", student, nil},
		{"super admin is not allowed", suAdmin, ptr.Str(service.NotAllowedErr.Error())},
		{"school admin is not allowed", schAdmin, ptr.Str(service.NotAllowedErr.Error())},
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			var resp response

			r := createRequest(t, fmt.Sprintf(operations, c.user.Username, c.user.Password), "{}")
			w := httptest.NewRecorder()

			srv.ServeHTTP(w, r)

			parseBody(t, w, &resp)

			if c.want == nil {
				require.Nil(t, resp.Errors)
				require.NotNil(t, resp.Data)
				require.NotNil(t, resp.Data.LoginUser)
				return
			}

			require.NotEmpty(t, resp.Errors)
			require.Equal(t, *c.want, resp.Errors[0].Message)
		})
	}
}
