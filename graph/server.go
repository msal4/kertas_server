package graph

import (
	"net/http"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/debug"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/lib/pq"
	"github.com/msal4/hassah_school_server/auth"
	"github.com/msal4/hassah_school_server/graph/generated"
	"github.com/msal4/hassah_school_server/service"
)

func NewServer(s *service.Service, debg bool) http.Handler {
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
