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
	"github.com/msal4/hassah_school_server/ent/stage"
	"github.com/msal4/hassah_school_server/ent/tuitionpayment"
	"github.com/msal4/hassah_school_server/ent/user"
)

// TuitionPaymentCreate is the builder for creating a TuitionPayment entity.
type TuitionPaymentCreate struct {
	config
	mutation *TuitionPaymentMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (tpc *TuitionPaymentCreate) SetCreatedAt(t time.Time) *TuitionPaymentCreate {
	tpc.mutation.SetCreatedAt(t)
	return tpc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tpc *TuitionPaymentCreate) SetNillableCreatedAt(t *time.Time) *TuitionPaymentCreate {
	if t != nil {
		tpc.SetCreatedAt(*t)
	}
	return tpc
}

// SetUpdatedAt sets the "updated_at" field.
func (tpc *TuitionPaymentCreate) SetUpdatedAt(t time.Time) *TuitionPaymentCreate {
	tpc.mutation.SetUpdatedAt(t)
	return tpc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tpc *TuitionPaymentCreate) SetNillableUpdatedAt(t *time.Time) *TuitionPaymentCreate {
	if t != nil {
		tpc.SetUpdatedAt(*t)
	}
	return tpc
}

// SetYear sets the "year" field.
func (tpc *TuitionPaymentCreate) SetYear(s string) *TuitionPaymentCreate {
	tpc.mutation.SetYear(s)
	return tpc
}

// SetPaidAmount sets the "paid_amount" field.
func (tpc *TuitionPaymentCreate) SetPaidAmount(i int) *TuitionPaymentCreate {
	tpc.mutation.SetPaidAmount(i)
	return tpc
}

// SetID sets the "id" field.
func (tpc *TuitionPaymentCreate) SetID(u uuid.UUID) *TuitionPaymentCreate {
	tpc.mutation.SetID(u)
	return tpc
}

// SetStudentID sets the "student" edge to the User entity by ID.
func (tpc *TuitionPaymentCreate) SetStudentID(id uuid.UUID) *TuitionPaymentCreate {
	tpc.mutation.SetStudentID(id)
	return tpc
}

// SetStudent sets the "student" edge to the User entity.
func (tpc *TuitionPaymentCreate) SetStudent(u *User) *TuitionPaymentCreate {
	return tpc.SetStudentID(u.ID)
}

// SetStageID sets the "stage" edge to the Stage entity by ID.
func (tpc *TuitionPaymentCreate) SetStageID(id uuid.UUID) *TuitionPaymentCreate {
	tpc.mutation.SetStageID(id)
	return tpc
}

// SetStage sets the "stage" edge to the Stage entity.
func (tpc *TuitionPaymentCreate) SetStage(s *Stage) *TuitionPaymentCreate {
	return tpc.SetStageID(s.ID)
}

// Mutation returns the TuitionPaymentMutation object of the builder.
func (tpc *TuitionPaymentCreate) Mutation() *TuitionPaymentMutation {
	return tpc.mutation
}

// Save creates the TuitionPayment in the database.
func (tpc *TuitionPaymentCreate) Save(ctx context.Context) (*TuitionPayment, error) {
	var (
		err  error
		node *TuitionPayment
	)
	tpc.defaults()
	if len(tpc.hooks) == 0 {
		if err = tpc.check(); err != nil {
			return nil, err
		}
		node, err = tpc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TuitionPaymentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tpc.check(); err != nil {
				return nil, err
			}
			tpc.mutation = mutation
			if node, err = tpc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tpc.hooks) - 1; i >= 0; i-- {
			if tpc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tpc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tpc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tpc *TuitionPaymentCreate) SaveX(ctx context.Context) *TuitionPayment {
	v, err := tpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tpc *TuitionPaymentCreate) Exec(ctx context.Context) error {
	_, err := tpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tpc *TuitionPaymentCreate) ExecX(ctx context.Context) {
	if err := tpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tpc *TuitionPaymentCreate) defaults() {
	if _, ok := tpc.mutation.CreatedAt(); !ok {
		v := tuitionpayment.DefaultCreatedAt()
		tpc.mutation.SetCreatedAt(v)
	}
	if _, ok := tpc.mutation.UpdatedAt(); !ok {
		v := tuitionpayment.DefaultUpdatedAt()
		tpc.mutation.SetUpdatedAt(v)
	}
	if _, ok := tpc.mutation.ID(); !ok {
		v := tuitionpayment.DefaultID()
		tpc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tpc *TuitionPaymentCreate) check() error {
	if _, ok := tpc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := tpc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := tpc.mutation.Year(); !ok {
		return &ValidationError{Name: "year", err: errors.New(`ent: missing required field "year"`)}
	}
	if v, ok := tpc.mutation.Year(); ok {
		if err := tuitionpayment.YearValidator(v); err != nil {
			return &ValidationError{Name: "year", err: fmt.Errorf(`ent: validator failed for field "year": %w`, err)}
		}
	}
	if _, ok := tpc.mutation.PaidAmount(); !ok {
		return &ValidationError{Name: "paid_amount", err: errors.New(`ent: missing required field "paid_amount"`)}
	}
	if _, ok := tpc.mutation.StudentID(); !ok {
		return &ValidationError{Name: "student", err: errors.New("ent: missing required edge \"student\"")}
	}
	if _, ok := tpc.mutation.StageID(); !ok {
		return &ValidationError{Name: "stage", err: errors.New("ent: missing required edge \"stage\"")}
	}
	return nil
}

func (tpc *TuitionPaymentCreate) sqlSave(ctx context.Context) (*TuitionPayment, error) {
	_node, _spec := tpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tpc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		_node.ID = _spec.ID.Value.(uuid.UUID)
	}
	return _node, nil
}

func (tpc *TuitionPaymentCreate) createSpec() (*TuitionPayment, *sqlgraph.CreateSpec) {
	var (
		_node = &TuitionPayment{config: tpc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: tuitionpayment.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: tuitionpayment.FieldID,
			},
		}
	)
	if id, ok := tpc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tpc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tuitionpayment.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := tpc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tuitionpayment.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := tpc.mutation.Year(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tuitionpayment.FieldYear,
		})
		_node.Year = value
	}
	if value, ok := tpc.mutation.PaidAmount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tuitionpayment.FieldPaidAmount,
		})
		_node.PaidAmount = value
	}
	if nodes := tpc.mutation.StudentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tuitionpayment.StudentTable,
			Columns: []string{tuitionpayment.StudentColumn},
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
		_node.user_payments = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tpc.mutation.StageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tuitionpayment.StageTable,
			Columns: []string{tuitionpayment.StageColumn},
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
		_node.stage_payments = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TuitionPaymentCreateBulk is the builder for creating many TuitionPayment entities in bulk.
type TuitionPaymentCreateBulk struct {
	config
	builders []*TuitionPaymentCreate
}

// Save creates the TuitionPayment entities in the database.
func (tpcb *TuitionPaymentCreateBulk) Save(ctx context.Context) ([]*TuitionPayment, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tpcb.builders))
	nodes := make([]*TuitionPayment, len(tpcb.builders))
	mutators := make([]Mutator, len(tpcb.builders))
	for i := range tpcb.builders {
		func(i int, root context.Context) {
			builder := tpcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TuitionPaymentMutation)
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
					_, err = mutators[i+1].Mutate(root, tpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tpcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tpcb *TuitionPaymentCreateBulk) SaveX(ctx context.Context) []*TuitionPayment {
	v, err := tpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tpcb *TuitionPaymentCreateBulk) Exec(ctx context.Context) error {
	_, err := tpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tpcb *TuitionPaymentCreateBulk) ExecX(ctx context.Context) {
	if err := tpcb.Exec(ctx); err != nil {
		panic(err)
	}
}
