package service_test

import (
	"strings"
	"testing"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/ent/group"
	"github.com/msal4/hassah_school_server/server/model"
	"github.com/msal4/hassah_school_server/service"
	"github.com/msal4/hassah_school_server/testutil"
	"github.com/msal4/hassah_school_server/util/ptr"
	"github.com/stretchr/testify/require"
)

func TestAssignments(t *testing.T) {
	s := newService(t)
	defer s.Close()

	sch := createSchool(ctx, s, "kd", "ks")
	schAdmin := createSchoolAdmin(ctx, s, "skdjfsd", sch)
	stdt := createStudent(ctx, s, "sjfkdj", sch)
	stg, err := stdt.Stage(ctx)
	require.NoError(t, err)
	stg2, err := createStudent(ctx, s, "sjfkdj2", sch).Stage(ctx)
	require.NoError(t, err)
	tchr := createTeacher(ctx, s, "sd", sch)
	tchr2 := createTeacher(ctx, s, "sddd", sch)
	grp1 := s.EC.Group.Create().SetName("mathgroup").SetGroupType(group.GroupTypeShared).SaveX(ctx)
	cls1 := s.EC.Class.Create().SetName("math").SetGroup(grp1).SetTeacher(tchr).SetStage(stg).
		SetCreatedAt(now(-time.Hour)).SaveX(ctx)
	grp2 := s.EC.Group.Create().SetName("geogroup").SetGroupType(group.GroupTypeShared).SaveX(ctx)
	cls2 := s.EC.Class.Create().SetName("geo").SetGroup(grp2).SetTeacher(tchr).
		SetCreatedAt(now(-time.Hour / 2)).SetStage(stg).SaveX(ctx)
	grp3 := s.EC.Group.Create().SetName("geogroup").SetGroupType(group.GroupTypeShared).SaveX(ctx)
	cls3 := s.EC.Class.Create().SetName("geo").SetGroup(grp3).SetTeacher(tchr2).
		SetCreatedAt(now(-time.Hour / 2)).SetStage(stg2).SaveX(ctx)

	ass1 := s.EC.Assignment.Create().SetCreatedAt(now(-time.Hour)).
		SetName("ass 1").SetClass(cls1).SetDueDate(now(time.Hour * 24)).SaveX(ctx)

	ass2 := s.EC.Assignment.Create().SetCreatedAt(now(-time.Hour / 2)).
		SetName("ass 2").SetClass(cls2).SetDueDate(now(time.Hour * 12)).SaveX(ctx)

	ass3 := s.EC.Assignment.Create().SetName("ass 3").SetCreatedAt(now(-time.Minute)).
		SetClass(cls2).SetDueDate(now(time.Hour * 19)).SaveX(ctx)

	ass4 := s.EC.Assignment.Create().SetName("ass 4").SetCreatedAt(now()).
		SetClass(cls3).SetDueDate(now(time.Hour * 19)).SaveX(ctx)
	_ = ass4

	t.Run("student", func(t *testing.T) {
		require := require.New(t)

		got, err := s.Assignments(ctx, service.AssignmentsOptions{
			UserID:  stdt.ID,
			OrderBy: &ent.AssignmentOrder{Field: ent.AssignmentOrderFieldCreatedAt, Direction: ent.OrderDirectionAsc},
		})
		require.NoError(err)
		require.NotNil(got)
		require.Len(got.Edges, 3)
		require.Equal(ass1.ID, got.Edges[0].Node.ID)
		require.Equal(ass2.ID, got.Edges[1].Node.ID)
		require.Equal(ass3.ID, got.Edges[2].Node.ID)
	})

	t.Run("student & class", func(t *testing.T) {
		require := require.New(t)

		got, err := s.Assignments(ctx, service.AssignmentsOptions{
			UserID:  stdt.ID,
			ClassID: &cls2.ID,
			OrderBy: &ent.AssignmentOrder{Field: ent.AssignmentOrderFieldCreatedAt, Direction: ent.OrderDirectionAsc},
		})
		require.NoError(err)
		require.NotNil(got)
		require.Len(got.Edges, 2)
		require.Equal(ass2.ID, got.Edges[0].Node.ID)
		require.Equal(ass3.ID, got.Edges[1].Node.ID)
	})

	t.Run("teacher", func(t *testing.T) {
		require := require.New(t)

		got, err := s.Assignments(ctx, service.AssignmentsOptions{
			UserID:  tchr.ID,
			OrderBy: &ent.AssignmentOrder{Field: ent.AssignmentOrderFieldCreatedAt, Direction: ent.OrderDirectionAsc},
		})
		require.NoError(err)
		require.NotNil(got)
		require.Len(got.Edges, 3)
		require.Equal(ass1.ID, got.Edges[0].Node.ID)
		require.Equal(ass2.ID, got.Edges[1].Node.ID)
		require.Equal(ass3.ID, got.Edges[2].Node.ID)
	})

	t.Run("teacher & class", func(t *testing.T) {
		require := require.New(t)

		got, err := s.Assignments(ctx, service.AssignmentsOptions{
			UserID:  tchr.ID,
			ClassID: &cls2.ID,
			OrderBy: &ent.AssignmentOrder{Field: ent.AssignmentOrderFieldCreatedAt, Direction: ent.OrderDirectionAsc},
		})
		require.NoError(err)
		require.NotNil(got)
		require.Len(got.Edges, 2)
		require.Equal(ass2.ID, got.Edges[0].Node.ID)
		require.Equal(ass3.ID, got.Edges[1].Node.ID)
	})

	t.Run("school admin", func(t *testing.T) {
		require := require.New(t)

		got, err := s.Assignments(ctx, service.AssignmentsOptions{
			UserID:  schAdmin.ID,
			OrderBy: &ent.AssignmentOrder{Field: ent.AssignmentOrderFieldCreatedAt, Direction: ent.OrderDirectionAsc},
		})
		require.NoError(err)
		require.NotNil(got)
		require.Len(got.Edges, 4)
		require.Equal(ass1.ID, got.Edges[0].Node.ID)
		require.Equal(ass2.ID, got.Edges[1].Node.ID)
		require.Equal(ass3.ID, got.Edges[2].Node.ID)
		require.Equal(ass4.ID, got.Edges[3].Node.ID)
	})

	t.Run("school admin & class", func(t *testing.T) {
		require := require.New(t)

		got, err := s.Assignments(ctx, service.AssignmentsOptions{
			UserID:  schAdmin.ID,
			ClassID: &cls1.ID,
			OrderBy: &ent.AssignmentOrder{Field: ent.AssignmentOrderFieldCreatedAt, Direction: ent.OrderDirectionAsc},
		})
		require.NoError(err)
		require.NotNil(got)
		require.Len(got.Edges, 1)
		require.Equal(ass1.ID, got.Edges[0].Node.ID)
	})
}

func TestAddAssignment(t *testing.T) {
	s := newService(t)
	defer s.Close()

	f := testutil.OpenFile(t, "../testfiles/file.txt")

	sch := createSchool(ctx, s, "kd", "ks")
	stdt := createStudent(ctx, s, "sjfkdj", sch)
	stg, err := stdt.Stage(ctx)
	require.NoError(t, err)
	tchr := createTeacher(ctx, s, "sd", sch)
	grp := s.EC.Group.Create().SetName("mathgroup").SetGroupType(group.GroupTypeShared).SaveX(ctx)
	cls := s.EC.Class.Create().SetName("math").SetGroup(grp).SetTeacher(tchr).SetStage(stg).
		SetCreatedAt(now(-time.Hour)).SaveX(ctx)

	t.Run("assignment", func(t *testing.T) {
		require := require.New(t)
		ass, err := s.AddAssignment(ctx, model.AddAssignmentInput{
			Name:        "first sass",
			Description: ptr.Str("desc"),
			File:        uploadFromFile(f),
			ClassID:     cls.ID,
			DueDate:     now(time.Hour * 48),
		})
		require.NoError(err)
		require.NotNil(ass)
		require.Equal("first sass", ass.Name)
		ass = s.EC.Assignment.GetX(ctx, ass.ID)
		require.Equal("first sass", ass.Name)
		require.Equal("desc", ass.Description)
		require.NotEmpty(ass.File)
		require.False(ass.IsExam)
		cls, err := ass.Class(ctx)
		require.NoError(err)
		require.NotNil(cls)
		require.True(strings.HasPrefix(ass.File, stg.Directory))
		stat, err := s.MC.StatObject(ctx, s.Config.RootBucket, ass.File, minio.StatObjectOptions{})
		require.NoError(err)
		require.Equal(f.Size(), stat.Size)
		require.True(strings.HasSuffix(stat.Key, ".txt"))
		require.Equal(now(time.Hour*48).Format(time.ANSIC), ass.DueDate.Format(time.ANSIC))
	})

	t.Run("exam", func(t *testing.T) {
		require := require.New(t)
		ass, err := s.AddAssignment(ctx, model.AddAssignmentInput{
			Name:        "first exam",
			Description: ptr.Str("exam desc"),
			File:        uploadFromFile(f),
			ClassID:     cls.ID,
			DueDate:     now(time.Hour * 48),
			Duration:    ptr.Duration(2 * time.Hour),
			IsExam:      true,
		})
		require.NoError(err)
		require.NotNil(ass)
		require.Equal("first exam", ass.Name)
		ass = s.EC.Assignment.GetX(ctx, ass.ID)
		require.Equal("first exam", ass.Name)
		require.Equal("exam desc", ass.Description)
		require.NotEmpty(ass.File)
		require.True(ass.IsExam)
		require.Equal(2*time.Hour, ass.Duration)
		cls, err := ass.Class(ctx)
		require.NoError(err)
		require.NotNil(cls)
		require.True(strings.HasPrefix(ass.File, stg.Directory))
		stat, err := s.MC.StatObject(ctx, s.Config.RootBucket, ass.File, minio.StatObjectOptions{})
		require.NoError(err)
		require.Equal(f.Size(), stat.Size)
		require.True(strings.HasSuffix(stat.Key, ".txt"))
	})
}

func TestUpdateAssignment(t *testing.T) {
	s := newService(t)
	defer s.Close()

	f := testutil.OpenFile(t, "../testfiles/file.txt")

	sch := createSchool(ctx, s, "kd", "ks")
	stdt := createStudent(ctx, s, "sjfkdj", sch)
	stg, err := stdt.Stage(ctx)
	require.NoError(t, err)
	tchr := createTeacher(ctx, s, "sd", sch)
	grp := s.EC.Group.Create().SetName("mathgroup").SetGroupType(group.GroupTypeShared).SaveX(ctx)
	cls := s.EC.Class.Create().SetName("math").SetGroup(grp).SetTeacher(tchr).SetStage(stg).
		SetCreatedAt(now(-time.Hour)).SaveX(ctx)
	ass := s.EC.Assignment.Create().SetName("first ass").SetDescription("desc").
		SetClass(cls).SetDueDate(now(time.Hour * 48)).SaveX(ctx)
	exam := s.EC.Assignment.Create().SetName("first exam").SetDescription("desc").
		SetClass(cls).SetDueDate(now(time.Hour * 48)).SetDuration(time.Hour).SetIsExam(true).SaveX(ctx)

	t.Run("assignment", func(t *testing.T) {
		require := require.New(t)
		ass, err := s.UpdateAssignment(ctx, ass.ID, model.UpdateAssignmentInput{
			Name:        ptr.Str("first sass updated"),
			Description: ptr.Str("new desc"),
			File:        uploadFromFile(f),
			DueDate:     ptr.Time(now(time.Hour * 28)),
		})
		require.NoError(err)
		require.NotNil(ass)
		require.Equal("first sass updated", ass.Name)

		ass = s.EC.Assignment.GetX(ctx, ass.ID)
		require.Equal("first sass updated", ass.Name)
		require.Equal("new desc", ass.Description)
		require.NotEmpty(ass.File)
		require.False(ass.IsExam)

		cls, err := ass.Class(ctx)
		require.NoError(err)
		require.NotNil(cls)
		require.True(strings.HasPrefix(ass.File, stg.Directory))

		stat, err := s.MC.StatObject(ctx, s.Config.RootBucket, ass.File, minio.StatObjectOptions{})
		require.NoError(err)
		require.Equal(f.Size(), stat.Size)
		require.True(strings.HasSuffix(stat.Key, ".txt"))
		require.Equal(now(time.Hour*28).Format(time.ANSIC), ass.DueDate.Format(time.ANSIC))
	})

	t.Run("exam", func(t *testing.T) {
		require := require.New(t)
		exam, err := s.UpdateAssignment(ctx, exam.ID, model.UpdateAssignmentInput{
			Name:        ptr.Str("first exam updated"),
			Description: ptr.Str("new exam desc"),
			File:        uploadFromFile(f),
			DueDate:     ptr.Time(now(time.Hour * 28)),
			Duration:    ptr.Duration(time.Hour * 3),
		})
		require.NoError(err)
		require.NotNil(exam)
		require.Equal("first exam updated", exam.Name)

		exam = s.EC.Assignment.GetX(ctx, exam.ID)
		require.Equal("first exam updated", exam.Name)
		require.Equal("new exam desc", exam.Description)
		require.NotEmpty(exam.File)
		require.True(exam.IsExam)
		require.Equal(3*time.Hour, exam.Duration)

		cls, err := exam.Class(ctx)
		require.NoError(err)
		require.NotNil(cls)
		require.True(strings.HasPrefix(exam.File, stg.Directory))

		stat, err := s.MC.StatObject(ctx, s.Config.RootBucket, exam.File, minio.StatObjectOptions{})
		require.NoError(err)
		require.Equal(f.Size(), stat.Size)
		require.True(strings.HasSuffix(stat.Key, ".txt"))
	})
}

func TestDeleteAssignment(t *testing.T) {
	s := newService(t)
	defer s.Close()

	require := require.New(t)

	sch := createSchool(ctx, s, "kd", "ks")
	stdt := createStudent(ctx, s, "sjfkdj", sch)
	stg, err := stdt.Stage(ctx)
	require.NoError(err)
	tchr := createTeacher(ctx, s, "sd", sch)
	grp := s.EC.Group.Create().SetName("mathgroup").SetGroupType(group.GroupTypeShared).SaveX(ctx)
	cls := s.EC.Class.Create().SetName("math").SetGroup(grp).SetTeacher(tchr).SetStage(stg).
		SetCreatedAt(now(-time.Hour)).SaveX(ctx)
	ass := s.EC.Assignment.Create().SetName("first ass").SetDescription("desc").
		SetClass(cls).SetDueDate(now(time.Hour * 48)).SaveX(ctx)
	anotherAss := s.EC.Assignment.Create().SetName("first exam").SetDescription("desc").
		SetClass(cls).SetDueDate(now(time.Hour * 48)).SetDuration(time.Hour).SetIsExam(true).SaveX(ctx)

	err = s.DeleteAssignment(ctx, ass.ID)
	require.NoError(err)

	ass = s.EC.Assignment.GetX(ctx, ass.ID)
	require.NotNil(ass.DeletedAt)

	anotherAss = s.EC.Assignment.GetX(ctx, anotherAss.ID)
	require.Nil(anotherAss.DeletedAt)
}
