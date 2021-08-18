package server_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"entgo.io/ent/dialect"
	"github.com/joho/godotenv"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/msal4/hassah_school_server/auth"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/ent/enttest"
	"github.com/msal4/hassah_school_server/ent/user"
	"github.com/msal4/hassah_school_server/server/model"
	"github.com/msal4/hassah_school_server/service"
	"github.com/msal4/hassah_school_server/util"
	"github.com/stretchr/testify/require"
)

var mc *minio.Client

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

type file struct {
	mapKey string
	*os.File
}

func createRequest(t *testing.T, query, variables string) *http.Request {
	b := bytes.NewBuffer([]byte(fmt.Sprintf(`{"query": %q, "variables": %s}`, query, variables)))
	r := httptest.NewRequest(http.MethodPost, "/graphql", b)
	r.Header.Set("content-type", "application/json")

	return r
}

func createMultipartRequest(t *testing.T, operations, mapData string, f file) *http.Request {
	b := new(bytes.Buffer)
	w := multipart.NewWriter(b)
	require.NoError(t, w.WriteField("operations", operations))
	require.NoError(t, w.WriteField("map", mapData))

	ff, err := w.CreateFormFile(f.mapKey, f.Name())
	require.NoError(t, err)
	_, err = io.Copy(ff, f)
	require.NoError(t, err)

	require.NoError(t, w.Close())

	r := httptest.NewRequest(http.MethodPost, "/graphql", b)

	r.Header.Set("content-type", w.FormDataContentType())

	return r
}

func parseBody(t testing.TB, w *httptest.ResponseRecorder, v interface{}) {
	require.NoError(t, json.NewDecoder(w.Body).Decode(v))
}

func setAuth(r *http.Request, ac string) {
	r.Header.Set("authorization", "Bearer "+ac)
}

func genTokens(t testing.TB, u *ent.User, s *service.Service) *model.AuthData {
	data, err := auth.GenerateTokens(*u, s.Config.AuthConfig)
	require.NoError(t, err)

	return data
}

func createSuperAdmin(ctx context.Context, s *service.Service, username string) *ent.User {
	return s.EC.User.Create().SetName("test userd" + username).SetUsername(username).
		SetPhone("077059333812").SetDirectory("diresss22" + username).SetRole(user.RoleSuperAdmin).SetImage("sss").SetPassword("mipassword22@@@@5").SaveX(ctx)
}

func createSchoolAdmin(ctx context.Context, s *service.Service, username string) *ent.User {
	sch := s.EC.School.Create().SetName("schooltest").SetDirectory("fsss").SetImage("fss").SaveX(ctx)
	return s.EC.User.Create().SetName("testu4serd" + username).SetUsername(username).
		SetPhone("077059333812").SetDirectory("diresss22" + username).SetPassword("mipassword22@@@@5").SetSchool(sch).SetRole(user.RoleSchoolAdmin).SaveX(ctx)
}

func createTeacher(ctx context.Context, s *service.Service, username string) *ent.User {
	sch := s.EC.School.Create().SetName("schooltest").SetDirectory("fsss").SetImage("fss").SaveX(ctx)
	return s.EC.User.Create().SetName("testu4serd" + username).SetUsername(username).
		SetPhone("077059333812").SetDirectory("diresss22" + username).SetPassword("mipassword22@@@@5").SetSchool(sch).SetRole(user.RoleTeacher).SaveX(ctx)
}

func createStudent(ctx context.Context, s *service.Service, username string) *ent.User {
	sch := s.EC.School.Create().SetName("schooltest").SetDirectory("fsss").SetImage("fss").SaveX(ctx)
	stage := s.EC.Stage.Create().SetName("2nd").SetDirectory("hello").SetTuitionAmount(122).SetSchool(sch).SaveX(ctx)
	return s.EC.User.Create().SetName("test userd" + username).SetUsername(username).
		SetPhone("077059333812").SetDirectory("diresss22" + username).SetPassword("mipassword22@@@@5").SetSchool(sch).SetStage(stage).SaveX(ctx)
}
