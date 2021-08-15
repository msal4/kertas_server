package graph_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/msal4/hassah_school_server/auth"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/ent/user"
	"github.com/msal4/hassah_school_server/graph"
	"github.com/msal4/hassah_school_server/graph/model"
	"github.com/msal4/hassah_school_server/service"
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
	srv := graph.NewServer(s, false)
	ctx := context.Background()

	u := createSuperAdmin(ctx, s, "hello23super")
	data, err := auth.GenerateTokens(*u, s.Config.AuthConfig)
	require.NoError(t, err)

	operations := `{ users { totalCount pageInfo { hasNextPage hasPreviousPage startCursor endCursor } edges { node { id } cursor } } }`

	t.Run("school admin authorized", func(t *testing.T) {
		var resp errsResponse

		r := createRequest(t, operations, "{}")
		w := httptest.NewRecorder()

		setAuth(r, data.AccessToken)

		srv.ServeHTTP(w, r)

		parseBody(t, w, &resp)

		require.Nil(t, resp.Errors)
	})

	t.Run("other roles unauthorized", func(t *testing.T) {
		var resp errsResponse

		r := createRequest(t, operations, "{}")
		w := httptest.NewRecorder()

		u.Role = user.RoleSchoolAdmin

		data = genTokens(t, u, s)

		setAuth(r, data.AccessToken)

		srv.ServeHTTP(w, r)

		parseBody(t, w, &resp)

		require.NotEmpty(t, resp.Errors)
		require.Equal(t, auth.UnauthorizedErr.Error(), resp.Errors[0].Message)

		r = createRequest(t, operations, "{}")
		w = httptest.NewRecorder()

		u.Role = user.RoleStudent

		data = genTokens(t, u, s)

		setAuth(r, data.AccessToken)

		srv.ServeHTTP(w, r)

		parseBody(t, w, &resp)

		require.NotEmpty(t, resp.Errors)
		require.Equal(t, auth.UnauthorizedErr.Error(), resp.Errors[0].Message)
	})
}

func setAuth(r *http.Request, ac string) {
	r.Header.Set("authorization", "Bearer "+ac)
}

func genTokens(t testing.TB, u *ent.User, s *service.Service) *model.AuthData {
	data, err := auth.GenerateTokens(*u, s.Config.AuthConfig)
	require.NoError(t, err)

	return data
}
