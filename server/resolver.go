package server

import (
	_ "image/png"

	"github.com/msal4/hassah_school_server/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	s *service.Service
}
