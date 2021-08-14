package service_test

import (
	"context"
	"testing"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/ent/schema"
	"github.com/msal4/hassah_school_server/ent/user"
	"github.com/msal4/hassah_school_server/graph/model"
	"github.com/msal4/hassah_school_server/service"
	"github.com/msal4/hassah_school_server/testutil"
	"github.com/msal4/hassah_school_server/util/ptr"
	"github.com/stretchr/testify/require"
)

func TestUserList(t *testing.T) {
	s := newService(t)
	defer s.EC.Close()
	ctx := context.Background()

	t.Run("empty", func(t *testing.T) {
		users, err := s.Users(ctx, service.UserListOptions{})
		require.NoError(t, err)
		require.NotNil(t, users)
		require.Empty(t, users.Edges)
	})

	t.Run("not empty", func(t *testing.T) {
		defer s.EC.User.Delete().ExecX(ctx)

		want := s.EC.User.Create().SetName("test name").SetUsername("msal").SetPassword("test password").SetPhone("test phone").
			SetDirectory("testdir").SaveX(ctx)

		users, err := s.Users(ctx, service.UserListOptions{})
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

		s.EC.User.Create().SetName("test name 1").SetUsername("msal1").SetPassword("test password").SetPhone("test phone 1").
			SetDirectory("testdir").SaveX(ctx)
		s.EC.User.Create().SetName("test name 2").SetUsername("msal2").SetPassword("test password").SetPhone("test phone 2").
			SetDirectory("testdir").SaveX(ctx)
		s.EC.User.Create().SetName("test name 3").SetUsername("msal3").SetPassword("test password").SetPhone("test phone 3").
			SetDirectory("testdir").SaveX(ctx)
		s.EC.User.Create().SetName("test name 4").SetUsername("msal4").SetPassword("test password").SetPhone("test phone 4").
			SetDirectory("testdir").SaveX(ctx)

		b := s.EC.User.Query().Order(ent.Asc(user.FieldCreatedAt))
		want := b.AllX(ctx)

		conn, err := s.Users(ctx, service.UserListOptions{
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

		conn, err = s.Users(ctx, service.UserListOptions{
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

func TestUserAdd(t *testing.T) {
	s := newService(t)
	defer s.EC.Close()
	ctx := context.Background()

	t.Run("invalid", func(t *testing.T) {
		got, err := s.AddUser(ctx, model.AddUserInput{Name: "test user"})
		require.Error(t, err)
		require.Nil(t, got)
	})

	t.Run("valid", func(t *testing.T) {
		defer s.EC.User.Delete().ExecX(ctx)

		f := testutil.OpenFile(t, "../testfiles/stanford.png")
		defer f.Close()
		got, err := s.AddUser(ctx, model.AddUserInput{
			Name:     "test user",
			Username: "testusner",
			Password: "testpassword",
			Phone:    "testphone",
			Role:     user.RoleSUPER_ADMIN,
			Image:    &graphql.Upload{File: f, Filename: f.File.Name(), ContentType: f.ContentType, Size: f.Size()},
			Status:   schema.StatusActive,
		})
		require.NoError(t, err)
		require.NotNil(t, got)
		require.NotEmpty(t, got.Image)
		info, err := s.MC.StatObject(ctx, s.Config.RootBucket, got.Image, minio.StatObjectOptions{})
		require.NoError(t, err)
		require.Equal(t, "image/jpeg", info.ContentType)
	})

	t.Run("with invalid image", func(t *testing.T) {
		defer s.EC.User.Delete().ExecX(ctx)

		f := testutil.OpenFile(t, "../testfiles/file.txt")
		defer f.Close()
		got, err := s.AddUser(ctx, model.AddUserInput{
			Name:     "test user",
			Image:    &graphql.Upload{File: f, Filename: f.File.Name(), ContentType: f.ContentType, Size: f.Size()},
			Status:   schema.StatusActive,
			Username: "testusner",
			Password: "testpassword",
			Phone:    "testphone",
			Role:     user.RoleSTUDENT,
		})
		require.Error(t, err)
		require.Nil(t, got)
	})

	t.Run("with invalid role & stage", func(t *testing.T) {
		defer s.EC.User.Delete().ExecX(ctx)

		got, err := s.AddUser(ctx, model.AddUserInput{
			Name:     "test user",
			Role:     user.RoleSTUDENT,
			Status:   schema.StatusActive,
			Username: "testusner",
			Password: "testpassword",
			Phone:    "testphone",
		})

		require.Error(t, err)
		require.Nil(t, got)

		got, err = s.AddUser(ctx, model.AddUserInput{
			Name:     "test user",
			Role:     user.RoleTEACHER,
			Status:   schema.StatusActive,
			Username: "testusner2",
			Password: "testpassword",
			Phone:    "testphone",
		})

		require.Error(t, err)
		require.Nil(t, got)

		got, err = s.AddUser(ctx, model.AddUserInput{
			Name:     "test user",
			Role:     user.RoleSCHOOL_ADMIN,
			Status:   schema.StatusActive,
			Username: "testusner2",
			Password: "testpassword",
			Phone:    "testphone",
		})

		require.Error(t, err)
		require.Nil(t, got)
	})

	t.Run("with valid role & stage", func(t *testing.T) {
		defer s.EC.User.Delete().ExecX(ctx)

		sch := s.EC.School.Create().SetName("hello").SetImage("hi").SetDirectory("testdir").SetStatus(schema.StatusActive).SaveX(ctx)
		stage := s.EC.Stage.Create().SetName("first stage").SetSchool(sch).SetTuitionAmount(299).SaveX(ctx)

		got, err := s.AddUser(ctx, model.AddUserInput{
			Name:     "test user",
			Role:     user.RoleSTUDENT,
			StageID:  &stage.ID,
			Status:   schema.StatusActive,
			Username: "testusner",
			Password: "testpassword",
			Phone:    "testphone",
		})

		require.NoError(t, err)
		require.NotNil(t, got)

		gotStage, err := got.Stage(ctx)
		require.NoError(t, err)
		require.Equal(t, stage.ID, gotStage.ID)
		require.Equal(t, stage.ID, gotStage.ID)

		gotSchool, err := got.School(ctx)
		require.Equal(t, sch.ID, gotSchool.ID)
		require.Equal(t, sch.ID, gotSchool.ID)
	})
}

func TestUserUpdate(t *testing.T) {
	s := newService(t)
	defer s.EC.Close()
	ctx := context.Background()

	u, err := s.AddUser(ctx, model.AddUserInput{
		Name:     "test user",
		Role:     user.RoleSUPER_ADMIN,
		Status:   schema.StatusActive,
		Username: "testusner",
		Password: "testpassword",
		Phone:    "testphone",
	})
	require.NoError(t, err)

	t.Run("name", func(t *testing.T) {
		updated, err := s.UpdateUser(ctx, u.ID, model.UpdateUserInput{Name: ptr.Str("new name")})
		require.NoError(t, err)
		require.NotNil(t, updated)
		require.Equal(t, u.ID, updated.ID)
		require.Equal(t, "new name", updated.Name)
		require.Equal(t, u.Username, updated.Username)
	})

	t.Run("non existing", func(t *testing.T) {
		updated, err := s.UpdateUser(ctx, uuid.New(), model.UpdateUserInput{Name: ptr.Str("new name 2")})
		require.Error(t, err)
		require.Nil(t, updated)
	})

	t.Run("image", func(t *testing.T) {
		f := testutil.OpenFile(t, "../testfiles/stanford.png")
		defer f.Close()

		updated, err := s.UpdateUser(ctx, u.ID, model.UpdateUserInput{Name: ptr.Str("new name 3"), Image: &graphql.Upload{
			File:        f,
			Filename:    f.File.Name(),
			Size:        f.Size(),
			ContentType: f.ContentType}},
		)
		require.NoError(t, err)
		require.NotNil(t, updated)
		require.Equal(t, u.ID, updated.ID)
		require.Equal(t, "new name 3", updated.Name)
		require.Equal(t, u.Username, updated.Username)
		require.Equal(t, u.Status, updated.Status)
		_, err = s.MC.StatObject(ctx, s.Config.RootBucket, updated.Image, minio.StatObjectOptions{})
		require.NoError(t, err)
	})

	t.Run("all", func(t *testing.T) {
		f := testutil.OpenFile(t, "../testfiles/stanford.png")
		defer f.Close()

		updated, err := s.UpdateUser(ctx, u.ID, model.UpdateUserInput{
			Name: ptr.Str("new name 3"),
			Image: &graphql.Upload{
				File:        f,
				Filename:    f.File.Name(),
				Size:        f.Size(),
				ContentType: f.ContentType,
			},
			Phone:    ptr.Str("12345677"),
			Password: ptr.Str("newpasswort"),
			Username: ptr.Str("222newusername"),
			Status:   ptr.Status(schema.StatusDisabled),
		})
		require.NoError(t, err)
		require.NotNil(t, updated)
		require.Equal(t, u.ID, updated.ID)
		require.Equal(t, "new name 3", updated.Name)
		require.Equal(t, "222newusername", updated.Username)
		require.Equal(t, schema.StatusDisabled, updated.Status)
		require.Equal(t, "12345677", updated.Phone)
		require.Equal(t, u.Role, updated.Role)

		_, err = s.MC.StatObject(ctx, s.Config.RootBucket, updated.Image, minio.StatObjectOptions{})
		require.NoError(t, err)
	})
}

func TestDeleteUser(t *testing.T) {
	s := newService(t)
	defer s.EC.Close()
	ctx := context.Background()

	t.Run("valid", func(t *testing.T) {
		f := testutil.OpenFile(t, "../testfiles/stanford.png")
		defer f.Close()

		u, err := s.AddUser(ctx, model.AddUserInput{
			Name:     "test user",
			Role:     user.RoleSUPER_ADMIN,
			Status:   schema.StatusActive,
			Username: "testusner",
			Password: "testpassword",
			Phone:    "testphone",
			Image: &graphql.Upload{
				File:        f,
				Filename:    f.File.Name(),
				Size:        f.Size(),
				ContentType: f.ContentType,
			},
		})
		require.NoError(t, err)
		require.NotNil(t, u)

		err = s.DeleteUser(ctx, u.ID)
		require.NoError(t, err)

		deleted, err := s.EC.User.Get(ctx, u.ID)
		require.Error(t, err)
		require.Nil(t, deleted)

		_, err = s.MC.StatObject(ctx, s.Config.RootBucket, u.Image, minio.StatObjectOptions{})
		require.Error(t, err)
	})

	t.Run("invalid non existing user", func(t *testing.T) {
		err := s.DeleteUser(ctx, uuid.New())
		require.Error(t, err)
	})
}
