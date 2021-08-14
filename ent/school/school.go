// Code generated by entc, DO NOT EDIT.

package school

import (
	"fmt"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
	"github.com/msal4/hassah_school_server/ent/schema"
)

const (
	// Label holds the string label denoting the school type in the database.
	Label = "school"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldImage holds the string denoting the image field in the database.
	FieldImage = "image"
	// FieldDirectory holds the string denoting the directory field in the database.
	FieldDirectory = "directory"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// EdgeStages holds the string denoting the stages edge name in mutations.
	EdgeStages = "stages"
	// Table holds the table name of the school in the database.
	Table = "schools"
	// UsersTable is the table that holds the users relation/edge.
	UsersTable = "users"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
	// UsersColumn is the table column denoting the users relation/edge.
	UsersColumn = "school_users"
	// StagesTable is the table that holds the stages relation/edge.
	StagesTable = "stages"
	// StagesInverseTable is the table name for the Stage entity.
	// It exists in this package in order to avoid circular dependency with the "stage" package.
	StagesInverseTable = "stages"
	// StagesColumn is the table column denoting the stages relation/edge.
	StagesColumn = "school_stages"
)

// Columns holds all SQL columns for school fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldImage,
	FieldDirectory,
	FieldStatus,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
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
	// ImageValidator is a validator for the "image" field. It is called by the builders before save.
	ImageValidator func(string) error
	// DirectoryValidator is a validator for the "directory" field. It is called by the builders before save.
	DirectoryValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

const DefaultStatus schema.Status = "ACTIVE"

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s schema.Status) error {
	switch s {
	case "DISABLED", "ACTIVE":
		return nil
	default:
		return fmt.Errorf("school: invalid enum value for status field: %q", s)
	}
}

var (
	// schema.Status must implement graphql.Marshaler.
	_ graphql.Marshaler = schema.Status("")
	// schema.Status must implement graphql.Unmarshaler.
	_ graphql.Unmarshaler = (*schema.Status)(nil)
)
