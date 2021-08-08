package main

import (
	"context"
	"flag"
	"log"

	"entgo.io/ent/dialect"
	_ "github.com/lib/pq"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/ent/migrate"
)

var debug *bool

func init() {
	debug = flag.Bool("debug", false, "Run migration in debug mode")
	flag.Parse()
}

func main() {
	client, err := ent.Open(dialect.Postgres, "postgres://postgres@localhost:5432/school?sslmode=disable")
	if err != nil {
		log.Fatalf("failed establishing connection: %v", err)
	}
	defer client.Close()

	if debug != nil && *debug {
		err = client.Debug().Schema.Create(context.Background(), migrate.WithDropIndex(true), migrate.WithDropColumn(true))
	} else {
		err = client.Schema.Create(context.Background())
	}
	if err != nil {
		log.Fatalf("failed creating schema: %v", err)
	}
}
