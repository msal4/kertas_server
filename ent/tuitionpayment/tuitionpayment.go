// Code generated by entc, DO NOT EDIT.

package tuitionpayment

import (
	"time"
)

const (
	// Label holds the string label denoting the tuitionpayment type in the database.
	Label = "tuition_payment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
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
	FieldCreateTime,
	FieldUpdateTime,
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
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
)
