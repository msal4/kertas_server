package graph

import (
	_ "image/png"

	"github.com/99designs/gqlgen/graphql"
	"github.com/msal4/hassah_school_server/graph/generated"
	"github.com/msal4/hassah_school_server/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	s *service.Service
}

func NewSchema(s *service.Service) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{s: s},
	})
}
