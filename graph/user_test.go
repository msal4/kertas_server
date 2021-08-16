package graph_test

import (
	"context"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/msal4/hassah_school_server/auth"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/ent/user"
	"github.com/msal4/hassah_school_server/graph"
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
	srv := graph.NewServer(s, false)
	ctx := context.Background()

	suAdmin := createSuperAdmin(ctx, s, "hello23super")
	schAdmin := *suAdmin
	schAdmin.Role = user.RoleSchoolAdmin
	teacher := schAdmin
	teacher.Role = user.RoleTeacher
	student := schAdmin
	student.Role = user.RoleStudent

	operations := `{ users { totalCount pageInfo { hasNextPage hasPreviousPage startCursor endCursor } edges { node { id } cursor } } }`

	cases := []struct {
		desc string
		user ent.User
		want *string
	}{
		{"super admin is authorized", *suAdmin, nil},
		{"school admin is not authorized", schAdmin, ptr.Str(auth.UnauthorizedErr.Error())},
		{"teacher is not authorized", teacher, ptr.Str(auth.UnauthorizedErr.Error())},
		{"student is not authorized", student, ptr.Str(auth.UnauthorizedErr.Error())},
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			var resp errsResponse

			r := createRequest(t, operations, "{}")
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

func TestAddUser(t *testing.T) {
	s := newService(t)
	defer s.EC.Close()
	srv := graph.NewServer(s, false)
	ctx := context.Background()

	suAdmin := createSuperAdmin(ctx, s, "hello23super")
	schAdmin := *suAdmin
	schAdmin.Role = user.RoleSchoolAdmin
	teacher := schAdmin
	teacher.Role = user.RoleTeacher
	student := schAdmin
	student.Role = user.RoleStudent

	sch := s.EC.School.Create().SetName("test school").SetDirectory("test_dir").SetImage("test/image").SaveX(ctx)
	stage := s.EC.Stage.Create().SetName("2nd").SetTuitionAmount(122).SetSchool(sch).SaveX(ctx)

	operations := `{
"query": "mutation ($image: Upload!) { addUser(input: {stage_id: \"%s\", name: \"a test user\", phone: \"077059333812\", username: \"minamo123%d\", password: \"helo234444488@@@@8\" image: $image}) { id name updatedAt }}",
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
