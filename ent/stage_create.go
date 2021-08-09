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
	"github.com/msal4/hassah_school_server/ent/schema"
	"github.com/msal4/hassah_school_server/ent/school"
	"github.com/msal4/hassah_school_server/ent/stage"
	"github.com/msal4/hassah_school_server/ent/tuitionpayment"
	"github.com/msal4/hassah_school_server/ent/user"
)

// StageCreate is the builder for creating a Stage entity.
type StageCreate struct {
	config
	mutation *StageMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (sc *StageCreate) SetCreatedAt(t time.Time) *StageCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *StageCreate) SetNillableCreatedAt(t *time.Time) *StageCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetUpdatedAt sets the "updated_at" field.
func (sc *StageCreate) SetUpdatedAt(t time.Time) *StageCreate {
	sc.mutation.SetUpdatedAt(t)
	return sc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sc *StageCreate) SetNillableUpdatedAt(t *time.Time) *StageCreate {
	if t != nil {
		sc.SetUpdatedAt(*t)
	}
	return sc
}

// SetName sets the "name" field.
func (sc *StageCreate) SetName(s string) *StageCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetTuitionAmount sets the "tuition_amount" field.
func (sc *StageCreate) SetTuitionAmount(i int) *StageCreate {
	sc.mutation.SetTuitionAmount(i)
	return sc
}

// SetStatus sets the "status" field.
func (sc *StageCreate) SetStatus(s schema.Status) *StageCreate {
	sc.mutation.SetStatus(s)
	return sc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (sc *StageCreate) SetNillableStatus(s *schema.Status) *StageCreate {
	if s != nil {
		sc.SetStatus(*s)
	}
	return sc
}

// SetID sets the "id" field.
func (sc *StageCreate) SetID(u uuid.UUID) *StageCreate {
	sc.mutation.SetID(u)
	return sc
}

// SetSchoolID sets the "school" edge to the School entity by ID.
func (sc *StageCreate) SetSchoolID(id uuid.UUID) *StageCreate {
	sc.mutation.SetSchoolID(id)
	return sc
}

// SetNillableSchoolID sets the "school" edge to the School entity by ID if the given value is not nil.
func (sc *StageCreate) SetNillableSchoolID(id *uuid.UUID) *StageCreate {
	if id != nil {
		sc = sc.SetSchoolID(*id)
	}
	return sc
}

// SetSchool sets the "school" edge to the School entity.
func (sc *StageCreate) SetSchool(s *School) *StageCreate {
	return sc.SetSchoolID(s.ID)
}

// AddClassIDs adds the "classes" edge to the Class entity by IDs.
func (sc *StageCreate) AddClassIDs(ids ...uuid.UUID) *StageCreate {
	sc.mutation.AddClassIDs(ids...)
	return sc
}

// AddClasses adds the "classes" edges to the Class entity.
func (sc *StageCreate) AddClasses(c ...*Class) *StageCreate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return sc.AddClassIDs(ids...)
}

// AddPaymentIDs adds the "payments" edge to the TuitionPayment entity by IDs.
func (sc *StageCreate) AddPaymentIDs(ids ...uuid.UUID) *StageCreate {
	sc.mutation.AddPaymentIDs(ids...)
	return sc
}

// AddPayments adds the "payments" edges to the TuitionPayment entity.
func (sc *StageCreate) AddPayments(t ...*TuitionPayment) *StageCreate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return sc.AddPaymentIDs(ids...)
}

// AddStudentIDs adds the "students" edge to the User entity by IDs.
func (sc *StageCreate) AddStudentIDs(ids ...uuid.UUID) *StageCreate {
	sc.mutation.AddStudentIDs(ids...)
	return sc
}

// AddStudents adds the "students" edges to the User entity.
func (sc *StageCreate) AddStudents(u ...*User) *StageCreate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return sc.AddStudentIDs(ids...)
}

// Mutation returns the StageMutation object of the builder.
func (sc *StageCreate) Mutation() *StageMutation {
	return sc.mutation
}

// Save creates the Stage in the database.
func (sc *StageCreate) Save(ctx context.Context) (*Stage, error) {
	var (
		err  error
		node *Stage
	)
	sc.defaults()
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StageMutation)
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
func (sc *StageCreate) SaveX(ctx context.Context) *Stage {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *StageCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *StageCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *StageCreate) defaults() {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		v := stage.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		v := stage.DefaultUpdatedAt()
		sc.mutation.SetUpdatedAt(v)
	}
	if _, ok := sc.mutation.Status(); !ok {
		v := stage.DefaultStatus
		sc.mutation.SetStatus(v)
	}
	if _, ok := sc.mutation.ID(); !ok {
		v := stage.DefaultID()
		sc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *StageCreate) check() error {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if v, ok := sc.mutation.Name(); ok {
		if err := stage.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "name": %w`, err)}
		}
	}
	if _, ok := sc.mutation.TuitionAmount(); !ok {
		return &ValidationError{Name: "tuition_amount", err: errors.New(`ent: missing required field "tuition_amount"`)}
	}
	if _, ok := sc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "status"`)}
	}
	if v, ok := sc.mutation.Status(); ok {
		if err := stage.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "status": %w`, err)}
		}
	}
	return nil
}

func (sc *StageCreate) sqlSave(ctx context.Context) (*Stage, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}

func (sc *StageCreate) createSpec() (*Stage, *sqlgraph.CreateSpec) {
	var (
		_node = &Stage{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: stage.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: stage.FieldID,
			},
		}
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: stage.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: stage.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := sc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: stage.FieldName,
		})
		_node.Name = value
	}
	if value, ok := sc.mutation.TuitionAmount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: stage.FieldTuitionAmount,
		})
		_node.TuitionAmount = value
	}
	if value, ok := sc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: stage.FieldStatus,
		})
		_node.Status = value
	}
	if nodes := sc.mutation.SchoolIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   stage.SchoolTable,
			Columns: []string{stage.SchoolColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: school.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.school_stages = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.ClassesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   stage.ClassesTable,
			Columns: []string{stage.ClassesColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.PaymentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   stage.PaymentsTable,
			Columns: []string{stage.PaymentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: tuitionpayment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.StudentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   stage.StudentsTable,
			Columns: []string{stage.StudentsColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// StageCreateBulk is the builder for creating many Stage entities in bulk.
type StageCreateBulk struct {
	config
	builders []*StageCreate
}

// Save creates the Stage entities in the database.
func (scb *StageCreateBulk) Save(ctx context.Context) ([]*Stage, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Stage, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StageMutation)
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
func (scb *StageCreateBulk) SaveX(ctx context.Context) []*Stage {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *StageCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *StageCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
