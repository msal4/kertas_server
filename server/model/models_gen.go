// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
	"github.com/msal4/hassah_school_server/ent/attendance"
	"github.com/msal4/hassah_school_server/ent/coursegrade"
	"github.com/msal4/hassah_school_server/ent/user"
)

type AddAssignmentInput struct {
	Name        string          `json:"name"`
	Description *string         `json:"description"`
	File        *graphql.Upload `json:"file"`
	ClassID     uuid.UUID       `json:"classID"`
	DueDate     time.Time       `json:"dueDate"`
	IsExam      bool            `json:"isExam"`
	Duration    *time.Duration  `json:"duration"`
}

type AddAssignmentSubmissionInput struct {
	AssignmentID uuid.UUID         `json:"assignmentID"`
	Files        []*graphql.Upload `json:"files"`
	SubmittedAt  *time.Time        `json:"submittedAt"`
}

type AddAttendanceInput struct {
	Date      time.Time        `json:"date"`
	State     attendance.State `json:"state"`
	ClassID   uuid.UUID        `json:"classID"`
	StudentID uuid.UUID        `json:"studentID"`
}

type AddClassInput struct {
	Name      string    `json:"name"`
	Active    bool      `json:"active"`
	TeacherID uuid.UUID `json:"teacherID"`
	StageID   uuid.UUID `json:"stageID"`
}

type AddCourseGradeInput struct {
	StudentID      uuid.UUID          `json:"studentID"`
	StageID        uuid.UUID          `json:"stageID"`
	ClassID        uuid.UUID          `json:"classID"`
	Course         coursegrade.Course `json:"course"`
	ActivityFirst  *int               `json:"activityFirst"`
	ActivitySecond *int               `json:"activitySecond"`
	WrittenFirst   *int               `json:"writtenFirst"`
	WrittenSecond  *int               `json:"writtenSecond"`
	CourseFinal    *int               `json:"courseFinal"`
	Year           string             `json:"year"`
}

type AddGroupInput struct {
	Name   string    `json:"name"`
	Active bool      `json:"active"`
	UserID uuid.UUID `json:"userID"`
}

type AddNotificationInput struct {
	Title   string          `json:"title"`
	Body    string          `json:"body"`
	Image   *graphql.Upload `json:"image"`
	Route   string          `json:"route"`
	Color   string          `json:"color"`
	StageID uuid.UUID       `json:"stageID"`
}

type AddScheduleInput struct {
	Weekday  time.Weekday  `json:"weekday"`
	Duration time.Duration `json:"duration"`
	StartsAt time.Time     `json:"startsAt"`
	ClassID  uuid.UUID     `json:"classID"`
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

type AddTuitionPaymentInput struct {
	StageID    uuid.UUID `json:"stageID"`
	StudentID  uuid.UUID `json:"studentID"`
	Year       string    `json:"year"`
	PaidAmount int       `json:"paidAmount"`
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
	Username  string  `json:"username"`
	Password  string  `json:"password"`
	PushToken *string `json:"pushToken"`
}

type PostMessageInput struct {
	GroupID    uuid.UUID       `json:"groupID"`
	Content    string          `json:"content"`
	Attachment *graphql.Upload `json:"attachment"`
}

type UpdateAssignmentInput struct {
	Name        *string         `json:"name"`
	Description *string         `json:"description"`
	File        *graphql.Upload `json:"file"`
	DueDate     *time.Time      `json:"dueDate"`
	Duration    *time.Duration  `json:"duration"`
}

type UpdateAssignmentSubmissionInput struct {
	Files       []*graphql.Upload `json:"files"`
	SubmittedAt *time.Time        `json:"submittedAt"`
}

type UpdateAttendanceInput struct {
	Date  *time.Time        `json:"date"`
	State *attendance.State `json:"state"`
}

type UpdateClassInput struct {
	Name      *string    `json:"name"`
	Active    *bool      `json:"active"`
	TeacherID *uuid.UUID `json:"teacherID"`
}

type UpdateCourseGradeInput struct {
	ActivityFirst  *int `json:"activityFirst"`
	ActivitySecond *int `json:"activitySecond"`
	WrittenFirst   *int `json:"writtenFirst"`
	WrittenSecond  *int `json:"writtenSecond"`
	CourseFinal    *int `json:"courseFinal"`
}

type UpdateGroupInput struct {
	Name   *string `json:"name"`
	Active *bool   `json:"active"`
}

type UpdateScheduleInput struct {
	Weekday  *time.Weekday  `json:"weekday"`
	Duration *time.Duration `json:"duration"`
	StartsAt *time.Time     `json:"startsAt"`
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

type UpdateTuitionPaymentInput struct {
	Year       *string `json:"year"`
	PaidAmount *int    `json:"paidAmount"`
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
