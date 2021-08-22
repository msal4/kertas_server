package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/ent/user"
	"github.com/msal4/hassah_school_server/server/model"
)

var UnauthorizedErr = errors.New("unauthorized")

type AccessClaims struct {
	*jwt.StandardClaims

	UserID uuid.UUID `json:"user_id"`
	Role   user.Role `json:"role"`
}

type RefreshClaims struct {
	*jwt.StandardClaims

	UserID       uuid.UUID `json:"user_id"`
	TokenVersion int       `json:"token_version"`
}

type userData struct {
	ID   uuid.UUID
	Role user.Role
}

type ctxKey string

const userCtxKey ctxKey = "user"

func UserForContext(ctx context.Context) (userData, bool) {
	data, ok := ctx.Value(userCtxKey).(userData)
	return data, ok
}

func IsAuthorized(ctx context.Context, allowedRoles ...user.Role) bool {
	u, ok := UserForContext(ctx)
	if !ok {
		return false
	}

	if len(allowedRoles) == 0 {
		return true
	}

	for _, r := range allowedRoles {
		if r == u.Role {
			return true
		}
	}

	return false
}

func IsAdmin(ctx context.Context) bool {
	return IsAuthorized(ctx, user.RoleSuperAdmin, user.RoleSchoolAdmin)
}

func IsSuperAdmin(ctx context.Context) bool {
	return IsAuthorized(ctx, user.RoleSuperAdmin)
}

func IsSchoolAdmin(ctx context.Context) bool {
	return IsAuthorized(ctx, user.RoleSchoolAdmin)
}

func IsUser(ctx context.Context) bool {
	return IsAuthorized(ctx, user.RoleTeacher, user.RoleStudent)
}

func IsTeacher(ctx context.Context) bool {
	return IsAuthorized(ctx, user.RoleTeacher)
}

func IsStudent(ctx context.Context) bool {
	return IsAuthorized(ctx, user.RoleStudent)
}

func ParseAuth(ctx context.Context, authHeader, accessKey string) (context.Context, error) {
	authHeader = strings.TrimSpace(authHeader)
	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenStr == "" {
		return ctx, nil
	}

	getSecret := func(token *jwt.Token) (interface{}, error) {
		return []byte(accessKey), nil
	}

	var claims AccessClaims
	t, err := jwt.ParseWithClaims(tokenStr, &claims, getSecret)
	if err != nil {
		return ctx, fmt.Errorf("invalid token: %v", err)
	}

	if !t.Valid {
		return ctx, errors.New("invalid token")
	}

	return context.WithValue(ctx, userCtxKey, userData{ID: claims.UserID, Role: claims.Role}), nil
}

func Middleware(h http.Handler, accessKey string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, err := ParseAuth(r.Context(), r.Header.Get("authorization"), accessKey)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
		}

		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

type AuthConfig struct {
	// AccessSecretKey is the key used to sign the access token.
	AccessSecretKey string `yaml:"access_secret_key" env:"ACCESS_SECRET_KEY"`

	// RefreshSecretKey is the key used to sign the refresh token.
	RefreshSecretKey string `yaml:"refresh_secret_key" env:"REFRESH_SECRET_KEY"`

	// AccessTokenLifetime is the duration used to determine the expiration date for the access token.
	AccessTokenLifetime time.Duration `yaml:"access_token_lifetime" env:"ACCESS_TOKEN_LIFETIME"`

	// RefreshTokenLifetime is the duration used to determine the expiration date for the refresh token.
	RefreshTokenLifetime time.Duration `yaml:"refresh_token_lifetime" env:"REFRESH_TOKEN_LIFETIME"`
}

func GenerateTokens(u ent.User, cfg AuthConfig) (*model.AuthData, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, AccessClaims{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(cfg.AccessTokenLifetime).Unix(),
		},
		UserID: u.ID,
		Role:   u.Role,
	})
	access, err := token.SignedString([]byte(cfg.AccessSecretKey))
	if err != nil {
		return nil, err
	}

	token = jwt.NewWithClaims(jwt.SigningMethodHS256, RefreshClaims{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(cfg.RefreshTokenLifetime).Unix(),
		},
		UserID:       u.ID,
		TokenVersion: u.TokenVersion,
	})

	refresh, err := token.SignedString([]byte(cfg.RefreshSecretKey))
	if err != nil {
		return nil, err
	}

	return &model.AuthData{AccessToken: access, RefreshToken: refresh}, nil
}
