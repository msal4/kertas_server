package service_test

import (
	"context"
	"testing"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/msal4/hassah_school_server/auth"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/ent/user"
	"github.com/msal4/hassah_school_server/server/model"
	"github.com/msal4/hassah_school_server/service"
	"github.com/msal4/hassah_school_server/testutil"
	"github.com/msal4/hassah_school_server/util/ptr"
	"github.com/stretchr/testify/require"
)

func TestUsers(t *testing.T) {
	s := newService(t)
	defer s.EC.Close()
	ctx := context.Background()

	t.Run("empty", func(t *testing.T) {
		users, err := s.Users(ctx, service.UsersOptions{})
		require.NoError(t, err)
		require.NotNil(t, users)
		require.Empty(t, users.Edges)
	})

	t.Run("not empty", func(t *testing.T) {
		defer s.EC.User.Delete().ExecX(ctx)

		want := s.EC.User.Create().SetName("test name").SetUsername("msal").SetPassword("test password").SetPhone("test phone").
			SetDirectory("testdir").SaveX(ctx)

		users, err := s.Users(ctx, service.UsersOptions{})
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

		conn, err := s.Users(ctx, service.UsersOptions{
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
			require.Equal(t, edge.Node.Active, want[i].Active)
			require.Equal(t, edge.Node.Image, want[i].Image)
			require.Equal(t, edge.Node.CreatedAt, want[i].CreatedAt)
			require.Equal(t, edge.Node.UpdatedAt, want[i].UpdatedAt)
		}

		want = b.Where(user.Active(false)).AllX(ctx)

		conn, err = s.Users(ctx, service.UsersOptions{
			Where:   &ent.UserWhereInput{Active: ptr.Bool(false)},
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
			require.Equal(t, e.Node.Active, w.Active)
			require.Equal(t, e.Node.Image, w.Image)
			require.Equal(t, e.Node.CreatedAt, w.CreatedAt)
		}

	})
}

func TestAddUser(t *testing.T) {
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
			Role:     user.RoleSuperAdmin,
			Image:    &graphql.Upload{File: f, Filename: f.File.Name(), ContentType: f.ContentType, Size: f.Size()},
			Active:   true,
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
			Active:   true,
			Username: "testusner",
			Password: "testpassword",
			Phone:    "testphone",
			Role:     user.RoleStudent,
		})
		require.Error(t, err)
		require.Nil(t, got)
	})

	t.Run("with invalid role & stage", func(t *testing.T) {
		defer s.EC.User.Delete().ExecX(ctx)

		got, err := s.AddUser(ctx, model.AddUserInput{
			Name:     "test user",
			Role:     user.RoleStudent,
			Active:   true,
			Username: "testusner",
			Password: "testpassword",
			Phone:    "testphone",
		})

		require.Error(t, err)
		require.Nil(t, got)

		got, err = s.AddUser(ctx, model.AddUserInput{
			Name:     "test user",
			Role:     user.RoleTeacher,
			Active:   true,
			Username: "testusner2",
			Password: "testpassword",
			Phone:    "testphone",
		})

		require.Error(t, err)
		require.Nil(t, got)

		got, err = s.AddUser(ctx, model.AddUserInput{
			Name:     "test user",
			Role:     user.RoleSchoolAdmin,
			Active:   true,
			Username: "testusner2",
			Password: "testpassword",
			Phone:    "testphone",
		})

		require.Error(t, err)
		require.Nil(t, got)
	})

	t.Run("with valid role & stage", func(t *testing.T) {
		defer s.EC.User.Delete().ExecX(ctx)

		sch := s.EC.School.Create().SetName("hello").SetImage("hi").SetDirectory("testdir").SetActive(true).SaveX(ctx)
		stage := s.EC.Stage.Create().SetName("first stage").SetSchool(sch).SetTuitionAmount(299).SetDirectory("testdir").SaveX(ctx)

		got, err := s.AddUser(ctx, model.AddUserInput{
			Name:     "test user",
			Role:     user.RoleStudent,
			StageID:  &stage.ID,
			Active:   true,
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

		gotSchool, _ := got.School(ctx)
		require.Equal(t, sch.ID, gotSchool.ID)
		require.Equal(t, sch.ID, gotSchool.ID)
	})
}

func TestUpdateUser(t *testing.T) {
	s := newService(t)
	defer s.EC.Close()
	ctx := context.Background()

	u, err := s.AddUser(ctx, model.AddUserInput{
		Name:     "test user",
		Role:     user.RoleSuperAdmin,
		Active:   true,
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
		require.Equal(t, u.Active, updated.Active)
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
			Active:   ptr.Bool(false),
		})
		require.NoError(t, err)
		require.NotNil(t, updated)
		require.Equal(t, u.ID, updated.ID)
		require.Equal(t, "new name 3", updated.Name)
		require.Equal(t, "222newusername", updated.Username)
		require.Equal(t, false, updated.Active)
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
			Role:     user.RoleSuperAdmin,
			Active:   true,
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
		require.NoError(t, err)
		require.NotNil(t, deleted.DeletedAt)

		_, err = s.MC.StatObject(ctx, s.Config.RootBucket, u.Image, minio.StatObjectOptions{})
		require.NoError(t, err)
	})

	t.Run("invalid non existing user", func(t *testing.T) {
		err := s.DeleteUser(ctx, uuid.New())
		require.Error(t, err)
	})
}

func TestDeleteUserPermanently(t *testing.T) {
	s := newService(t)
	defer s.EC.Close()
	ctx := context.Background()

	t.Run("valid", func(t *testing.T) {
		f := testutil.OpenFile(t, "../testfiles/stanford.png")
		defer f.Close()

		u, err := s.AddUser(ctx, model.AddUserInput{
			Name:     "test user",
			Role:     user.RoleSuperAdmin,
			Active:   true,
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

		err = s.DeleteUserPermanently(ctx, u.ID)
		require.NoError(t, err)

		deleted, err := s.EC.User.Get(ctx, u.ID)
		require.Error(t, err)
		require.Nil(t, deleted)

		_, err = s.MC.StatObject(ctx, s.Config.RootBucket, u.Image, minio.StatObjectOptions{})
		require.Error(t, err)
	})

	t.Run("invalid non existing user", func(t *testing.T) {
		err := s.DeleteUserPermanently(ctx, uuid.New())
		require.Error(t, err)
	})
}

func TestLoginAdmin(t *testing.T) {
	s := newService(t)
	defer s.EC.Close()
	ctx := context.Background()

	input := model.AddUserInput{
		Name:     "test name",
		Username: "testusername",
		Password: "testpassword",
		Phone:    "07705983835",
		Role:     user.RoleSuperAdmin,
		Active:   true,
	}

	u, err := s.AddUser(ctx, input)
	require.NoError(t, err)
	require.NotNil(t, u)

	t.Run("invalid password", func(t *testing.T) {
		resp, err := s.LoginAdmin(ctx, model.LoginInput{
			Username: input.Username,
			Password: "wrongpassword",
		})
		require.Error(t, err)
		require.Nil(t, resp)
		require.ErrorIs(t, err, service.ErrInvalidCreds)
	})

	t.Run("invalid username", func(t *testing.T) {
		resp, err := s.LoginAdmin(ctx, model.LoginInput{
			Username: input.Username + "whatever",
			Password: input.Password,
		})
		require.Error(t, err)
		require.Nil(t, resp)
		require.ErrorIs(t, err, service.ErrNotFound)
	})

	t.Run("disabled user", func(t *testing.T) {
		input := model.AddUserInput{
			Name:     "test name",
			Username: "testusername22",
			Password: "testpassword",
			Phone:    "07705983835",
			Role:     user.RoleSuperAdmin,
		}

		u, err := s.AddUser(ctx, input)
		require.NoError(t, err)
		require.NotNil(t, u)

		resp, err := s.LoginAdmin(ctx, model.LoginInput{
			Username: input.Username,
			Password: input.Password,
		})
		require.Error(t, err)
		require.Nil(t, resp)
		require.ErrorIs(t, err, service.ErrUserDisabled)
	})

	t.Run("deleted user", func(t *testing.T) {
		input := model.AddUserInput{
			Name:     "test name",
			Username: "testusername223",
			Password: "testpassword",
			Phone:    "07705983835",
			Role:     user.RoleSuperAdmin,
			Active:   true,
		}

		u, err := s.AddUser(ctx, input)
		u = u.Update().SetDeletedAt(time.Now()).SaveX(ctx)
		require.NoError(t, err)
		require.NotNil(t, u)

		resp, err := s.LoginAdmin(ctx, model.LoginInput{
			Username: input.Username,
			Password: input.Password,
		})
		require.Error(t, err)
		require.Nil(t, resp)
		require.ErrorIs(t, err, service.ErrNotFound)
	})

	t.Run("super admin", func(t *testing.T) {
		resp, err := s.LoginAdmin(ctx, model.LoginInput{
			Username: input.Username,
			Password: input.Password,
		})
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, resp.AccessToken)
		require.NotEmpty(t, resp.RefreshToken)

		var accessClaims auth.AccessClaims
		token, err := jwt.ParseWithClaims(resp.AccessToken, &accessClaims, func(t *jwt.Token) (interface{}, error) {
			return []byte(s.Config.AccessSecretKey), nil
		})
		require.NoError(t, err)
		require.NotNil(t, token)
		require.True(t, token.Valid)
		require.Equal(t, u.ID, accessClaims.UserID)
		require.Equal(t, u.Role, accessClaims.Role)

		var refreshClaims auth.RefreshClaims
		token, err = jwt.ParseWithClaims(resp.RefreshToken, &refreshClaims, func(t *jwt.Token) (interface{}, error) {
			return []byte(s.Config.RefreshSecretKey), nil
		})
		require.NoError(t, err)
		require.NotNil(t, token)
		require.True(t, token.Valid)
		require.Equal(t, u.ID, refreshClaims.UserID)
		require.Equal(t, u.TokenVersion, refreshClaims.TokenVersion)
	})

	t.Run("school admin", func(t *testing.T) {
		sch := s.EC.School.Create().SetName("test school").SetActive(true).SetImage("hello").SetDirectory("whatev").SaveX(ctx)

		input := model.AddUserInput{
			Name:     "test name",
			Username: "testusernamesss",
			Password: "testpassword",
			Phone:    "07705983835",
			Role:     user.RoleSchoolAdmin,
			Active:   true,
			SchoolID: &sch.ID,
		}

		u, err := s.AddUser(ctx, input)
		require.NoError(t, err)
		require.NotNil(t, u)

		resp, err := s.LoginAdmin(ctx, model.LoginInput{
			Username: input.Username,
			Password: input.Password,
		})
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, resp.AccessToken)
		require.NotEmpty(t, resp.RefreshToken)

		var accessClaims auth.AccessClaims
		token, err := jwt.ParseWithClaims(resp.AccessToken, &accessClaims, func(t *jwt.Token) (interface{}, error) {
			return []byte(s.Config.AccessSecretKey), nil
		})
		require.NoError(t, err)
		require.NotNil(t, token)
		require.True(t, token.Valid)
		require.Equal(t, u.ID, accessClaims.UserID)
		require.Equal(t, u.Role, accessClaims.Role)

		var refreshClaims auth.RefreshClaims
		token, err = jwt.ParseWithClaims(resp.RefreshToken, &refreshClaims, func(t *jwt.Token) (interface{}, error) {
			return []byte(s.Config.RefreshSecretKey), nil
		})
		require.NoError(t, err)
		require.NotNil(t, token)
		require.True(t, token.Valid)
		require.Equal(t, u.ID, refreshClaims.UserID)
		require.Equal(t, u.TokenVersion, refreshClaims.TokenVersion)
	})

	t.Run("disabled school", func(t *testing.T) {
		sch := s.EC.School.Create().SetName("test school").SetActive(false).SetImage("hello").SetDirectory("whatev").SaveX(ctx)

		input := model.AddUserInput{
			Name:     "test name",
			Username: "testusernam3242",
			Password: "testpassword",
			Phone:    "07705983835",
			Role:     user.RoleSchoolAdmin,
			Active:   true,
			SchoolID: &sch.ID,
		}

		u, err := s.AddUser(ctx, input)
		require.NoError(t, err)
		require.NotNil(t, u)

		resp, err := s.LoginAdmin(ctx, model.LoginInput{
			Username: input.Username,
			Password: input.Password,
		})
		require.Error(t, err)
		require.Nil(t, resp)
		require.ErrorIs(t, err, service.ErrSchoolDisabled)

		sch.Update().SetDeletedAt(time.Now()).SetActive(true).SaveX(ctx)

		resp, err = s.LoginAdmin(ctx, model.LoginInput{
			Username: input.Username,
			Password: input.Password,
		})
		require.Error(t, err)
		require.Nil(t, resp)
		require.ErrorIs(t, err, service.ErrNotFound)
	})

	t.Run("teacher & student", func(t *testing.T) {
		sch := s.EC.School.Create().SetName("test school").SetActive(true).SetImage("hello").SetDirectory("whatev").SaveX(ctx)

		input := model.AddUserInput{
			Name:     "test name",
			Username: "22testusernam3242",
			Password: "testpassword",
			Phone:    "07705983835",
			Role:     user.RoleTeacher,
			Active:   true,
			SchoolID: &sch.ID,
		}

		u, err := s.AddUser(ctx, input)
		require.NoError(t, err)
		require.NotNil(t, u)

		resp, err := s.LoginAdmin(ctx, model.LoginInput{
			Username: input.Username,
			Password: input.Password,
		})
		require.Error(t, err)
		require.Nil(t, resp)
		require.ErrorIs(t, err, service.ErrNotAllowed)

		stage := s.EC.Stage.Create().SetName("1st").SetTuitionAmount(1000).SetActive(true).SetSchool(sch).SetDirectory("ifidsfksd").SaveX(ctx)
		u.Update().SetStage(stage).SaveX(ctx)

		resp, err = s.LoginAdmin(ctx, model.LoginInput{
			Username: input.Username,
			Password: input.Password,
		})
		require.Error(t, err)
		require.Nil(t, resp)
		require.ErrorIs(t, err, service.ErrNotAllowed)
	})
}

func TestLoginUser(t *testing.T) {
	s := newService(t)
	defer s.EC.Close()
	ctx := context.Background()

	sch := s.EC.School.Create().SetName("test school").SetActive(true).SetImage("hello").SetDirectory("whatev").SaveX(ctx)
	stage := s.EC.Stage.Create().SetName("1st").SetTuitionAmount(1000).SetActive(true).SetSchool(sch).SetDirectory("iifidfjs").SaveX(ctx)

	input := model.AddUserInput{
		Name:     "test name",
		Username: "testusername",
		Password: "testpassword",
		Phone:    "07705983835",
		Role:     user.RoleStudent,
		Active:   true,
		SchoolID: &sch.ID,
		StageID:  &stage.ID,
	}

	u, err := s.AddUser(ctx, input)
	require.NoError(t, err)
	require.NotNil(t, u)

	t.Run("invalid password", func(t *testing.T) {
		resp, err := s.LoginUser(ctx, model.LoginInput{
			Username: input.Username,
			Password: "wrongpassword",
		})
		require.Error(t, err)
		require.Nil(t, resp)
		require.ErrorIs(t, err, service.ErrInvalidCreds)
	})

	t.Run("invalid username", func(t *testing.T) {
		resp, err := s.LoginUser(ctx, model.LoginInput{
			Username: input.Username + "whatever",
			Password: input.Password,
		})
		require.Error(t, err)
		require.Nil(t, resp)
		require.ErrorIs(t, err, service.ErrNotFound)
	})

	t.Run("disabled user", func(t *testing.T) {
		input := model.AddUserInput{
			Name:     "test name",
			Username: "stestusername22",
			Password: "testpassword",
			Phone:    "07705983835",
			Role:     user.RoleStudent,
			StageID:  &stage.ID,
		}

		u, err := s.AddUser(ctx, input)
		require.NoError(t, err)
		require.NotNil(t, u)

		resp, err := s.LoginUser(ctx, model.LoginInput{
			Username: input.Username,
			Password: input.Password,
		})
		require.Error(t, err)
		require.Nil(t, resp)
		require.ErrorIs(t, err, service.ErrUserDisabled)
	})

	t.Run("deleted user", func(t *testing.T) {
		input := model.AddUserInput{
			Name:     "test name",
			Username: "test4username223",
			Password: "testpassword",
			Phone:    "07705983835",
			Role:     user.RoleTeacher,
			SchoolID: &sch.ID,
			Active:   true,
		}

		u, err := s.AddUser(ctx, input)
		u = u.Update().SetDeletedAt(time.Now()).SaveX(ctx)
		require.NoError(t, err)
		require.NotNil(t, u)

		resp, err := s.LoginUser(ctx, model.LoginInput{
			Username: input.Username,
			Password: input.Password,
		})
		require.Error(t, err)
		require.Nil(t, resp)
		require.ErrorIs(t, err, service.ErrNotFound)
	})

	t.Run("student", func(t *testing.T) {
		resp, err := s.LoginUser(ctx, model.LoginInput{
			Username: input.Username,
			Password: input.Password,
		})
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, resp.AccessToken)
		require.NotEmpty(t, resp.RefreshToken)

		var accessClaims auth.AccessClaims
		token, err := jwt.ParseWithClaims(resp.AccessToken, &accessClaims, func(t *jwt.Token) (interface{}, error) {
			return []byte(s.Config.AccessSecretKey), nil
		})
		require.NoError(t, err)
		require.NotNil(t, token)
		require.True(t, token.Valid)
		require.Equal(t, u.ID, accessClaims.UserID)
		require.Equal(t, u.Role, accessClaims.Role)

		var refreshClaims auth.RefreshClaims
		token, err = jwt.ParseWithClaims(resp.RefreshToken, &refreshClaims, func(t *jwt.Token) (interface{}, error) {
			return []byte(s.Config.RefreshSecretKey), nil
		})
		require.NoError(t, err)
		require.NotNil(t, token)
		require.True(t, token.Valid)
		require.Equal(t, u.ID, refreshClaims.UserID)
		require.Equal(t, u.TokenVersion, refreshClaims.TokenVersion)
	})

	t.Run("teacher", func(t *testing.T) {
		input := model.AddUserInput{
			Name:     "test name",
			Username: "testusernamesss",
			Password: "testpassword",
			Phone:    "07705983835",
			Role:     user.RoleTeacher,
			Active:   true,
			SchoolID: &sch.ID,
		}

		u, err := s.AddUser(ctx, input)
		require.NoError(t, err)
		require.NotNil(t, u)

		resp, err := s.LoginUser(ctx, model.LoginInput{
			Username: input.Username,
			Password: input.Password,
		})
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, resp.AccessToken)
		require.NotEmpty(t, resp.RefreshToken)

		var accessClaims auth.AccessClaims
		token, err := jwt.ParseWithClaims(resp.AccessToken, &accessClaims, func(t *jwt.Token) (interface{}, error) {
			return []byte(s.Config.AccessSecretKey), nil
		})
		require.NoError(t, err)
		require.NotNil(t, token)
		require.True(t, token.Valid)
		require.Equal(t, u.ID, accessClaims.UserID)
		require.Equal(t, u.Role, accessClaims.Role)

		var refreshClaims auth.RefreshClaims
		token, err = jwt.ParseWithClaims(resp.RefreshToken, &refreshClaims, func(t *jwt.Token) (interface{}, error) {
			return []byte(s.Config.RefreshSecretKey), nil
		})
		require.NoError(t, err)
		require.NotNil(t, token)
		require.True(t, token.Valid)
		require.Equal(t, u.ID, refreshClaims.UserID)
		require.Equal(t, u.TokenVersion, refreshClaims.TokenVersion)
	})

	t.Run("disabled school", func(t *testing.T) {
		sch := s.EC.School.Create().SetName("test school").SetActive(false).SetImage("hello").SetDirectory("whatev").SaveX(ctx)

		input := model.AddUserInput{
			Name:     "test name",
			Username: "testusernam3242",
			Password: "testpassword",
			Phone:    "07705983835",
			Role:     user.RoleTeacher,
			Active:   true,
			SchoolID: &sch.ID,
		}

		u, err := s.AddUser(ctx, input)
		require.NoError(t, err)
		require.NotNil(t, u)

		resp, err := s.LoginUser(ctx, model.LoginInput{
			Username: input.Username,
			Password: input.Password,
		})
		require.Error(t, err)
		require.Nil(t, resp)
		require.ErrorIs(t, err, service.ErrSchoolDisabled)

		sch.Update().SetDeletedAt(time.Now()).SetActive(true).SaveX(ctx)

		resp, err = s.LoginUser(ctx, model.LoginInput{
			Username: input.Username,
			Password: input.Password,
		})
		require.Error(t, err)
		require.Nil(t, resp)
		require.ErrorIs(t, err, service.ErrNotFound)
	})

	t.Run("disabled stage", func(t *testing.T) {
		stage := s.EC.Stage.Create().SetName("test naem").SetActive(false).SetSchool(sch).SetTuitionAmount(100).SetDirectory("k").SaveX(ctx)
		input := model.AddUserInput{
			Name:     "test name",
			Username: "s3stestusernam3242",
			Password: "testpassword",
			Phone:    "07705983835",
			Role:     user.RoleStudent,
			Active:   true,
			StageID:  &stage.ID,
		}
		u, err := s.AddUser(ctx, input)
		require.NoError(t, err)
		require.NotNil(t, u)

		resp, err := s.LoginUser(ctx, model.LoginInput{
			Username: input.Username,
			Password: input.Password,
		})
		require.Error(t, err)
		require.Nil(t, resp)
		require.ErrorIs(t, err, service.ErrStageDisabled)

		stage.Update().SetDeletedAt(time.Now()).SetActive(true).SaveX(ctx)

		resp, err = s.LoginUser(ctx, model.LoginInput{
			Username: input.Username,
			Password: input.Password,
		})
		require.Error(t, err)
		require.Nil(t, resp)
		require.ErrorIs(t, err, service.ErrNotFound)
	})

	t.Run("super admin & school admin", func(t *testing.T) {
		sch := s.EC.School.Create().SetName("test school").SetActive(true).SetImage("hello").SetDirectory("whatev").SaveX(ctx)

		input := model.AddUserInput{
			Name:     "test name",
			Username: "22testusernam3242",
			Password: "testpassword",
			Phone:    "07705983835",
			Role:     user.RoleSuperAdmin,
			Active:   true,
		}

		u, err := s.AddUser(ctx, input)
		require.NoError(t, err)
		require.NotNil(t, u)

		resp, err := s.LoginUser(ctx, model.LoginInput{
			Username: input.Username,
			Password: input.Password,
		})
		require.Error(t, err)
		require.Nil(t, resp)
		require.ErrorIs(t, err, service.ErrNotAllowed)

		u.Update().SetRole(user.RoleSuperAdmin).SetSchool(sch).SaveX(ctx)

		resp, err = s.LoginUser(ctx, model.LoginInput{
			Username: input.Username,
			Password: input.Password,
		})
		require.Error(t, err)
		require.Nil(t, resp)
		require.ErrorIs(t, err, service.ErrNotAllowed)
	})
}

func TestRefreshTokens(t *testing.T) {
	s := newService(t)
	ctx := context.Background()

	t.Run("valid", func(t *testing.T) {
		input := model.AddUserInput{
			Name:     "test name",
			Username: "testusernamesss",
			Password: "testpassword",
			Phone:    "07705983835",
			Role:     user.RoleSuperAdmin,
			Active:   true,
		}

		u, err := s.AddUser(ctx, input)
		require.NoError(t, err)
		require.NotNil(t, u)

		resp, err := s.LoginAdmin(ctx, model.LoginInput{
			Username: input.Username,
			Password: input.Password,
		})
		require.NoError(t, err)
		require.NotNil(t, resp)

		data, err := s.RefreshTokens(ctx, resp.RefreshToken)
		require.NoError(t, err)
		require.NotNil(t, data)
		require.NotEmpty(t, data.AccessToken)
		require.NotEmpty(t, data.RefreshToken)

		var accessClaims auth.AccessClaims
		token, err := jwt.ParseWithClaims(data.AccessToken, &accessClaims, func(t *jwt.Token) (interface{}, error) {
			return []byte(s.Config.AccessSecretKey), nil
		})
		require.NoError(t, err)
		require.NotNil(t, token)
		require.True(t, token.Valid)
		require.Equal(t, u.ID, accessClaims.UserID)
		require.Equal(t, u.Role, accessClaims.Role)

		var refreshClaims auth.RefreshClaims
		token, err = jwt.ParseWithClaims(resp.RefreshToken, &refreshClaims, func(t *jwt.Token) (interface{}, error) {
			return []byte(s.Config.RefreshSecretKey), nil
		})
		require.NoError(t, err)
		require.NotNil(t, token)
		require.True(t, token.Valid)
		require.Equal(t, u.ID, refreshClaims.UserID)
		require.Equal(t, u.TokenVersion, refreshClaims.TokenVersion)
	})

	t.Run("expired token", func(t *testing.T) {
		s := newService(t)
		s.Config.RefreshTokenLifetime = -time.Second

		input := model.AddUserInput{
			Name:     "test name",
			Username: "testusernamess4s2",
			Password: "testpassword",
			Phone:    "07705983835",
			Role:     user.RoleSuperAdmin,
			Active:   true,
		}

		u, err := s.AddUser(ctx, input)
		require.NoError(t, err)
		require.NotNil(t, u)

		resp, err := s.LoginAdmin(ctx, model.LoginInput{
			Username: input.Username,
			Password: input.Password,
		})
		require.NoError(t, err)
		require.NotNil(t, resp)

		data, err := s.RefreshTokens(ctx, resp.RefreshToken)
		require.Error(t, err)
		require.Nil(t, data)
	})

	t.Run("disabled user", func(t *testing.T) {
		s := newService(t)
		input := model.AddUserInput{
			Name:     "test name",
			Username: "sstestusernamess4s2",
			Password: "testpassword",
			Phone:    "07705983835",
			Role:     user.RoleSuperAdmin,
			Active:   true,
		}

		u, err := s.AddUser(ctx, input)
		require.NoError(t, err)
		require.NotNil(t, u)

		resp, err := s.LoginAdmin(ctx, model.LoginInput{
			Username: input.Username,
			Password: input.Password,
		})
		require.NoError(t, err)
		require.NotNil(t, resp)

		u.Update().SetActive(false).SaveX(ctx)

		data, err := s.RefreshTokens(ctx, resp.RefreshToken)
		require.Error(t, err)
		require.Nil(t, data)
		require.ErrorIs(t, err, service.ErrUserDisabled)

		u.Update().SetActive(true).SetDeletedAt(time.Now()).SaveX(ctx)
		data, err = s.RefreshTokens(ctx, resp.RefreshToken)
		require.Error(t, err)
		require.Nil(t, data)
		require.ErrorIs(t, err, service.ErrNotFound)
	})

	t.Run("token version mismatch", func(t *testing.T) {
		s := newService(t)

		input := model.AddUserInput{
			Name:     "test name",
			Username: "testusernamess4s2",
			Password: "testpassword",
			Phone:    "07705983835",
			Role:     user.RoleSuperAdmin,
			Active:   true,
		}

		u, err := s.AddUser(ctx, input)
		require.NoError(t, err)
		require.NotNil(t, u)

		resp, err := s.LoginAdmin(ctx, model.LoginInput{
			Username: input.Username,
			Password: input.Password,
		})
		require.NoError(t, err)
		require.NotNil(t, resp)

		u.Update().SetTokenVersion(4).SaveX(ctx)

		data, err := s.RefreshTokens(ctx, resp.RefreshToken)
		require.Error(t, err)
		require.Nil(t, data)
	})
}
