// Code generated by entc, DO NOT EDIT.

package schedule

const (
	// Label holds the string label denoting the schedule type in the database.
	Label = "schedule"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldWeekday holds the string denoting the weekday field in the database.
	FieldWeekday = "weekday"
	// FieldStartsAt holds the string denoting the starts_at field in the database.
	FieldStartsAt = "starts_at"
	// FieldDuration holds the string denoting the duration field in the database.
	FieldDuration = "duration"
	// EdgeClass holds the string denoting the class edge name in mutations.
	EdgeClass = "class"
	// Table holds the table name of the schedule in the database.
	Table = "schedules"
	// ClassTable is the table that holds the class relation/edge.
	ClassTable = "schedules"
	// ClassInverseTable is the table name for the Class entity.
	// It exists in this package in order to avoid circular dependency with the "class" package.
	ClassInverseTable = "classes"
	// ClassColumn is the table column denoting the class relation/edge.
	ClassColumn = "class_schedules"
)

// Columns holds all SQL columns for schedule fields.
var Columns = []string{
	FieldID,
	FieldWeekday,
	FieldStartsAt,
	FieldDuration,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "schedules"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"class_schedules",
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
	// WeekdayValidator is a validator for the "weekday" field. It is called by the builders before save.
	WeekdayValidator func(uint8) error
	// DefaultDuration holds the default value on creation for the "duration" field.
	DefaultDuration int
)
