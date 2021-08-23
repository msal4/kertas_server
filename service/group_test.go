package service_test

import (
	"context"
	"testing"
	"time"

	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/service"
	"github.com/stretchr/testify/require"
)

func TestGroups(t *testing.T) {
	s := newService(t)
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
