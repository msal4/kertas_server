package graph_test

import (
	"context"
	"log"
	"math/rand"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	_ "github.com/mattn/go-sqlite3"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/ent/enttest"
	"github.com/msal4/hassah_school_server/graph"
)

func TestSchools_List(t *testing.T) {
	ec := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", enttest.WithOptions(ent.Log(t.Log)))
	defer ec.Close()

	ctx := context.Background()

	ec.School.Create().SetName("schoola").SetImage("migmo").SaveX(ctx)

	mc, err := minio.New("localhost:9000", &minio.Options{
		Creds: credentials.NewStaticV4("minioadmin", "minioadmin", ""),
	})
	if err != nil {
		t.Fatalf("instantiating minio client: %v", err)
	}
	if _, err := mc.ListBuckets(ctx); err != nil {
		t.Fatalf("connecting to minio: %v", err)
	}

	gc := client.New(handler.NewDefaultServer(graph.NewSchema(ec, mc, rand.NewSource(0))))

	t.Run("list", func(t *testing.T) {
		var resp struct {
			Schools ent.SchoolConnection
		}
		gc.MustPost(`{
  schools {
    totalCount
    edges {
      node {
        id
        status
        name
        image
      }
    }
  }
}`, &resp)

		log.Println(resp)
	})
	t.Run("order", func(t *testing.T) {})
	t.Run("filter", func(t *testing.T) {})
	t.Run("search", func(t *testing.T) {})
	t.Run("not authorized", func(t *testing.T) {})
}
