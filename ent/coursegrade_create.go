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
	"github.com/msal4/hassah_school_server/ent/coursegrade"
	"github.com/msal4/hassah_school_server/ent/stage"
	"github.com/msal4/hassah_school_server/ent/user"
)

// CourseGradeCreate is the builder for creating a CourseGrade entity.
type CourseGradeCreate struct {
	config
	mutation *CourseGradeMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (cgc *CourseGradeCreate) SetCreatedAt(t time.Time) *CourseGradeCreate {
	cgc.mutation.SetCreatedAt(t)
	return cgc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cgc *CourseGradeCreate) SetNillableCreatedAt(t *time.Time) *CourseGradeCreate {
	if t != nil {
		cgc.SetCreatedAt(*t)
	}
	return cgc
}

// SetUpdatedAt sets the "updated_at" field.
func (cgc *CourseGradeCreate) SetUpdatedAt(t time.Time) *CourseGradeCreate {
	cgc.mutation.SetUpdatedAt(t)
	return cgc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cgc *CourseGradeCreate) SetNillableUpdatedAt(t *time.Time) *CourseGradeCreate {
	if t != nil {
		cgc.SetUpdatedAt(*t)
	}
	return cgc
}

// SetCourse sets the "course" field.
func (cgc *CourseGradeCreate) SetCourse(c coursegrade.Course) *CourseGradeCreate {
	cgc.mutation.SetCourse(c)
	return cgc
}

// SetActivityFirst sets the "activity_first" field.
func (cgc *CourseGradeCreate) SetActivityFirst(i int) *CourseGradeCreate {
	cgc.mutation.SetActivityFirst(i)
	return cgc
}

// SetNillableActivityFirst sets the "activity_first" field if the given value is not nil.
func (cgc *CourseGradeCreate) SetNillableActivityFirst(i *int) *CourseGradeCreate {
	if i != nil {
		cgc.SetActivityFirst(*i)
	}
	return cgc
}

// SetActivitySecond sets the "activity_second" field.
func (cgc *CourseGradeCreate) SetActivitySecond(i int) *CourseGradeCreate {
	cgc.mutation.SetActivitySecond(i)
	return cgc
}

// SetNillableActivitySecond sets the "activity_second" field if the given value is not nil.
func (cgc *CourseGradeCreate) SetNillableActivitySecond(i *int) *CourseGradeCreate {
	if i != nil {
		cgc.SetActivitySecond(*i)
	}
	return cgc
}

// SetWrittenFirst sets the "written_first" field.
func (cgc *CourseGradeCreate) SetWrittenFirst(i int) *CourseGradeCreate {
	cgc.mutation.SetWrittenFirst(i)
	return cgc
}

// SetNillableWrittenFirst sets the "written_first" field if the given value is not nil.
func (cgc *CourseGradeCreate) SetNillableWrittenFirst(i *int) *CourseGradeCreate {
	if i != nil {
		cgc.SetWrittenFirst(*i)
	}
	return cgc
}

// SetWrittenSecond sets the "written_second" field.
func (cgc *CourseGradeCreate) SetWrittenSecond(i int) *CourseGradeCreate {
	cgc.mutation.SetWrittenSecond(i)
	return cgc
}

// SetNillableWrittenSecond sets the "written_second" field if the given value is not nil.
func (cgc *CourseGradeCreate) SetNillableWrittenSecond(i *int) *CourseGradeCreate {
	if i != nil {
		cgc.SetWrittenSecond(*i)
	}
	return cgc
}

// SetCourseFinal sets the "course_final" field.
func (cgc *CourseGradeCreate) SetCourseFinal(i int) *CourseGradeCreate {
	cgc.mutation.SetCourseFinal(i)
	return cgc
}

// SetNillableCourseFinal sets the "course_final" field if the given value is not nil.
func (cgc *CourseGradeCreate) SetNillableCourseFinal(i *int) *CourseGradeCreate {
	if i != nil {
		cgc.SetCourseFinal(*i)
	}
	return cgc
}

// SetYear sets the "year" field.
func (cgc *CourseGradeCreate) SetYear(s string) *CourseGradeCreate {
	cgc.mutation.SetYear(s)
	return cgc
}

// SetID sets the "id" field.
func (cgc *CourseGradeCreate) SetID(u uuid.UUID) *CourseGradeCreate {
	cgc.mutation.SetID(u)
	return cgc
}

// SetStudentID sets the "student" edge to the User entity by ID.
func (cgc *CourseGradeCreate) SetStudentID(id uuid.UUID) *CourseGradeCreate {
	cgc.mutation.SetStudentID(id)
	return cgc
}

// SetStudent sets the "student" edge to the User entity.
func (cgc *CourseGradeCreate) SetStudent(u *User) *CourseGradeCreate {
	return cgc.SetStudentID(u.ID)
}

// SetClassID sets the "class" edge to the Class entity by ID.
func (cgc *CourseGradeCreate) SetClassID(id uuid.UUID) *CourseGradeCreate {
	cgc.mutation.SetClassID(id)
	return cgc
}

// SetClass sets the "class" edge to the Class entity.
func (cgc *CourseGradeCreate) SetClass(c *Class) *CourseGradeCreate {
	return cgc.SetClassID(c.ID)
}

// SetStageID sets the "stage" edge to the Stage entity by ID.
func (cgc *CourseGradeCreate) SetStageID(id uuid.UUID) *CourseGradeCreate {
	cgc.mutation.SetStageID(id)
	return cgc
}

// SetStage sets the "stage" edge to the Stage entity.
func (cgc *CourseGradeCreate) SetStage(s *Stage) *CourseGradeCreate {
	return cgc.SetStageID(s.ID)
}

// Mutation returns the CourseGradeMutation object of the builder.
func (cgc *CourseGradeCreate) Mutation() *CourseGradeMutation {
	return cgc.mutation
}

// Save creates the CourseGrade in the database.
func (cgc *CourseGradeCreate) Save(ctx context.Context) (*CourseGrade, error) {
	var (
		err  error
		node *CourseGrade
	)
	cgc.defaults()
	if len(cgc.hooks) == 0 {
		if err = cgc.check(); err != nil {
			return nil, err
		}
		node, err = cgc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CourseGradeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cgc.check(); err != nil {
				return nil, err
			}
			cgc.mutation = mutation
			if node, err = cgc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cgc.hooks) - 1; i >= 0; i-- {
			if cgc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cgc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cgc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cgc *CourseGradeCreate) SaveX(ctx context.Context) *CourseGrade {
	v, err := cgc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cgc *CourseGradeCreate) Exec(ctx context.Context) error {
	_, err := cgc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cgc *CourseGradeCreate) ExecX(ctx context.Context) {
	if err := cgc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cgc *CourseGradeCreate) defaults() {
	if _, ok := cgc.mutation.CreatedAt(); !ok {
		v := coursegrade.DefaultCreatedAt()
		cgc.mutation.SetCreatedAt(v)
	}
	if _, ok := cgc.mutation.UpdatedAt(); !ok {
		v := coursegrade.DefaultUpdatedAt()
		cgc.mutation.SetUpdatedAt(v)
	}
	if _, ok := cgc.mutation.ID(); !ok {
		v := coursegrade.DefaultID()
		cgc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cgc *CourseGradeCreate) check() error {
	if _, ok := cgc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := cgc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := cgc.mutation.Course(); !ok {
		return &ValidationError{Name: "course", err: errors.New(`ent: missing required field "course"`)}
	}
	if v, ok := cgc.mutation.Course(); ok {
		if err := coursegrade.CourseValidator(v); err != nil {
			return &ValidationError{Name: "course", err: fmt.Errorf(`ent: validator failed for field "course": %w`, err)}
		}
	}
	if v, ok := cgc.mutation.ActivityFirst(); ok {
		if err := coursegrade.ActivityFirstValidator(v); err != nil {
			return &ValidationError{Name: "activity_first", err: fmt.Errorf(`ent: validator failed for field "activity_first": %w`, err)}
		}
	}
	if v, ok := cgc.mutation.ActivitySecond(); ok {
		if err := coursegrade.ActivitySecondValidator(v); err != nil {
			return &ValidationError{Name: "activity_second", err: fmt.Errorf(`ent: validator failed for field "activity_second": %w`, err)}
		}
	}
	if v, ok := cgc.mutation.WrittenFirst(); ok {
		if err := coursegrade.WrittenFirstValidator(v); err != nil {
			return &ValidationError{Name: "written_first", err: fmt.Errorf(`ent: validator failed for field "written_first": %w`, err)}
		}
	}
	if v, ok := cgc.mutation.WrittenSecond(); ok {
		if err := coursegrade.WrittenSecondValidator(v); err != nil {
			return &ValidationError{Name: "written_second", err: fmt.Errorf(`ent: validator failed for field "written_second": %w`, err)}
		}
	}
	if v, ok := cgc.mutation.CourseFinal(); ok {
		if err := coursegrade.CourseFinalValidator(v); err != nil {
			return &ValidationError{Name: "course_final", err: fmt.Errorf(`ent: validator failed for field "course_final": %w`, err)}
		}
	}
	if _, ok := cgc.mutation.Year(); !ok {
		return &ValidationError{Name: "year", err: errors.New(`ent: missing required field "year"`)}
	}
	if v, ok := cgc.mutation.Year(); ok {
		if err := coursegrade.YearValidator(v); err != nil {
			return &ValidationError{Name: "year", err: fmt.Errorf(`ent: validator failed for field "year": %w`, err)}
		}
	}
	if _, ok := cgc.mutation.StudentID(); !ok {
		return &ValidationError{Name: "student", err: errors.New("ent: missing required edge \"student\"")}
	}
	if _, ok := cgc.mutation.ClassID(); !ok {
		return &ValidationError{Name: "class", err: errors.New("ent: missing required edge \"class\"")}
	}
	if _, ok := cgc.mutation.StageID(); !ok {
		return &ValidationError{Name: "stage", err: errors.New("ent: missing required edge \"stage\"")}
	}
	return nil
}

func (cgc *CourseGradeCreate) sqlSave(ctx context.Context) (*CourseGrade, error) {
	_node, _spec := cgc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cgc.driver, _spec); err != nil {
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

func (cgc *CourseGradeCreate) createSpec() (*CourseGrade, *sqlgraph.CreateSpec) {
	var (
		_node = &CourseGrade{config: cgc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: coursegrade.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: coursegrade.FieldID,
			},
		}
	)
	if id, ok := cgc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cgc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: coursegrade.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := cgc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: coursegrade.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := cgc.mutation.Course(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: coursegrade.FieldCourse,
		})
		_node.Course = value
	}
	if value, ok := cgc.mutation.ActivityFirst(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coursegrade.FieldActivityFirst,
		})
		_node.ActivityFirst = &value
	}
	if value, ok := cgc.mutation.ActivitySecond(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coursegrade.FieldActivitySecond,
		})
		_node.ActivitySecond = &value
	}
	if value, ok := cgc.mutation.WrittenFirst(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coursegrade.FieldWrittenFirst,
		})
		_node.WrittenFirst = &value
	}
	if value, ok := cgc.mutation.WrittenSecond(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coursegrade.FieldWrittenSecond,
		})
		_node.WrittenSecond = &value
	}
	if value, ok := cgc.mutation.CourseFinal(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coursegrade.FieldCourseFinal,
		})
		_node.CourseFinal = &value
	}
	if value, ok := cgc.mutation.Year(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coursegrade.FieldYear,
		})
		_node.Year = value
	}
	if nodes := cgc.mutation.StudentIDs(); len(nodes) > 0 {
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
		_node.user_course_grades = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cgc.mutation.ClassIDs(); len(nodes) > 0 {
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
		_node.class_course_grades = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cgc.mutation.StageIDs(); len(nodes) > 0 {
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
		_node.stage_course_grades = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CourseGradeCreateBulk is the builder for creating many CourseGrade entities in bulk.
type CourseGradeCreateBulk struct {
	config
	builders []*CourseGradeCreate
}

// Save creates the CourseGrade entities in the database.
func (cgcb *CourseGradeCreateBulk) Save(ctx context.Context) ([]*CourseGrade, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cgcb.builders))
	nodes := make([]*CourseGrade, len(cgcb.builders))
	mutators := make([]Mutator, len(cgcb.builders))
	for i := range cgcb.builders {
		func(i int, root context.Context) {
			builder := cgcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CourseGradeMutation)
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
					_, err = mutators[i+1].Mutate(root, cgcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cgcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, cgcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cgcb *CourseGradeCreateBulk) SaveX(ctx context.Context) []*CourseGrade {
	v, err := cgcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cgcb *CourseGradeCreateBulk) Exec(ctx context.Context) error {
	_, err := cgcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cgcb *CourseGradeCreateBulk) ExecX(ctx context.Context) {
	if err := cgcb.Exec(ctx); err != nil {
		panic(err)
	}
}
