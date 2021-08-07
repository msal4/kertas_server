// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/msal4/hassah_school_server/ent/schema"
	"github.com/msal4/hassah_school_server/ent/school"
)

// School is the model entity for the School schema.
type School struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Image holds the value of the "image" field.
	Image string `json:"image,omitempty"`
	// Status holds the value of the "status" field.
	Status schema.Status `json:"status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SchoolQuery when eager-loading is set.
	Edges SchoolEdges `json:"edges"`
}

// SchoolEdges holds the relations/edges for other nodes in the graph.
type SchoolEdges struct {
	// Users holds the value of the users edge.
	Users []*User `json:"users,omitempty"`
	// Stages holds the value of the stages edge.
	Stages []*Stage `json:"stages,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e SchoolEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// StagesOrErr returns the Stages value or an error if the edge
// was not loaded in eager-loading.
func (e SchoolEdges) StagesOrErr() ([]*Stage, error) {
	if e.loadedTypes[1] {
		return e.Stages, nil
	}
	return nil, &NotLoadedError{edge: "stages"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*School) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case school.FieldID:
			values[i] = new(sql.NullInt64)
		case school.FieldName, school.FieldImage, school.FieldStatus:
			values[i] = new(sql.NullString)
		case school.FieldCreateTime, school.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type School", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the School fields.
func (s *School) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case school.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case school.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				s.CreateTime = value.Time
			}
		case school.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				s.UpdateTime = value.Time
			}
		case school.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case school.FieldImage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image", values[i])
			} else if value.Valid {
				s.Image = value.String
			}
		case school.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				s.Status = schema.Status(value.String)
			}
		}
	}
	return nil
}

// QueryUsers queries the "users" edge of the School entity.
func (s *School) QueryUsers() *UserQuery {
	return (&SchoolClient{config: s.config}).QueryUsers(s)
}

// QueryStages queries the "stages" edge of the School entity.
func (s *School) QueryStages() *StageQuery {
	return (&SchoolClient{config: s.config}).QueryStages(s)
}

// Update returns a builder for updating this School.
// Note that you need to call School.Unwrap() before calling this method if this School
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *School) Update() *SchoolUpdateOne {
	return (&SchoolClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the School entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *School) Unwrap() *School {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: School is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *School) String() string {
	var builder strings.Builder
	builder.WriteString("School(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(s.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(s.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", name=")
	builder.WriteString(s.Name)
	builder.WriteString(", image=")
	builder.WriteString(s.Image)
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", s.Status))
	builder.WriteByte(')')
	return builder.String()
}

// Schools is a parsable slice of School.
type Schools []*School

func (s Schools) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
