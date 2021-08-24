package service_test

import (
	"testing"
	"time"

	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/ent/group"
	"github.com/msal4/hassah_school_server/server/model"
	"github.com/msal4/hassah_school_server/service"
	"github.com/msal4/hassah_school_server/util/ptr"
	"github.com/stretchr/testify/require"
)

func TestClasses(t *testing.T) {
	s := newService(t)
	defer s.EC.Close()

	sch := createSchool(ctx, s, "jlfs", "k'sdfs")
	stdt := createStudent(ctx, s, "fkdsj2222", sch)
	tchr := createTeacher(ctx, s, "jlfsdjfkd", sch)
	stg, err := stdt.Stage(ctx)
	require.NoError(t, err)
	anotherStg, err := createStudent(ctx, s, "idkjfs3", sch).Stage(ctx)
	require.NoError(t, err)
	anotherSchStg := createStage(ctx, s, "jl", 12)
	cls1 := s.EC.Class.Create().SetStage(stg).SetGroup(createGroup(ctx, s)).
		SetTeacher(tchr).SetName("geo").SetCreatedAt(time.Now().Add(2 * -time.Hour)).SaveX(ctx)
	cls2 := s.EC.Class.Create().SetStage(stg).SetGroup(createGroup(ctx, s)).
		SetTeacher(tchr).SetName("geo").SetCreatedAt(time.Now().Add(-time.Hour)).SaveX(ctx)
	cls3 := s.EC.Class.Create().SetStage(anotherStg).SetGroup(createGroup(ctx, s)).
		SetTeacher(tchr).SetName("geo").SaveX(ctx)
	cls4 := s.EC.Class.Create().SetStage(anotherSchStg).SetGroup(createGroup(ctx, s)).
		SetTeacher(tchr).SetName("geo").SetCreatedAt(time.Now().Add(time.Minute)).SaveX(ctx)

	t.Run("student classes", func(t *testing.T) {
		require := require.New(t)

		classes, err := s.Classes(ctx, service.ClassesOptions{
			UserID:  &stdt.ID,
			OrderBy: &ent.ClassOrder{Field: ent.ClassOrderFieldCreatedAt, Direction: ent.OrderDirectionAsc},
		})
		require.NoError(err)
		require.NotNil(classes)
		require.Len(classes.Edges, 2)
		require.Equal(cls1.ID, classes.Edges[0].Node.ID)
		require.Equal(cls2.ID, classes.Edges[1].Node.ID)
	})

	t.Run("stage classes", func(t *testing.T) {
		require := require.New(t)

		classes, err := s.Classes(ctx, service.ClassesOptions{
			StageID: &stg.ID,
			OrderBy: &ent.ClassOrder{Field: ent.ClassOrderFieldCreatedAt, Direction: ent.OrderDirectionAsc},
		})
		require.NoError(err)
		require.NotNil(classes)
		require.Len(classes.Edges, 2)
		require.Equal(cls1.ID, classes.Edges[0].Node.ID)
		require.Equal(cls2.ID, classes.Edges[1].Node.ID)
	})

	t.Run("school classes", func(t *testing.T) {
		require := require.New(t)

		classes, err := s.Classes(ctx, service.ClassesOptions{
			SchoolID: &sch.ID,
			OrderBy:  &ent.ClassOrder{Field: ent.ClassOrderFieldCreatedAt, Direction: ent.OrderDirectionAsc},
		})
		require.NoError(err)
		require.NotNil(classes)
		require.Len(classes.Edges, 3)
		require.Equal(cls1.ID, classes.Edges[0].Node.ID)
		require.Equal(cls2.ID, classes.Edges[1].Node.ID)
		require.Equal(cls3.ID, classes.Edges[2].Node.ID)
	})

	t.Run("all", func(t *testing.T) {
		require := require.New(t)

		classes, err := s.Classes(ctx, service.ClassesOptions{
			OrderBy: &ent.ClassOrder{Field: ent.ClassOrderFieldCreatedAt, Direction: ent.OrderDirectionAsc},
		})
		require.NoError(err)
		require.NotNil(classes)
		require.Len(classes.Edges, 4)
		require.Equal(cls1.ID, classes.Edges[0].Node.ID)
		require.Equal(cls2.ID, classes.Edges[1].Node.ID)
		require.Equal(cls3.ID, classes.Edges[2].Node.ID)
		require.Equal(cls4.ID, classes.Edges[3].Node.ID)
	})
}

func TestAddClass(t *testing.T) {
	s := newService(t)
	defer s.Close()

	sch := createSchool(ctx, s, "fkds", "fjdks")
	tchr := createTeacher(ctx, s, "ikdfs2", sch)
	stg, err := createStudent(ctx, s, "kdsjfkd", sch).Stage(ctx)
	require.NoError(t, err)

	t.Run("mssing teacher id/missing stage id", func(t *testing.T) {
		// missing teacher id & stage id
		cls, err := s.AddClass(ctx, model.AddClassInput{Name: "new group", Active: true})
		require.Error(t, err)
		require.Nil(t, cls)

		// missing stage
		cls, err = s.AddClass(ctx, model.AddClassInput{Name: "new group", Active: true, TeacherID: tchr.ID})
		require.Error(t, err)
		require.Nil(t, cls)

		// missing teacher id
		cls, err = s.AddClass(ctx, model.AddClassInput{Name: "new group", Active: true, StageID: stg.ID})
		require.Error(t, err)
		require.Nil(t, cls)
	})

	t.Run("with teacher & stage", func(t *testing.T) {
		require := require.New(t)
		cls, err := s.AddClass(ctx, model.AddClassInput{Name: "new group", Active: true, TeacherID: tchr.ID, StageID: stg.ID})
		require.NoError(err)
		require.NotNil(cls)

		cls = s.EC.Class.GetX(ctx, cls.ID)
		gotTchr, err := cls.Teacher(ctx)
		require.NoError(err)
		require.NotNil(gotTchr)
		require.Equal(tchr.ID, gotTchr.ID)
		require.Equal("new group", cls.Name)
		grp, err := cls.Group(ctx)
		require.NoError(err)
		require.NotNil(grp)
		require.Equal(group.GroupTypeShared, grp.GroupType)
		require.Equal(cls.Name, grp.Name)
	})
}

func TestUpdateClass(t *testing.T) {
	s := newService(t)
	defer s.Close()

	require := require.New(t)

	sch := createSchool(ctx, s, "ks", "kd")
	grp := createGroup(ctx, s)
	grp2 := createGroup(ctx, s)
	stg, err := createStudent(ctx, s, "ksd", sch).Stage(ctx)
	require.NoError(err)
	tchr := createTeacher(ctx, s, "skdfjsdjfkds", sch)

	cls := s.EC.Class.Create().SetName("hello").SetGroup(grp).SetStage(stg).SetTeacher(tchr).SaveX(ctx)
	anotherCls := s.EC.Class.Create().SetName("hello 2").SetGroup(grp2).SetStage(stg).SetTeacher(tchr).SaveX(ctx)

	got, err := s.UpdateClass(ctx, cls.ID, model.UpdateClassInput{Name: ptr.Str("new name")})
	require.NoError(err)
	require.NotNil(got)
	require.Equal("new name", got.Name)

	got = s.EC.Class.GetX(ctx, cls.ID)
	require.Equal("new name", got.Name)
	require.Equal(cls.Active, got.Active)

	got, err = s.UpdateClass(ctx, cls.ID, model.UpdateClassInput{Active: ptr.Bool(false)})
	require.NoError(err)
	require.NotNil(got)
	require.Equal("new name", got.Name)
	require.Equal(false, got.Active)
	got = s.EC.Class.GetX(ctx, cls.ID)
	require.Equal("new name", got.Name)
	require.Equal(false, got.Active)

	// all
	anotherTchr := createTeacher(ctx, s, "kdsjd", sch)
	got, err = s.UpdateClass(ctx, cls.ID, model.UpdateClassInput{
		Active:    ptr.Bool(true),
		Name:      ptr.Str("new name 2"),
		TeacherID: &anotherTchr.ID,
	})
	require.NoError(err)
	require.NotNil(got)
	require.Equal("new name 2", got.Name)
	require.Equal(true, got.Active)
	got = s.EC.Class.GetX(ctx, cls.ID)
	gotTchr, err := got.Teacher(ctx)
	require.NoError(err)
	require.NotNil(gotTchr)
	require.Equal("new name 2", got.Name)
	require.Equal(true, got.Active)
	require.Equal(anotherTchr.ID, gotTchr.ID)

	gotAnotherCls := s.EC.Class.GetX(ctx, anotherCls.ID)
	require.Equal(anotherCls.Name, gotAnotherCls.Name)
	require.Equal(anotherCls.Active, gotAnotherCls.Active)
}

func TestDeleteClass(t *testing.T) {
	s := newService(t)
	defer s.Close()

	require := require.New(t)

	sch := createSchool(ctx, s, "ks", "kd")
	grp := createGroup(ctx, s)
	grp2 := createGroup(ctx, s)
	stg, err := createStudent(ctx, s, "ksd", sch).Stage(ctx)
	require.NoError(err)
	tchr := createTeacher(ctx, s, "skdfjsdjfkds", sch)

	cls := s.EC.Class.Create().SetName("hello").SetGroup(grp).SetStage(stg).SetTeacher(tchr).SaveX(ctx)
	cls2 := s.EC.Class.Create().SetName("2hello").SetGroup(grp2).SetStage(stg).SetTeacher(tchr).SaveX(ctx)

	err = s.DeleteClass(ctx, cls.ID)
	require.NoError(err)

	cls = s.EC.Class.GetX(ctx, cls.ID)
	require.NotNil(cls.DeletedAt)

	cls2 = s.EC.Class.GetX(ctx, cls2.ID)
	require.Nil(cls2.DeletedAt)
}
