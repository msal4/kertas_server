// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/msal4/hassah_school_server/ent/schema"
	"github.com/msal4/hassah_school_server/ent/school"
	"github.com/msal4/hassah_school_server/ent/stage"
	"github.com/msal4/hassah_school_server/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// Image holds the value of the "image" field.
	Image string `json:"image,omitempty"`
	// TokenVersion holds the value of the "token_version" field.
	TokenVersion int `json:"token_version,omitempty"`
	// Role holds the value of the "role" field.
	Role user.Role `json:"role,omitempty"`
	// Status holds the value of the "status" field.
	Status schema.Status `json:"status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges          UserEdges `json:"edges"`
	school_users   *int
	stage_students *int
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Stage holds the value of the stage edge.
	Stage *Stage `json:"stage,omitempty"`
	// School holds the value of the school edge.
	School *School `json:"school,omitempty"`
	// Classes holds the value of the classes edge.
	Classes []*Class `json:"classes,omitempty"`
	// Messages holds the value of the messages edge.
	Messages []*Message `json:"messages,omitempty"`
	// Submissions holds the value of the submissions edge.
	Submissions []*AssignmentSubmission `json:"submissions,omitempty"`
	// Attendances holds the value of the attendances edge.
	Attendances []*Attendance `json:"attendances,omitempty"`
	// Payments holds the value of the payments edge.
	Payments []*TuitionPayment `json:"payments,omitempty"`
	// Grades holds the value of the grades edge.
	Grades []*Grade `json:"grades,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [8]bool
}

// StageOrErr returns the Stage value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) StageOrErr() (*Stage, error) {
	if e.loadedTypes[0] {
		if e.Stage == nil {
			// The edge stage was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: stage.Label}
		}
		return e.Stage, nil
	}
	return nil, &NotLoadedError{edge: "stage"}
}

// SchoolOrErr returns the School value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) SchoolOrErr() (*School, error) {
	if e.loadedTypes[1] {
		if e.School == nil {
			// The edge school was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: school.Label}
		}
		return e.School, nil
	}
	return nil, &NotLoadedError{edge: "school"}
}

// ClassesOrErr returns the Classes value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ClassesOrErr() ([]*Class, error) {
	if e.loadedTypes[2] {
		return e.Classes, nil
	}
	return nil, &NotLoadedError{edge: "classes"}
}

// MessagesOrErr returns the Messages value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) MessagesOrErr() ([]*Message, error) {
	if e.loadedTypes[3] {
		return e.Messages, nil
	}
	return nil, &NotLoadedError{edge: "messages"}
}

// SubmissionsOrErr returns the Submissions value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) SubmissionsOrErr() ([]*AssignmentSubmission, error) {
	if e.loadedTypes[4] {
		return e.Submissions, nil
	}
	return nil, &NotLoadedError{edge: "submissions"}
}

// AttendancesOrErr returns the Attendances value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) AttendancesOrErr() ([]*Attendance, error) {
	if e.loadedTypes[5] {
		return e.Attendances, nil
	}
	return nil, &NotLoadedError{edge: "attendances"}
}

// PaymentsOrErr returns the Payments value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) PaymentsOrErr() ([]*TuitionPayment, error) {
	if e.loadedTypes[6] {
		return e.Payments, nil
	}
	return nil, &NotLoadedError{edge: "payments"}
}

// GradesOrErr returns the Grades value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) GradesOrErr() ([]*Grade, error) {
	if e.loadedTypes[7] {
		return e.Grades, nil
	}
	return nil, &NotLoadedError{edge: "grades"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID, user.FieldTokenVersion:
			values[i] = new(sql.NullInt64)
		case user.FieldName, user.FieldUsername, user.FieldPassword, user.FieldPhone, user.FieldImage, user.FieldRole, user.FieldStatus:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt, user.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case user.ForeignKeys[0]: // school_users
			values[i] = new(sql.NullInt64)
		case user.ForeignKeys[1]: // stage_students
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		case user.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				u.Name = value.String
			}
		case user.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				u.Username = value.String
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case user.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				u.Phone = value.String
			}
		case user.FieldImage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image", values[i])
			} else if value.Valid {
				u.Image = value.String
			}
		case user.FieldTokenVersion:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field token_version", values[i])
			} else if value.Valid {
				u.TokenVersion = int(value.Int64)
			}
		case user.FieldRole:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field role", values[i])
			} else if value.Valid {
				u.Role = user.Role(value.String)
			}
		case user.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				u.Status = schema.Status(value.String)
			}
		case user.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field school_users", value)
			} else if value.Valid {
				u.school_users = new(int)
				*u.school_users = int(value.Int64)
			}
		case user.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field stage_students", value)
			} else if value.Valid {
				u.stage_students = new(int)
				*u.stage_students = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryStage queries the "stage" edge of the User entity.
func (u *User) QueryStage() *StageQuery {
	return (&UserClient{config: u.config}).QueryStage(u)
}

// QuerySchool queries the "school" edge of the User entity.
func (u *User) QuerySchool() *SchoolQuery {
	return (&UserClient{config: u.config}).QuerySchool(u)
}

// QueryClasses queries the "classes" edge of the User entity.
func (u *User) QueryClasses() *ClassQuery {
	return (&UserClient{config: u.config}).QueryClasses(u)
}

// QueryMessages queries the "messages" edge of the User entity.
func (u *User) QueryMessages() *MessageQuery {
	return (&UserClient{config: u.config}).QueryMessages(u)
}

// QuerySubmissions queries the "submissions" edge of the User entity.
func (u *User) QuerySubmissions() *AssignmentSubmissionQuery {
	return (&UserClient{config: u.config}).QuerySubmissions(u)
}

// QueryAttendances queries the "attendances" edge of the User entity.
func (u *User) QueryAttendances() *AttendanceQuery {
	return (&UserClient{config: u.config}).QueryAttendances(u)
}

// QueryPayments queries the "payments" edge of the User entity.
func (u *User) QueryPayments() *TuitionPaymentQuery {
	return (&UserClient{config: u.config}).QueryPayments(u)
}

// QueryGrades queries the "grades" edge of the User entity.
func (u *User) QueryGrades() *GradeQuery {
	return (&UserClient{config: u.config}).QueryGrades(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", name=")
	builder.WriteString(u.Name)
	builder.WriteString(", username=")
	builder.WriteString(u.Username)
	builder.WriteString(", password=")
	builder.WriteString(u.Password)
	builder.WriteString(", phone=")
	builder.WriteString(u.Phone)
	builder.WriteString(", image=")
	builder.WriteString(u.Image)
	builder.WriteString(", token_version=")
	builder.WriteString(fmt.Sprintf("%v", u.TokenVersion))
	builder.WriteString(", role=")
	builder.WriteString(fmt.Sprintf("%v", u.Role))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", u.Status))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
