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
	"github.com/msal4/hassah_school_server/ent/schema"
	"github.com/msal4/hassah_school_server/ent/school"
	"github.com/msal4/hassah_school_server/ent/stage"
	"github.com/msal4/hassah_school_server/ent/user"
)

// SchoolCreate is the builder for creating a School entity.
type SchoolCreate struct {
	config
	mutation *SchoolMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (sc *SchoolCreate) SetCreatedAt(t time.Time) *SchoolCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *SchoolCreate) SetNillableCreatedAt(t *time.Time) *SchoolCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetUpdatedAt sets the "updated_at" field.
func (sc *SchoolCreate) SetUpdatedAt(t time.Time) *SchoolCreate {
	sc.mutation.SetUpdatedAt(t)
	return sc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sc *SchoolCreate) SetNillableUpdatedAt(t *time.Time) *SchoolCreate {
	if t != nil {
		sc.SetUpdatedAt(*t)
	}
	return sc
}

// SetName sets the "name" field.
func (sc *SchoolCreate) SetName(s string) *SchoolCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetImage sets the "image" field.
func (sc *SchoolCreate) SetImage(s string) *SchoolCreate {
	sc.mutation.SetImage(s)
	return sc
}

// SetDirectory sets the "directory" field.
func (sc *SchoolCreate) SetDirectory(s string) *SchoolCreate {
	sc.mutation.SetDirectory(s)
	return sc
}

// SetStatus sets the "status" field.
func (sc *SchoolCreate) SetStatus(s schema.Status) *SchoolCreate {
	sc.mutation.SetStatus(s)
	return sc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (sc *SchoolCreate) SetNillableStatus(s *schema.Status) *SchoolCreate {
	if s != nil {
		sc.SetStatus(*s)
	}
	return sc
}

// SetID sets the "id" field.
func (sc *SchoolCreate) SetID(u uuid.UUID) *SchoolCreate {
	sc.mutation.SetID(u)
	return sc
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (sc *SchoolCreate) AddUserIDs(ids ...uuid.UUID) *SchoolCreate {
	sc.mutation.AddUserIDs(ids...)
	return sc
}

// AddUsers adds the "users" edges to the User entity.
func (sc *SchoolCreate) AddUsers(u ...*User) *SchoolCreate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return sc.AddUserIDs(ids...)
}

// AddStageIDs adds the "stages" edge to the Stage entity by IDs.
func (sc *SchoolCreate) AddStageIDs(ids ...uuid.UUID) *SchoolCreate {
	sc.mutation.AddStageIDs(ids...)
	return sc
}

// AddStages adds the "stages" edges to the Stage entity.
func (sc *SchoolCreate) AddStages(s ...*Stage) *SchoolCreate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sc.AddStageIDs(ids...)
}

// Mutation returns the SchoolMutation object of the builder.
func (sc *SchoolCreate) Mutation() *SchoolMutation {
	return sc.mutation
}

// Save creates the School in the database.
func (sc *SchoolCreate) Save(ctx context.Context) (*School, error) {
	var (
		err  error
		node *School
	)
	sc.defaults()
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SchoolMutation)
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
func (sc *SchoolCreate) SaveX(ctx context.Context) *School {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SchoolCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SchoolCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SchoolCreate) defaults() {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		v := school.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		v := school.DefaultUpdatedAt()
		sc.mutation.SetUpdatedAt(v)
	}
	if _, ok := sc.mutation.Status(); !ok {
		v := school.DefaultStatus
		sc.mutation.SetStatus(v)
	}
	if _, ok := sc.mutation.ID(); !ok {
		v := school.DefaultID()
		sc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SchoolCreate) check() error {
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
		if err := school.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "name": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Image(); !ok {
		return &ValidationError{Name: "image", err: errors.New(`ent: missing required field "image"`)}
	}
	if v, ok := sc.mutation.Image(); ok {
		if err := school.ImageValidator(v); err != nil {
			return &ValidationError{Name: "image", err: fmt.Errorf(`ent: validator failed for field "image": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Directory(); !ok {
		return &ValidationError{Name: "directory", err: errors.New(`ent: missing required field "directory"`)}
	}
	if v, ok := sc.mutation.Directory(); ok {
		if err := school.DirectoryValidator(v); err != nil {
			return &ValidationError{Name: "directory", err: fmt.Errorf(`ent: validator failed for field "directory": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "status"`)}
	}
	if v, ok := sc.mutation.Status(); ok {
		if err := school.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "status": %w`, err)}
		}
	}
	return nil
}

func (sc *SchoolCreate) sqlSave(ctx context.Context) (*School, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}

func (sc *SchoolCreate) createSpec() (*School, *sqlgraph.CreateSpec) {
	var (
		_node = &School{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: school.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: school.FieldID,
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
			Column: school.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: school.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := sc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: school.FieldName,
		})
		_node.Name = value
	}
	if value, ok := sc.mutation.Image(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: school.FieldImage,
		})
		_node.Image = value
	}
	if value, ok := sc.mutation.Directory(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: school.FieldDirectory,
		})
		_node.Directory = value
	}
	if value, ok := sc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: school.FieldStatus,
		})
		_node.Status = value
	}
	if nodes := sc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   school.UsersTable,
			Columns: []string{school.UsersColumn},
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
	if nodes := sc.mutation.StagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   school.StagesTable,
			Columns: []string{school.StagesColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SchoolCreateBulk is the builder for creating many School entities in bulk.
type SchoolCreateBulk struct {
	config
	builders []*SchoolCreate
}

// Save creates the School entities in the database.
func (scb *SchoolCreateBulk) Save(ctx context.Context) ([]*School, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*School, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SchoolMutation)
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
func (scb *SchoolCreateBulk) SaveX(ctx context.Context) []*School {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SchoolCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SchoolCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
