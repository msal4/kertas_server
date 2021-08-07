// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/msal4/hassah_school_server/ent/assignment"
	"github.com/msal4/hassah_school_server/ent/grade"
	"github.com/msal4/hassah_school_server/ent/user"
)

// GradeCreate is the builder for creating a Grade entity.
type GradeCreate struct {
	config
	mutation *GradeMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (gc *GradeCreate) SetCreateTime(t time.Time) *GradeCreate {
	gc.mutation.SetCreateTime(t)
	return gc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (gc *GradeCreate) SetNillableCreateTime(t *time.Time) *GradeCreate {
	if t != nil {
		gc.SetCreateTime(*t)
	}
	return gc
}

// SetUpdateTime sets the "update_time" field.
func (gc *GradeCreate) SetUpdateTime(t time.Time) *GradeCreate {
	gc.mutation.SetUpdateTime(t)
	return gc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (gc *GradeCreate) SetNillableUpdateTime(t *time.Time) *GradeCreate {
	if t != nil {
		gc.SetUpdateTime(*t)
	}
	return gc
}

// SetExamGrade sets the "exam_grade" field.
func (gc *GradeCreate) SetExamGrade(f float64) *GradeCreate {
	gc.mutation.SetExamGrade(f)
	return gc
}

// SetStudentID sets the "student" edge to the User entity by ID.
func (gc *GradeCreate) SetStudentID(id int) *GradeCreate {
	gc.mutation.SetStudentID(id)
	return gc
}

// SetStudent sets the "student" edge to the User entity.
func (gc *GradeCreate) SetStudent(u *User) *GradeCreate {
	return gc.SetStudentID(u.ID)
}

// SetExamID sets the "exam" edge to the Assignment entity by ID.
func (gc *GradeCreate) SetExamID(id int) *GradeCreate {
	gc.mutation.SetExamID(id)
	return gc
}

// SetExam sets the "exam" edge to the Assignment entity.
func (gc *GradeCreate) SetExam(a *Assignment) *GradeCreate {
	return gc.SetExamID(a.ID)
}

// Mutation returns the GradeMutation object of the builder.
func (gc *GradeCreate) Mutation() *GradeMutation {
	return gc.mutation
}

// Save creates the Grade in the database.
func (gc *GradeCreate) Save(ctx context.Context) (*Grade, error) {
	var (
		err  error
		node *Grade
	)
	gc.defaults()
	if len(gc.hooks) == 0 {
		if err = gc.check(); err != nil {
			return nil, err
		}
		node, err = gc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GradeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gc.check(); err != nil {
				return nil, err
			}
			gc.mutation = mutation
			if node, err = gc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(gc.hooks) - 1; i >= 0; i-- {
			if gc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (gc *GradeCreate) SaveX(ctx context.Context) *Grade {
	v, err := gc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gc *GradeCreate) Exec(ctx context.Context) error {
	_, err := gc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gc *GradeCreate) ExecX(ctx context.Context) {
	if err := gc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gc *GradeCreate) defaults() {
	if _, ok := gc.mutation.CreateTime(); !ok {
		v := grade.DefaultCreateTime()
		gc.mutation.SetCreateTime(v)
	}
	if _, ok := gc.mutation.UpdateTime(); !ok {
		v := grade.DefaultUpdateTime()
		gc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gc *GradeCreate) check() error {
	if _, ok := gc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "create_time"`)}
	}
	if _, ok := gc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "update_time"`)}
	}
	if _, ok := gc.mutation.ExamGrade(); !ok {
		return &ValidationError{Name: "exam_grade", err: errors.New(`ent: missing required field "exam_grade"`)}
	}
	if v, ok := gc.mutation.ExamGrade(); ok {
		if err := grade.ExamGradeValidator(v); err != nil {
			return &ValidationError{Name: "exam_grade", err: fmt.Errorf(`ent: validator failed for field "exam_grade": %w`, err)}
		}
	}
	if _, ok := gc.mutation.StudentID(); !ok {
		return &ValidationError{Name: "student", err: errors.New("ent: missing required edge \"student\"")}
	}
	if _, ok := gc.mutation.ExamID(); !ok {
		return &ValidationError{Name: "exam", err: errors.New("ent: missing required edge \"exam\"")}
	}
	return nil
}

func (gc *GradeCreate) sqlSave(ctx context.Context) (*Grade, error) {
	_node, _spec := gc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (gc *GradeCreate) createSpec() (*Grade, *sqlgraph.CreateSpec) {
	var (
		_node = &Grade{config: gc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: grade.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: grade.FieldID,
			},
		}
	)
	if value, ok := gc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: grade.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := gc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: grade.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := gc.mutation.ExamGrade(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: grade.FieldExamGrade,
		})
		_node.ExamGrade = value
	}
	if nodes := gc.mutation.StudentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   grade.StudentTable,
			Columns: []string{grade.StudentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_grades = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := gc.mutation.ExamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   grade.ExamTable,
			Columns: []string{grade.ExamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: assignment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.assignment_grades = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// GradeCreateBulk is the builder for creating many Grade entities in bulk.
type GradeCreateBulk struct {
	config
	builders []*GradeCreate
}

// Save creates the Grade entities in the database.
func (gcb *GradeCreateBulk) Save(ctx context.Context) ([]*Grade, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gcb.builders))
	nodes := make([]*Grade, len(gcb.builders))
	mutators := make([]Mutator, len(gcb.builders))
	for i := range gcb.builders {
		func(i int, root context.Context) {
			builder := gcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GradeMutation)
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
					_, err = mutators[i+1].Mutate(root, gcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gcb.driver, spec); err != nil {
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
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, gcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gcb *GradeCreateBulk) SaveX(ctx context.Context) []*Grade {
	v, err := gcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gcb *GradeCreateBulk) Exec(ctx context.Context) error {
	_, err := gcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gcb *GradeCreateBulk) ExecX(ctx context.Context) {
	if err := gcb.Exec(ctx); err != nil {
		panic(err)
	}
}
