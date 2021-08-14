package service_test

import (
	"context"
	"testing"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/ent/schema"
	"github.com/msal4/hassah_school_server/ent/school"
	"github.com/msal4/hassah_school_server/graph/model"
	"github.com/msal4/hassah_school_server/service"
	"github.com/msal4/hassah_school_server/testutil"
	"github.com/msal4/hassah_school_server/util/ptr"
	"github.com/stretchr/testify/require"
)

func TestSchoolList(t *testing.T) {
	s := newService(t)
	defer s.EC.Close()
	ctx := context.Background()

	t.Run("empty", func(t *testing.T) {
		defer s.EC.School.Delete().ExecX(ctx)

		conn, err := s.Schools(ctx, service.SchoolListOptions{})
		require.NoError(t, err)
		require.NotNil(t, conn)
		require.Zero(t, conn.TotalCount)
		require.Empty(t, conn.Edges)
	})

	t.Run("not empty", func(t *testing.T) {
		defer s.EC.School.Delete().ExecX(ctx)

		want := s.EC.School.Create().SetName("school 1").SetImage("s/image").SetDirectory("test_dir").SaveX(ctx)

		conn, err := s.Schools(ctx, service.SchoolListOptions{})
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
		s.EC.School.Create().SetName("school 3").SetImage("s/image3").SetDirectory("test_dir").SaveX(ctx)
		s.EC.School.Create().SetName("school 4").SetImage("s/image4").SetDirectory("test_dir").SetStatus(schema.StatusDisabled).SaveX(ctx)

		b := s.EC.School.Query().Order(ent.Asc(school.FieldCreatedAt))
		want := b.AllX(ctx)

		conn, err := s.Schools(ctx, service.SchoolListOptions{
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

		conn, err = s.Schools(ctx, service.SchoolListOptions{
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
	s := newService(t)
	defer s.EC.Close()
	ctx := context.Background()

	t.Run("without image", func(t *testing.T) {
		got, err := s.AddSchool(ctx, model.AddSchoolInput{Name: "test school"})
		require.Error(t, err)
		require.Nil(t, got)
	})

	t.Run("with image", func(t *testing.T) {
		f := testutil.OpenFile(t, "../testfiles/stanford.png")
		defer f.Close()
		got, err := s.AddSchool(ctx, model.AddSchoolInput{
			Name:   "test school",
			Image:  graphql.Upload{File: f, Filename: f.File.Name(), ContentType: f.ContentType, Size: f.Size()},
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
		f := testutil.OpenFile(t, "../testfiles/file.txt")
		defer f.Close()
		got, err := s.AddSchool(ctx, model.AddSchoolInput{
			Name:   "test school",
			Image:  graphql.Upload{File: f, Filename: f.File.Name(), ContentType: f.ContentType, Size: f.Size()},
			Status: schema.StatusActive,
		})
		require.Error(t, err)
		require.Nil(t, got)
	})
}

func TestSchoolDelete(t *testing.T) {
	s := newService(t)
	defer s.EC.Close()

	ctx := context.Background()

	f := testutil.OpenFile(t, "../testfiles/stanford.png")
	defer f.Close()

	sch, err := s.AddSchool(ctx, model.AddSchoolInput{
		Name:   "test school",
		Image:  graphql.Upload{File: f, Filename: f.File.Name(), ContentType: f.ContentType, Size: f.Size()},
		Status: schema.StatusActive,
	})
	require.NoError(t, err)
	require.NotNil(t, sch)

	t.Run("non-existing school", func(t *testing.T) {
		err := s.DeleteSchool(ctx, uuid.MustParse("2710c203-7842-4356-8d9f-12f9da4722a2"))
		require.Error(t, err)
		_, err = s.EC.School.Query().Where(school.ID(sch.ID)).Only(ctx)
		require.NoError(t, err)
		_, err = s.MC.StatObject(ctx, s.Config.RootBucket, sch.Image, minio.StatObjectOptions{})
		require.NoError(t, err)
	})

	t.Run("existing school", func(t *testing.T) {
		err := s.DeleteSchool(ctx, sch.ID)
		require.NoError(t, err)
		_, err = s.EC.School.Query().Where(school.ID(sch.ID)).Only(ctx)
		require.Error(t, err)

		_, err = s.MC.StatObject(ctx, s.Config.RootBucket, sch.Image, minio.StatObjectOptions{})
		require.Error(t, err)
	})
}

func TestSchoolUpdate(t *testing.T) {
	s := newService(t)
	defer s.EC.Close()

	ctx := context.Background()

	f := testutil.OpenFile(t, "../testfiles/stanford.png")
	defer f.Close()

	sch, err := s.AddSchool(ctx, model.AddSchoolInput{
		Name:   "test school",
		Image:  graphql.Upload{File: f, Filename: f.File.Name(), ContentType: f.ContentType, Size: f.Size()},
		Status: schema.StatusActive,
	})
	require.NoError(t, err)
	require.NotNil(t, sch)

	t.Run("name", func(t *testing.T) {
		newSch, err := s.UpdateSchool(ctx, sch.ID, model.UpdateSchoolInput{Name: ptr.Str("new name")})
		require.NoError(t, err)
		require.NotNil(t, newSch)
		require.Equal(t, newSch.Name, "new name")
		require.Equal(t, newSch.Status, sch.Status)
		require.Equal(t, newSch.Image, sch.Image)
		require.Equal(t, newSch.Directory, sch.Directory)
	})

	t.Run("image", func(t *testing.T) {
		f := testutil.OpenFile(t, "../testfiles/harvard.jpg")
		defer f.Close()

		newSch, err := s.UpdateSchool(ctx, sch.ID, model.UpdateSchoolInput{Image: &graphql.Upload{
			File:        f,
			Filename:    f.File.Name(),
			ContentType: f.ContentType,
			Size:        f.Size(),
		}})
		require.NoError(t, err)
		require.NotNil(t, newSch)
		require.Equal(t, sch.Image, newSch.Image)
		require.Equal(t, sch.Directory, newSch.Directory)
		stat, err := s.MC.StatObject(ctx, s.Config.RootBucket, newSch.Image, minio.StatObjectOptions{})
		require.Equal(t, "image/jpeg", stat.ContentType)
	})

	t.Run("invalid image file", func(t *testing.T) {
		f := testutil.OpenFile(t, "../testfiles/file.txt")
		defer f.Close()

		newSch, err := s.UpdateSchool(ctx, sch.ID, model.UpdateSchoolInput{Image: &graphql.Upload{
			File:        f,
			Filename:    f.File.Name(),
			ContentType: f.ContentType,
			Size:        f.Size(),
		}})
		require.Error(t, err)
		require.Nil(t, newSch)
	})

	t.Run("all", func(t *testing.T) {
		f := testutil.OpenFile(t, "../testfiles/harvard.jpg")
		defer f.Close()

		newSch, err := s.UpdateSchool(ctx, sch.ID, model.UpdateSchoolInput{
			Name:   ptr.Str("name 2"),
			Status: ptr.Status(schema.StatusDisabled),
			Image: &graphql.Upload{
				File:        f,
				Filename:    f.File.Name(),
				ContentType: f.ContentType,
				Size:        f.Size(),
			},
		})
		require.NoError(t, err)
		require.NotNil(t, newSch)
		require.Equal(t, "name 2", newSch.Name)
		require.Equal(t, schema.StatusDisabled, newSch.Status)
		require.Equal(t, sch.Image, newSch.Image)
		require.Equal(t, sch.Directory, newSch.Directory)
	})
}
