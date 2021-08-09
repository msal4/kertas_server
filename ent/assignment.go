// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/msal4/hassah_school_server/ent/assignment"
	"github.com/msal4/hassah_school_server/ent/class"
)

// Assignment is the model entity for the Assignment schema.
type Assignment struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// IsExam holds the value of the "is_exam" field.
	IsExam bool `json:"is_exam,omitempty"`
	// DueDate holds the value of the "due_date" field.
	DueDate time.Time `json:"due_date,omitempty"`
	// Duration holds the value of the "duration" field.
	Duration int `json:"duration,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AssignmentQuery when eager-loading is set.
	Edges             AssignmentEdges `json:"edges"`
	class_assignments *uuid.UUID
}

// AssignmentEdges holds the relations/edges for other nodes in the graph.
type AssignmentEdges struct {
	// Class holds the value of the class edge.
	Class *Class `json:"class,omitempty"`
	// Submissions holds the value of the submissions edge.
	Submissions []*AssignmentSubmission `json:"submissions,omitempty"`
	// Grades holds the value of the grades edge.
	Grades []*Grade `json:"grades,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ClassOrErr returns the Class value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AssignmentEdges) ClassOrErr() (*Class, error) {
	if e.loadedTypes[0] {
		if e.Class == nil {
			// The edge class was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: class.Label}
		}
		return e.Class, nil
	}
	return nil, &NotLoadedError{edge: "class"}
}

// SubmissionsOrErr returns the Submissions value or an error if the edge
// was not loaded in eager-loading.
func (e AssignmentEdges) SubmissionsOrErr() ([]*AssignmentSubmission, error) {
	if e.loadedTypes[1] {
		return e.Submissions, nil
	}
	return nil, &NotLoadedError{edge: "submissions"}
}

// GradesOrErr returns the Grades value or an error if the edge
// was not loaded in eager-loading.
func (e AssignmentEdges) GradesOrErr() ([]*Grade, error) {
	if e.loadedTypes[2] {
		return e.Grades, nil
	}
	return nil, &NotLoadedError{edge: "grades"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Assignment) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case assignment.FieldIsExam:
			values[i] = new(sql.NullBool)
		case assignment.FieldDuration:
			values[i] = new(sql.NullInt64)
		case assignment.FieldName, assignment.FieldDescription:
			values[i] = new(sql.NullString)
		case assignment.FieldCreatedAt, assignment.FieldUpdatedAt, assignment.FieldDueDate, assignment.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case assignment.FieldID:
			values[i] = new(uuid.UUID)
		case assignment.ForeignKeys[0]: // class_assignments
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Assignment", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Assignment fields.
func (a *Assignment) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case assignment.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				a.ID = *value
			}
		case assignment.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case assignment.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.Time
			}
		case assignment.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				a.Name = value.String
			}
		case assignment.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				a.Description = value.String
			}
		case assignment.FieldIsExam:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_exam", values[i])
			} else if value.Valid {
				a.IsExam = value.Bool
			}
		case assignment.FieldDueDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field due_date", values[i])
			} else if value.Valid {
				a.DueDate = value.Time
			}
		case assignment.FieldDuration:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field duration", values[i])
			} else if value.Valid {
				a.Duration = int(value.Int64)
			}
		case assignment.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				a.DeletedAt = new(time.Time)
				*a.DeletedAt = value.Time
			}
		case assignment.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field class_assignments", values[i])
			} else if value.Valid {
				a.class_assignments = new(uuid.UUID)
				*a.class_assignments = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryClass queries the "class" edge of the Assignment entity.
func (a *Assignment) QueryClass() *ClassQuery {
	return (&AssignmentClient{config: a.config}).QueryClass(a)
}

// QuerySubmissions queries the "submissions" edge of the Assignment entity.
func (a *Assignment) QuerySubmissions() *AssignmentSubmissionQuery {
	return (&AssignmentClient{config: a.config}).QuerySubmissions(a)
}

// QueryGrades queries the "grades" edge of the Assignment entity.
func (a *Assignment) QueryGrades() *GradeQuery {
	return (&AssignmentClient{config: a.config}).QueryGrades(a)
}

// Update returns a builder for updating this Assignment.
// Note that you need to call Assignment.Unwrap() before calling this method if this Assignment
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Assignment) Update() *AssignmentUpdateOne {
	return (&AssignmentClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Assignment entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Assignment) Unwrap() *Assignment {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Assignment is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Assignment) String() string {
	var builder strings.Builder
	builder.WriteString("Assignment(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", name=")
	builder.WriteString(a.Name)
	builder.WriteString(", description=")
	builder.WriteString(a.Description)
	builder.WriteString(", is_exam=")
	builder.WriteString(fmt.Sprintf("%v", a.IsExam))
	builder.WriteString(", due_date=")
	builder.WriteString(a.DueDate.Format(time.ANSIC))
	builder.WriteString(", duration=")
	builder.WriteString(fmt.Sprintf("%v", a.Duration))
	if v := a.DeletedAt; v != nil {
		builder.WriteString(", deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Assignments is a parsable slice of Assignment.
type Assignments []*Assignment

func (a Assignments) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
