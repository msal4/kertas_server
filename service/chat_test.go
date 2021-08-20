package service_test

import (
	"context"
	"strings"
	"testing"

	"github.com/99designs/gqlgen/graphql"
	"github.com/minio/minio-go/v7"
	"github.com/msal4/hassah_school_server/ent/group"
	"github.com/msal4/hassah_school_server/server/model"
	"github.com/msal4/hassah_school_server/testutil"
	"github.com/stretchr/testify/require"
)

func TestPostMessage(t *testing.T) {
	s := newService(t)
	defer s.EC.Close()

	ctx := context.Background()

	sch := createSchool(ctx, s, "myuniqueschoolname", "myimage")
	stdt := createStudent(ctx, s, "myuniqueusername", sch)
	tchr := createTeacher(ctx, s, "myuniqueteachernaem", sch)

	outsiderStdt := createStudent(ctx, s, "sdsmyuniqueusername", sch)

	grp := s.EC.Group.Create().SetName("test group name").AddUsers(tchr, stdt).SetGroupType(group.GroupTypePrivate).SaveX(ctx)

	t.Run("without attachment", func(t *testing.T) {
		require := require.New(t)

		input := model.PostMessageInput{GroupID: grp.ID, Content: "message test content"}
		got, err := s.PostMessage(ctx, stdt, input)
		require.NoError(err)
		require.NotNil(got)

		msg, err := s.EC.Message.Get(ctx, got.ID)
		require.NoError(err)
		require.Equal(msg.Content, got.Content)

		require.Equal(got.Content, input.Content)
		gotGrp, err := got.Group(ctx)
		require.NoError(err)
		require.Equal(grp.ID, gotGrp.ID)
	})

	t.Run("with attachment", func(t *testing.T) {
		require := require.New(t)

		f := testutil.OpenFile(t, "../testfiles/file.txt")

		input := model.PostMessageInput{GroupID: grp.ID, Content: "message test content", Attachment: &graphql.Upload{
			File:        f,
			Filename:    f.File.Name(),
			Size:        f.Size(),
			ContentType: f.ContentType,
		}}

		got, err := s.PostMessage(ctx, stdt, input)
		require.NoError(err)
		require.NotNil(got)

		require.Equal(got.Content, input.Content)
		gotGrp, err := got.Group(ctx)
		require.NoError(err)
		require.Equal(grp.ID, gotGrp.ID)
		require.NotEmpty(got.Attachment)

		_, err = s.MC.StatObject(ctx, s.Config.RootBucket, got.Attachment, minio.GetObjectOptions{})
		require.NoError(err)

		require.True(strings.HasPrefix(got.Attachment, stdt.Directory))
		require.True(strings.HasSuffix(got.Attachment, ".txt"))
	})

	t.Run("outsider", func(t *testing.T) {
		require := require.New(t)

		f := testutil.OpenFile(t, "../testfiles/file.txt")

		input := model.PostMessageInput{GroupID: grp.ID, Content: "message test content", Attachment: &graphql.Upload{
			File:        f,
			Filename:    f.File.Name(),
			Size:        f.Size(),
			ContentType: f.ContentType,
		}}

		got, err := s.PostMessage(ctx, outsiderStdt, input)
		require.Error(err)
		require.Nil(got)
	})
}

func TestRegisterGroupListener(t *testing.T) {
	s := newService(t)
	defer s.EC.Close()
	ctx := context.Background()

	sch := createSchool(ctx, s, "myuniqueschoolname", "myimage")
	stdt := createStudent(ctx, s, "myuniqueusername", sch)
	tchr := createTeacher(ctx, s, "myuniqueteachernaem", sch)

	grp := s.EC.Group.Create().SetName("test group name").AddUsers(tchr, stdt).SetGroupType(group.GroupTypePrivate).SaveX(ctx)

	t.Run("private", func(t *testing.T) {
		require := require.New(t)
		cancelableCtx, cancel := context.WithCancel(context.Background())

		msgCh, err := s.RegisterGroupListener(cancelableCtx, grp.ID, stdt.ID)
		require.NoError(err)

		input := model.PostMessageInput{GroupID: grp.ID, Content: "message test content"}
		msg, err := s.PostMessage(ctx, stdt, input)
		require.NoError(err)
		require.NotNil(msg)

		got := <-msgCh

		require.NotNil(got)
		require.Equal(msg.Content, got.Content)
		require.Equal(msg.ID, got.ID)
		require.Equal(msg.Attachment, got.Attachment)

		input = model.PostMessageInput{GroupID: grp.ID, Content: "message test content 2"}
		msg, err = s.PostMessage(ctx, tchr, input)
		require.NoError(err)
		require.NotNil(msg)

		got = <-msgCh

		require.NotNil(got)
		require.Equal(msg.Content, got.Content)
		require.Equal(msg.ID, got.ID)
		require.Equal(msg.Attachment, got.Attachment)

		cancel()

		input = model.PostMessageInput{GroupID: grp.ID, Content: "message test content 3"}
		msg, err = s.PostMessage(ctx, tchr, input)
		require.NoError(err)
		require.NotNil(msg)

		got = <-msgCh

		require.Nil(got)
	})
}
