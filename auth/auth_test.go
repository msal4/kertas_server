package auth_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/msal4/hassah_school_server/auth"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/ent/user"
	"github.com/stretchr/testify/require"
)

func TestGenerateToken(t *testing.T) {
	cfg := auth.AuthConfig{}

	u := ent.User{
		ID:           uuid.New(),
		Role:         user.RoleSuperAdmin,
		TokenVersion: 20,
	}

	data, err := auth.GenerateTokens(u, cfg)
	require.NoError(t, err)
	require.NotEmpty(t, data.AccessToken)
	require.NotEmpty(t, data.RefreshToken)

	var accessClaims auth.AccessClaims
	_, err = jwt.ParseWithClaims(data.AccessToken, &accessClaims, func(t *jwt.Token) (interface{}, error) {
		return cfg.AccessSecretKey, nil
	})
	require.NoError(t, err)

	require.Equal(t, u.ID, accessClaims.UserID)
	require.Equal(t, u.Role, accessClaims.Role)

	var refreshClaims auth.RefreshClaims
	_, err = jwt.ParseWithClaims(data.AccessToken, &refreshClaims, func(t *jwt.Token) (interface{}, error) {
		return cfg.RefreshSecretKey, nil
	})
	require.NoError(t, err)

	require.Equal(t, u.ID, refreshClaims.UserID)
	require.Equal(t, u.TokenVersion, refreshClaims.TokenVersion)
}

func TestMiddleware(t *testing.T) {
	t.Run("expired token", func(t *testing.T) {
		cfg := auth.AuthConfig{
			AccessSecretKey:      []byte("secret"),
			RefreshSecretKey:     []byte("secret"),
			AccessTokenLifetime:  -time.Hour,
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

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/whatever", nil)

		r.Header.Set("authorization", "Bearer "+data.AccessToken)

		srv := auth.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}), cfg.AccessSecretKey)

		srv.ServeHTTP(w, r)

		require.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("no token", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/whatever", nil)

		srv := auth.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}), []byte(""))

		srv.ServeHTTP(w, r)

		require.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("authenticated", func(t *testing.T) {
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

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/whatever", nil)

		r.Header.Set("authorization", "Bearer "+data.AccessToken)

		srv := auth.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			got, ok := auth.UserForContext(r.Context())

			require.True(t, ok)

			require.Equal(t, u.ID, got.ID)
			require.Equal(t, u.Role, got.Role)
		}), cfg.AccessSecretKey)

		srv.ServeHTTP(w, r)
	})
}
