// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/msal4/hassah_school_server/ent/class"
	"github.com/msal4/hassah_school_server/ent/schedule"
)

// Schedule is the model entity for the Schedule schema.
type Schedule struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Weekday holds the value of the "weekday" field.
	Weekday time.Weekday `json:"weekday,omitempty"`
	// StartsAt holds the value of the "starts_at" field.
	StartsAt time.Time `json:"starts_at,omitempty"`
	// Duration holds the value of the "duration" field.
	Duration time.Duration `json:"duration,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ScheduleQuery when eager-loading is set.
	Edges           ScheduleEdges `json:"edges"`
	class_schedules *uuid.UUID
}

// ScheduleEdges holds the relations/edges for other nodes in the graph.
type ScheduleEdges struct {
	// Class holds the value of the class edge.
	Class *Class `json:"class,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ClassOrErr returns the Class value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ScheduleEdges) ClassOrErr() (*Class, error) {
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

// scanValues returns the types for scanning values from sql.Rows.
func (*Schedule) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case schedule.FieldWeekday, schedule.FieldDuration:
			values[i] = new(sql.NullInt64)
		case schedule.FieldStartsAt:
			values[i] = new(sql.NullTime)
		case schedule.FieldID:
			values[i] = new(uuid.UUID)
		case schedule.ForeignKeys[0]: // class_schedules
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Schedule", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Schedule fields.
func (s *Schedule) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case schedule.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				s.ID = *value
			}
		case schedule.FieldWeekday:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field weekday", values[i])
			} else if value.Valid {
				s.Weekday = time.Weekday(value.Int64)
			}
		case schedule.FieldStartsAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field starts_at", values[i])
			} else if value.Valid {
				s.StartsAt = value.Time
			}
		case schedule.FieldDuration:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field duration", values[i])
			} else if value.Valid {
				s.Duration = time.Duration(value.Int64)
			}
		case schedule.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field class_schedules", values[i])
			} else if value.Valid {
				s.class_schedules = new(uuid.UUID)
				*s.class_schedules = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryClass queries the "class" edge of the Schedule entity.
func (s *Schedule) QueryClass() *ClassQuery {
	return (&ScheduleClient{config: s.config}).QueryClass(s)
}

// Update returns a builder for updating this Schedule.
// Note that you need to call Schedule.Unwrap() before calling this method if this Schedule
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Schedule) Update() *ScheduleUpdateOne {
	return (&ScheduleClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Schedule entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Schedule) Unwrap() *Schedule {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Schedule is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Schedule) String() string {
	var builder strings.Builder
	builder.WriteString("Schedule(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", weekday=")
	builder.WriteString(fmt.Sprintf("%v", s.Weekday))
	builder.WriteString(", starts_at=")
	builder.WriteString(s.StartsAt.Format(time.ANSIC))
	builder.WriteString(", duration=")
	builder.WriteString(fmt.Sprintf("%v", s.Duration))
	builder.WriteByte(')')
	return builder.String()
}

// Schedules is a parsable slice of Schedule.
type Schedules []*Schedule

func (s Schedules) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
