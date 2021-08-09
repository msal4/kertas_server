// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/msal4/hassah_school_server/ent/class"
	"github.com/msal4/hassah_school_server/ent/predicate"
	"github.com/msal4/hassah_school_server/ent/schedule"
)

// ScheduleUpdate is the builder for updating Schedule entities.
type ScheduleUpdate struct {
	config
	hooks    []Hook
	mutation *ScheduleMutation
}

// Where appends a list predicates to the ScheduleUpdate builder.
func (su *ScheduleUpdate) Where(ps ...predicate.Schedule) *ScheduleUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetWeekday sets the "weekday" field.
func (su *ScheduleUpdate) SetWeekday(i int) *ScheduleUpdate {
	su.mutation.ResetWeekday()
	su.mutation.SetWeekday(i)
	return su
}

// AddWeekday adds i to the "weekday" field.
func (su *ScheduleUpdate) AddWeekday(i int) *ScheduleUpdate {
	su.mutation.AddWeekday(i)
	return su
}

// SetStartsAt sets the "starts_at" field.
func (su *ScheduleUpdate) SetStartsAt(t time.Time) *ScheduleUpdate {
	su.mutation.SetStartsAt(t)
	return su
}

// SetDuration sets the "duration" field.
func (su *ScheduleUpdate) SetDuration(i int) *ScheduleUpdate {
	su.mutation.ResetDuration()
	su.mutation.SetDuration(i)
	return su
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (su *ScheduleUpdate) SetNillableDuration(i *int) *ScheduleUpdate {
	if i != nil {
		su.SetDuration(*i)
	}
	return su
}

// AddDuration adds i to the "duration" field.
func (su *ScheduleUpdate) AddDuration(i int) *ScheduleUpdate {
	su.mutation.AddDuration(i)
	return su
}

// SetClassID sets the "class" edge to the Class entity by ID.
func (su *ScheduleUpdate) SetClassID(id uuid.UUID) *ScheduleUpdate {
	su.mutation.SetClassID(id)
	return su
}

// SetClass sets the "class" edge to the Class entity.
func (su *ScheduleUpdate) SetClass(c *Class) *ScheduleUpdate {
	return su.SetClassID(c.ID)
}

// Mutation returns the ScheduleMutation object of the builder.
func (su *ScheduleUpdate) Mutation() *ScheduleMutation {
	return su.mutation
}

// ClearClass clears the "class" edge to the Class entity.
func (su *ScheduleUpdate) ClearClass() *ScheduleUpdate {
	su.mutation.ClearClass()
	return su
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *ScheduleUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(su.hooks) == 0 {
		if err = su.check(); err != nil {
			return 0, err
		}
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ScheduleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = su.check(); err != nil {
				return 0, err
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			if su.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *ScheduleUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *ScheduleUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *ScheduleUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *ScheduleUpdate) check() error {
	if v, ok := su.mutation.Weekday(); ok {
		if err := schedule.WeekdayValidator(v); err != nil {
			return &ValidationError{Name: "weekday", err: fmt.Errorf("ent: validator failed for field \"weekday\": %w", err)}
		}
	}
	if _, ok := su.mutation.ClassID(); su.mutation.ClassCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"class\"")
	}
	return nil
}

func (su *ScheduleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   schedule.Table,
			Columns: schedule.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: schedule.FieldID,
			},
		},
	}
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Weekday(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: schedule.FieldWeekday,
		})
	}
	if value, ok := su.mutation.AddedWeekday(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: schedule.FieldWeekday,
		})
	}
	if value, ok := su.mutation.StartsAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: schedule.FieldStartsAt,
		})
	}
	if value, ok := su.mutation.Duration(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: schedule.FieldDuration,
		})
	}
	if value, ok := su.mutation.AddedDuration(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: schedule.FieldDuration,
		})
	}
	if su.mutation.ClassCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.ClassIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{schedule.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ScheduleUpdateOne is the builder for updating a single Schedule entity.
type ScheduleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ScheduleMutation
}

// SetWeekday sets the "weekday" field.
func (suo *ScheduleUpdateOne) SetWeekday(i int) *ScheduleUpdateOne {
	suo.mutation.ResetWeekday()
	suo.mutation.SetWeekday(i)
	return suo
}

// AddWeekday adds i to the "weekday" field.
func (suo *ScheduleUpdateOne) AddWeekday(i int) *ScheduleUpdateOne {
	suo.mutation.AddWeekday(i)
	return suo
}

// SetStartsAt sets the "starts_at" field.
func (suo *ScheduleUpdateOne) SetStartsAt(t time.Time) *ScheduleUpdateOne {
	suo.mutation.SetStartsAt(t)
	return suo
}

// SetDuration sets the "duration" field.
func (suo *ScheduleUpdateOne) SetDuration(i int) *ScheduleUpdateOne {
	suo.mutation.ResetDuration()
	suo.mutation.SetDuration(i)
	return suo
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (suo *ScheduleUpdateOne) SetNillableDuration(i *int) *ScheduleUpdateOne {
	if i != nil {
		suo.SetDuration(*i)
	}
	return suo
}

// AddDuration adds i to the "duration" field.
func (suo *ScheduleUpdateOne) AddDuration(i int) *ScheduleUpdateOne {
	suo.mutation.AddDuration(i)
	return suo
}

// SetClassID sets the "class" edge to the Class entity by ID.
func (suo *ScheduleUpdateOne) SetClassID(id uuid.UUID) *ScheduleUpdateOne {
	suo.mutation.SetClassID(id)
	return suo
}

// SetClass sets the "class" edge to the Class entity.
func (suo *ScheduleUpdateOne) SetClass(c *Class) *ScheduleUpdateOne {
	return suo.SetClassID(c.ID)
}

// Mutation returns the ScheduleMutation object of the builder.
func (suo *ScheduleUpdateOne) Mutation() *ScheduleMutation {
	return suo.mutation
}

// ClearClass clears the "class" edge to the Class entity.
func (suo *ScheduleUpdateOne) ClearClass() *ScheduleUpdateOne {
	suo.mutation.ClearClass()
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *ScheduleUpdateOne) Select(field string, fields ...string) *ScheduleUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Schedule entity.
func (suo *ScheduleUpdateOne) Save(ctx context.Context) (*Schedule, error) {
	var (
		err  error
		node *Schedule
	)
	if len(suo.hooks) == 0 {
		if err = suo.check(); err != nil {
			return nil, err
		}
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ScheduleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = suo.check(); err != nil {
				return nil, err
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			if suo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = suo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *ScheduleUpdateOne) SaveX(ctx context.Context) *Schedule {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *ScheduleUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *ScheduleUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *ScheduleUpdateOne) check() error {
	if v, ok := suo.mutation.Weekday(); ok {
		if err := schedule.WeekdayValidator(v); err != nil {
			return &ValidationError{Name: "weekday", err: fmt.Errorf("ent: validator failed for field \"weekday\": %w", err)}
		}
	}
	if _, ok := suo.mutation.ClassID(); suo.mutation.ClassCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"class\"")
	}
	return nil
}

func (suo *ScheduleUpdateOne) sqlSave(ctx context.Context) (_node *Schedule, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   schedule.Table,
			Columns: schedule.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: schedule.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Schedule.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, schedule.FieldID)
		for _, f := range fields {
			if !schedule.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != schedule.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Weekday(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: schedule.FieldWeekday,
		})
	}
	if value, ok := suo.mutation.AddedWeekday(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: schedule.FieldWeekday,
		})
	}
	if value, ok := suo.mutation.StartsAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: schedule.FieldStartsAt,
		})
	}
	if value, ok := suo.mutation.Duration(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: schedule.FieldDuration,
		})
	}
	if value, ok := suo.mutation.AddedDuration(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: schedule.FieldDuration,
		})
	}
	if suo.mutation.ClassCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.ClassIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Schedule{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{schedule.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
