package main

import (
	"context"
	"flag"
	"log"
	"math/rand"
	"net/http"
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/debug"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/lib/pq"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/graph"
)

var debg bool

func init() {
	d := flag.Bool("debug", false, "Run server in debug mode")
	flag.Parse()

	debg = d != nil && *d
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

	ctx := context.Background()
	if _, err := mc.ListBuckets(ctx); err != nil {
		log.Fatalf("connecting to minio: %v", err)
	}
	exists, err := mc.BucketExists(ctx, "images")
	if err != nil {
		log.Fatalf("checking if images bucket exists: %v", err)
	}
	if !exists {
		log.Println(`bucket "images" does not exist, creating one...`)
		err := mc.MakeBucket(ctx, "images", minio.MakeBucketOptions{})
		if err != nil {
			log.Fatalf(`making "images" bucket: %v`, err)
		}
		log.Println(`created bucket "images".`)
	}

	srv := handler.NewDefaultServer(graph.NewSchema(ec, mc, rand.NewSource(time.Now().Unix())))
	srv.Use(entgql.Transactioner{TxOpener: ec})
	if debg {
		srv.Use(&debug.Tracer{})
	}
	http.Handle("/", playground.Handler("School", "/graphql"))
	http.Handle("/graphql", srv)

	log.Println("listening on :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
