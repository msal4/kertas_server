// Code generated by entc, DO NOT EDIT.

package coursegrade

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the coursegrade type in the database.
	Label = "course_grade"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldCourse holds the string denoting the course field in the database.
	FieldCourse = "course"
	// FieldActivityFirst holds the string denoting the activity_first field in the database.
	FieldActivityFirst = "activity_first"
	// FieldActivitySecond holds the string denoting the activity_second field in the database.
	FieldActivitySecond = "activity_second"
	// FieldWrittenFirst holds the string denoting the written_first field in the database.
	FieldWrittenFirst = "written_first"
	// FieldWrittenSecond holds the string denoting the written_second field in the database.
	FieldWrittenSecond = "written_second"
	// FieldCourseFinal holds the string denoting the course_final field in the database.
	FieldCourseFinal = "course_final"
	// FieldYear holds the string denoting the year field in the database.
	FieldYear = "year"
	// EdgeStudent holds the string denoting the student edge name in mutations.
	EdgeStudent = "student"
	// EdgeClass holds the string denoting the class edge name in mutations.
	EdgeClass = "class"
	// Table holds the table name of the coursegrade in the database.
	Table = "course_grades"
	// StudentTable is the table that holds the student relation/edge.
	StudentTable = "course_grades"
	// StudentInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	StudentInverseTable = "users"
	// StudentColumn is the table column denoting the student relation/edge.
	StudentColumn = "user_course_grades"
	// ClassTable is the table that holds the class relation/edge.
	ClassTable = "course_grades"
	// ClassInverseTable is the table name for the Class entity.
	// It exists in this package in order to avoid circular dependency with the "class" package.
	ClassInverseTable = "classes"
	// ClassColumn is the table column denoting the class relation/edge.
	ClassColumn = "class_course_grades"
)

// Columns holds all SQL columns for coursegrade fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldCourse,
	FieldActivityFirst,
	FieldActivitySecond,
	FieldWrittenFirst,
	FieldWrittenSecond,
	FieldCourseFinal,
	FieldYear,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "course_grades"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"class_course_grades",
	"user_course_grades",
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
	// ActivityFirstValidator is a validator for the "activity_first" field. It is called by the builders before save.
	ActivityFirstValidator func(int) error
	// ActivitySecondValidator is a validator for the "activity_second" field. It is called by the builders before save.
	ActivitySecondValidator func(int) error
	// WrittenFirstValidator is a validator for the "written_first" field. It is called by the builders before save.
	WrittenFirstValidator func(int) error
	// WrittenSecondValidator is a validator for the "written_second" field. It is called by the builders before save.
	WrittenSecondValidator func(int) error
	// CourseFinalValidator is a validator for the "course_final" field. It is called by the builders before save.
	CourseFinalValidator func(int) error
	// YearValidator is a validator for the "year" field. It is called by the builders before save.
	YearValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Course defines the type for the "course" enum field.
type Course string

// Course values.
const (
	CourseFirst  Course = "FIRST"
	CourseSecond Course = "SECOND"
)

func (c Course) String() string {
	return string(c)
}

// CourseValidator is a validator for the "course" field enum values. It is called by the builders before save.
func CourseValidator(c Course) error {
	switch c {
	case CourseFirst, CourseSecond:
		return nil
	default:
		return fmt.Errorf("coursegrade: invalid enum value for course field: %q", c)
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (c Course) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(c.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (c *Course) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*c = Course(str)
	if err := CourseValidator(*c); err != nil {
		return fmt.Errorf("%s is not a valid Course", str)
	}
	return nil
}
