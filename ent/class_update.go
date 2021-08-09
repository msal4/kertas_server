// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/msal4/hassah_school_server/ent/assignment"
	"github.com/msal4/hassah_school_server/ent/attendance"
	"github.com/msal4/hassah_school_server/ent/class"
	"github.com/msal4/hassah_school_server/ent/group"
	"github.com/msal4/hassah_school_server/ent/predicate"
	"github.com/msal4/hassah_school_server/ent/schedule"
	"github.com/msal4/hassah_school_server/ent/schema"
	"github.com/msal4/hassah_school_server/ent/stage"
	"github.com/msal4/hassah_school_server/ent/user"
)

// ClassUpdate is the builder for updating Class entities.
type ClassUpdate struct {
	config
	hooks    []Hook
	mutation *ClassMutation
}

// Where appends a list predicates to the ClassUpdate builder.
func (cu *ClassUpdate) Where(ps ...predicate.Class) *ClassUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetName sets the "name" field.
func (cu *ClassUpdate) SetName(s string) *ClassUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetStatus sets the "status" field.
func (cu *ClassUpdate) SetStatus(s schema.Status) *ClassUpdate {
	cu.mutation.SetStatus(s)
	return cu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cu *ClassUpdate) SetNillableStatus(s *schema.Status) *ClassUpdate {
	if s != nil {
		cu.SetStatus(*s)
	}
	return cu
}

// SetStageID sets the "stage" edge to the Stage entity by ID.
func (cu *ClassUpdate) SetStageID(id uuid.UUID) *ClassUpdate {
	cu.mutation.SetStageID(id)
	return cu
}

// SetStage sets the "stage" edge to the Stage entity.
func (cu *ClassUpdate) SetStage(s *Stage) *ClassUpdate {
	return cu.SetStageID(s.ID)
}

// SetTeacherID sets the "teacher" edge to the User entity by ID.
func (cu *ClassUpdate) SetTeacherID(id uuid.UUID) *ClassUpdate {
	cu.mutation.SetTeacherID(id)
	return cu
}

// SetTeacher sets the "teacher" edge to the User entity.
func (cu *ClassUpdate) SetTeacher(u *User) *ClassUpdate {
	return cu.SetTeacherID(u.ID)
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (cu *ClassUpdate) SetGroupID(id uuid.UUID) *ClassUpdate {
	cu.mutation.SetGroupID(id)
	return cu
}

// SetGroup sets the "group" edge to the Group entity.
func (cu *ClassUpdate) SetGroup(g *Group) *ClassUpdate {
	return cu.SetGroupID(g.ID)
}

// AddAssignmentIDs adds the "assignments" edge to the Assignment entity by IDs.
func (cu *ClassUpdate) AddAssignmentIDs(ids ...uuid.UUID) *ClassUpdate {
	cu.mutation.AddAssignmentIDs(ids...)
	return cu
}

// AddAssignments adds the "assignments" edges to the Assignment entity.
func (cu *ClassUpdate) AddAssignments(a ...*Assignment) *ClassUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cu.AddAssignmentIDs(ids...)
}

// AddAttendanceIDs adds the "attendances" edge to the Attendance entity by IDs.
func (cu *ClassUpdate) AddAttendanceIDs(ids ...uuid.UUID) *ClassUpdate {
	cu.mutation.AddAttendanceIDs(ids...)
	return cu
}

// AddAttendances adds the "attendances" edges to the Attendance entity.
func (cu *ClassUpdate) AddAttendances(a ...*Attendance) *ClassUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cu.AddAttendanceIDs(ids...)
}

// AddScheduleIDs adds the "schedules" edge to the Schedule entity by IDs.
func (cu *ClassUpdate) AddScheduleIDs(ids ...uuid.UUID) *ClassUpdate {
	cu.mutation.AddScheduleIDs(ids...)
	return cu
}

// AddSchedules adds the "schedules" edges to the Schedule entity.
func (cu *ClassUpdate) AddSchedules(s ...*Schedule) *ClassUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cu.AddScheduleIDs(ids...)
}

// Mutation returns the ClassMutation object of the builder.
func (cu *ClassUpdate) Mutation() *ClassMutation {
	return cu.mutation
}

// ClearStage clears the "stage" edge to the Stage entity.
func (cu *ClassUpdate) ClearStage() *ClassUpdate {
	cu.mutation.ClearStage()
	return cu
}

// ClearTeacher clears the "teacher" edge to the User entity.
func (cu *ClassUpdate) ClearTeacher() *ClassUpdate {
	cu.mutation.ClearTeacher()
	return cu
}

// ClearGroup clears the "group" edge to the Group entity.
func (cu *ClassUpdate) ClearGroup() *ClassUpdate {
	cu.mutation.ClearGroup()
	return cu
}

// ClearAssignments clears all "assignments" edges to the Assignment entity.
func (cu *ClassUpdate) ClearAssignments() *ClassUpdate {
	cu.mutation.ClearAssignments()
	return cu
}

// RemoveAssignmentIDs removes the "assignments" edge to Assignment entities by IDs.
func (cu *ClassUpdate) RemoveAssignmentIDs(ids ...uuid.UUID) *ClassUpdate {
	cu.mutation.RemoveAssignmentIDs(ids...)
	return cu
}

// RemoveAssignments removes "assignments" edges to Assignment entities.
func (cu *ClassUpdate) RemoveAssignments(a ...*Assignment) *ClassUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cu.RemoveAssignmentIDs(ids...)
}

// ClearAttendances clears all "attendances" edges to the Attendance entity.
func (cu *ClassUpdate) ClearAttendances() *ClassUpdate {
	cu.mutation.ClearAttendances()
	return cu
}

// RemoveAttendanceIDs removes the "attendances" edge to Attendance entities by IDs.
func (cu *ClassUpdate) RemoveAttendanceIDs(ids ...uuid.UUID) *ClassUpdate {
	cu.mutation.RemoveAttendanceIDs(ids...)
	return cu
}

// RemoveAttendances removes "attendances" edges to Attendance entities.
func (cu *ClassUpdate) RemoveAttendances(a ...*Attendance) *ClassUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cu.RemoveAttendanceIDs(ids...)
}

// ClearSchedules clears all "schedules" edges to the Schedule entity.
func (cu *ClassUpdate) ClearSchedules() *ClassUpdate {
	cu.mutation.ClearSchedules()
	return cu
}

// RemoveScheduleIDs removes the "schedules" edge to Schedule entities by IDs.
func (cu *ClassUpdate) RemoveScheduleIDs(ids ...uuid.UUID) *ClassUpdate {
	cu.mutation.RemoveScheduleIDs(ids...)
	return cu
}

// RemoveSchedules removes "schedules" edges to Schedule entities.
func (cu *ClassUpdate) RemoveSchedules(s ...*Schedule) *ClassUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cu.RemoveScheduleIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ClassUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	cu.defaults()
	if len(cu.hooks) == 0 {
		if err = cu.check(); err != nil {
			return 0, err
		}
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ClassMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cu.check(); err != nil {
				return 0, err
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ClassUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ClassUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ClassUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *ClassUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := class.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *ClassUpdate) check() error {
	if v, ok := cu.mutation.Name(); ok {
		if err := class.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := cu.mutation.Status(); ok {
		if err := class.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	if _, ok := cu.mutation.StageID(); cu.mutation.StageCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"stage\"")
	}
	if _, ok := cu.mutation.TeacherID(); cu.mutation.TeacherCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"teacher\"")
	}
	if _, ok := cu.mutation.GroupID(); cu.mutation.GroupCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"group\"")
	}
	return nil
}

func (cu *ClassUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   class.Table,
			Columns: class.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: class.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: class.FieldUpdatedAt,
		})
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: class.FieldName,
		})
	}
	if value, ok := cu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: class.FieldStatus,
		})
	}
	if cu.mutation.StageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   class.StageTable,
			Columns: []string{class.StageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: stage.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.StageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   class.StageTable,
			Columns: []string{class.StageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: stage.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.TeacherCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   class.TeacherTable,
			Columns: []string{class.TeacherColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.TeacherIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   class.TeacherTable,
			Columns: []string{class.TeacherColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.GroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   class.GroupTable,
			Columns: []string{class.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: group.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   class.GroupTable,
			Columns: []string{class.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.AssignmentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.AssignmentsTable,
			Columns: []string{class.AssignmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: assignment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedAssignmentsIDs(); len(nodes) > 0 && !cu.mutation.AssignmentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.AssignmentsTable,
			Columns: []string{class.AssignmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: assignment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.AssignmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.AssignmentsTable,
			Columns: []string{class.AssignmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: assignment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.AttendancesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.AttendancesTable,
			Columns: []string{class.AttendancesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: attendance.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedAttendancesIDs(); len(nodes) > 0 && !cu.mutation.AttendancesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.AttendancesTable,
			Columns: []string{class.AttendancesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: attendance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.AttendancesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.AttendancesTable,
			Columns: []string{class.AttendancesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: attendance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.SchedulesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.SchedulesTable,
			Columns: []string{class.SchedulesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: schedule.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedSchedulesIDs(); len(nodes) > 0 && !cu.mutation.SchedulesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.SchedulesTable,
			Columns: []string{class.SchedulesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: schedule.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.SchedulesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.SchedulesTable,
			Columns: []string{class.SchedulesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: schedule.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{class.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ClassUpdateOne is the builder for updating a single Class entity.
type ClassUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ClassMutation
}

// SetName sets the "name" field.
func (cuo *ClassUpdateOne) SetName(s string) *ClassUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetStatus sets the "status" field.
func (cuo *ClassUpdateOne) SetStatus(s schema.Status) *ClassUpdateOne {
	cuo.mutation.SetStatus(s)
	return cuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cuo *ClassUpdateOne) SetNillableStatus(s *schema.Status) *ClassUpdateOne {
	if s != nil {
		cuo.SetStatus(*s)
	}
	return cuo
}

// SetStageID sets the "stage" edge to the Stage entity by ID.
func (cuo *ClassUpdateOne) SetStageID(id uuid.UUID) *ClassUpdateOne {
	cuo.mutation.SetStageID(id)
	return cuo
}

// SetStage sets the "stage" edge to the Stage entity.
func (cuo *ClassUpdateOne) SetStage(s *Stage) *ClassUpdateOne {
	return cuo.SetStageID(s.ID)
}

// SetTeacherID sets the "teacher" edge to the User entity by ID.
func (cuo *ClassUpdateOne) SetTeacherID(id uuid.UUID) *ClassUpdateOne {
	cuo.mutation.SetTeacherID(id)
	return cuo
}

// SetTeacher sets the "teacher" edge to the User entity.
func (cuo *ClassUpdateOne) SetTeacher(u *User) *ClassUpdateOne {
	return cuo.SetTeacherID(u.ID)
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (cuo *ClassUpdateOne) SetGroupID(id uuid.UUID) *ClassUpdateOne {
	cuo.mutation.SetGroupID(id)
	return cuo
}

// SetGroup sets the "group" edge to the Group entity.
func (cuo *ClassUpdateOne) SetGroup(g *Group) *ClassUpdateOne {
	return cuo.SetGroupID(g.ID)
}

// AddAssignmentIDs adds the "assignments" edge to the Assignment entity by IDs.
func (cuo *ClassUpdateOne) AddAssignmentIDs(ids ...uuid.UUID) *ClassUpdateOne {
	cuo.mutation.AddAssignmentIDs(ids...)
	return cuo
}

// AddAssignments adds the "assignments" edges to the Assignment entity.
func (cuo *ClassUpdateOne) AddAssignments(a ...*Assignment) *ClassUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cuo.AddAssignmentIDs(ids...)
}

// AddAttendanceIDs adds the "attendances" edge to the Attendance entity by IDs.
func (cuo *ClassUpdateOne) AddAttendanceIDs(ids ...uuid.UUID) *ClassUpdateOne {
	cuo.mutation.AddAttendanceIDs(ids...)
	return cuo
}

// AddAttendances adds the "attendances" edges to the Attendance entity.
func (cuo *ClassUpdateOne) AddAttendances(a ...*Attendance) *ClassUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cuo.AddAttendanceIDs(ids...)
}

// AddScheduleIDs adds the "schedules" edge to the Schedule entity by IDs.
func (cuo *ClassUpdateOne) AddScheduleIDs(ids ...uuid.UUID) *ClassUpdateOne {
	cuo.mutation.AddScheduleIDs(ids...)
	return cuo
}

// AddSchedules adds the "schedules" edges to the Schedule entity.
func (cuo *ClassUpdateOne) AddSchedules(s ...*Schedule) *ClassUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cuo.AddScheduleIDs(ids...)
}

// Mutation returns the ClassMutation object of the builder.
func (cuo *ClassUpdateOne) Mutation() *ClassMutation {
	return cuo.mutation
}

// ClearStage clears the "stage" edge to the Stage entity.
func (cuo *ClassUpdateOne) ClearStage() *ClassUpdateOne {
	cuo.mutation.ClearStage()
	return cuo
}

// ClearTeacher clears the "teacher" edge to the User entity.
func (cuo *ClassUpdateOne) ClearTeacher() *ClassUpdateOne {
	cuo.mutation.ClearTeacher()
	return cuo
}

// ClearGroup clears the "group" edge to the Group entity.
func (cuo *ClassUpdateOne) ClearGroup() *ClassUpdateOne {
	cuo.mutation.ClearGroup()
	return cuo
}

// ClearAssignments clears all "assignments" edges to the Assignment entity.
func (cuo *ClassUpdateOne) ClearAssignments() *ClassUpdateOne {
	cuo.mutation.ClearAssignments()
	return cuo
}

// RemoveAssignmentIDs removes the "assignments" edge to Assignment entities by IDs.
func (cuo *ClassUpdateOne) RemoveAssignmentIDs(ids ...uuid.UUID) *ClassUpdateOne {
	cuo.mutation.RemoveAssignmentIDs(ids...)
	return cuo
}

// RemoveAssignments removes "assignments" edges to Assignment entities.
func (cuo *ClassUpdateOne) RemoveAssignments(a ...*Assignment) *ClassUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cuo.RemoveAssignmentIDs(ids...)
}

// ClearAttendances clears all "attendances" edges to the Attendance entity.
func (cuo *ClassUpdateOne) ClearAttendances() *ClassUpdateOne {
	cuo.mutation.ClearAttendances()
	return cuo
}

// RemoveAttendanceIDs removes the "attendances" edge to Attendance entities by IDs.
func (cuo *ClassUpdateOne) RemoveAttendanceIDs(ids ...uuid.UUID) *ClassUpdateOne {
	cuo.mutation.RemoveAttendanceIDs(ids...)
	return cuo
}

// RemoveAttendances removes "attendances" edges to Attendance entities.
func (cuo *ClassUpdateOne) RemoveAttendances(a ...*Attendance) *ClassUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cuo.RemoveAttendanceIDs(ids...)
}

// ClearSchedules clears all "schedules" edges to the Schedule entity.
func (cuo *ClassUpdateOne) ClearSchedules() *ClassUpdateOne {
	cuo.mutation.ClearSchedules()
	return cuo
}

// RemoveScheduleIDs removes the "schedules" edge to Schedule entities by IDs.
func (cuo *ClassUpdateOne) RemoveScheduleIDs(ids ...uuid.UUID) *ClassUpdateOne {
	cuo.mutation.RemoveScheduleIDs(ids...)
	return cuo
}

// RemoveSchedules removes "schedules" edges to Schedule entities.
func (cuo *ClassUpdateOne) RemoveSchedules(s ...*Schedule) *ClassUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cuo.RemoveScheduleIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ClassUpdateOne) Select(field string, fields ...string) *ClassUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Class entity.
func (cuo *ClassUpdateOne) Save(ctx context.Context) (*Class, error) {
	var (
		err  error
		node *Class
	)
	cuo.defaults()
	if len(cuo.hooks) == 0 {
		if err = cuo.check(); err != nil {
			return nil, err
		}
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ClassMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cuo.check(); err != nil {
				return nil, err
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ClassUpdateOne) SaveX(ctx context.Context) *Class {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ClassUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ClassUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *ClassUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := class.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *ClassUpdateOne) check() error {
	if v, ok := cuo.mutation.Name(); ok {
		if err := class.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := cuo.mutation.Status(); ok {
		if err := class.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	if _, ok := cuo.mutation.StageID(); cuo.mutation.StageCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"stage\"")
	}
	if _, ok := cuo.mutation.TeacherID(); cuo.mutation.TeacherCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"teacher\"")
	}
	if _, ok := cuo.mutation.GroupID(); cuo.mutation.GroupCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"group\"")
	}
	return nil
}

func (cuo *ClassUpdateOne) sqlSave(ctx context.Context) (_node *Class, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   class.Table,
			Columns: class.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: class.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Class.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, class.FieldID)
		for _, f := range fields {
			if !class.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != class.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: class.FieldUpdatedAt,
		})
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: class.FieldName,
		})
	}
	if value, ok := cuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: class.FieldStatus,
		})
	}
	if cuo.mutation.StageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   class.StageTable,
			Columns: []string{class.StageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: stage.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.StageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   class.StageTable,
			Columns: []string{class.StageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: stage.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.TeacherCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   class.TeacherTable,
			Columns: []string{class.TeacherColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.TeacherIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   class.TeacherTable,
			Columns: []string{class.TeacherColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.GroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   class.GroupTable,
			Columns: []string{class.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: group.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   class.GroupTable,
			Columns: []string{class.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.AssignmentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.AssignmentsTable,
			Columns: []string{class.AssignmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: assignment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedAssignmentsIDs(); len(nodes) > 0 && !cuo.mutation.AssignmentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.AssignmentsTable,
			Columns: []string{class.AssignmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: assignment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.AssignmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.AssignmentsTable,
			Columns: []string{class.AssignmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: assignment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.AttendancesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.AttendancesTable,
			Columns: []string{class.AttendancesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: attendance.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedAttendancesIDs(); len(nodes) > 0 && !cuo.mutation.AttendancesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.AttendancesTable,
			Columns: []string{class.AttendancesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: attendance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.AttendancesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.AttendancesTable,
			Columns: []string{class.AttendancesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: attendance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.SchedulesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.SchedulesTable,
			Columns: []string{class.SchedulesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: schedule.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedSchedulesIDs(); len(nodes) > 0 && !cuo.mutation.SchedulesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.SchedulesTable,
			Columns: []string{class.SchedulesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: schedule.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.SchedulesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.SchedulesTable,
			Columns: []string{class.SchedulesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: schedule.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Class{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{class.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
