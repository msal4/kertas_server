package service_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"entgo.io/ent/dialect"
	_ "github.com/mattn/go-sqlite3"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/ent/enttest"
	"github.com/msal4/hassah_school_server/service"
	"github.com/stretchr/testify/require"
)

var mc *minio.Client

func TestMain(m *testing.M) {
	var err error
	mc, err = minio.New("localhost:9000", &minio.Options{
		Creds: credentials.NewStaticV4("minioadmin", "minioadmin", ""),
	})
	if err != nil {
		log.Fatalf("initializing minio client: %v", err)
	}
	if _, err := mc.ListBuckets(context.Background()); err != nil {
		log.Fatalf("connecting to minio: %v", err)
	}

	os.Exit(m.Run())
}

func newService(t *testing.T, db string) *service.Service {
	ec := enttest.Open(t, dialect.SQLite, fmt.Sprintf("file:%s?mode=memory&cache=shared&_fk=1", db), enttest.WithOptions(ent.Log(t.Log)))
	s, err := service.New(ec, mc, nil)
	require.NoError(t, err)
	return s
}
