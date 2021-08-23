package service_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/server/model"
	"github.com/msal4/hassah_school_server/service"
	"github.com/msal4/hassah_school_server/util/ptr"
	"github.com/stretchr/testify/require"
)

func TestGroups(t *testing.T) {
	s := newService(t)
	defer s.EC.Close()
	ctx := context.Background()

	t.Run("list empty groups list", func(t *testing.T) {
		require := require.New(t)
		got, err := s.Groups(ctx, service.GroupsOptions{})
		require.NoError(err)
		require.NotNil(got)
		require.Zero(got.TotalCount)
	})

	t.Run("list user's groups", func(t *testing.T) {
		require := require.New(t)
		sch := createSchool(ctx, s, "jskdfj", "fkdsjk")
		stdt := createStudent(ctx, s, "unigjsd", sch)
		stg, err := stdt.Stage(ctx)
		require.NoError(err)
		otherStdt := createStudent(ctx, s, "fksd342", sch)
		tchr := createTeacher(ctx, s, "idsjfksde", sch)
		anotherTchr := createTeacher(ctx, s, "ksfkdsjfksd22", sch)
		grp1 := s.EC.Group.Create().SetName("1").AddUsers(stdt, tchr).SetCreatedAt(time.Now().Add(-time.Hour)).SaveX(ctx)
		grp2 := s.EC.Group.Create().SetName("2").AddUsers(stdt, anotherTchr).SaveX(ctx)
		grp3 := s.EC.Group.Create().SetName("3").AddUsers(otherStdt, tchr).SetCreatedAt(time.Now().Add(time.Hour)).SaveX(ctx)

		// student
		got, err := s.Groups(ctx, service.GroupsOptions{UserID: &stdt.ID, OrderBy: &ent.GroupOrder{Field: ent.GroupOrderFieldCreatedAt, Direction: ent.OrderDirectionAsc}})
		require.NoError(err)
		require.NotNil(got)
		require.Len(got.Edges, 2)
		require.Equal(grp1.ID, got.Edges[0].Node.ID)
		require.Equal(grp2.ID, got.Edges[1].Node.ID)

		// teacher
		got, err = s.Groups(ctx, service.GroupsOptions{UserID: &tchr.ID, OrderBy: &ent.GroupOrder{Field: ent.GroupOrderFieldCreatedAt, Direction: ent.OrderDirectionAsc}})
		require.NoError(err)
		require.NotNil(got)
		require.Len(got.Edges, 2)
		require.Equal(grp1.ID, got.Edges[0].Node.ID)
		require.Equal(grp3.ID, got.Edges[1].Node.ID)

		grp4 := s.EC.Group.Create().SetName("math group").SetCreatedAt(time.Now().Add(2 * time.Hour)).SaveX(ctx)
		s.EC.Class.Create().SetName("math").SetGroup(grp4).SetStage(stg).SetTeacher(tchr).SaveX(ctx)

		// student with shared groups
		got, err = s.Groups(ctx, service.GroupsOptions{UserID: &stdt.ID, OrderBy: &ent.GroupOrder{Field: ent.GroupOrderFieldCreatedAt, Direction: ent.OrderDirectionAsc}})
		require.NoError(err)
		require.NotNil(got)
		require.Len(got.Edges, 3)
		require.Equal(grp1.ID, got.Edges[0].Node.ID)
		require.Equal(grp2.ID, got.Edges[1].Node.ID)
		require.Equal(grp4.ID, got.Edges[2].Node.ID)

		// teacher with shared groups
		got, err = s.Groups(ctx, service.GroupsOptions{UserID: &tchr.ID, OrderBy: &ent.GroupOrder{Field: ent.GroupOrderFieldCreatedAt, Direction: ent.OrderDirectionAsc}})
		require.NoError(err)
		require.NotNil(got)
		require.Len(got.Edges, 3)
		require.Equal(grp1.ID, got.Edges[0].Node.ID)
		require.Equal(grp3.ID, got.Edges[1].Node.ID)
		require.Equal(grp4.ID, got.Edges[2].Node.ID)

		schAdmin := createSchoolAdmin(ctx, s, "ksdjfksd3222455", sch)
		anotherSch := createSchool(ctx, s, "kjfs", "jfa")
		anotherStdt := createStudent(ctx, s, "jlfsss44", anotherSch)
		anotherStg, err := anotherStdt.Stage(ctx)
		grp5 := s.EC.Group.Create().SetName("math idfsroup").SetCreatedAt(time.Now().Add(3 * time.Hour)).SaveX(ctx)
		s.EC.Class.Create().SetName("skdfjslmath").SetGroup(grp5).SetStage(anotherStg).SetTeacher(tchr).SaveX(ctx)

		// school admin with shared groups
		got, err = s.Groups(ctx, service.GroupsOptions{UserID: &schAdmin.ID, OrderBy: &ent.GroupOrder{Field: ent.GroupOrderFieldCreatedAt, Direction: ent.OrderDirectionAsc}})
		require.NoError(err)
		require.NotNil(got)
		require.Len(got.Edges, 4)
		require.Equal(grp1.ID, got.Edges[0].Node.ID)
		require.Equal(grp2.ID, got.Edges[1].Node.ID)
		require.Equal(grp3.ID, got.Edges[2].Node.ID)
		require.Equal(grp4.ID, got.Edges[3].Node.ID)

		suAdmin := createSuperAdmin(ctx, s, "ksdkdsfksdk4444")
		// super admin
		got, err = s.Groups(ctx, service.GroupsOptions{UserID: &suAdmin.ID, OrderBy: &ent.GroupOrder{Field: ent.GroupOrderFieldCreatedAt, Direction: ent.OrderDirectionAsc}})
		require.NoError(err)
		require.NotNil(got)
		require.Len(got.Edges, 5)
		require.Equal(grp1.ID, got.Edges[0].Node.ID)
		require.Equal(grp2.ID, got.Edges[1].Node.ID)
		require.Equal(grp3.ID, got.Edges[2].Node.ID)
		require.Equal(grp4.ID, got.Edges[3].Node.ID)
		require.Equal(grp5.ID, got.Edges[4].Node.ID)
	})
}

func TestAddGroup(t *testing.T) {
	s := newService(t)
	defer s.EC.Close()
	ctx := context.Background()

	t.Run("empty users", func(t *testing.T) {
		require := require.New(t)
		got, err := s.AddGroup(ctx, service.AddGroupInput{
			Name:    "Our Group",
			Active:  true,
			UserIDs: []uuid.UUID{},
		})
		require.Error(err)
		require.Nil(got)

		got, err = s.AddGroup(ctx, service.AddGroupInput{
			Name:   "Our Group",
			Active: true,
		})
		require.Error(err)
		require.Nil(got)
	})

	t.Run("a group with one user", func(t *testing.T) {
		require := require.New(t)
		sch := createSchool(ctx, s, "fkdjs", "fjdks")
		stdt1 := createStudent(ctx, s, "skfjsldkfkjsdfld", sch)
		got, err := s.AddGroup(ctx, service.AddGroupInput{
			Name:    "Our Group",
			Active:  true,
			UserIDs: []uuid.UUID{stdt1.ID},
		})
		require.Error(err)
		require.Nil(got)

		// duplicate user
		got, err = s.AddGroup(ctx, service.AddGroupInput{
			Name:    "Our Group",
			Active:  true,
			UserIDs: []uuid.UUID{stdt1.ID, stdt1.ID},
		})
		require.Error(err)
		require.Nil(got)
	})

	t.Run("add users", func(t *testing.T) {
		require := require.New(t)
		sch := createSchool(ctx, s, "fkdjs", "fjdks")
		stdt1 := createStudent(ctx, s, "fkjsdfld", sch)
		stdt2 := createStudent(ctx, s, "djfkjsdfld", sch)
		createStudent(ctx, s, "sdjfkjsdfld", sch)

		got, err := s.AddGroup(ctx, service.AddGroupInput{
			Name:    "Our Group",
			Active:  true,
			UserIDs: []uuid.UUID{stdt1.ID, stdt2.ID},
		})

		require.NoError(err)
		require.NotNil(got)
		grp := s.EC.Group.GetX(ctx, got.ID)
		require.Equal(grp.Name, got.Name)
		users, err := grp.Users(ctx)
		require.Len(users, 2)
	})
}

func TestUpdateGroup(t *testing.T) {
	s := newService(t)
	defer s.EC.Close()
	ctx := context.Background()

	t.Run("name", func(t *testing.T) {
		require := require.New(t)
		grp := createGroup(ctx, s)

		got, err := s.UpdateGroup(ctx, grp.ID, model.UpdateGroupInput{Name: ptr.Str("new name")})
		require.NoError(err)
		require.NotNil(got)
		require.Equal("new name", got.Name)
		require.Equal(grp.Active, got.Active)

		g := s.EC.Group.GetX(ctx, got.ID)
		require.Equal("new name", g.Name)
		require.Equal(grp.Active, g.Active)
	})

	t.Run("active", func(t *testing.T) {
		require := require.New(t)
		grp := createGroup(ctx, s)

		got, err := s.UpdateGroup(ctx, grp.ID, model.UpdateGroupInput{Active: ptr.Bool(false)})
		require.NoError(err)
		require.NotNil(got)
		require.Equal(grp.Name, got.Name)
		require.Equal(false, got.Active)

		g := s.EC.Group.GetX(ctx, got.ID)
		require.Equal(grp.Name, g.Name)
		require.Equal(false, g.Active)
	})

	t.Run("all", func(t *testing.T) {
		require := require.New(t)
		grp := createGroup(ctx, s)

		got, err := s.UpdateGroup(ctx, grp.ID, model.UpdateGroupInput{Active: ptr.Bool(false), Name: ptr.Str("new name")})
		require.NoError(err)
		require.NotNil(got)
		require.Equal("new name", got.Name)
		require.Equal(false, got.Active)

		g := s.EC.Group.GetX(ctx, got.ID)
		require.Equal("new name", g.Name)
		require.Equal(false, g.Active)
	})
}

var ctx = context.Background()

func TestDeleteGroup(t *testing.T) {
	s := newService(t)
	defer s.EC.Close()
	require := require.New(t)

	grp := createGroup(ctx, s)
	grp2 := createGroup(ctx, s)

	require.Nil(grp.DeletedAt)

	err := s.DeleteGroup(ctx, grp.ID)
	require.NoError(err)

	grp = s.EC.Group.GetX(ctx, grp.ID)
	require.NotNil(grp.DeletedAt)

	grp2 = s.EC.Group.GetX(ctx, grp2.ID)
	require.Nil(grp.DeletedAt)
}
