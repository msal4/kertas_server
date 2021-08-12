// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/json"
	"fmt"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
	"github.com/hashicorp/go-multierror"
	"github.com/msal4/hassah_school_server/ent/assignment"
	"github.com/msal4/hassah_school_server/ent/assignmentsubmission"
	"github.com/msal4/hassah_school_server/ent/attendance"
	"github.com/msal4/hassah_school_server/ent/class"
	"github.com/msal4/hassah_school_server/ent/grade"
	"github.com/msal4/hassah_school_server/ent/group"
	"github.com/msal4/hassah_school_server/ent/message"
	"github.com/msal4/hassah_school_server/ent/schedule"
	"github.com/msal4/hassah_school_server/ent/school"
	"github.com/msal4/hassah_school_server/ent/stage"
	"github.com/msal4/hassah_school_server/ent/tuitionpayment"
	"github.com/msal4/hassah_school_server/ent/user"
)

// Noder wraps the basic Node method.
type Noder interface {
	Node(context.Context) (*Node, error)
}

// Node in the graph.
type Node struct {
	ID     uuid.UUID `json:"id,omitempty"`     // node id.
	Type   string    `json:"type,omitempty"`   // node type.
	Fields []*Field  `json:"fields,omitempty"` // node fields.
	Edges  []*Edge   `json:"edges,omitempty"`  // node edges.
}

// Field of a node.
type Field struct {
	Type  string `json:"type,omitempty"`  // field type.
	Name  string `json:"name,omitempty"`  // field name (as in struct).
	Value string `json:"value,omitempty"` // stringified value.
}

// Edges between two nodes.
type Edge struct {
	Type string      `json:"type,omitempty"` // edge type.
	Name string      `json:"name,omitempty"` // edge name.
	IDs  []uuid.UUID `json:"ids,omitempty"`  // node ids (where this edge point to).
}

func (a *Assignment) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     a.ID,
		Type:   "Assignment",
		Fields: make([]*Field, 8),
		Edges:  make([]*Edge, 3),
	}
	var buf []byte
	if buf, err = json.Marshal(a.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.Name); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.Description); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "string",
		Name:  "description",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.IsExam); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "bool",
		Name:  "is_exam",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.DueDate); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "time.Time",
		Name:  "due_date",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.Duration); err != nil {
		return nil, err
	}
	node.Fields[6] = &Field{
		Type:  "int",
		Name:  "duration",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.DeletedAt); err != nil {
		return nil, err
	}
	node.Fields[7] = &Field{
		Type:  "time.Time",
		Name:  "deleted_at",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Class",
		Name: "class",
	}
	err = a.QueryClass().
		Select(class.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "AssignmentSubmission",
		Name: "submissions",
	}
	err = a.QuerySubmissions().
		Select(assignmentsubmission.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[2] = &Edge{
		Type: "Grade",
		Name: "grades",
	}
	err = a.QueryGrades().
		Select(grade.FieldID).
		Scan(ctx, &node.Edges[2].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (as *AssignmentSubmission) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     as.ID,
		Type:   "AssignmentSubmission",
		Fields: make([]*Field, 4),
		Edges:  make([]*Edge, 2),
	}
	var buf []byte
	if buf, err = json.Marshal(as.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(as.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(as.Files); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "[]string",
		Name:  "files",
		Value: string(buf),
	}
	if buf, err = json.Marshal(as.SubmittedAt); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "time.Time",
		Name:  "submitted_at",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "User",
		Name: "student",
	}
	err = as.QueryStudent().
		Select(user.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "Assignment",
		Name: "assignment",
	}
	err = as.QueryAssignment().
		Select(assignment.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (a *Attendance) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     a.ID,
		Type:   "Attendance",
		Fields: make([]*Field, 4),
		Edges:  make([]*Edge, 2),
	}
	var buf []byte
	if buf, err = json.Marshal(a.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.Date); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "time.Time",
		Name:  "date",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.State); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "attendance.State",
		Name:  "state",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Class",
		Name: "class",
	}
	err = a.QueryClass().
		Select(class.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "User",
		Name: "student",
	}
	err = a.QueryStudent().
		Select(user.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (c *Class) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     c.ID,
		Type:   "Class",
		Fields: make([]*Field, 4),
		Edges:  make([]*Edge, 6),
	}
	var buf []byte
	if buf, err = json.Marshal(c.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(c.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(c.Name); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(c.Status); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "schema.Status",
		Name:  "status",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Stage",
		Name: "stage",
	}
	err = c.QueryStage().
		Select(stage.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "User",
		Name: "teacher",
	}
	err = c.QueryTeacher().
		Select(user.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[2] = &Edge{
		Type: "Group",
		Name: "group",
	}
	err = c.QueryGroup().
		Select(group.FieldID).
		Scan(ctx, &node.Edges[2].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[3] = &Edge{
		Type: "Assignment",
		Name: "assignments",
	}
	err = c.QueryAssignments().
		Select(assignment.FieldID).
		Scan(ctx, &node.Edges[3].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[4] = &Edge{
		Type: "Attendance",
		Name: "attendances",
	}
	err = c.QueryAttendances().
		Select(attendance.FieldID).
		Scan(ctx, &node.Edges[4].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[5] = &Edge{
		Type: "Schedule",
		Name: "schedules",
	}
	err = c.QuerySchedules().
		Select(schedule.FieldID).
		Scan(ctx, &node.Edges[5].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (gr *Grade) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     gr.ID,
		Type:   "Grade",
		Fields: make([]*Field, 3),
		Edges:  make([]*Edge, 2),
	}
	var buf []byte
	if buf, err = json.Marshal(gr.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(gr.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(gr.ExamGrade); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "int",
		Name:  "exam_grade",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "User",
		Name: "student",
	}
	err = gr.QueryStudent().
		Select(user.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "Assignment",
		Name: "exam",
	}
	err = gr.QueryExam().
		Select(assignment.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (gr *Group) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     gr.ID,
		Type:   "Group",
		Fields: make([]*Field, 5),
		Edges:  make([]*Edge, 2),
	}
	var buf []byte
	if buf, err = json.Marshal(gr.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(gr.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(gr.Name); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(gr.GroupType); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "group.GroupType",
		Name:  "group_type",
		Value: string(buf),
	}
	if buf, err = json.Marshal(gr.Status); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "schema.Status",
		Name:  "status",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Class",
		Name: "class",
	}
	err = gr.QueryClass().
		Select(class.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "Message",
		Name: "messages",
	}
	err = gr.QueryMessages().
		Select(message.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (m *Message) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     m.ID,
		Type:   "Message",
		Fields: make([]*Field, 5),
		Edges:  make([]*Edge, 2),
	}
	var buf []byte
	if buf, err = json.Marshal(m.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(m.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(m.Content); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "string",
		Name:  "content",
		Value: string(buf),
	}
	if buf, err = json.Marshal(m.Attachment); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "string",
		Name:  "attachment",
		Value: string(buf),
	}
	if buf, err = json.Marshal(m.DeletedAt); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "time.Time",
		Name:  "deleted_at",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Group",
		Name: "group",
	}
	err = m.QueryGroup().
		Select(group.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "User",
		Name: "owner",
	}
	err = m.QueryOwner().
		Select(user.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (s *Schedule) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     s.ID,
		Type:   "Schedule",
		Fields: make([]*Field, 3),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(s.Weekday); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "int",
		Name:  "weekday",
		Value: string(buf),
	}
	if buf, err = json.Marshal(s.StartsAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "starts_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(s.Duration); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "int",
		Name:  "duration",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Class",
		Name: "class",
	}
	err = s.QueryClass().
		Select(class.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (s *School) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     s.ID,
		Type:   "School",
		Fields: make([]*Field, 6),
		Edges:  make([]*Edge, 2),
	}
	var buf []byte
	if buf, err = json.Marshal(s.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(s.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(s.Name); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(s.Image); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "string",
		Name:  "image",
		Value: string(buf),
	}
	if buf, err = json.Marshal(s.Directory); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "string",
		Name:  "directory",
		Value: string(buf),
	}
	if buf, err = json.Marshal(s.Status); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "schema.Status",
		Name:  "status",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "User",
		Name: "users",
	}
	err = s.QueryUsers().
		Select(user.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "Stage",
		Name: "stages",
	}
	err = s.QueryStages().
		Select(stage.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (s *Stage) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     s.ID,
		Type:   "Stage",
		Fields: make([]*Field, 5),
		Edges:  make([]*Edge, 4),
	}
	var buf []byte
	if buf, err = json.Marshal(s.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(s.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(s.Name); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(s.TuitionAmount); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "int",
		Name:  "tuition_amount",
		Value: string(buf),
	}
	if buf, err = json.Marshal(s.Status); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "schema.Status",
		Name:  "status",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "School",
		Name: "school",
	}
	err = s.QuerySchool().
		Select(school.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "Class",
		Name: "classes",
	}
	err = s.QueryClasses().
		Select(class.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[2] = &Edge{
		Type: "TuitionPayment",
		Name: "payments",
	}
	err = s.QueryPayments().
		Select(tuitionpayment.FieldID).
		Scan(ctx, &node.Edges[2].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[3] = &Edge{
		Type: "User",
		Name: "students",
	}
	err = s.QueryStudents().
		Select(user.FieldID).
		Scan(ctx, &node.Edges[3].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (tp *TuitionPayment) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     tp.ID,
		Type:   "TuitionPayment",
		Fields: make([]*Field, 3),
		Edges:  make([]*Edge, 2),
	}
	var buf []byte
	if buf, err = json.Marshal(tp.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(tp.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(tp.PaidAmount); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "int",
		Name:  "paid_amount",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "User",
		Name: "student",
	}
	err = tp.QueryStudent().
		Select(user.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "Stage",
		Name: "stage",
	}
	err = tp.QueryStage().
		Select(stage.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (u *User) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     u.ID,
		Type:   "User",
		Fields: make([]*Field, 10),
		Edges:  make([]*Edge, 8),
	}
	var buf []byte
	if buf, err = json.Marshal(u.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(u.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(u.Name); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(u.Username); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "string",
		Name:  "username",
		Value: string(buf),
	}
	if buf, err = json.Marshal(u.Password); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "string",
		Name:  "password",
		Value: string(buf),
	}
	if buf, err = json.Marshal(u.Phone); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "string",
		Name:  "phone",
		Value: string(buf),
	}
	if buf, err = json.Marshal(u.Image); err != nil {
		return nil, err
	}
	node.Fields[6] = &Field{
		Type:  "string",
		Name:  "image",
		Value: string(buf),
	}
	if buf, err = json.Marshal(u.TokenVersion); err != nil {
		return nil, err
	}
	node.Fields[7] = &Field{
		Type:  "int",
		Name:  "token_version",
		Value: string(buf),
	}
	if buf, err = json.Marshal(u.Role); err != nil {
		return nil, err
	}
	node.Fields[8] = &Field{
		Type:  "user.Role",
		Name:  "role",
		Value: string(buf),
	}
	if buf, err = json.Marshal(u.Status); err != nil {
		return nil, err
	}
	node.Fields[9] = &Field{
		Type:  "schema.Status",
		Name:  "status",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Stage",
		Name: "stage",
	}
	err = u.QueryStage().
		Select(stage.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "School",
		Name: "school",
	}
	err = u.QuerySchool().
		Select(school.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[2] = &Edge{
		Type: "Class",
		Name: "classes",
	}
	err = u.QueryClasses().
		Select(class.FieldID).
		Scan(ctx, &node.Edges[2].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[3] = &Edge{
		Type: "Message",
		Name: "messages",
	}
	err = u.QueryMessages().
		Select(message.FieldID).
		Scan(ctx, &node.Edges[3].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[4] = &Edge{
		Type: "AssignmentSubmission",
		Name: "submissions",
	}
	err = u.QuerySubmissions().
		Select(assignmentsubmission.FieldID).
		Scan(ctx, &node.Edges[4].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[5] = &Edge{
		Type: "Attendance",
		Name: "attendances",
	}
	err = u.QueryAttendances().
		Select(attendance.FieldID).
		Scan(ctx, &node.Edges[5].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[6] = &Edge{
		Type: "TuitionPayment",
		Name: "payments",
	}
	err = u.QueryPayments().
		Select(tuitionpayment.FieldID).
		Scan(ctx, &node.Edges[6].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[7] = &Edge{
		Type: "Grade",
		Name: "grades",
	}
	err = u.QueryGrades().
		Select(grade.FieldID).
		Scan(ctx, &node.Edges[7].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (c *Client) Node(ctx context.Context, id uuid.UUID) (*Node, error) {
	n, err := c.Noder(ctx, id)
	if err != nil {
		return nil, err
	}
	return n.Node(ctx)
}

var errNodeInvalidID = &NotFoundError{"node"}

// NodeOption allows configuring the Noder execution using functional options.
type NodeOption func(*nodeOptions)

// WithNodeType sets the node Type resolver function (i.e. the table to query).
// If was not provided, the table will be derived from the universal-id
// configuration as described in: https://entgo.io/docs/migrate/#universal-ids.
func WithNodeType(f func(context.Context, uuid.UUID) (string, error)) NodeOption {
	return func(o *nodeOptions) {
		o.nodeType = f
	}
}

// WithFixedNodeType sets the Type of the node to a fixed value.
func WithFixedNodeType(t string) NodeOption {
	return WithNodeType(func(context.Context, uuid.UUID) (string, error) {
		return t, nil
	})
}

type nodeOptions struct {
	nodeType func(context.Context, uuid.UUID) (string, error)
}

func (c *Client) newNodeOpts(opts []NodeOption) *nodeOptions {
	nopts := &nodeOptions{}
	for _, opt := range opts {
		opt(nopts)
	}
	if nopts.nodeType == nil {
		nopts.nodeType = func(ctx context.Context, id uuid.UUID) (string, error) {
			return "", fmt.Errorf("cannot resolve noder (%v) without its type", id)
		}
	}
	return nopts
}

// Noder returns a Node by its id. If the NodeType was not provided, it will
// be derived from the id value according to the universal-id configuration.
//
//		c.Noder(ctx, id)
//		c.Noder(ctx, id, ent.WithNodeType(pet.Table))
//
func (c *Client) Noder(ctx context.Context, id uuid.UUID, opts ...NodeOption) (_ Noder, err error) {
	defer func() {
		if IsNotFound(err) {
			err = multierror.Append(err, entgql.ErrNodeNotFound(id))
		}
	}()
	table, err := c.newNodeOpts(opts).nodeType(ctx, id)
	if err != nil {
		return nil, err
	}
	return c.noder(ctx, table, id)
}

func (c *Client) noder(ctx context.Context, table string, id uuid.UUID) (Noder, error) {
	switch table {
	case assignment.Table:
		n, err := c.Assignment.Query().
			Where(assignment.ID(id)).
			CollectFields(ctx, "Assignment").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case assignmentsubmission.Table:
		n, err := c.AssignmentSubmission.Query().
			Where(assignmentsubmission.ID(id)).
			CollectFields(ctx, "AssignmentSubmission").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case attendance.Table:
		n, err := c.Attendance.Query().
			Where(attendance.ID(id)).
			CollectFields(ctx, "Attendance").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case class.Table:
		n, err := c.Class.Query().
			Where(class.ID(id)).
			CollectFields(ctx, "Class").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case grade.Table:
		n, err := c.Grade.Query().
			Where(grade.ID(id)).
			CollectFields(ctx, "Grade").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case group.Table:
		n, err := c.Group.Query().
			Where(group.ID(id)).
			CollectFields(ctx, "Group").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case message.Table:
		n, err := c.Message.Query().
			Where(message.ID(id)).
			CollectFields(ctx, "Message").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case schedule.Table:
		n, err := c.Schedule.Query().
			Where(schedule.ID(id)).
			CollectFields(ctx, "Schedule").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case school.Table:
		n, err := c.School.Query().
			Where(school.ID(id)).
			CollectFields(ctx, "School").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case stage.Table:
		n, err := c.Stage.Query().
			Where(stage.ID(id)).
			CollectFields(ctx, "Stage").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case tuitionpayment.Table:
		n, err := c.TuitionPayment.Query().
			Where(tuitionpayment.ID(id)).
			CollectFields(ctx, "TuitionPayment").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case user.Table:
		n, err := c.User.Query().
			Where(user.ID(id)).
			CollectFields(ctx, "User").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	default:
		return nil, fmt.Errorf("cannot resolve noder from table %q: %w", table, errNodeInvalidID)
	}
}

func (c *Client) Noders(ctx context.Context, ids []uuid.UUID, opts ...NodeOption) ([]Noder, error) {
	switch len(ids) {
	case 1:
		noder, err := c.Noder(ctx, ids[0], opts...)
		if err != nil {
			return nil, err
		}
		return []Noder{noder}, nil
	case 0:
		return []Noder{}, nil
	}

	noders := make([]Noder, len(ids))
	errors := make([]error, len(ids))
	tables := make(map[string][]uuid.UUID)
	id2idx := make(map[uuid.UUID][]int, len(ids))
	nopts := c.newNodeOpts(opts)
	for i, id := range ids {
		table, err := nopts.nodeType(ctx, id)
		if err != nil {
			errors[i] = err
			continue
		}
		tables[table] = append(tables[table], id)
		id2idx[id] = append(id2idx[id], i)
	}

	for table, ids := range tables {
		nodes, err := c.noders(ctx, table, ids)
		if err != nil {
			for _, id := range ids {
				for _, idx := range id2idx[id] {
					errors[idx] = err
				}
			}
		} else {
			for i, id := range ids {
				for _, idx := range id2idx[id] {
					noders[idx] = nodes[i]
				}
			}
		}
	}

	for i, id := range ids {
		if errors[i] == nil {
			if noders[i] != nil {
				continue
			}
			errors[i] = entgql.ErrNodeNotFound(id)
		} else if IsNotFound(errors[i]) {
			errors[i] = multierror.Append(errors[i], entgql.ErrNodeNotFound(id))
		}
		ctx := graphql.WithPathContext(ctx,
			graphql.NewPathWithIndex(i),
		)
		graphql.AddError(ctx, errors[i])
	}
	return noders, nil
}

func (c *Client) noders(ctx context.Context, table string, ids []uuid.UUID) ([]Noder, error) {
	noders := make([]Noder, len(ids))
	idmap := make(map[uuid.UUID][]*Noder, len(ids))
	for i, id := range ids {
		idmap[id] = append(idmap[id], &noders[i])
	}
	switch table {
	case assignment.Table:
		nodes, err := c.Assignment.Query().
			Where(assignment.IDIn(ids...)).
			CollectFields(ctx, "Assignment").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case assignmentsubmission.Table:
		nodes, err := c.AssignmentSubmission.Query().
			Where(assignmentsubmission.IDIn(ids...)).
			CollectFields(ctx, "AssignmentSubmission").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case attendance.Table:
		nodes, err := c.Attendance.Query().
			Where(attendance.IDIn(ids...)).
			CollectFields(ctx, "Attendance").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case class.Table:
		nodes, err := c.Class.Query().
			Where(class.IDIn(ids...)).
			CollectFields(ctx, "Class").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case grade.Table:
		nodes, err := c.Grade.Query().
			Where(grade.IDIn(ids...)).
			CollectFields(ctx, "Grade").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case group.Table:
		nodes, err := c.Group.Query().
			Where(group.IDIn(ids...)).
			CollectFields(ctx, "Group").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case message.Table:
		nodes, err := c.Message.Query().
			Where(message.IDIn(ids...)).
			CollectFields(ctx, "Message").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case schedule.Table:
		nodes, err := c.Schedule.Query().
			Where(schedule.IDIn(ids...)).
			CollectFields(ctx, "Schedule").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case school.Table:
		nodes, err := c.School.Query().
			Where(school.IDIn(ids...)).
			CollectFields(ctx, "School").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case stage.Table:
		nodes, err := c.Stage.Query().
			Where(stage.IDIn(ids...)).
			CollectFields(ctx, "Stage").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case tuitionpayment.Table:
		nodes, err := c.TuitionPayment.Query().
			Where(tuitionpayment.IDIn(ids...)).
			CollectFields(ctx, "TuitionPayment").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case user.Table:
		nodes, err := c.User.Query().
			Where(user.IDIn(ids...)).
			CollectFields(ctx, "User").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	default:
		return nil, fmt.Errorf("cannot resolve noders from table %q: %w", table, errNodeInvalidID)
	}
	return noders, nil
}
