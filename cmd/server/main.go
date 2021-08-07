package main

import (
	"context"
	"log"

	"entgo.io/ent/dialect"
	_ "github.com/lib/pq"
	"github.com/msal4/hassah_school_server/ent"
)

func main() {
	client, err := ent.Open(dialect.Postgres, "postgres://postgres@localhost:5432/school_rewrite?sslmode=disable")
	if err != nil {
		log.Fatalf("failed establishing connection: %v", err)
	}
	ctx := context.Background()
	client.User.Create().SetName("Mohammed").SetUsername("msal").SetPassword("password").SetPhone("770").SaveX(ctx)
}
