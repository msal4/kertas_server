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
	"github.com/msal4/hassah_school_server/ent/assignmentsubmission"
	"github.com/msal4/hassah_school_server/ent/user"
)

// AssignmentSubmissionCreate is the builder for creating a AssignmentSubmission entity.
type AssignmentSubmissionCreate struct {
	config
	mutation *AssignmentSubmissionMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (asc *AssignmentSubmissionCreate) SetCreatedAt(t time.Time) *AssignmentSubmissionCreate {
	asc.mutation.SetCreatedAt(t)
	return asc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (asc *AssignmentSubmissionCreate) SetNillableCreatedAt(t *time.Time) *AssignmentSubmissionCreate {
	if t != nil {
		asc.SetCreatedAt(*t)
	}
	return asc
}

// SetUpdatedAt sets the "updated_at" field.
func (asc *AssignmentSubmissionCreate) SetUpdatedAt(t time.Time) *AssignmentSubmissionCreate {
	asc.mutation.SetUpdatedAt(t)
	return asc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (asc *AssignmentSubmissionCreate) SetNillableUpdatedAt(t *time.Time) *AssignmentSubmissionCreate {
	if t != nil {
		asc.SetUpdatedAt(*t)
	}
	return asc
}

// SetFiles sets the "files" field.
func (asc *AssignmentSubmissionCreate) SetFiles(s []string) *AssignmentSubmissionCreate {
	asc.mutation.SetFiles(s)
	return asc
}

// SetSubmittedAt sets the "submitted_at" field.
func (asc *AssignmentSubmissionCreate) SetSubmittedAt(t time.Time) *AssignmentSubmissionCreate {
	asc.mutation.SetSubmittedAt(t)
	return asc
}

// SetNillableSubmittedAt sets the "submitted_at" field if the given value is not nil.
func (asc *AssignmentSubmissionCreate) SetNillableSubmittedAt(t *time.Time) *AssignmentSubmissionCreate {
	if t != nil {
		asc.SetSubmittedAt(*t)
	}
	return asc
}

// SetID sets the "id" field.
func (asc *AssignmentSubmissionCreate) SetID(u uuid.UUID) *AssignmentSubmissionCreate {
	asc.mutation.SetID(u)
	return asc
}

// SetStudentID sets the "student" edge to the User entity by ID.
func (asc *AssignmentSubmissionCreate) SetStudentID(id uuid.UUID) *AssignmentSubmissionCreate {
	asc.mutation.SetStudentID(id)
	return asc
}

// SetStudent sets the "student" edge to the User entity.
func (asc *AssignmentSubmissionCreate) SetStudent(u *User) *AssignmentSubmissionCreate {
	return asc.SetStudentID(u.ID)
}

// SetAssignmentID sets the "assignment" edge to the Assignment entity by ID.
func (asc *AssignmentSubmissionCreate) SetAssignmentID(id uuid.UUID) *AssignmentSubmissionCreate {
	asc.mutation.SetAssignmentID(id)
	return asc
}

// SetAssignment sets the "assignment" edge to the Assignment entity.
func (asc *AssignmentSubmissionCreate) SetAssignment(a *Assignment) *AssignmentSubmissionCreate {
	return asc.SetAssignmentID(a.ID)
}

// Mutation returns the AssignmentSubmissionMutation object of the builder.
func (asc *AssignmentSubmissionCreate) Mutation() *AssignmentSubmissionMutation {
	return asc.mutation
}

// Save creates the AssignmentSubmission in the database.
func (asc *AssignmentSubmissionCreate) Save(ctx context.Context) (*AssignmentSubmission, error) {
	var (
		err  error
		node *AssignmentSubmission
	)
	asc.defaults()
	if len(asc.hooks) == 0 {
		if err = asc.check(); err != nil {
			return nil, err
		}
		node, err = asc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AssignmentSubmissionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = asc.check(); err != nil {
				return nil, err
			}
			asc.mutation = mutation
			if node, err = asc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(asc.hooks) - 1; i >= 0; i-- {
			if asc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = asc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, asc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (asc *AssignmentSubmissionCreate) SaveX(ctx context.Context) *AssignmentSubmission {
	v, err := asc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (asc *AssignmentSubmissionCreate) Exec(ctx context.Context) error {
	_, err := asc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (asc *AssignmentSubmissionCreate) ExecX(ctx context.Context) {
	if err := asc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (asc *AssignmentSubmissionCreate) defaults() {
	if _, ok := asc.mutation.CreatedAt(); !ok {
		v := assignmentsubmission.DefaultCreatedAt()
		asc.mutation.SetCreatedAt(v)
	}
	if _, ok := asc.mutation.UpdatedAt(); !ok {
		v := assignmentsubmission.DefaultUpdatedAt()
		asc.mutation.SetUpdatedAt(v)
	}
	if _, ok := asc.mutation.ID(); !ok {
		v := assignmentsubmission.DefaultID()
		asc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (asc *AssignmentSubmissionCreate) check() error {
	if _, ok := asc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := asc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := asc.mutation.Files(); !ok {
		return &ValidationError{Name: "files", err: errors.New(`ent: missing required field "files"`)}
	}
	if _, ok := asc.mutation.StudentID(); !ok {
		return &ValidationError{Name: "student", err: errors.New("ent: missing required edge \"student\"")}
	}
	if _, ok := asc.mutation.AssignmentID(); !ok {
		return &ValidationError{Name: "assignment", err: errors.New("ent: missing required edge \"assignment\"")}
	}
	return nil
}

func (asc *AssignmentSubmissionCreate) sqlSave(ctx context.Context) (*AssignmentSubmission, error) {
	_node, _spec := asc.createSpec()
	if err := sqlgraph.CreateNode(ctx, asc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}

func (asc *AssignmentSubmissionCreate) createSpec() (*AssignmentSubmission, *sqlgraph.CreateSpec) {
	var (
		_node = &AssignmentSubmission{config: asc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: assignmentsubmission.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: assignmentsubmission.FieldID,
			},
		}
	)
	if id, ok := asc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := asc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: assignmentsubmission.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := asc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: assignmentsubmission.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := asc.mutation.Files(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: assignmentsubmission.FieldFiles,
		})
		_node.Files = value
	}
	if value, ok := asc.mutation.SubmittedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: assignmentsubmission.FieldSubmittedAt,
		})
		_node.SubmittedAt = &value
	}
	if nodes := asc.mutation.StudentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   assignmentsubmission.StudentTable,
			Columns: []string{assignmentsubmission.StudentColumn},
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
		_node.user_submissions = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := asc.mutation.AssignmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   assignmentsubmission.AssignmentTable,
			Columns: []string{assignmentsubmission.AssignmentColumn},
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
		_node.assignment_submissions = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AssignmentSubmissionCreateBulk is the builder for creating many AssignmentSubmission entities in bulk.
type AssignmentSubmissionCreateBulk struct {
	config
	builders []*AssignmentSubmissionCreate
}

// Save creates the AssignmentSubmission entities in the database.
func (ascb *AssignmentSubmissionCreateBulk) Save(ctx context.Context) ([]*AssignmentSubmission, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ascb.builders))
	nodes := make([]*AssignmentSubmission, len(ascb.builders))
	mutators := make([]Mutator, len(ascb.builders))
	for i := range ascb.builders {
		func(i int, root context.Context) {
			builder := ascb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AssignmentSubmissionMutation)
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
					_, err = mutators[i+1].Mutate(root, ascb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ascb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ascb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ascb *AssignmentSubmissionCreateBulk) SaveX(ctx context.Context) []*AssignmentSubmission {
	v, err := ascb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ascb *AssignmentSubmissionCreateBulk) Exec(ctx context.Context) error {
	_, err := ascb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ascb *AssignmentSubmissionCreateBulk) ExecX(ctx context.Context) {
	if err := ascb.Exec(ctx); err != nil {
		panic(err)
	}
}
