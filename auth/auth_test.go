package auth_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/msal4/hassah_school_server/auth"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/ent/user"
	"github.com/stretchr/testify/require"
)

func TestGenerateToken(t *testing.T) {
	u := ent.User{
		ID:   uuid.New(),
		Role: user.RoleSuperAdmin,
	}

	data, err := auth.GenerateTokens(u, auth.AuthConfig{})
	require.NoError(t, err)
	require.NotEmpty(t, data.AccessToken)
	require.NotEmpty(t, data.RefreshToken)
}

func TestMiddleware(t *testing.T) {
	cfg := auth.AuthConfig{
		AccessSecretKey:      []byte("secret"),
		RefreshSecretKey:     []byte("secret"),
		AccessTokenLifetime:  time.Hour,
		RefreshTokenLifetime: time.Hour,
	}

	u := ent.User{
		ID:   uuid.New(),
		Role: user.RoleSuperAdmin,
	}

	data, err := auth.GenerateTokens(u, cfg)
	require.NoError(t, err)
	require.NotEmpty(t, data.AccessToken)
	require.NotEmpty(t, data.RefreshToken)

	srv := auth.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got, ok := auth.UserForContext(r.Context())

		require.True(t, ok)

		require.Equal(t, u.ID, got.ID)
		require.Equal(t, u.Role, got.Role)
	}), cfg.AccessSecretKey)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/whatever", nil)

	ctx := context.Background()
	r = r.WithContext(ctx)

	r.Header.Set("authorization", "Bearer "+data.AccessToken)

	srv.ServeHTTP(w, r)
}
