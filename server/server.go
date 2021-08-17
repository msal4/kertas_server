package server

import (
	"fmt"
	"net/http"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/debug"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/lib/pq"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/msal4/hassah_school_server/auth"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/server/generated"
	"github.com/msal4/hassah_school_server/service"
)

type Config struct {
	DatabaseURL     string `yaml:"database_url" env:"DATABASE_URL"`
	DatabaseDialect string `yam:"database_dialect" env:"DATABASE_DIALECT,default=postgres"`
	Port            int    `yaml:"port" env:"PORT"`
	Debug           bool   `yaml:"debug" env:"DEBUG"`

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

	ec, err := ent.Open(cfg.DatabaseDialect, cfg.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("establishing db connection: %v", err)
	}

	mc, err := minio.New(cfg.Minio.Endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(cfg.Minio.AccessKey, cfg.Minio.Token, ""),
	})
	if err != nil {
		return nil, fmt.Errorf("instantiating minio client: %v", err)
	}

	s, err := service.New(ec, mc, nil)
	if err != nil {
		return nil, fmt.Errorf("initializing service: %v", err)
	}

	return NewServer(s, cfg.Debug), nil
}

func NewServer(s *service.Service, debg bool) *http.ServeMux {
	r := http.NewServeMux()
	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(generated.Config{
			Resolvers: &Resolver{s: s},
		}),
	)
	srv.Use(entgql.Transactioner{TxOpener: s.EC})
	if debg {
		srv.Use(&debug.Tracer{})
	}
	r.Handle("/", playground.Handler("Hassah School", "/graphql"))
	r.Handle("/graphql", auth.Middleware(srv, s.Config.AccessSecretKey))

	return r
}
