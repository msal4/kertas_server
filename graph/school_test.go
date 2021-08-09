package graph_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"

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

	srv := handler.NewDefaultServer(graph.NewSchema(ec, mc, rand.NewSource(0)))

	w := httptest.NewRecorder()

	reqbody := struct {
		OperationName *string  `json:"operationName"`
		Query         string   `json:"query"`
		Variables     struct{} `json:"variables"`
	}{
		Query: `query {
  schools {
    totalCount
    pageInfo {
      hasNextPage
      hasPreviousPage
      startCursor
      endCursor
    }
    edges {
      node {
        id
        name
        image
        status
        createdAt
        updatedAt
      }
      cursor
    }
  }
}`,
	}

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(reqbody)

	r := httptest.NewRequest(http.MethodPost, "/graphql", b)
	r.Header.Set("content-type", "application/json")

	srv.ServeHTTP(w, r)

	fmt.Println(w.Body.String())
}
