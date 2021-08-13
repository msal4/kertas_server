package service_test

import (
	"context"
	"testing"

	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/ent/schema"
	"github.com/msal4/hassah_school_server/ent/user"
	"github.com/msal4/hassah_school_server/service"
	"github.com/msal4/hassah_school_server/util/ptr"
	"github.com/stretchr/testify/require"
)

func TestUserList(t *testing.T) {
	s := newService(t)
	defer s.EC.Close()
	ctx := context.Background()

	t.Run("empty", func(t *testing.T) {
		users, err := s.UserList(ctx, service.UserListOptions{})
		require.NoError(t, err)
		require.NotNil(t, users)
		require.Empty(t, users.Edges)
	})

	t.Run("not empty", func(t *testing.T) {
		defer s.EC.User.Delete().ExecX(ctx)

		want := s.EC.User.Create().SetName("test name").SetUsername("msal").SetPassword("test password").SetPhone("test phone").SaveX(ctx)

		users, err := s.UserList(ctx, service.UserListOptions{})
		require.NoError(t, err)
		require.NotNil(t, users)
		require.Len(t, users.Edges, 1)
		got := users.Edges[0].Node
		require.Equal(t, want.Username, got.Username)
		require.Equal(t, want.ID, got.ID)
		require.Equal(t, want.Name, got.Name)
	})

	t.Run("order & filter", func(t *testing.T) {
		defer s.EC.School.Delete().ExecX(ctx)

		s.EC.User.Create().SetName("test name 1").SetUsername("msal1").SetPassword("test password").SetPhone("test phone 1").SaveX(ctx)
		s.EC.User.Create().SetName("test name 2").SetUsername("msal2").SetPassword("test password").SetPhone("test phone 2").SaveX(ctx)
		s.EC.User.Create().SetName("test name 3").SetUsername("msal3").SetPassword("test password").SetPhone("test phone 3").SaveX(ctx)
		s.EC.User.Create().SetName("test name 4").SetUsername("msal4").SetPassword("test password").SetPhone("test phone 4").SaveX(ctx)

		b := s.EC.User.Query().Order(ent.Asc(user.FieldCreatedAt))
		want := b.AllX(ctx)

		conn, err := s.UserList(ctx, service.UserListOptions{
			OrderBy: &ent.UserOrder{Field: ent.UserOrderFieldCreatedAt, Direction: ent.OrderDirectionAsc},
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
			require.Equal(t, edge.Node.Username, want[i].Username)
			require.Equal(t, edge.Node.Phone, want[i].Phone)
			require.Equal(t, edge.Node.Status, want[i].Status)
			require.Equal(t, edge.Node.Image, want[i].Image)
			require.Equal(t, edge.Node.CreatedAt, want[i].CreatedAt)
			require.Equal(t, edge.Node.UpdatedAt, want[i].UpdatedAt)
		}

		want = b.Where(user.StatusEQ(schema.StatusDisabled)).AllX(ctx)

		conn, err = s.UserList(ctx, service.UserListOptions{
			Where:   &ent.UserWhereInput{Status: ptr.Status(schema.StatusDisabled)},
			OrderBy: &ent.UserOrder{Field: ent.UserOrderFieldCreatedAt, Direction: ent.OrderDirectionAsc},
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
