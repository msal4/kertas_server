package server_test

import (
	"context"
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/msal4/hassah_school_server/auth"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/ent/group"
	"github.com/msal4/hassah_school_server/server"
	"github.com/msal4/hassah_school_server/service"
	"github.com/msal4/hassah_school_server/testutil"
	"github.com/msal4/hassah_school_server/util/ptr"
	"github.com/stretchr/testify/require"
)

func TestMessages(t *testing.T) {
	s := newService(t)
	defer s.EC.Close()

	srv := server.NewServer(s, false)
	ctx := context.Background()

	suAdmin := createSuperAdmin(ctx, s, "hello23super")
	schAdmin := createSchoolAdmin(ctx, s, "jfskfkdsj444")
	teacher := createTeacher(ctx, s, "fjfskfkdsj444ksdjfkd")
	student := createStudent(ctx, s, "jksdfjdk331")
	outsiderStudent := createStudent(ctx, s, "jflksjdskd2342k")

	grp := s.EC.Group.Create().AddUsers(suAdmin, schAdmin, teacher, student).
		SetName("group").SetGroupType(group.GroupTypePrivate).SaveX(ctx)

	operations := fmt.Sprintf(`{ messages(groupID: %q) { totalCount pageInfo { hasNextPage hasPreviousPage startCursor endCursor } edges { node { id } cursor } } }`, grp.ID)

	cases := []struct {
		desc string
		user *ent.User
		err  *string
	}{
		{"super admin is allowed", suAdmin, nil},
		{"school admin is allowed", schAdmin, nil},
		{"teacher is allowed", teacher, nil},
		{"student is allowed", student, nil},
		{"outsider student is not allowd", outsiderStudent, ptr.Str(service.ErrNotAllowed.Error())},
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			var resp errsResponse

			r := createRequest(t, operations, "{}")
			w := httptest.NewRecorder()

			data := genTokens(t, c.user, s)

			setAuth(r, data.AccessToken)

			srv.ServeHTTP(w, r)

			parseBody(t, w, &resp)

			if c.err == nil {
				require.Nil(t, resp.Errors)
				return
			}

			require.NotEmpty(t, resp.Errors)
			require.Equal(t, *c.err, resp.Errors[0].Message)
		})
	}
}

func TestPostMessage(t *testing.T) {
	s := newService(t)
	defer s.EC.Close()

	srv := server.NewServer(s, false)
	ctx := context.Background()

	suAdmin := createSuperAdmin(ctx, s, "hello23super")
	schAdmin := createSchoolAdmin(ctx, s, "jfskfkdsj444")
	teacher := createTeacher(ctx, s, "fjfskfkdsj444ksdjfkd")
	student := createStudent(ctx, s, "jksdfjdk331")
	outsiderStudent := createStudent(ctx, s, "jflksjdskd2342k")

	grp := s.EC.Group.Create().AddUsers(suAdmin, schAdmin, teacher, student).
		SetName("group").SetGroupType(group.GroupTypePrivate).SaveX(ctx)

	operations := `{
	"query": "mutation ($image: Upload!) { postMessage(input: {groupID: \"%s\", content: \"a test message\", attachment: $image}) { id } }",
	"variables": {"image": null}
		}`

	cases := []struct {
		desc string
		user *ent.User
		err  *string
	}{
		{"super admin is allowed", suAdmin, nil},
		{"school admin is allowed", schAdmin, nil},
		{"teacher is allowed", teacher, nil},
		{"student is allowed", student, nil},
		{"outsider student is not allowed", outsiderStudent, ptr.Str(service.ErrNotAllowed.Error())},
	}

	f := testutil.OpenFile(t, "../testfiles/harvard.jpg")

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			var resp errsResponse

			f.Seek(0, 0)

			r := createMultipartRequest(t, fmt.Sprintf(operations, grp.ID), `{"0": ["variables.image"]}`, file{File: f.File, mapKey: "0"})
			w := httptest.NewRecorder()

			data := genTokens(t, c.user, s)

			setAuth(r, data.AccessToken)

			srv.ServeHTTP(w, r)

			parseBody(t, w, &resp)

			if c.err == nil {
				require.Nil(t, resp.Errors)
				return
			}

			require.NotEmpty(t, resp.Errors)
			require.Equal(t, *c.err, resp.Errors[0].Message)
		})
	}
}

type msgResponse struct {
	Type    string        `json:"type"`
	Payload *errsResponse `json:"payload,omitempty"`
}

func TestMessagePosted(t *testing.T) {

	s := newService(t)
	defer s.EC.Close()

	srv := server.NewServer(s, false)
	ctx := context.Background()

	suAdmin := createSuperAdmin(ctx, s, "hello23super")
	schAdmin := createSchoolAdmin(ctx, s, "jfskfkdsj444")
	teacher := createTeacher(ctx, s, "fjfskfkdsj444ksdjfkd")
	student := createStudent(ctx, s, "jksdfjdk331")
	outsiderStudent := createStudent(ctx, s, "jflksjdskd2342k")

	grp := s.EC.Group.Create().AddUsers(suAdmin, schAdmin, teacher, student).
		SetName("group").SetGroupType(group.GroupTypePrivate).SaveX(ctx)

	cases := []struct {
		desc string
		user *ent.User
		err  *string
	}{
		{"super admin is allowed", suAdmin, nil},
		{"school admin is allowed", schAdmin, nil},
		{"teacher is allowed", teacher, nil},
		{"student is allowed", student, nil},
		{"outsider student is not allowed", outsiderStudent, ptr.Str(service.ErrNotAllowed.Error())},
	}

	operationsMsg := []byte(fmt.Sprintf(`{
    "id": "1",
    "type": "start",
    "payload": {
        "variables": {},
        "extensions": {},
        "operationName": null,
        "query": "subscription {\n  messagePosted(groupID: \"%s\") {\n    id\n    content\n  }\n}\n"
    }
}`, grp.ID))

	t.Run("unauthenticated", func(t *testing.T) {
		require := require.New(t)

		server := httptest.NewServer(srv)
		defer server.Close()

		wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/graphql"

		ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
		require.NoError(err)
		defer ws.Close()

		msg := []byte(fmt.Sprintf(`{
    "type": "connection_init",
    "payload": {
        "authorization": %q
    }
}`, ""))

		err = ws.WriteMessage(websocket.TextMessage, msg)
		require.NoError(err)

		var data msgResponse
		err = ws.ReadJSON(&data)
		require.NoError(err)
		require.Equal("connection_ack", data.Type)

		err = ws.WriteMessage(websocket.TextMessage, operationsMsg)
		require.NoError(err)

		err = ws.ReadJSON(&data)
		require.NoError(err)
		require.Equal("ka", data.Type)

		err = ws.ReadJSON(&data)
		require.NoError(err)
		require.Equal("data", data.Type)
		require.NotNil(data.Payload)
		require.NotEmpty(data.Payload.Errors)
		require.Equal(data.Payload.Errors[0].Message, auth.UnauthorizedErr.Error())
	})

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			require := require.New(t)

			server := httptest.NewServer(srv)
			defer server.Close()

			wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/graphql"
			authData := genTokens(t, c.user, s)

			ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
			require.NoError(err)
			defer ws.Close()

			msg := []byte(fmt.Sprintf(`{
    "type": "connection_init",
    "payload": {
        "authorization": %q
    }
}`, authData.AccessToken))

			err = ws.WriteMessage(websocket.TextMessage, msg)
			require.NoError(err)

			var data msgResponse
			err = ws.ReadJSON(&data)
			require.Equal("connection_ack", data.Type)

			err = ws.WriteMessage(websocket.TextMessage, operationsMsg)
			require.NoError(err)

			err = ws.ReadJSON(&data)
			require.NoError(err)
			require.Equal("ka", data.Type)
		})
	}
}
