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
	"github.com/msal4/hassah_school_server/ent/class"
	"github.com/msal4/hassah_school_server/ent/schedule"
)

// ScheduleCreate is the builder for creating a Schedule entity.
type ScheduleCreate struct {
	config
	mutation *ScheduleMutation
	hooks    []Hook
}

// SetWeekday sets the "weekday" field.
func (sc *ScheduleCreate) SetWeekday(i int) *ScheduleCreate {
	sc.mutation.SetWeekday(i)
	return sc
}

// SetStartsAt sets the "starts_at" field.
func (sc *ScheduleCreate) SetStartsAt(t time.Time) *ScheduleCreate {
	sc.mutation.SetStartsAt(t)
	return sc
}

// SetDuration sets the "duration" field.
func (sc *ScheduleCreate) SetDuration(i int) *ScheduleCreate {
	sc.mutation.SetDuration(i)
	return sc
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (sc *ScheduleCreate) SetNillableDuration(i *int) *ScheduleCreate {
	if i != nil {
		sc.SetDuration(*i)
	}
	return sc
}

// SetID sets the "id" field.
func (sc *ScheduleCreate) SetID(u uuid.UUID) *ScheduleCreate {
	sc.mutation.SetID(u)
	return sc
}

// SetClassID sets the "class" edge to the Class entity by ID.
func (sc *ScheduleCreate) SetClassID(id uuid.UUID) *ScheduleCreate {
	sc.mutation.SetClassID(id)
	return sc
}

// SetClass sets the "class" edge to the Class entity.
func (sc *ScheduleCreate) SetClass(c *Class) *ScheduleCreate {
	return sc.SetClassID(c.ID)
}

// Mutation returns the ScheduleMutation object of the builder.
func (sc *ScheduleCreate) Mutation() *ScheduleMutation {
	return sc.mutation
}

// Save creates the Schedule in the database.
func (sc *ScheduleCreate) Save(ctx context.Context) (*Schedule, error) {
	var (
		err  error
		node *Schedule
	)
	sc.defaults()
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ScheduleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			if node, err = sc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			if sc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *ScheduleCreate) SaveX(ctx context.Context) *Schedule {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *ScheduleCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *ScheduleCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *ScheduleCreate) defaults() {
	if _, ok := sc.mutation.Duration(); !ok {
		v := schedule.DefaultDuration
		sc.mutation.SetDuration(v)
	}
	if _, ok := sc.mutation.ID(); !ok {
		v := schedule.DefaultID()
		sc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *ScheduleCreate) check() error {
	if _, ok := sc.mutation.Weekday(); !ok {
		return &ValidationError{Name: "weekday", err: errors.New(`ent: missing required field "weekday"`)}
	}
	if v, ok := sc.mutation.Weekday(); ok {
		if err := schedule.WeekdayValidator(v); err != nil {
			return &ValidationError{Name: "weekday", err: fmt.Errorf(`ent: validator failed for field "weekday": %w`, err)}
		}
	}
	if _, ok := sc.mutation.StartsAt(); !ok {
		return &ValidationError{Name: "starts_at", err: errors.New(`ent: missing required field "starts_at"`)}
	}
	if _, ok := sc.mutation.Duration(); !ok {
		return &ValidationError{Name: "duration", err: errors.New(`ent: missing required field "duration"`)}
	}
	if _, ok := sc.mutation.ClassID(); !ok {
		return &ValidationError{Name: "class", err: errors.New("ent: missing required edge \"class\"")}
	}
	return nil
}

func (sc *ScheduleCreate) sqlSave(ctx context.Context) (*Schedule, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}

func (sc *ScheduleCreate) createSpec() (*Schedule, *sqlgraph.CreateSpec) {
	var (
		_node = &Schedule{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: schedule.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: schedule.FieldID,
			},
		}
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sc.mutation.Weekday(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: schedule.FieldWeekday,
		})
		_node.Weekday = value
	}
	if value, ok := sc.mutation.StartsAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: schedule.FieldStartsAt,
		})
		_node.StartsAt = value
	}
	if value, ok := sc.mutation.Duration(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: schedule.FieldDuration,
		})
		_node.Duration = value
	}
	if nodes := sc.mutation.ClassIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   schedule.ClassTable,
			Columns: []string{schedule.ClassColumn},
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
		_node.class_schedules = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ScheduleCreateBulk is the builder for creating many Schedule entities in bulk.
type ScheduleCreateBulk struct {
	config
	builders []*ScheduleCreate
}

// Save creates the Schedule entities in the database.
func (scb *ScheduleCreateBulk) Save(ctx context.Context) ([]*Schedule, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Schedule, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ScheduleMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *ScheduleCreateBulk) SaveX(ctx context.Context) []*Schedule {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *ScheduleCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *ScheduleCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
