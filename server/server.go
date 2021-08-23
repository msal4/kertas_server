package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/debug"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/websocket"
	_ "github.com/lib/pq"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/msal4/hassah_school_server/auth"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/ent/user"
	"github.com/msal4/hassah_school_server/server/generated"
	"github.com/msal4/hassah_school_server/server/model"
	"github.com/msal4/hassah_school_server/service"
)

type Config struct {
	DatabaseURL        string `yaml:"database_url" env:"DATABASE_URL"`
	DatabaseDialect    string `yam:"database_dialect" env:"DATABASE_DIALECT,default=postgres"`
	Port               int    `yaml:"port" env:"PORT"`
	Debug              bool   `yaml:"debug" env:"DEBUG"`
	SuperAdminUsername string `yaml:"super_admin_username" env:"SUPER_ADMIN_USERNAME"`
	SuperAdminPassword string `yaml:"super_admin_password" env:"SUPER_ADMIN_PASSWORD"`

	Minio struct {
		Endpoint  string `yaml:"endpoint" env:"MINIO_ENDPOINT"`
		AccessKey string `yaml:"access_key" env:"MINIO_ACCESS_KEY"`
		Token     string `yaml:"token" env:"MINIO_TOKEN"`
	} `yaml:"minio"`

	service.Config `yaml:"service"`
}

func NewDefaultServer(cfg Config) (*http.ServeMux, error) {
	if cfg.DatabaseDialect == "" {
		cfg.DatabaseDialect = dialect.Postgres
	}

	entOpts := []ent.Option{}
	if cfg.Debug {
		entOpts = append(entOpts, ent.Debug())
	}

	ec, err := ent.Open(cfg.DatabaseDialect, cfg.DatabaseURL, entOpts...)
	if err != nil {
		return nil, fmt.Errorf("establishing db connection: %v", err)
	}

	mc, err := minio.New(cfg.Minio.Endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(cfg.Minio.AccessKey, cfg.Minio.Token, ""),
	})
	if err != nil {
		return nil, fmt.Errorf("instantiating minio client: %v", err)
	}

	s, err := service.New(ec, mc, &cfg.Config)
	if err != nil {
		return nil, fmt.Errorf("initializing service: %v", err)
	}

	if err := createDefaultAdminIfNotExists(context.Background(), s, cfg); err != nil {
		return nil, fmt.Errorf("creating super admin: %v", err)
	}

	return NewServer(s, cfg.Debug), nil
}

func createDefaultAdminIfNotExists(ctx context.Context, s *service.Service, cfg Config) error {
	if s.EC.User.Query().Where(user.RoleEQ(user.RoleSuperAdmin)).CountX(context.Background()) > 0 {
		return nil
	}

	log.Println("No super admins found, creating the default admin...")
	u, err := s.AddUser(ctx, model.AddUserInput{
		Name:     "Admin",
		Username: cfg.SuperAdminUsername,
		Password: cfg.SuperAdminPassword,
		Phone:    "07712345678",
		Role:     user.RoleSuperAdmin,
		Active:   true,
	})
	if err != nil {
		return err
	}
	log.Printf("Default super admin created with username: %q and password: %q.\n", u.Username, u.Password)

	return nil
}

func NewServer(s *service.Service, debg bool) *http.ServeMux {
	r := http.NewServeMux()
	srv := handler.New(
		generated.NewExecutableSchema(generated.Config{
			Resolvers: &Resolver{s: s},
		}),
	)

	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		InitFunc: func(ctx context.Context, initPayload transport.InitPayload) (context.Context, error) {
			return auth.ParseAuth(ctx, initPayload.Authorization(), s.Config.AccessSecretKey)
		},
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

	srv.SetQueryCache(lru.New(1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	srv.Use(entgql.Transactioner{TxOpener: s.EC})
	if debg {
		srv.Use(&debug.Tracer{})
	}
	r.Handle("/", playground.Handler("Hassah School", "/graphql"))
	r.Handle("/graphql", auth.Middleware(srv, s.Config.AccessSecretKey))

	return r
}
