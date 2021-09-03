// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/msal4/hassah_school_server/ent/class"
	"github.com/msal4/hassah_school_server/ent/coursegrade"
	"github.com/msal4/hassah_school_server/ent/predicate"
	"github.com/msal4/hassah_school_server/ent/stage"
	"github.com/msal4/hassah_school_server/ent/user"
)

// CourseGradeUpdate is the builder for updating CourseGrade entities.
type CourseGradeUpdate struct {
	config
	hooks    []Hook
	mutation *CourseGradeMutation
}

// Where appends a list predicates to the CourseGradeUpdate builder.
func (cgu *CourseGradeUpdate) Where(ps ...predicate.CourseGrade) *CourseGradeUpdate {
	cgu.mutation.Where(ps...)
	return cgu
}

// SetActivityFirst sets the "activity_first" field.
func (cgu *CourseGradeUpdate) SetActivityFirst(i int) *CourseGradeUpdate {
	cgu.mutation.ResetActivityFirst()
	cgu.mutation.SetActivityFirst(i)
	return cgu
}

// SetNillableActivityFirst sets the "activity_first" field if the given value is not nil.
func (cgu *CourseGradeUpdate) SetNillableActivityFirst(i *int) *CourseGradeUpdate {
	if i != nil {
		cgu.SetActivityFirst(*i)
	}
	return cgu
}

// AddActivityFirst adds i to the "activity_first" field.
func (cgu *CourseGradeUpdate) AddActivityFirst(i int) *CourseGradeUpdate {
	cgu.mutation.AddActivityFirst(i)
	return cgu
}

// ClearActivityFirst clears the value of the "activity_first" field.
func (cgu *CourseGradeUpdate) ClearActivityFirst() *CourseGradeUpdate {
	cgu.mutation.ClearActivityFirst()
	return cgu
}

// SetActivitySecond sets the "activity_second" field.
func (cgu *CourseGradeUpdate) SetActivitySecond(i int) *CourseGradeUpdate {
	cgu.mutation.ResetActivitySecond()
	cgu.mutation.SetActivitySecond(i)
	return cgu
}

// SetNillableActivitySecond sets the "activity_second" field if the given value is not nil.
func (cgu *CourseGradeUpdate) SetNillableActivitySecond(i *int) *CourseGradeUpdate {
	if i != nil {
		cgu.SetActivitySecond(*i)
	}
	return cgu
}

// AddActivitySecond adds i to the "activity_second" field.
func (cgu *CourseGradeUpdate) AddActivitySecond(i int) *CourseGradeUpdate {
	cgu.mutation.AddActivitySecond(i)
	return cgu
}

// ClearActivitySecond clears the value of the "activity_second" field.
func (cgu *CourseGradeUpdate) ClearActivitySecond() *CourseGradeUpdate {
	cgu.mutation.ClearActivitySecond()
	return cgu
}

// SetWrittenFirst sets the "written_first" field.
func (cgu *CourseGradeUpdate) SetWrittenFirst(i int) *CourseGradeUpdate {
	cgu.mutation.ResetWrittenFirst()
	cgu.mutation.SetWrittenFirst(i)
	return cgu
}

// SetNillableWrittenFirst sets the "written_first" field if the given value is not nil.
func (cgu *CourseGradeUpdate) SetNillableWrittenFirst(i *int) *CourseGradeUpdate {
	if i != nil {
		cgu.SetWrittenFirst(*i)
	}
	return cgu
}

// AddWrittenFirst adds i to the "written_first" field.
func (cgu *CourseGradeUpdate) AddWrittenFirst(i int) *CourseGradeUpdate {
	cgu.mutation.AddWrittenFirst(i)
	return cgu
}

// ClearWrittenFirst clears the value of the "written_first" field.
func (cgu *CourseGradeUpdate) ClearWrittenFirst() *CourseGradeUpdate {
	cgu.mutation.ClearWrittenFirst()
	return cgu
}

// SetWrittenSecond sets the "written_second" field.
func (cgu *CourseGradeUpdate) SetWrittenSecond(i int) *CourseGradeUpdate {
	cgu.mutation.ResetWrittenSecond()
	cgu.mutation.SetWrittenSecond(i)
	return cgu
}

// SetNillableWrittenSecond sets the "written_second" field if the given value is not nil.
func (cgu *CourseGradeUpdate) SetNillableWrittenSecond(i *int) *CourseGradeUpdate {
	if i != nil {
		cgu.SetWrittenSecond(*i)
	}
	return cgu
}

// AddWrittenSecond adds i to the "written_second" field.
func (cgu *CourseGradeUpdate) AddWrittenSecond(i int) *CourseGradeUpdate {
	cgu.mutation.AddWrittenSecond(i)
	return cgu
}

// ClearWrittenSecond clears the value of the "written_second" field.
func (cgu *CourseGradeUpdate) ClearWrittenSecond() *CourseGradeUpdate {
	cgu.mutation.ClearWrittenSecond()
	return cgu
}

// SetCourseFinal sets the "course_final" field.
func (cgu *CourseGradeUpdate) SetCourseFinal(i int) *CourseGradeUpdate {
	cgu.mutation.ResetCourseFinal()
	cgu.mutation.SetCourseFinal(i)
	return cgu
}

// SetNillableCourseFinal sets the "course_final" field if the given value is not nil.
func (cgu *CourseGradeUpdate) SetNillableCourseFinal(i *int) *CourseGradeUpdate {
	if i != nil {
		cgu.SetCourseFinal(*i)
	}
	return cgu
}

// AddCourseFinal adds i to the "course_final" field.
func (cgu *CourseGradeUpdate) AddCourseFinal(i int) *CourseGradeUpdate {
	cgu.mutation.AddCourseFinal(i)
	return cgu
}

// ClearCourseFinal clears the value of the "course_final" field.
func (cgu *CourseGradeUpdate) ClearCourseFinal() *CourseGradeUpdate {
	cgu.mutation.ClearCourseFinal()
	return cgu
}

// SetStudentID sets the "student" edge to the User entity by ID.
func (cgu *CourseGradeUpdate) SetStudentID(id uuid.UUID) *CourseGradeUpdate {
	cgu.mutation.SetStudentID(id)
	return cgu
}

// SetStudent sets the "student" edge to the User entity.
func (cgu *CourseGradeUpdate) SetStudent(u *User) *CourseGradeUpdate {
	return cgu.SetStudentID(u.ID)
}

// SetClassID sets the "class" edge to the Class entity by ID.
func (cgu *CourseGradeUpdate) SetClassID(id uuid.UUID) *CourseGradeUpdate {
	cgu.mutation.SetClassID(id)
	return cgu
}

// SetClass sets the "class" edge to the Class entity.
func (cgu *CourseGradeUpdate) SetClass(c *Class) *CourseGradeUpdate {
	return cgu.SetClassID(c.ID)
}

// SetStageID sets the "stage" edge to the Stage entity by ID.
func (cgu *CourseGradeUpdate) SetStageID(id uuid.UUID) *CourseGradeUpdate {
	cgu.mutation.SetStageID(id)
	return cgu
}

// SetStage sets the "stage" edge to the Stage entity.
func (cgu *CourseGradeUpdate) SetStage(s *Stage) *CourseGradeUpdate {
	return cgu.SetStageID(s.ID)
}

// Mutation returns the CourseGradeMutation object of the builder.
func (cgu *CourseGradeUpdate) Mutation() *CourseGradeMutation {
	return cgu.mutation
}

// ClearStudent clears the "student" edge to the User entity.
func (cgu *CourseGradeUpdate) ClearStudent() *CourseGradeUpdate {
	cgu.mutation.ClearStudent()
	return cgu
}

// ClearClass clears the "class" edge to the Class entity.
func (cgu *CourseGradeUpdate) ClearClass() *CourseGradeUpdate {
	cgu.mutation.ClearClass()
	return cgu
}

// ClearStage clears the "stage" edge to the Stage entity.
func (cgu *CourseGradeUpdate) ClearStage() *CourseGradeUpdate {
	cgu.mutation.ClearStage()
	return cgu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cgu *CourseGradeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	cgu.defaults()
	if len(cgu.hooks) == 0 {
		if err = cgu.check(); err != nil {
			return 0, err
		}
		affected, err = cgu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CourseGradeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cgu.check(); err != nil {
				return 0, err
			}
			cgu.mutation = mutation
			affected, err = cgu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cgu.hooks) - 1; i >= 0; i-- {
			if cgu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cgu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cgu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cgu *CourseGradeUpdate) SaveX(ctx context.Context) int {
	affected, err := cgu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cgu *CourseGradeUpdate) Exec(ctx context.Context) error {
	_, err := cgu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cgu *CourseGradeUpdate) ExecX(ctx context.Context) {
	if err := cgu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cgu *CourseGradeUpdate) defaults() {
	if _, ok := cgu.mutation.UpdatedAt(); !ok {
		v := coursegrade.UpdateDefaultUpdatedAt()
		cgu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cgu *CourseGradeUpdate) check() error {
	if v, ok := cgu.mutation.ActivityFirst(); ok {
		if err := coursegrade.ActivityFirstValidator(v); err != nil {
			return &ValidationError{Name: "activity_first", err: fmt.Errorf("ent: validator failed for field \"activity_first\": %w", err)}
		}
	}
	if v, ok := cgu.mutation.ActivitySecond(); ok {
		if err := coursegrade.ActivitySecondValidator(v); err != nil {
			return &ValidationError{Name: "activity_second", err: fmt.Errorf("ent: validator failed for field \"activity_second\": %w", err)}
		}
	}
	if v, ok := cgu.mutation.WrittenFirst(); ok {
		if err := coursegrade.WrittenFirstValidator(v); err != nil {
			return &ValidationError{Name: "written_first", err: fmt.Errorf("ent: validator failed for field \"written_first\": %w", err)}
		}
	}
	if v, ok := cgu.mutation.WrittenSecond(); ok {
		if err := coursegrade.WrittenSecondValidator(v); err != nil {
			return &ValidationError{Name: "written_second", err: fmt.Errorf("ent: validator failed for field \"written_second\": %w", err)}
		}
	}
	if v, ok := cgu.mutation.CourseFinal(); ok {
		if err := coursegrade.CourseFinalValidator(v); err != nil {
			return &ValidationError{Name: "course_final", err: fmt.Errorf("ent: validator failed for field \"course_final\": %w", err)}
		}
	}
	if _, ok := cgu.mutation.StudentID(); cgu.mutation.StudentCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"student\"")
	}
	if _, ok := cgu.mutation.ClassID(); cgu.mutation.ClassCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"class\"")
	}
	if _, ok := cgu.mutation.StageID(); cgu.mutation.StageCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"stage\"")
	}
	return nil
}

func (cgu *CourseGradeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   coursegrade.Table,
			Columns: coursegrade.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: coursegrade.FieldID,
			},
		},
	}
	if ps := cgu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cgu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: coursegrade.FieldUpdatedAt,
		})
	}
	if value, ok := cgu.mutation.ActivityFirst(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coursegrade.FieldActivityFirst,
		})
	}
	if value, ok := cgu.mutation.AddedActivityFirst(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coursegrade.FieldActivityFirst,
		})
	}
	if cgu.mutation.ActivityFirstCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: coursegrade.FieldActivityFirst,
		})
	}
	if value, ok := cgu.mutation.ActivitySecond(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coursegrade.FieldActivitySecond,
		})
	}
	if value, ok := cgu.mutation.AddedActivitySecond(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coursegrade.FieldActivitySecond,
		})
	}
	if cgu.mutation.ActivitySecondCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: coursegrade.FieldActivitySecond,
		})
	}
	if value, ok := cgu.mutation.WrittenFirst(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coursegrade.FieldWrittenFirst,
		})
	}
	if value, ok := cgu.mutation.AddedWrittenFirst(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coursegrade.FieldWrittenFirst,
		})
	}
	if cgu.mutation.WrittenFirstCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: coursegrade.FieldWrittenFirst,
		})
	}
	if value, ok := cgu.mutation.WrittenSecond(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coursegrade.FieldWrittenSecond,
		})
	}
	if value, ok := cgu.mutation.AddedWrittenSecond(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coursegrade.FieldWrittenSecond,
		})
	}
	if cgu.mutation.WrittenSecondCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: coursegrade.FieldWrittenSecond,
		})
	}
	if value, ok := cgu.mutation.CourseFinal(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coursegrade.FieldCourseFinal,
		})
	}
	if value, ok := cgu.mutation.AddedCourseFinal(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coursegrade.FieldCourseFinal,
		})
	}
	if cgu.mutation.CourseFinalCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: coursegrade.FieldCourseFinal,
		})
	}
	if cgu.mutation.StudentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   coursegrade.StudentTable,
			Columns: []string{coursegrade.StudentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cgu.mutation.StudentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   coursegrade.StudentTable,
			Columns: []string{coursegrade.StudentColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cgu.mutation.ClassCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   coursegrade.ClassTable,
			Columns: []string{coursegrade.ClassColumn},
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
	if nodes := cgu.mutation.ClassIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   coursegrade.ClassTable,
			Columns: []string{coursegrade.ClassColumn},
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
	if cgu.mutation.StageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   coursegrade.StageTable,
			Columns: []string{coursegrade.StageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: stage.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cgu.mutation.StageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   coursegrade.StageTable,
			Columns: []string{coursegrade.StageColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cgu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{coursegrade.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// CourseGradeUpdateOne is the builder for updating a single CourseGrade entity.
type CourseGradeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CourseGradeMutation
}

// SetActivityFirst sets the "activity_first" field.
func (cguo *CourseGradeUpdateOne) SetActivityFirst(i int) *CourseGradeUpdateOne {
	cguo.mutation.ResetActivityFirst()
	cguo.mutation.SetActivityFirst(i)
	return cguo
}

// SetNillableActivityFirst sets the "activity_first" field if the given value is not nil.
func (cguo *CourseGradeUpdateOne) SetNillableActivityFirst(i *int) *CourseGradeUpdateOne {
	if i != nil {
		cguo.SetActivityFirst(*i)
	}
	return cguo
}

// AddActivityFirst adds i to the "activity_first" field.
func (cguo *CourseGradeUpdateOne) AddActivityFirst(i int) *CourseGradeUpdateOne {
	cguo.mutation.AddActivityFirst(i)
	return cguo
}

// ClearActivityFirst clears the value of the "activity_first" field.
func (cguo *CourseGradeUpdateOne) ClearActivityFirst() *CourseGradeUpdateOne {
	cguo.mutation.ClearActivityFirst()
	return cguo
}

// SetActivitySecond sets the "activity_second" field.
func (cguo *CourseGradeUpdateOne) SetActivitySecond(i int) *CourseGradeUpdateOne {
	cguo.mutation.ResetActivitySecond()
	cguo.mutation.SetActivitySecond(i)
	return cguo
}

// SetNillableActivitySecond sets the "activity_second" field if the given value is not nil.
func (cguo *CourseGradeUpdateOne) SetNillableActivitySecond(i *int) *CourseGradeUpdateOne {
	if i != nil {
		cguo.SetActivitySecond(*i)
	}
	return cguo
}

// AddActivitySecond adds i to the "activity_second" field.
func (cguo *CourseGradeUpdateOne) AddActivitySecond(i int) *CourseGradeUpdateOne {
	cguo.mutation.AddActivitySecond(i)
	return cguo
}

// ClearActivitySecond clears the value of the "activity_second" field.
func (cguo *CourseGradeUpdateOne) ClearActivitySecond() *CourseGradeUpdateOne {
	cguo.mutation.ClearActivitySecond()
	return cguo
}

// SetWrittenFirst sets the "written_first" field.
func (cguo *CourseGradeUpdateOne) SetWrittenFirst(i int) *CourseGradeUpdateOne {
	cguo.mutation.ResetWrittenFirst()
	cguo.mutation.SetWrittenFirst(i)
	return cguo
}

// SetNillableWrittenFirst sets the "written_first" field if the given value is not nil.
func (cguo *CourseGradeUpdateOne) SetNillableWrittenFirst(i *int) *CourseGradeUpdateOne {
	if i != nil {
		cguo.SetWrittenFirst(*i)
	}
	return cguo
}

// AddWrittenFirst adds i to the "written_first" field.
func (cguo *CourseGradeUpdateOne) AddWrittenFirst(i int) *CourseGradeUpdateOne {
	cguo.mutation.AddWrittenFirst(i)
	return cguo
}

// ClearWrittenFirst clears the value of the "written_first" field.
func (cguo *CourseGradeUpdateOne) ClearWrittenFirst() *CourseGradeUpdateOne {
	cguo.mutation.ClearWrittenFirst()
	return cguo
}

// SetWrittenSecond sets the "written_second" field.
func (cguo *CourseGradeUpdateOne) SetWrittenSecond(i int) *CourseGradeUpdateOne {
	cguo.mutation.ResetWrittenSecond()
	cguo.mutation.SetWrittenSecond(i)
	return cguo
}

// SetNillableWrittenSecond sets the "written_second" field if the given value is not nil.
func (cguo *CourseGradeUpdateOne) SetNillableWrittenSecond(i *int) *CourseGradeUpdateOne {
	if i != nil {
		cguo.SetWrittenSecond(*i)
	}
	return cguo
}

// AddWrittenSecond adds i to the "written_second" field.
func (cguo *CourseGradeUpdateOne) AddWrittenSecond(i int) *CourseGradeUpdateOne {
	cguo.mutation.AddWrittenSecond(i)
	return cguo
}

// ClearWrittenSecond clears the value of the "written_second" field.
func (cguo *CourseGradeUpdateOne) ClearWrittenSecond() *CourseGradeUpdateOne {
	cguo.mutation.ClearWrittenSecond()
	return cguo
}

// SetCourseFinal sets the "course_final" field.
func (cguo *CourseGradeUpdateOne) SetCourseFinal(i int) *CourseGradeUpdateOne {
	cguo.mutation.ResetCourseFinal()
	cguo.mutation.SetCourseFinal(i)
	return cguo
}

// SetNillableCourseFinal sets the "course_final" field if the given value is not nil.
func (cguo *CourseGradeUpdateOne) SetNillableCourseFinal(i *int) *CourseGradeUpdateOne {
	if i != nil {
		cguo.SetCourseFinal(*i)
	}
	return cguo
}

// AddCourseFinal adds i to the "course_final" field.
func (cguo *CourseGradeUpdateOne) AddCourseFinal(i int) *CourseGradeUpdateOne {
	cguo.mutation.AddCourseFinal(i)
	return cguo
}

// ClearCourseFinal clears the value of the "course_final" field.
func (cguo *CourseGradeUpdateOne) ClearCourseFinal() *CourseGradeUpdateOne {
	cguo.mutation.ClearCourseFinal()
	return cguo
}

// SetStudentID sets the "student" edge to the User entity by ID.
func (cguo *CourseGradeUpdateOne) SetStudentID(id uuid.UUID) *CourseGradeUpdateOne {
	cguo.mutation.SetStudentID(id)
	return cguo
}

// SetStudent sets the "student" edge to the User entity.
func (cguo *CourseGradeUpdateOne) SetStudent(u *User) *CourseGradeUpdateOne {
	return cguo.SetStudentID(u.ID)
}

// SetClassID sets the "class" edge to the Class entity by ID.
func (cguo *CourseGradeUpdateOne) SetClassID(id uuid.UUID) *CourseGradeUpdateOne {
	cguo.mutation.SetClassID(id)
	return cguo
}

// SetClass sets the "class" edge to the Class entity.
func (cguo *CourseGradeUpdateOne) SetClass(c *Class) *CourseGradeUpdateOne {
	return cguo.SetClassID(c.ID)
}

// SetStageID sets the "stage" edge to the Stage entity by ID.
func (cguo *CourseGradeUpdateOne) SetStageID(id uuid.UUID) *CourseGradeUpdateOne {
	cguo.mutation.SetStageID(id)
	return cguo
}

// SetStage sets the "stage" edge to the Stage entity.
func (cguo *CourseGradeUpdateOne) SetStage(s *Stage) *CourseGradeUpdateOne {
	return cguo.SetStageID(s.ID)
}

// Mutation returns the CourseGradeMutation object of the builder.
func (cguo *CourseGradeUpdateOne) Mutation() *CourseGradeMutation {
	return cguo.mutation
}

// ClearStudent clears the "student" edge to the User entity.
func (cguo *CourseGradeUpdateOne) ClearStudent() *CourseGradeUpdateOne {
	cguo.mutation.ClearStudent()
	return cguo
}

// ClearClass clears the "class" edge to the Class entity.
func (cguo *CourseGradeUpdateOne) ClearClass() *CourseGradeUpdateOne {
	cguo.mutation.ClearClass()
	return cguo
}

// ClearStage clears the "stage" edge to the Stage entity.
func (cguo *CourseGradeUpdateOne) ClearStage() *CourseGradeUpdateOne {
	cguo.mutation.ClearStage()
	return cguo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cguo *CourseGradeUpdateOne) Select(field string, fields ...string) *CourseGradeUpdateOne {
	cguo.fields = append([]string{field}, fields...)
	return cguo
}

// Save executes the query and returns the updated CourseGrade entity.
func (cguo *CourseGradeUpdateOne) Save(ctx context.Context) (*CourseGrade, error) {
	var (
		err  error
		node *CourseGrade
	)
	cguo.defaults()
	if len(cguo.hooks) == 0 {
		if err = cguo.check(); err != nil {
			return nil, err
		}
		node, err = cguo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CourseGradeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cguo.check(); err != nil {
				return nil, err
			}
			cguo.mutation = mutation
			node, err = cguo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cguo.hooks) - 1; i >= 0; i-- {
			if cguo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cguo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cguo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cguo *CourseGradeUpdateOne) SaveX(ctx context.Context) *CourseGrade {
	node, err := cguo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cguo *CourseGradeUpdateOne) Exec(ctx context.Context) error {
	_, err := cguo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cguo *CourseGradeUpdateOne) ExecX(ctx context.Context) {
	if err := cguo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cguo *CourseGradeUpdateOne) defaults() {
	if _, ok := cguo.mutation.UpdatedAt(); !ok {
		v := coursegrade.UpdateDefaultUpdatedAt()
		cguo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cguo *CourseGradeUpdateOne) check() error {
	if v, ok := cguo.mutation.ActivityFirst(); ok {
		if err := coursegrade.ActivityFirstValidator(v); err != nil {
			return &ValidationError{Name: "activity_first", err: fmt.Errorf("ent: validator failed for field \"activity_first\": %w", err)}
		}
	}
	if v, ok := cguo.mutation.ActivitySecond(); ok {
		if err := coursegrade.ActivitySecondValidator(v); err != nil {
			return &ValidationError{Name: "activity_second", err: fmt.Errorf("ent: validator failed for field \"activity_second\": %w", err)}
		}
	}
	if v, ok := cguo.mutation.WrittenFirst(); ok {
		if err := coursegrade.WrittenFirstValidator(v); err != nil {
			return &ValidationError{Name: "written_first", err: fmt.Errorf("ent: validator failed for field \"written_first\": %w", err)}
		}
	}
	if v, ok := cguo.mutation.WrittenSecond(); ok {
		if err := coursegrade.WrittenSecondValidator(v); err != nil {
			return &ValidationError{Name: "written_second", err: fmt.Errorf("ent: validator failed for field \"written_second\": %w", err)}
		}
	}
	if v, ok := cguo.mutation.CourseFinal(); ok {
		if err := coursegrade.CourseFinalValidator(v); err != nil {
			return &ValidationError{Name: "course_final", err: fmt.Errorf("ent: validator failed for field \"course_final\": %w", err)}
		}
	}
	if _, ok := cguo.mutation.StudentID(); cguo.mutation.StudentCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"student\"")
	}
	if _, ok := cguo.mutation.ClassID(); cguo.mutation.ClassCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"class\"")
	}
	if _, ok := cguo.mutation.StageID(); cguo.mutation.StageCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"stage\"")
	}
	return nil
}

func (cguo *CourseGradeUpdateOne) sqlSave(ctx context.Context) (_node *CourseGrade, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   coursegrade.Table,
			Columns: coursegrade.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: coursegrade.FieldID,
			},
		},
	}
	id, ok := cguo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing CourseGrade.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := cguo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, coursegrade.FieldID)
		for _, f := range fields {
			if !coursegrade.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != coursegrade.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cguo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cguo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: coursegrade.FieldUpdatedAt,
		})
	}
	if value, ok := cguo.mutation.ActivityFirst(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coursegrade.FieldActivityFirst,
		})
	}
	if value, ok := cguo.mutation.AddedActivityFirst(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coursegrade.FieldActivityFirst,
		})
	}
	if cguo.mutation.ActivityFirstCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: coursegrade.FieldActivityFirst,
		})
	}
	if value, ok := cguo.mutation.ActivitySecond(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coursegrade.FieldActivitySecond,
		})
	}
	if value, ok := cguo.mutation.AddedActivitySecond(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coursegrade.FieldActivitySecond,
		})
	}
	if cguo.mutation.ActivitySecondCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: coursegrade.FieldActivitySecond,
		})
	}
	if value, ok := cguo.mutation.WrittenFirst(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coursegrade.FieldWrittenFirst,
		})
	}
	if value, ok := cguo.mutation.AddedWrittenFirst(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coursegrade.FieldWrittenFirst,
		})
	}
	if cguo.mutation.WrittenFirstCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: coursegrade.FieldWrittenFirst,
		})
	}
	if value, ok := cguo.mutation.WrittenSecond(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coursegrade.FieldWrittenSecond,
		})
	}
	if value, ok := cguo.mutation.AddedWrittenSecond(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coursegrade.FieldWrittenSecond,
		})
	}
	if cguo.mutation.WrittenSecondCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: coursegrade.FieldWrittenSecond,
		})
	}
	if value, ok := cguo.mutation.CourseFinal(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coursegrade.FieldCourseFinal,
		})
	}
	if value, ok := cguo.mutation.AddedCourseFinal(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coursegrade.FieldCourseFinal,
		})
	}
	if cguo.mutation.CourseFinalCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: coursegrade.FieldCourseFinal,
		})
	}
	if cguo.mutation.StudentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   coursegrade.StudentTable,
			Columns: []string{coursegrade.StudentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cguo.mutation.StudentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   coursegrade.StudentTable,
			Columns: []string{coursegrade.StudentColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cguo.mutation.ClassCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   coursegrade.ClassTable,
			Columns: []string{coursegrade.ClassColumn},
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
	if nodes := cguo.mutation.ClassIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   coursegrade.ClassTable,
			Columns: []string{coursegrade.ClassColumn},
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
	if cguo.mutation.StageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   coursegrade.StageTable,
			Columns: []string{coursegrade.StageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: stage.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cguo.mutation.StageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   coursegrade.StageTable,
			Columns: []string{coursegrade.StageColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &CourseGrade{config: cguo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cguo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{coursegrade.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
