// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
	"github.com/msal4/hassah_school_server/ent/user"
)

type AddClassInput struct {
	Name      string    `json:"name"`
	Active    bool      `json:"active"`
	TeacherID uuid.UUID `json:"teacherID"`
	StageID   uuid.UUID `json:"stageID"`
}

type AddGroupInput struct {
	Name   string    `json:"name"`
	Active bool      `json:"active"`
	UserID uuid.UUID `json:"userID"`
}

type AddSchoolInput struct {
	Name   string         `json:"name"`
	Image  graphql.Upload `json:"image"`
	Active bool           `json:"active"`
}

type AddStageInput struct {
	Name          string    `json:"name"`
	Active        bool      `json:"active"`
	TuitionAmount int       `json:"tuitionAmount"`
	SchoolID      uuid.UUID `json:"schoolID"`
}

type AddUserInput struct {
	Name     string          `json:"name"`
	Username string          `json:"username"`
	Password string          `json:"password"`
	Phone    string          `json:"phone"`
	Image    *graphql.Upload `json:"image"`
	Role     user.Role       `json:"role"`
	Active   bool            `json:"active"`
	SchoolID *uuid.UUID      `json:"schoolID"`
	StageID  *uuid.UUID      `json:"stageID"`
}

type AuthData struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type PostMessageInput struct {
	GroupID    uuid.UUID       `json:"groupID"`
	Content    string          `json:"content"`
	Attachment *graphql.Upload `json:"attachment"`
}

type UpdateClassInput struct {
	Name      *string    `json:"name"`
	Active    *bool      `json:"active"`
	TeacherID *uuid.UUID `json:"teacherID"`
}

type UpdateGroupInput struct {
	Name   *string `json:"name"`
	Active *bool   `json:"active"`
}

type UpdateSchoolInput struct {
	Name   *string         `json:"name"`
	Image  *graphql.Upload `json:"image"`
	Active *bool           `json:"active"`
}

type UpdateStageInput struct {
	Name          *string `json:"name"`
	Active        *bool   `json:"active"`
	TuitionAmount *int    `json:"tuitionAmount"`
}

type UpdateUserInput struct {
	Name     *string         `json:"name"`
	Username *string         `json:"username"`
	Password *string         `json:"password"`
	Phone    *string         `json:"phone"`
	Image    *graphql.Upload `json:"image"`
	Active   *bool           `json:"active"`
	StageID  *uuid.UUID      `json:"stageID"`
}
