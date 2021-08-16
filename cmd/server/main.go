package main

import (
	"flag"
	"log"
	"net/http"

	"entgo.io/ent/dialect"
	_ "github.com/lib/pq"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/graph"
	"github.com/msal4/hassah_school_server/service"
)

var debg *bool

func init() {
	debg = flag.Bool("debug", false, "Run server in debug mode")
	flag.Parse()
}

// minio credentials
const (
	endpoint        = "localhost:9000"
	accessKeyID     = "minioadmin"
	secretAccessKey = "minioadmin"
)

func main() {
	ec, err := ent.Open(dialect.Postgres, "postgres://postgres@localhost:5432/school?sslmode=disable")
	if err != nil {
		log.Fatalf("establishing db connection: %v", err)
	}
	mc, err := minio.New(endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
	})
	if err != nil {
		log.Fatalf("instantiating minio client: %v", err)
	}

	s, err := service.New(ec, mc, nil)
	if err != nil {
		log.Fatalf("initializing service: %v", err)
	}

	srv := graph.NewServer(s, *debg)

	log.Println("listening on :3000")
	log.Fatal(http.ListenAndServe(":3000", srv))
}
