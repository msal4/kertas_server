package service_test

import (
	"context"
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/99designs/gqlgen/graphql"
	"github.com/minio/minio-go/v7"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/ent/schema"
	"github.com/msal4/hassah_school_server/ent/school"
	"github.com/msal4/hassah_school_server/graph/model"
	"github.com/msal4/hassah_school_server/service"
	"github.com/msal4/hassah_school_server/util/ptr"
	"github.com/stretchr/testify/require"
)

func TestSchoolList(t *testing.T) {
	s := newService(t, "1")
	ctx := context.Background()

	t.Run("empty", func(t *testing.T) {
		defer s.EC.School.Delete().ExecX(ctx)

		conn, err := s.SchoolList(ctx, service.SchoolListOptions{})
		require.NoError(t, err)
		require.NotNil(t, conn)
		require.Zero(t, conn.TotalCount)
		require.Empty(t, conn.Edges)
	})

	t.Run("not empty", func(t *testing.T) {
		defer s.EC.School.Delete().ExecX(ctx)

		want := s.EC.School.Create().SetName("school 1").SetImage("s/image").SetDirectory("test_dir").SaveX(ctx)

		conn, err := s.SchoolList(ctx, service.SchoolListOptions{})
		require.NoError(t, err)
		require.NotNil(t, conn)
		require.Equal(t, 1, conn.TotalCount)
		require.Len(t, conn.Edges, 1)
		edge := conn.Edges[0]
		require.NotNil(t, edge)
		require.NotNil(t, edge.Cursor)
		require.NotNil(t, edge.Node)
		require.Equal(t, edge.Node.ID, want.ID)
		require.Equal(t, edge.Node.Name, want.Name)
	})

	t.Run("order & filter", func(t *testing.T) {
		defer s.EC.School.Delete().ExecX(ctx)

		s.EC.School.Create().SetName("school 1").SetImage("s/image1").SetDirectory("test_dir").SaveX(ctx)
		s.EC.School.Create().SetName("school 2").SetImage("s/image2").SetDirectory("test_dir").SetStatus(schema.StatusDisabled).SaveX(ctx)
		s.EC.School.Create().SetName("school 3").SetImage("s/image3").SetDirectory("test_dir").SetStatus(schema.StatusDeleted).SaveX(ctx)
		s.EC.School.Create().SetName("school 4").SetImage("s/image4").SetDirectory("test_dir").SetStatus(schema.StatusDisabled).SaveX(ctx)

		b := s.EC.School.Query().Order(ent.Asc(school.FieldCreatedAt))
		want := b.AllX(ctx)

		conn, err := s.SchoolList(ctx, service.SchoolListOptions{
			OrderBy: &ent.SchoolOrder{Field: ent.SchoolOrderFieldCreatedAt, Direction: ent.OrderDirectionAsc},
		})
		require.NoError(t, err)
		require.NotNil(t, conn)
		require.Equal(t, len(want), conn.TotalCount)
		require.Len(t, conn.Edges, len(want))
		for i, edge := range conn.Edges {
			require.NotNil(t, edge)
			require.NotNil(t, edge.Cursor)
			require.NotNil(t, edge.Node)
			require.Equal(t, edge.Node.ID, want[i].ID)
			require.Equal(t, edge.Node.Name, want[i].Name)
			require.Equal(t, edge.Node.Status, want[i].Status)
			require.Equal(t, edge.Node.Image, want[i].Image)
			require.Equal(t, edge.Node.CreatedAt, want[i].CreatedAt)
		}

		want = b.Where(school.StatusEQ(schema.StatusDisabled)).AllX(ctx)

		conn, err = s.SchoolList(ctx, service.SchoolListOptions{
			Where:   &ent.SchoolWhereInput{Status: ptr.Status(schema.StatusDisabled)},
			OrderBy: &ent.SchoolOrder{Field: ent.SchoolOrderFieldCreatedAt, Direction: ent.OrderDirectionAsc},
		})
		require.NoError(t, err)

		require.Equal(t, len(want), conn.TotalCount)
		require.Len(t, conn.Edges, len(want))

		for i, w := range want {
			e := conn.Edges[i]

			require.NotNil(t, e)
			require.NotNil(t, e.Cursor)
			require.NotNil(t, e.Node)
			require.Equal(t, e.Node.ID, w.ID)
			require.Equal(t, e.Node.Name, w.Name)
			require.Equal(t, e.Node.Status, w.Status)
			require.Equal(t, e.Node.Image, w.Image)
			require.Equal(t, e.Node.CreatedAt, w.CreatedAt)
		}

	})
}

func TestSchoolAdd(t *testing.T) {
	s := newService(t, "2")
	ctx := context.Background()

	t.Run("without image", func(t *testing.T) {
		got, err := s.SchoolAdd(ctx, model.CreateSchoolInput{Name: "test school"})
		require.Error(t, err)
		require.Nil(t, got)
	})

	t.Run("with image", func(t *testing.T) {
		f := openFile(t, "../testfiles/stanford.png")
		defer f.Close()
		got, err := s.SchoolAdd(ctx, model.CreateSchoolInput{
			Name:   "test school",
			Image:  graphql.Upload{File: f, Filename: f.File.Name(), ContentType: f.contentType, Size: f.Size()},
			Status: schema.StatusActive,
		})
		require.NoError(t, err)
		require.NotNil(t, got)
		require.NotEmpty(t, got.Image)
		info, err := s.MC.StatObject(ctx, s.Config.RootBucket, got.Image, minio.StatObjectOptions{})
		require.NoError(t, err)
		require.Equal(t, "image/jpeg", info.ContentType)
	})

	t.Run("with invalid image", func(t *testing.T) {
		f := openFile(t, "../testfiles/file.txt")
		defer f.Close()
		got, err := s.SchoolAdd(ctx, model.CreateSchoolInput{
			Name:   "test school",
			Image:  graphql.Upload{File: f, Filename: f.File.Name(), ContentType: f.contentType, Size: f.Size()},
			Status: schema.StatusActive,
		})
		require.Error(t, err)
		require.Nil(t, got)
	})
}

type file struct {
	*os.File
	fs.FileInfo
	contentType string
}

func openFile(t *testing.T, name string) *file {
	f, err := os.Open(name)
	require.NoError(t, err)
	contentType, err := getFileContentType(f)
	require.NoError(t, err)
	stat, err := f.Stat()
	require.NoError(t, err)
	f.Seek(0, 0)

	return &file{File: f, contentType: contentType, FileInfo: stat}
}

func getFileContentType(f *os.File) (string, error) {
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	return http.DetectContentType(data), nil
}

func TestSchoolDelete(t *testing.T) {
}
