// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/msal4/hassah_school_server/ent/assignment"
	"github.com/msal4/hassah_school_server/ent/attendance"
	"github.com/msal4/hassah_school_server/ent/class"
	"github.com/msal4/hassah_school_server/ent/group"
	"github.com/msal4/hassah_school_server/ent/schedule"
	"github.com/msal4/hassah_school_server/ent/schema"
	"github.com/msal4/hassah_school_server/ent/stage"
	"github.com/msal4/hassah_school_server/ent/user"
)

// ClassCreate is the builder for creating a Class entity.
type ClassCreate struct {
	config
	mutation *ClassMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (cc *ClassCreate) SetCreatedAt(t time.Time) *ClassCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *ClassCreate) SetNillableCreatedAt(t *time.Time) *ClassCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *ClassCreate) SetUpdatedAt(t time.Time) *ClassCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *ClassCreate) SetNillableUpdatedAt(t *time.Time) *ClassCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetName sets the "name" field.
func (cc *ClassCreate) SetName(s string) *ClassCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetStatus sets the "status" field.
func (cc *ClassCreate) SetStatus(s schema.Status) *ClassCreate {
	cc.mutation.SetStatus(s)
	return cc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cc *ClassCreate) SetNillableStatus(s *schema.Status) *ClassCreate {
	if s != nil {
		cc.SetStatus(*s)
	}
	return cc
}

// SetDeletedAt sets the "deleted_at" field.
func (cc *ClassCreate) SetDeletedAt(t time.Time) *ClassCreate {
	cc.mutation.SetDeletedAt(t)
	return cc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cc *ClassCreate) SetNillableDeletedAt(t *time.Time) *ClassCreate {
	if t != nil {
		cc.SetDeletedAt(*t)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *ClassCreate) SetID(u uuid.UUID) *ClassCreate {
	cc.mutation.SetID(u)
	return cc
}

// SetStageID sets the "stage" edge to the Stage entity by ID.
func (cc *ClassCreate) SetStageID(id uuid.UUID) *ClassCreate {
	cc.mutation.SetStageID(id)
	return cc
}

// SetStage sets the "stage" edge to the Stage entity.
func (cc *ClassCreate) SetStage(s *Stage) *ClassCreate {
	return cc.SetStageID(s.ID)
}

// SetTeacherID sets the "teacher" edge to the User entity by ID.
func (cc *ClassCreate) SetTeacherID(id uuid.UUID) *ClassCreate {
	cc.mutation.SetTeacherID(id)
	return cc
}

// SetTeacher sets the "teacher" edge to the User entity.
func (cc *ClassCreate) SetTeacher(u *User) *ClassCreate {
	return cc.SetTeacherID(u.ID)
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (cc *ClassCreate) SetGroupID(id uuid.UUID) *ClassCreate {
	cc.mutation.SetGroupID(id)
	return cc
}

// SetGroup sets the "group" edge to the Group entity.
func (cc *ClassCreate) SetGroup(g *Group) *ClassCreate {
	return cc.SetGroupID(g.ID)
}

// AddAssignmentIDs adds the "assignments" edge to the Assignment entity by IDs.
func (cc *ClassCreate) AddAssignmentIDs(ids ...uuid.UUID) *ClassCreate {
	cc.mutation.AddAssignmentIDs(ids...)
	return cc
}

// AddAssignments adds the "assignments" edges to the Assignment entity.
func (cc *ClassCreate) AddAssignments(a ...*Assignment) *ClassCreate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cc.AddAssignmentIDs(ids...)
}

// AddAttendanceIDs adds the "attendances" edge to the Attendance entity by IDs.
func (cc *ClassCreate) AddAttendanceIDs(ids ...uuid.UUID) *ClassCreate {
	cc.mutation.AddAttendanceIDs(ids...)
	return cc
}

// AddAttendances adds the "attendances" edges to the Attendance entity.
func (cc *ClassCreate) AddAttendances(a ...*Attendance) *ClassCreate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cc.AddAttendanceIDs(ids...)
}

// AddScheduleIDs adds the "schedules" edge to the Schedule entity by IDs.
func (cc *ClassCreate) AddScheduleIDs(ids ...uuid.UUID) *ClassCreate {
	cc.mutation.AddScheduleIDs(ids...)
	return cc
}

// AddSchedules adds the "schedules" edges to the Schedule entity.
func (cc *ClassCreate) AddSchedules(s ...*Schedule) *ClassCreate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cc.AddScheduleIDs(ids...)
}

// Mutation returns the ClassMutation object of the builder.
func (cc *ClassCreate) Mutation() *ClassMutation {
	return cc.mutation
}

// Save creates the Class in the database.
func (cc *ClassCreate) Save(ctx context.Context) (*Class, error) {
	var (
		err  error
		node *Class
	)
	cc.defaults()
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ClassMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ClassCreate) SaveX(ctx context.Context) *Class {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ClassCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ClassCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *ClassCreate) defaults() {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := class.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		v := class.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
	if _, ok := cc.mutation.Status(); !ok {
		v := class.DefaultStatus
		cc.mutation.SetStatus(v)
	}
	if _, ok := cc.mutation.ID(); !ok {
		v := class.DefaultID()
		cc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *ClassCreate) check() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if v, ok := cc.mutation.Name(); ok {
		if err := class.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "name": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "status"`)}
	}
	if v, ok := cc.mutation.Status(); ok {
		if err := class.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "status": %w`, err)}
		}
	}
	if _, ok := cc.mutation.StageID(); !ok {
		return &ValidationError{Name: "stage", err: errors.New("ent: missing required edge \"stage\"")}
	}
	if _, ok := cc.mutation.TeacherID(); !ok {
		return &ValidationError{Name: "teacher", err: errors.New("ent: missing required edge \"teacher\"")}
	}
	if _, ok := cc.mutation.GroupID(); !ok {
		return &ValidationError{Name: "group", err: errors.New("ent: missing required edge \"group\"")}
	}
	return nil
}

func (cc *ClassCreate) sqlSave(ctx context.Context) (*Class, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}

func (cc *ClassCreate) createSpec() (*Class, *sqlgraph.CreateSpec) {
	var (
		_node = &Class{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: class.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: class.FieldID,
			},
		}
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: class.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: class.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: class.FieldName,
		})
		_node.Name = value
	}
	if value, ok := cc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: class.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := cc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: class.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if nodes := cc.mutation.StageIDs(); len(nodes) > 0 {
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
		_node.stage_classes = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.TeacherIDs(); len(nodes) > 0 {
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
		_node.user_classes = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.GroupIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.AssignmentsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.AttendancesIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.SchedulesIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ClassCreateBulk is the builder for creating many Class entities in bulk.
type ClassCreateBulk struct {
	config
	builders []*ClassCreate
}

// Save creates the Class entities in the database.
func (ccb *ClassCreateBulk) Save(ctx context.Context) ([]*Class, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Class, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ClassMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ClassCreateBulk) SaveX(ctx context.Context) []*Class {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ClassCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ClassCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
