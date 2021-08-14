// Code generated by entc, DO NOT EDIT.

package group

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
	// Label holds the string label denoting the group type in the database.
	Label = "group"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldGroupType holds the string denoting the group_type field in the database.
	FieldGroupType = "group_type"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// EdgeClass holds the string denoting the class edge name in mutations.
	EdgeClass = "class"
	// EdgeMessages holds the string denoting the messages edge name in mutations.
	EdgeMessages = "messages"
	// Table holds the table name of the group in the database.
	Table = "groups"
	// ClassTable is the table that holds the class relation/edge.
	ClassTable = "groups"
	// ClassInverseTable is the table name for the Class entity.
	// It exists in this package in order to avoid circular dependency with the "class" package.
	ClassInverseTable = "classes"
	// ClassColumn is the table column denoting the class relation/edge.
	ClassColumn = "class_group"
	// MessagesTable is the table that holds the messages relation/edge.
	MessagesTable = "messages"
	// MessagesInverseTable is the table name for the Message entity.
	// It exists in this package in order to avoid circular dependency with the "message" package.
	MessagesInverseTable = "messages"
	// MessagesColumn is the table column denoting the messages relation/edge.
	MessagesColumn = "group_messages"
)

// Columns holds all SQL columns for group fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldGroupType,
	FieldStatus,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "groups"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"class_group",
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
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// GroupType defines the type for the "group_type" enum field.
type GroupType string

// GroupTypeSHARED is the default value of the GroupType enum.
const DefaultGroupType = GroupTypeSHARED

// GroupType values.
const (
	GroupTypePRIVATE GroupType = "PRIVATE"
	GroupTypeSHARED  GroupType = "SHARED"
)

func (gt GroupType) String() string {
	return string(gt)
}

// GroupTypeValidator is a validator for the "group_type" field enum values. It is called by the builders before save.
func GroupTypeValidator(gt GroupType) error {
	switch gt {
	case GroupTypePRIVATE, GroupTypeSHARED:
		return nil
	default:
		return fmt.Errorf("group: invalid enum value for group_type field: %q", gt)
	}
}

const DefaultStatus schema.Status = "ACTIVE"

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s schema.Status) error {
	switch s {
	case "DISABLED", "ACTIVE":
		return nil
	default:
		return fmt.Errorf("group: invalid enum value for status field: %q", s)
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (gt GroupType) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(gt.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (gt *GroupType) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*gt = GroupType(str)
	if err := GroupTypeValidator(*gt); err != nil {
		return fmt.Errorf("%s is not a valid GroupType", str)
	}
	return nil
}

var (
	// schema.Status must implement graphql.Marshaler.
	_ graphql.Marshaler = schema.Status("")
	// schema.Status must implement graphql.Unmarshaler.
	_ graphql.Unmarshaler = (*schema.Status)(nil)
)
