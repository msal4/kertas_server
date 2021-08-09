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
	"github.com/msal4/hassah_school_server/ent/attendance"
	"github.com/msal4/hassah_school_server/ent/class"
	"github.com/msal4/hassah_school_server/ent/user"
)

// AttendanceCreate is the builder for creating a Attendance entity.
type AttendanceCreate struct {
	config
	mutation *AttendanceMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ac *AttendanceCreate) SetCreatedAt(t time.Time) *AttendanceCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AttendanceCreate) SetNillableCreatedAt(t *time.Time) *AttendanceCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *AttendanceCreate) SetUpdatedAt(t time.Time) *AttendanceCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *AttendanceCreate) SetNillableUpdatedAt(t *time.Time) *AttendanceCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetDate sets the "date" field.
func (ac *AttendanceCreate) SetDate(t time.Time) *AttendanceCreate {
	ac.mutation.SetDate(t)
	return ac
}

// SetState sets the "state" field.
func (ac *AttendanceCreate) SetState(a attendance.State) *AttendanceCreate {
	ac.mutation.SetState(a)
	return ac
}

// SetNillableState sets the "state" field if the given value is not nil.
func (ac *AttendanceCreate) SetNillableState(a *attendance.State) *AttendanceCreate {
	if a != nil {
		ac.SetState(*a)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *AttendanceCreate) SetID(u uuid.UUID) *AttendanceCreate {
	ac.mutation.SetID(u)
	return ac
}

// SetClassID sets the "class" edge to the Class entity by ID.
func (ac *AttendanceCreate) SetClassID(id uuid.UUID) *AttendanceCreate {
	ac.mutation.SetClassID(id)
	return ac
}

// SetClass sets the "class" edge to the Class entity.
func (ac *AttendanceCreate) SetClass(c *Class) *AttendanceCreate {
	return ac.SetClassID(c.ID)
}

// SetStudentID sets the "student" edge to the User entity by ID.
func (ac *AttendanceCreate) SetStudentID(id uuid.UUID) *AttendanceCreate {
	ac.mutation.SetStudentID(id)
	return ac
}

// SetStudent sets the "student" edge to the User entity.
func (ac *AttendanceCreate) SetStudent(u *User) *AttendanceCreate {
	return ac.SetStudentID(u.ID)
}

// Mutation returns the AttendanceMutation object of the builder.
func (ac *AttendanceCreate) Mutation() *AttendanceMutation {
	return ac.mutation
}

// Save creates the Attendance in the database.
func (ac *AttendanceCreate) Save(ctx context.Context) (*Attendance, error) {
	var (
		err  error
		node *Attendance
	)
	ac.defaults()
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AttendanceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ac.check(); err != nil {
				return nil, err
			}
			ac.mutation = mutation
			if node, err = ac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			if ac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AttendanceCreate) SaveX(ctx context.Context) *Attendance {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AttendanceCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AttendanceCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AttendanceCreate) defaults() {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := attendance.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		v := attendance.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
	if _, ok := ac.mutation.State(); !ok {
		v := attendance.DefaultState
		ac.mutation.SetState(v)
	}
	if _, ok := ac.mutation.ID(); !ok {
		v := attendance.DefaultID()
		ac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AttendanceCreate) check() error {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := ac.mutation.Date(); !ok {
		return &ValidationError{Name: "date", err: errors.New(`ent: missing required field "date"`)}
	}
	if _, ok := ac.mutation.State(); !ok {
		return &ValidationError{Name: "state", err: errors.New(`ent: missing required field "state"`)}
	}
	if v, ok := ac.mutation.State(); ok {
		if err := attendance.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`ent: validator failed for field "state": %w`, err)}
		}
	}
	if _, ok := ac.mutation.ClassID(); !ok {
		return &ValidationError{Name: "class", err: errors.New("ent: missing required edge \"class\"")}
	}
	if _, ok := ac.mutation.StudentID(); !ok {
		return &ValidationError{Name: "student", err: errors.New("ent: missing required edge \"student\"")}
	}
	return nil
}

func (ac *AttendanceCreate) sqlSave(ctx context.Context) (*Attendance, error) {
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}

func (ac *AttendanceCreate) createSpec() (*Attendance, *sqlgraph.CreateSpec) {
	var (
		_node = &Attendance{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: attendance.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: attendance.FieldID,
			},
		}
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: attendance.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: attendance.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := ac.mutation.Date(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: attendance.FieldDate,
		})
		_node.Date = value
	}
	if value, ok := ac.mutation.State(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: attendance.FieldState,
		})
		_node.State = value
	}
	if nodes := ac.mutation.ClassIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attendance.ClassTable,
			Columns: []string{attendance.ClassColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: class.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.class_attendances = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.StudentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attendance.StudentTable,
			Columns: []string{attendance.StudentColumn},
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
		_node.user_attendances = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AttendanceCreateBulk is the builder for creating many Attendance entities in bulk.
type AttendanceCreateBulk struct {
	config
	builders []*AttendanceCreate
}

// Save creates the Attendance entities in the database.
func (acb *AttendanceCreateBulk) Save(ctx context.Context) ([]*Attendance, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Attendance, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AttendanceMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AttendanceCreateBulk) SaveX(ctx context.Context) []*Attendance {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AttendanceCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AttendanceCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
