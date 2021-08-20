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
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/require"
)

func TestPostMessage(t *testing.T) {
	s := newService(t)
	defer s.EC.Close()

	ctx := context.Background()

	sch := createSchool(ctx, s, "myuniqueschoolname", "myimage")
	stdt := createStudent(ctx, s, "skdfjdskmyuniqueusername", sch)
	tchr := createTeacher(ctx, s, "myuniqkfjdkueteachernaem", sch)

	outsiderStdt := createStudent(ctx, s, "sdsmyuniqueusername", sch)

	grp := s.EC.Group.Create().SetName("test group name").AddUsers(tchr, stdt).SetGroupType(group.GroupTypePrivate).SaveX(ctx)

	t.Run("student message without attachment", func(t *testing.T) {
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

	t.Run("student message with attachment", func(t *testing.T) {
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

	t.Run("outsider student", func(t *testing.T) {
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

	t.Run("shared", func(t *testing.T) {
		require := require.New(t)

		stg := s.EC.Stage.Create().SetName("stage 1").SetDirectory("testdir").SetTuitionAmount(23423).SetSchool(sch).SaveX(ctx)
		grp := s.EC.Group.Create().SetName("test group name").AddUsers(tchr, stdt).SetGroupType(group.GroupTypeShared).SaveX(ctx)
		s.EC.Class.Create().SetName("math").SetStage(stg).SetTeacher(tchr).SetGroup(grp).SaveX(ctx)

		stdt := s.EC.User.Create().SetName("test userd").SetUsername("teststsdfksdjudent1").
			SetPhone("077059333812").SetDirectory("diresss22").SetPassword("mipassword22@@@@5").SetSchool(sch).SetStage(stg).SaveX(ctx)

		// outsider teacher
		sch := createSchool(ctx, s, "tesskd", "fkjds")
		outsiderTchr := createTeacher(ctx, s, ksuid.New().String(), sch)
		input := model.PostMessageInput{GroupID: grp.ID, Content: "message test content"}
		got, err := s.PostMessage(ctx, outsiderTchr, input)
		require.Error(err)
		require.Nil(got)

		// owner
		input = model.PostMessageInput{GroupID: grp.ID, Content: "message test content"}
		got, err = s.PostMessage(ctx, tchr, input)
		require.NoError(err)
		require.NotNil(got)
		require.Equal(input.Content, got.Content)

		// student
		input = model.PostMessageInput{GroupID: grp.ID, Content: "message test content"}
		got, err = s.PostMessage(ctx, stdt, input)
		require.NoError(err)
		require.NotNil(got)
		require.Equal(input.Content, got.Content)
	})
}

func TestRegisterGroupListener(t *testing.T) {
	s := newService(t)
	defer s.EC.Close()
	ctx := context.Background()

	sch := createSchool(ctx, s, "myuniqueschoolname", "myimage")
	stdt := createStudent(ctx, s, "sjkdfjsdmyuniqueusername", sch)
	tchr := createTeacher(ctx, s, "myuniqujjkfldseteachernaem", sch)
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
		<-cancelableCtx.Done()

		got, ok := <-msgCh

		require.False(ok)
	})

	t.Run("shared", func(t *testing.T) {
		require := require.New(t)
		stg := s.EC.Stage.Create().SetName("stage 1").SetDirectory("testdir").SetTuitionAmount(23423).SetSchool(sch).SaveX(ctx)
		grp := s.EC.Group.Create().SetName("test group name").AddUsers(tchr, stdt).SetGroupType(group.GroupTypeShared).SaveX(ctx)
		s.EC.Class.Create().SetName("math").SetStage(stg).SetTeacher(tchr).SetGroup(grp).SaveX(ctx)

		stdt := s.EC.User.Create().SetName("test userd").SetUsername("teststsdfksdjudent1").
			SetPhone("077059333812").SetDirectory("diresss22").SetPassword("mipassword22@@@@5").SetSchool(sch).SetStage(stg).SaveX(ctx)
		s.EC.User.Create().SetName("test userd 2").SetUsername("ksdjfklsdjteststudent2").
			SetPhone("077059333812").SetDirectory("diresss22").SetPassword("mipassword22@@@@5").SetSchool(sch).SetStage(stg).SaveX(ctx)
		s.EC.User.Create().SetName("test userd 3").SetUsername("teststudejfksdjfkldnt3").
			SetPhone("077059333812").SetDirectory("diresss22").SetPassword("mipassword22@@@@5").SetSchool(sch).SetStage(stg).SaveX(ctx)

		stg2 := s.EC.Stage.Create().SetName("stage 2").SetDirectory("testdir").SetTuitionAmount(23423).SetSchool(sch).SaveX(ctx)
		differentStageStdt := s.EC.User.Create().SetName("test userd 4").SetUsername("tefjdskfjdklsststudent4").
			SetPhone("077059333812").SetDirectory("diresss22").SetPassword("mipassword22@@@@5").SetSchool(sch).SetStage(stg2).SaveX(ctx)

		// student
		msgCh, err := s.RegisterGroupListener(ctx, grp.ID, stdt.ID)
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

		// teacher
		msgCh, err = s.RegisterGroupListener(ctx, grp.ID, tchr.ID)
		require.NoError(err)
		input = model.PostMessageInput{GroupID: grp.ID, Content: "message test content"}
		msg, err = s.PostMessage(ctx, stdt, input)
		require.NoError(err)
		require.NotNil(msg)
		got = <-msgCh
		require.NotNil(got)
		require.Equal(msg.Content, got.Content)
		require.Equal(msg.ID, got.ID)
		require.Equal(msg.Attachment, got.Attachment)

		// outside student
		msgCh, err = s.RegisterGroupListener(ctx, grp.ID, differentStageStdt.ID)
		require.Error(err)
		require.Nil(msgCh)

		anotherSch := createSchool(ctx, s, "tet", "kdfjs")
		// outside teacher
		outsideTchr := createTeacher(ctx, s, "skdmyuniqueteachernaem", anotherSch)
		msgCh, err = s.RegisterGroupListener(ctx, grp.ID, outsideTchr.ID)
		require.Error(err)
		require.Nil(msgCh)
	})

	t.Run("multiple devices", func(t *testing.T) {
		require := require.New(t)

		// device 1
		msgCh, err := s.RegisterGroupListener(ctx, grp.ID, stdt.ID)
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

		// device 2
		msgCh2, err := s.RegisterGroupListener(ctx, grp.ID, stdt.ID)
		require.NoError(err)
		input = model.PostMessageInput{GroupID: grp.ID, Content: "message test content"}
		msg, err = s.PostMessage(ctx, stdt, input)
		require.NoError(err)
		require.NotNil(msg)
		got = <-msgCh2

		require.NotNil(got)
		require.Equal(msg.Content, got.Content)
		require.Equal(msg.ID, got.ID)
		require.Equal(msg.Attachment, got.Attachment)

		got = <-msgCh
		require.NotNil(got)
		require.Equal(msg.Content, got.Content)
		require.Equal(msg.ID, got.ID)
		require.Equal(msg.Attachment, got.Attachment)
	})
}
