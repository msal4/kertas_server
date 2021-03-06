// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/msal4/hassah_school_server/ent/class"
	"github.com/msal4/hassah_school_server/ent/coursegrade"
	"github.com/msal4/hassah_school_server/ent/user"
)

// CourseGrade is the model entity for the CourseGrade schema.
type CourseGrade struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Course holds the value of the "course" field.
	Course coursegrade.Course `json:"course,omitempty"`
	// ActivityFirst holds the value of the "activity_first" field.
	ActivityFirst *int `json:"activity_first,omitempty"`
	// ActivitySecond holds the value of the "activity_second" field.
	ActivitySecond *int `json:"activity_second,omitempty"`
	// WrittenFirst holds the value of the "written_first" field.
	WrittenFirst *int `json:"written_first,omitempty"`
	// WrittenSecond holds the value of the "written_second" field.
	WrittenSecond *int `json:"written_second,omitempty"`
	// CourseFinal holds the value of the "course_final" field.
	CourseFinal *int `json:"course_final,omitempty"`
	// Year holds the value of the "year" field.
	Year string `json:"year,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CourseGradeQuery when eager-loading is set.
	Edges               CourseGradeEdges `json:"edges"`
	class_course_grades *uuid.UUID
	user_course_grades  *uuid.UUID
}

// CourseGradeEdges holds the relations/edges for other nodes in the graph.
type CourseGradeEdges struct {
	// Student holds the value of the student edge.
	Student *User `json:"student,omitempty"`
	// Class holds the value of the class edge.
	Class *Class `json:"class,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// StudentOrErr returns the Student value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CourseGradeEdges) StudentOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Student == nil {
			// The edge student was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Student, nil
	}
	return nil, &NotLoadedError{edge: "student"}
}

// ClassOrErr returns the Class value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CourseGradeEdges) ClassOrErr() (*Class, error) {
	if e.loadedTypes[1] {
		if e.Class == nil {
			// The edge class was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: class.Label}
		}
		return e.Class, nil
	}
	return nil, &NotLoadedError{edge: "class"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CourseGrade) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case coursegrade.FieldActivityFirst, coursegrade.FieldActivitySecond, coursegrade.FieldWrittenFirst, coursegrade.FieldWrittenSecond, coursegrade.FieldCourseFinal:
			values[i] = new(sql.NullInt64)
		case coursegrade.FieldCourse, coursegrade.FieldYear:
			values[i] = new(sql.NullString)
		case coursegrade.FieldCreatedAt, coursegrade.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case coursegrade.FieldID:
			values[i] = new(uuid.UUID)
		case coursegrade.ForeignKeys[0]: // class_course_grades
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case coursegrade.ForeignKeys[1]: // user_course_grades
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type CourseGrade", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CourseGrade fields.
func (cg *CourseGrade) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case coursegrade.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				cg.ID = *value
			}
		case coursegrade.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				cg.CreatedAt = value.Time
			}
		case coursegrade.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				cg.UpdatedAt = value.Time
			}
		case coursegrade.FieldCourse:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field course", values[i])
			} else if value.Valid {
				cg.Course = coursegrade.Course(value.String)
			}
		case coursegrade.FieldActivityFirst:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field activity_first", values[i])
			} else if value.Valid {
				cg.ActivityFirst = new(int)
				*cg.ActivityFirst = int(value.Int64)
			}
		case coursegrade.FieldActivitySecond:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field activity_second", values[i])
			} else if value.Valid {
				cg.ActivitySecond = new(int)
				*cg.ActivitySecond = int(value.Int64)
			}
		case coursegrade.FieldWrittenFirst:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field written_first", values[i])
			} else if value.Valid {
				cg.WrittenFirst = new(int)
				*cg.WrittenFirst = int(value.Int64)
			}
		case coursegrade.FieldWrittenSecond:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field written_second", values[i])
			} else if value.Valid {
				cg.WrittenSecond = new(int)
				*cg.WrittenSecond = int(value.Int64)
			}
		case coursegrade.FieldCourseFinal:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field course_final", values[i])
			} else if value.Valid {
				cg.CourseFinal = new(int)
				*cg.CourseFinal = int(value.Int64)
			}
		case coursegrade.FieldYear:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field year", values[i])
			} else if value.Valid {
				cg.Year = value.String
			}
		case coursegrade.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field class_course_grades", values[i])
			} else if value.Valid {
				cg.class_course_grades = new(uuid.UUID)
				*cg.class_course_grades = *value.S.(*uuid.UUID)
			}
		case coursegrade.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_course_grades", values[i])
			} else if value.Valid {
				cg.user_course_grades = new(uuid.UUID)
				*cg.user_course_grades = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryStudent queries the "student" edge of the CourseGrade entity.
func (cg *CourseGrade) QueryStudent() *UserQuery {
	return (&CourseGradeClient{config: cg.config}).QueryStudent(cg)
}

// QueryClass queries the "class" edge of the CourseGrade entity.
func (cg *CourseGrade) QueryClass() *ClassQuery {
	return (&CourseGradeClient{config: cg.config}).QueryClass(cg)
}

// Update returns a builder for updating this CourseGrade.
// Note that you need to call CourseGrade.Unwrap() before calling this method if this CourseGrade
// was returned from a transaction, and the transaction was committed or rolled back.
func (cg *CourseGrade) Update() *CourseGradeUpdateOne {
	return (&CourseGradeClient{config: cg.config}).UpdateOne(cg)
}

// Unwrap unwraps the CourseGrade entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cg *CourseGrade) Unwrap() *CourseGrade {
	tx, ok := cg.config.driver.(*txDriver)
	if !ok {
		panic("ent: CourseGrade is not a transactional entity")
	}
	cg.config.driver = tx.drv
	return cg
}

// String implements the fmt.Stringer.
func (cg *CourseGrade) String() string {
	var builder strings.Builder
	builder.WriteString("CourseGrade(")
	builder.WriteString(fmt.Sprintf("id=%v", cg.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(cg.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(cg.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", course=")
	builder.WriteString(fmt.Sprintf("%v", cg.Course))
	if v := cg.ActivityFirst; v != nil {
		builder.WriteString(", activity_first=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	if v := cg.ActivitySecond; v != nil {
		builder.WriteString(", activity_second=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	if v := cg.WrittenFirst; v != nil {
		builder.WriteString(", written_first=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	if v := cg.WrittenSecond; v != nil {
		builder.WriteString(", written_second=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	if v := cg.CourseFinal; v != nil {
		builder.WriteString(", course_final=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", year=")
	builder.WriteString(cg.Year)
	builder.WriteByte(')')
	return builder.String()
}

// CourseGrades is a parsable slice of CourseGrade.
type CourseGrades []*CourseGrade

func (cg CourseGrades) config(cfg config) {
	for _i := range cg {
		cg[_i].config = cfg
	}
}
