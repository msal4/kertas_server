package service_test

import (
	"bufio"
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/ent/stage"
	"github.com/msal4/hassah_school_server/server/model"
	"github.com/msal4/hassah_school_server/service"
	"github.com/msal4/hassah_school_server/util/ptr"
	"github.com/stretchr/testify/require"
)

func TestStages(t *testing.T) {
	s := newService(t)
	defer s.EC.Close()
	ctx := context.Background()

	t.Run("empty", func(t *testing.T) {
		defer s.EC.School.Delete().ExecX(ctx)
		require := require.New(t)

		conn, err := s.Stages(ctx, service.StagesOptions{})
		require.NoError(err)
		require.NotNil(t, conn)
		require.Zero(conn.TotalCount)
		require.Empty(conn.Edges)
	})

	t.Run("not empty", func(t *testing.T) {
		defer s.EC.School.Delete().ExecX(ctx)
		require := require.New(t)

		want := createStage(ctx, s, "test stage", 20334)

		conn, err := s.Stages(ctx, service.StagesOptions{})
		require.NoError(err)
		require.NotNil(conn)
		require.Equal(1, conn.TotalCount)
		require.Len(conn.Edges, 1)
		edge := conn.Edges[0]
		require.NotNil(edge)
		require.NotNil(edge.Cursor)
		require.NotNil(edge.Node)
		require.Equal(edge.Node.ID, want.ID)
		require.Equal(edge.Node.Name, want.Name)
		sch, err := edge.Node.School(ctx)
		require.NoError(err)
		require.NotNil(sch)
		wantSch, err := want.School(ctx)
		require.NoError(err)
		require.NotNil(sch)
		require.Equal(wantSch.ID, sch.ID)
		require.Equal(wantSch.Name, sch.Name)
	})

	t.Run("order & filter", func(t *testing.T) {
		defer s.EC.School.Delete().ExecX(ctx)
		require := require.New(t)

		createStage(ctx, s, "stage 3", 300)
		createStage(ctx, s, "stage 1", 100)
		createStage(ctx, s, "stage 4", 400)
		createStage(ctx, s, "stage 2", 300)

		b := s.EC.Stage.Query().Order(ent.Asc(stage.FieldCreatedAt))
		want := b.AllX(ctx)

		conn, err := s.Stages(ctx, service.StagesOptions{
			OrderBy: &ent.StageOrder{Field: ent.StageOrderFieldCreatedAt, Direction: ent.OrderDirectionAsc},
		})
		require.NoError(err)
		require.NotNil(conn)
		require.Equal(len(want), conn.TotalCount)
		require.Len(conn.Edges, len(want))
		for i, edge := range conn.Edges {
			require.NotNil(edge)
			require.NotNil(edge.Cursor)
			require.NotNil(edge.Node)
			require.Equal(edge.Node.ID, want[i].ID)
			require.Equal(edge.Node.Name, want[i].Name)
			require.Equal(edge.Node.Active, want[i].Active)
			require.Equal(edge.Node.TuitionAmount, want[i].TuitionAmount)
			require.Equal(edge.Node.Directory, want[i].Directory)
			require.Equal(edge.Node.CreatedAt, want[i].CreatedAt)
		}

		want = b.Where(stage.Active(false)).AllX(ctx)

		conn, err = s.Stages(ctx, service.StagesOptions{
			Where:   &ent.StageWhereInput{Active: ptr.Bool(false)},
			OrderBy: &ent.StageOrder{Field: ent.StageOrderFieldCreatedAt, Direction: ent.OrderDirectionAsc},
		})
		require.NoError(err)

		require.Equal(len(want), conn.TotalCount)
		require.Len(conn.Edges, len(want))

		for i, w := range want {
			e := conn.Edges[i]

			require.NotNil(e)
			require.NotNil(e.Cursor)
			require.NotNil(e.Node)
			require.Equal(e.Node.ID, w.ID)
			require.Equal(e.Node.Name, w.Name)
			require.Equal(e.Node.Active, w.Active)
			require.Equal(e.Node.TuitionAmount, w.TuitionAmount)
			require.Equal(e.Node.Directory, w.Directory)
			require.Equal(e.Node.CreatedAt, w.CreatedAt)
		}

	})
}

func TestAddStage(t *testing.T) {
	s := newService(t)
	defer s.EC.Close()
	ctx := context.Background()

	t.Run("without school", func(t *testing.T) {
		got, err := s.AddStage(ctx, model.AddStageInput{Name: "test school", TuitionAmount: 200, Active: true})
		require.Error(t, err)
		require.Nil(t, got)
	})

	t.Run("with school", func(t *testing.T) {
		require := require.New(t)

		sch := createSchool(ctx, s, "test schoo", "he/image.png")

		got, err := s.AddStage(ctx, model.AddStageInput{
			Name:          "test stage",
			SchoolID:      sch.ID,
			Active:        false,
			TuitionAmount: 22222,
		})
		require.NoError(err)
		require.NotNil(got)
		require.NotEmpty(got.Name)
		require.NotEmpty(got.Directory)

		gotSch, err := got.School(ctx)
		require.NoError(err)
		require.NotNil(gotSch)
		require.Equal(sch.ID, gotSch.ID)

		st := s.EC.Stage.GetX(ctx, got.ID)
		require.Equal("test stage", st.Name)
		require.Equal(got.Directory, st.Directory)
		require.Equal(22222, st.TuitionAmount)
	})

	t.Run("with invalid school", func(t *testing.T) {
		require := require.New(t)

		got, err := s.AddStage(ctx, model.AddStageInput{
			Name:          "test stage",
			SchoolID:      uuid.New(),
			Active:        false,
			TuitionAmount: 22222,
		})
		require.Error(err)
		require.Nil(got)
	})
}

func TestUpdateStage(t *testing.T) {
	s := newService(t)
	defer s.EC.Close()

	ctx := context.Background()

	st := createStage(ctx, s, "test stage name 1", 234)

	t.Run("name", func(t *testing.T) {
		require := require.New(t)

		got, err := s.UpdateStage(ctx, st.ID, model.UpdateStageInput{Name: ptr.Str("new name")})
		require.NoError(err)
		require.NotNil(got)
		require.Equal(got.Name, "new name")
		require.Equal(st.Active, got.Active)
		require.Equal(st.Directory, got.Directory)
		require.Equal(st.TuitionAmount, got.TuitionAmount)
	})

	t.Run("active", func(t *testing.T) {
		got, err := s.UpdateStage(ctx, st.ID, model.UpdateStageInput{Active: ptr.Bool(false)})
		require.NoError(t, err)
		require.NotNil(t, got)
		require.Equal(t, st.Directory, got.Directory)
		require.Equal(t, false, got.Active)
	})

	t.Run("negative tuition", func(t *testing.T) {
		got, err := s.UpdateStage(ctx, st.ID, model.UpdateStageInput{TuitionAmount: ptr.Int(-100)})
		require.Error(t, err)
		require.Nil(t, got)
	})

	t.Run("all", func(t *testing.T) {
		got, err := s.UpdateStage(ctx, st.ID, model.UpdateStageInput{
			Name:          ptr.Str("name 2"),
			Active:        ptr.Bool(true),
			TuitionAmount: ptr.Int(300),
		})
		require.NoError(t, err)
		require.NotNil(t, got)
		require.Equal(t, "name 2", got.Name)
		require.Equal(t, true, got.Active)
		require.Equal(t, st.Directory, got.Directory)
		require.Equal(t, 300, got.TuitionAmount)
	})
}

func TestDeleteStage(t *testing.T) {
	s := newService(t)
	defer s.EC.Close()

	ctx := context.Background()

	st := createStage(ctx, s, "test stage", 32423)

	t.Run("non-existing stage", func(t *testing.T) {
		err := s.DeleteStage(ctx, uuid.New())
		require.Error(t, err)
		got, err := s.EC.Stage.Get(ctx, st.ID)
		require.NoError(t, err)
		require.Nil(t, got.DeletedAt)
	})

	t.Run("existing stage", func(t *testing.T) {
		err := s.DeleteStage(ctx, st.ID)
		require.NoError(t, err)
		got, err := s.EC.Stage.Get(ctx, st.ID)
		require.NoError(t, err)
		require.NotNil(t, got.DeletedAt)
	})
}

func TestDeleteStagePermanently(t *testing.T) {
	s := newService(t)
	defer s.EC.Close()

	ctx := context.Background()

	st := createStage(ctx, s, "test stage name 1", 234)

	t.Run("non-existing stage", func(t *testing.T) {
		err := s.DeleteStagePermanently(ctx, uuid.New())
		require.Error(t, err)
		_, err = s.EC.Stage.Get(ctx, st.ID)
		require.NoError(t, err)
	})

	t.Run("stage with empty directory", func(t *testing.T) {
		err := s.DeleteStagePermanently(ctx, st.ID)
		require.NoError(t, err)
		_, err = s.EC.Stage.Get(ctx, st.ID)
		require.Error(t, err)
		require.True(t, ent.IsNotFound(err))
	})

	t.Run("stage with directory", func(t *testing.T) {
		st := createStage(ctx, s, "testfs", 12)
		st2 := createStage(ctx, s, "testfs23", 22312)

		filename1 := st.Directory + "/testfile1.txt"
		_, err := s.MC.PutObject(ctx, s.Config.RootBucket, filename1, bufio.NewReader(nil), 0, minio.PutObjectOptions{})
		require.NoError(t, err)

		filename2 := st.Directory + "/testfile2.txt"
		_, err = s.MC.PutObject(ctx, s.Config.RootBucket, filename2, bufio.NewReader(nil), 0, minio.PutObjectOptions{})
		require.NoError(t, err)

		filename3 := st.Directory + "/hello/testfile3.txt"
		_, err = s.MC.PutObject(ctx, s.Config.RootBucket, filename2, bufio.NewReader(nil), 0, minio.PutObjectOptions{})
		require.NoError(t, err)

		err = s.DeleteStagePermanently(ctx, st.ID)
		require.NoError(t, err)

		_, err = s.EC.Stage.Get(ctx, st.ID)
		require.Error(t, err)
		require.True(t, ent.IsNotFound(err))

		_, err = s.EC.Stage.Get(ctx, st2.ID)
		require.NoError(t, err)

		_, err = s.MC.StatObject(ctx, s.Config.RootBucket, filename1, minio.StatObjectOptions{})
		require.Error(t, err)
		_, err = s.MC.StatObject(ctx, s.Config.RootBucket, filename2, minio.StatObjectOptions{})
		require.Error(t, err)
		_, err = s.MC.StatObject(ctx, s.Config.RootBucket, filename3, minio.StatObjectOptions{})
		require.Error(t, err)
	})
}

func createSchool(ctx context.Context, s *service.Service, name, image string) *ent.School {
	return s.EC.School.Create().SetName(name).SetImage(image).SetDirectory("test_dir").SaveX(ctx)
}

func createStage(ctx context.Context, s *service.Service, name string, tuition int) *ent.Stage {
	sch := createSchool(ctx, s, "school for"+name, "image/"+name)
	return s.EC.Stage.Create().SetName(name).SetDirectory("testdir" + name).SetTuitionAmount(tuition).SetSchool(sch).SaveX(ctx)
}
