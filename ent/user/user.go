// Code generated by entc, DO NOT EDIT.

package user

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
	"github.com/msal4/hassah_school_server/ent/schema"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldImage holds the string denoting the image field in the database.
	FieldImage = "image"
	// FieldDirectory holds the string denoting the directory field in the database.
	FieldDirectory = "directory"
	// FieldTokenVersion holds the string denoting the token_version field in the database.
	FieldTokenVersion = "token_version"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// EdgeStage holds the string denoting the stage edge name in mutations.
	EdgeStage = "stage"
	// EdgeSchool holds the string denoting the school edge name in mutations.
	EdgeSchool = "school"
	// EdgeClasses holds the string denoting the classes edge name in mutations.
	EdgeClasses = "classes"
	// EdgeMessages holds the string denoting the messages edge name in mutations.
	EdgeMessages = "messages"
	// EdgeSubmissions holds the string denoting the submissions edge name in mutations.
	EdgeSubmissions = "submissions"
	// EdgeAttendances holds the string denoting the attendances edge name in mutations.
	EdgeAttendances = "attendances"
	// EdgePayments holds the string denoting the payments edge name in mutations.
	EdgePayments = "payments"
	// EdgeGrades holds the string denoting the grades edge name in mutations.
	EdgeGrades = "grades"
	// Table holds the table name of the user in the database.
	Table = "users"
	// StageTable is the table that holds the stage relation/edge.
	StageTable = "users"
	// StageInverseTable is the table name for the Stage entity.
	// It exists in this package in order to avoid circular dependency with the "stage" package.
	StageInverseTable = "stages"
	// StageColumn is the table column denoting the stage relation/edge.
	StageColumn = "stage_students"
	// SchoolTable is the table that holds the school relation/edge.
	SchoolTable = "users"
	// SchoolInverseTable is the table name for the School entity.
	// It exists in this package in order to avoid circular dependency with the "school" package.
	SchoolInverseTable = "schools"
	// SchoolColumn is the table column denoting the school relation/edge.
	SchoolColumn = "school_users"
	// ClassesTable is the table that holds the classes relation/edge.
	ClassesTable = "classes"
	// ClassesInverseTable is the table name for the Class entity.
	// It exists in this package in order to avoid circular dependency with the "class" package.
	ClassesInverseTable = "classes"
	// ClassesColumn is the table column denoting the classes relation/edge.
	ClassesColumn = "user_classes"
	// MessagesTable is the table that holds the messages relation/edge.
	MessagesTable = "messages"
	// MessagesInverseTable is the table name for the Message entity.
	// It exists in this package in order to avoid circular dependency with the "message" package.
	MessagesInverseTable = "messages"
	// MessagesColumn is the table column denoting the messages relation/edge.
	MessagesColumn = "user_messages"
	// SubmissionsTable is the table that holds the submissions relation/edge.
	SubmissionsTable = "assignment_submissions"
	// SubmissionsInverseTable is the table name for the AssignmentSubmission entity.
	// It exists in this package in order to avoid circular dependency with the "assignmentsubmission" package.
	SubmissionsInverseTable = "assignment_submissions"
	// SubmissionsColumn is the table column denoting the submissions relation/edge.
	SubmissionsColumn = "user_submissions"
	// AttendancesTable is the table that holds the attendances relation/edge.
	AttendancesTable = "attendances"
	// AttendancesInverseTable is the table name for the Attendance entity.
	// It exists in this package in order to avoid circular dependency with the "attendance" package.
	AttendancesInverseTable = "attendances"
	// AttendancesColumn is the table column denoting the attendances relation/edge.
	AttendancesColumn = "user_attendances"
	// PaymentsTable is the table that holds the payments relation/edge.
	PaymentsTable = "tuition_payments"
	// PaymentsInverseTable is the table name for the TuitionPayment entity.
	// It exists in this package in order to avoid circular dependency with the "tuitionpayment" package.
	PaymentsInverseTable = "tuition_payments"
	// PaymentsColumn is the table column denoting the payments relation/edge.
	PaymentsColumn = "user_payments"
	// GradesTable is the table that holds the grades relation/edge.
	GradesTable = "grades"
	// GradesInverseTable is the table name for the Grade entity.
	// It exists in this package in order to avoid circular dependency with the "grade" package.
	GradesInverseTable = "grades"
	// GradesColumn is the table column denoting the grades relation/edge.
	GradesColumn = "user_grades"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldUsername,
	FieldPassword,
	FieldPhone,
	FieldImage,
	FieldDirectory,
	FieldTokenVersion,
	FieldRole,
	FieldStatus,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "users"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"school_users",
	"stage_students",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func(string) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
	// PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	PhoneValidator func(string) error
	// DirectoryValidator is a validator for the "directory" field. It is called by the builders before save.
	DirectoryValidator func(string) error
	// DefaultTokenVersion holds the default value on creation for the "token_version" field.
	DefaultTokenVersion int
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Role defines the type for the "role" enum field.
type Role string

// RoleStudent is the default value of the Role enum.
const DefaultRole = RoleStudent

// Role values.
const (
	RoleSuperAdmin  Role = "SUPER_ADMIN"
	RoleSchoolAdmin Role = "SCHOOL_ADMIN"
	RoleTeacher     Role = "TEACHER"
	RoleStudent     Role = "STUDENT"
)

func (r Role) String() string {
	return string(r)
}

// RoleValidator is a validator for the "role" field enum values. It is called by the builders before save.
func RoleValidator(r Role) error {
	switch r {
	case RoleSuperAdmin, RoleSchoolAdmin, RoleTeacher, RoleStudent:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for role field: %q", r)
	}
}

const DefaultStatus schema.Status = "ACTIVE"

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s schema.Status) error {
	switch s {
	case "DISABLED", "ACTIVE":
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for status field: %q", s)
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (r Role) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(r.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (r *Role) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*r = Role(str)
	if err := RoleValidator(*r); err != nil {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

var (
	// schema.Status must implement graphql.Marshaler.
	_ graphql.Marshaler = schema.Status("")
	// schema.Status must implement graphql.Unmarshaler.
	_ graphql.Unmarshaler = (*schema.Status)(nil)
)
