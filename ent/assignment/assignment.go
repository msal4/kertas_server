// Code generated by entc, DO NOT EDIT.

package assignment

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the assignment type in the database.
	Label = "assignment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldIsExam holds the string denoting the is_exam field in the database.
	FieldIsExam = "is_exam"
	// FieldDueDate holds the string denoting the due_date field in the database.
	FieldDueDate = "due_date"
	// FieldDuration holds the string denoting the duration field in the database.
	FieldDuration = "duration"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// EdgeClass holds the string denoting the class edge name in mutations.
	EdgeClass = "class"
	// EdgeSubmissions holds the string denoting the submissions edge name in mutations.
	EdgeSubmissions = "submissions"
	// EdgeGrades holds the string denoting the grades edge name in mutations.
	EdgeGrades = "grades"
	// Table holds the table name of the assignment in the database.
	Table = "assignments"
	// ClassTable is the table that holds the class relation/edge.
	ClassTable = "assignments"
	// ClassInverseTable is the table name for the Class entity.
	// It exists in this package in order to avoid circular dependency with the "class" package.
	ClassInverseTable = "classes"
	// ClassColumn is the table column denoting the class relation/edge.
	ClassColumn = "class_assignments"
	// SubmissionsTable is the table that holds the submissions relation/edge.
	SubmissionsTable = "assignment_submissions"
	// SubmissionsInverseTable is the table name for the AssignmentSubmission entity.
	// It exists in this package in order to avoid circular dependency with the "assignmentsubmission" package.
	SubmissionsInverseTable = "assignment_submissions"
	// SubmissionsColumn is the table column denoting the submissions relation/edge.
	SubmissionsColumn = "assignment_submissions"
	// GradesTable is the table that holds the grades relation/edge.
	GradesTable = "grades"
	// GradesInverseTable is the table name for the Grade entity.
	// It exists in this package in order to avoid circular dependency with the "grade" package.
	GradesInverseTable = "grades"
	// GradesColumn is the table column denoting the grades relation/edge.
	GradesColumn = "assignment_grades"
)

// Columns holds all SQL columns for assignment fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldDescription,
	FieldIsExam,
	FieldDueDate,
	FieldDuration,
	FieldDeletedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "assignments"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"class_assignments",
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
	// DefaultIsExam holds the default value on creation for the "is_exam" field.
	DefaultIsExam bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
