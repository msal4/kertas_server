package graph_test

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"

	"entgo.io/ent/dialect"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/ent/enttest"
	"github.com/msal4/hassah_school_server/service"
	"github.com/msal4/hassah_school_server/util"
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

var randomSource = rand.NewSource(time.Now().Unix())

func newService(t *testing.T) *service.Service {
	db := util.RandomString(randomSource, 6) + time.Now().Format("04-05")

	ec := enttest.Open(t, dialect.SQLite, fmt.Sprintf("file:%s?mode=memory&cache=shared&_fk=1", db), enttest.WithOptions(ent.Log(t.Log)))
	s, err := service.New(ec, mc, nil)
	require.NoError(t, err)
	return s
}
