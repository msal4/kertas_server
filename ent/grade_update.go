// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/msal4/hassah_school_server/ent/assignment"
	"github.com/msal4/hassah_school_server/ent/grade"
	"github.com/msal4/hassah_school_server/ent/predicate"
	"github.com/msal4/hassah_school_server/ent/user"
)

// GradeUpdate is the builder for updating Grade entities.
type GradeUpdate struct {
	config
	hooks    []Hook
	mutation *GradeMutation
}

// Where appends a list predicates to the GradeUpdate builder.
func (gu *GradeUpdate) Where(ps ...predicate.Grade) *GradeUpdate {
	gu.mutation.Where(ps...)
	return gu
}

// SetExamGrade sets the "exam_grade" field.
func (gu *GradeUpdate) SetExamGrade(f float64) *GradeUpdate {
	gu.mutation.ResetExamGrade()
	gu.mutation.SetExamGrade(f)
	return gu
}

// AddExamGrade adds f to the "exam_grade" field.
func (gu *GradeUpdate) AddExamGrade(f float64) *GradeUpdate {
	gu.mutation.AddExamGrade(f)
	return gu
}

// SetStudentID sets the "student" edge to the User entity by ID.
func (gu *GradeUpdate) SetStudentID(id int) *GradeUpdate {
	gu.mutation.SetStudentID(id)
	return gu
}

// SetStudent sets the "student" edge to the User entity.
func (gu *GradeUpdate) SetStudent(u *User) *GradeUpdate {
	return gu.SetStudentID(u.ID)
}

// SetExamID sets the "exam" edge to the Assignment entity by ID.
func (gu *GradeUpdate) SetExamID(id int) *GradeUpdate {
	gu.mutation.SetExamID(id)
	return gu
}

// SetExam sets the "exam" edge to the Assignment entity.
func (gu *GradeUpdate) SetExam(a *Assignment) *GradeUpdate {
	return gu.SetExamID(a.ID)
}

// Mutation returns the GradeMutation object of the builder.
func (gu *GradeUpdate) Mutation() *GradeMutation {
	return gu.mutation
}

// ClearStudent clears the "student" edge to the User entity.
func (gu *GradeUpdate) ClearStudent() *GradeUpdate {
	gu.mutation.ClearStudent()
	return gu
}

// ClearExam clears the "exam" edge to the Assignment entity.
func (gu *GradeUpdate) ClearExam() *GradeUpdate {
	gu.mutation.ClearExam()
	return gu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gu *GradeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	gu.defaults()
	if len(gu.hooks) == 0 {
		if err = gu.check(); err != nil {
			return 0, err
		}
		affected, err = gu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GradeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gu.check(); err != nil {
				return 0, err
			}
			gu.mutation = mutation
			affected, err = gu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gu.hooks) - 1; i >= 0; i-- {
			if gu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GradeUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GradeUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GradeUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gu *GradeUpdate) defaults() {
	if _, ok := gu.mutation.UpdateTime(); !ok {
		v := grade.UpdateDefaultUpdateTime()
		gu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gu *GradeUpdate) check() error {
	if v, ok := gu.mutation.ExamGrade(); ok {
		if err := grade.ExamGradeValidator(v); err != nil {
			return &ValidationError{Name: "exam_grade", err: fmt.Errorf("ent: validator failed for field \"exam_grade\": %w", err)}
		}
	}
	if _, ok := gu.mutation.StudentID(); gu.mutation.StudentCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"student\"")
	}
	if _, ok := gu.mutation.ExamID(); gu.mutation.ExamCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"exam\"")
	}
	return nil
}

func (gu *GradeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   grade.Table,
			Columns: grade.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: grade.FieldID,
			},
		},
	}
	if ps := gu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: grade.FieldUpdateTime,
		})
	}
	if value, ok := gu.mutation.ExamGrade(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: grade.FieldExamGrade,
		})
	}
	if value, ok := gu.mutation.AddedExamGrade(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: grade.FieldExamGrade,
		})
	}
	if gu.mutation.StudentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.StudentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gu.mutation.ExamCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.ExamIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{grade.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// GradeUpdateOne is the builder for updating a single Grade entity.
type GradeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GradeMutation
}

// SetExamGrade sets the "exam_grade" field.
func (guo *GradeUpdateOne) SetExamGrade(f float64) *GradeUpdateOne {
	guo.mutation.ResetExamGrade()
	guo.mutation.SetExamGrade(f)
	return guo
}

// AddExamGrade adds f to the "exam_grade" field.
func (guo *GradeUpdateOne) AddExamGrade(f float64) *GradeUpdateOne {
	guo.mutation.AddExamGrade(f)
	return guo
}

// SetStudentID sets the "student" edge to the User entity by ID.
func (guo *GradeUpdateOne) SetStudentID(id int) *GradeUpdateOne {
	guo.mutation.SetStudentID(id)
	return guo
}

// SetStudent sets the "student" edge to the User entity.
func (guo *GradeUpdateOne) SetStudent(u *User) *GradeUpdateOne {
	return guo.SetStudentID(u.ID)
}

// SetExamID sets the "exam" edge to the Assignment entity by ID.
func (guo *GradeUpdateOne) SetExamID(id int) *GradeUpdateOne {
	guo.mutation.SetExamID(id)
	return guo
}

// SetExam sets the "exam" edge to the Assignment entity.
func (guo *GradeUpdateOne) SetExam(a *Assignment) *GradeUpdateOne {
	return guo.SetExamID(a.ID)
}

// Mutation returns the GradeMutation object of the builder.
func (guo *GradeUpdateOne) Mutation() *GradeMutation {
	return guo.mutation
}

// ClearStudent clears the "student" edge to the User entity.
func (guo *GradeUpdateOne) ClearStudent() *GradeUpdateOne {
	guo.mutation.ClearStudent()
	return guo
}

// ClearExam clears the "exam" edge to the Assignment entity.
func (guo *GradeUpdateOne) ClearExam() *GradeUpdateOne {
	guo.mutation.ClearExam()
	return guo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (guo *GradeUpdateOne) Select(field string, fields ...string) *GradeUpdateOne {
	guo.fields = append([]string{field}, fields...)
	return guo
}

// Save executes the query and returns the updated Grade entity.
func (guo *GradeUpdateOne) Save(ctx context.Context) (*Grade, error) {
	var (
		err  error
		node *Grade
	)
	guo.defaults()
	if len(guo.hooks) == 0 {
		if err = guo.check(); err != nil {
			return nil, err
		}
		node, err = guo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GradeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = guo.check(); err != nil {
				return nil, err
			}
			guo.mutation = mutation
			node, err = guo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(guo.hooks) - 1; i >= 0; i-- {
			if guo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = guo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, guo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (guo *GradeUpdateOne) SaveX(ctx context.Context) *Grade {
	node, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (guo *GradeUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GradeUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (guo *GradeUpdateOne) defaults() {
	if _, ok := guo.mutation.UpdateTime(); !ok {
		v := grade.UpdateDefaultUpdateTime()
		guo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (guo *GradeUpdateOne) check() error {
	if v, ok := guo.mutation.ExamGrade(); ok {
		if err := grade.ExamGradeValidator(v); err != nil {
			return &ValidationError{Name: "exam_grade", err: fmt.Errorf("ent: validator failed for field \"exam_grade\": %w", err)}
		}
	}
	if _, ok := guo.mutation.StudentID(); guo.mutation.StudentCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"student\"")
	}
	if _, ok := guo.mutation.ExamID(); guo.mutation.ExamCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"exam\"")
	}
	return nil
}

func (guo *GradeUpdateOne) sqlSave(ctx context.Context) (_node *Grade, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   grade.Table,
			Columns: grade.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: grade.FieldID,
			},
		},
	}
	id, ok := guo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Grade.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := guo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, grade.FieldID)
		for _, f := range fields {
			if !grade.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != grade.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := guo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := guo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: grade.FieldUpdateTime,
		})
	}
	if value, ok := guo.mutation.ExamGrade(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: grade.FieldExamGrade,
		})
	}
	if value, ok := guo.mutation.AddedExamGrade(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: grade.FieldExamGrade,
		})
	}
	if guo.mutation.StudentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.StudentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if guo.mutation.ExamCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.ExamIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Grade{config: guo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, guo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{grade.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
