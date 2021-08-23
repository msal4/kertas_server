package service_test

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"

	"entgo.io/ent/dialect"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/ent/enttest"
	"github.com/msal4/hassah_school_server/ent/user"
	"github.com/msal4/hassah_school_server/service"
	"github.com/msal4/hassah_school_server/util"
	"github.com/stretchr/testify/require"
)

var mc *minio.Client
var ctx = context.Background()

func TestMain(m *testing.M) {
	godotenv.Load("../.env")

	var err error
	mc, err = minio.New(os.Getenv("MINIO_ENDPOINT"), &minio.Options{
		Creds: credentials.NewStaticV4(os.Getenv("MINIO_ACCESS_KEY"), os.Getenv("MINIO_TOKEN"), ""),
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

func createSuperAdmin(ctx context.Context, s *service.Service, username string) *ent.User {
	return s.EC.User.Create().SetName("testu4serd" + username).SetUsername(username).
		SetPhone("077059333812").SetDirectory("diresss22" + username).SetPassword("mipassword22@@@@5").
		SetRole(user.RoleSuperAdmin).SaveX(ctx)
}

func createSchoolAdmin(ctx context.Context, s *service.Service, username string, sch *ent.School) *ent.User {
	return s.EC.User.Create().SetName("testu4serd" + username).SetUsername(username).
		SetPhone("077059333812").SetDirectory("diresss22" + username).SetPassword("mipassword22@@@@5").SetSchool(sch).SetRole(user.RoleSchoolAdmin).SaveX(ctx)
}

func createTeacher(ctx context.Context, s *service.Service, username string, sch *ent.School) *ent.User {
	return s.EC.User.Create().SetName("testu4serd" + username).SetUsername(username).
		SetPhone("077059333812").SetDirectory("diresss22" + username).SetPassword("mipassword22@@@@5").SetSchool(sch).SetRole(user.RoleTeacher).SaveX(ctx)
}

func createStudent(ctx context.Context, s *service.Service, username string, sch *ent.School, stg ...*ent.Stage) *ent.User {
	if len(stg) == 0 {
		stg = append(stg, s.EC.Stage.Create().SetName("2nd").SetDirectory("hello").
			SetTuitionAmount(122).SetSchool(sch).SaveX(ctx))
	}
	return s.EC.User.Create().SetName("test userd" + username).SetUsername(username).
		SetPhone("077059333812").SetDirectory("diresss22" + username).SetPassword("mipassword22@@@@5").SetSchool(sch).SetStage(stg[0]).SaveX(ctx)
}

func createSchool(ctx context.Context, s *service.Service, name, image string) *ent.School {
	return s.EC.School.Create().SetName(name).SetImage(image).SetDirectory("test_dir").SaveX(ctx)
}

func createStage(ctx context.Context, s *service.Service, name string, tuition int) *ent.Stage {
	sch := createSchool(ctx, s, "school for"+name, "image/"+name)
	return s.EC.Stage.Create().SetName(name).SetDirectory("testdir" + name).SetTuitionAmount(tuition).
		SetSchool(sch).SaveX(ctx)
}
