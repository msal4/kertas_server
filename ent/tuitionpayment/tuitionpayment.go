// Code generated by entc, DO NOT EDIT.

package tuitionpayment

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the tuitionpayment type in the database.
	Label = "tuition_payment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldYear holds the string denoting the year field in the database.
	FieldYear = "year"
	// FieldPaidAmount holds the string denoting the paid_amount field in the database.
	FieldPaidAmount = "paid_amount"
	// EdgeStudent holds the string denoting the student edge name in mutations.
	EdgeStudent = "student"
	// EdgeStage holds the string denoting the stage edge name in mutations.
	EdgeStage = "stage"
	// Table holds the table name of the tuitionpayment in the database.
	Table = "tuition_payments"
	// StudentTable is the table that holds the student relation/edge.
	StudentTable = "tuition_payments"
	// StudentInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	StudentInverseTable = "users"
	// StudentColumn is the table column denoting the student relation/edge.
	StudentColumn = "user_payments"
	// StageTable is the table that holds the stage relation/edge.
	StageTable = "tuition_payments"
	// StageInverseTable is the table name for the Stage entity.
	// It exists in this package in order to avoid circular dependency with the "stage" package.
	StageInverseTable = "stages"
	// StageColumn is the table column denoting the stage relation/edge.
	StageColumn = "stage_payments"
)

// Columns holds all SQL columns for tuitionpayment fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldYear,
	FieldPaidAmount,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "tuition_payments"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"stage_payments",
	"user_payments",
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
	// YearValidator is a validator for the "year" field. It is called by the builders before save.
	YearValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
