// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/msal4/hassah_school_server/ent/assignment"
	"github.com/msal4/hassah_school_server/ent/attendance"
	"github.com/msal4/hassah_school_server/ent/class"
	"github.com/msal4/hassah_school_server/ent/coursegrade"
	"github.com/msal4/hassah_school_server/ent/group"
	"github.com/msal4/hassah_school_server/ent/schedule"
	"github.com/msal4/hassah_school_server/ent/stage"
	"github.com/msal4/hassah_school_server/ent/user"
)

// ClassCreate is the builder for creating a Class entity.
type ClassCreate struct {
	config
	mutation *ClassMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (cc *ClassCreate) SetCreatedAt(t time.Time) *ClassCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *ClassCreate) SetNillableCreatedAt(t *time.Time) *ClassCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *ClassCreate) SetUpdatedAt(t time.Time) *ClassCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *ClassCreate) SetNillableUpdatedAt(t *time.Time) *ClassCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetName sets the "name" field.
func (cc *ClassCreate) SetName(s string) *ClassCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetActive sets the "active" field.
func (cc *ClassCreate) SetActive(b bool) *ClassCreate {
	cc.mutation.SetActive(b)
	return cc
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (cc *ClassCreate) SetNillableActive(b *bool) *ClassCreate {
	if b != nil {
		cc.SetActive(*b)
	}
	return cc
}

// SetDeletedAt sets the "deleted_at" field.
func (cc *ClassCreate) SetDeletedAt(t time.Time) *ClassCreate {
	cc.mutation.SetDeletedAt(t)
	return cc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cc *ClassCreate) SetNillableDeletedAt(t *time.Time) *ClassCreate {
	if t != nil {
		cc.SetDeletedAt(*t)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *ClassCreate) SetID(u uuid.UUID) *ClassCreate {
	cc.mutation.SetID(u)
	return cc
}

// SetStageID sets the "stage" edge to the Stage entity by ID.
func (cc *ClassCreate) SetStageID(id uuid.UUID) *ClassCreate {
	cc.mutation.SetStageID(id)
	return cc
}

// SetStage sets the "stage" edge to the Stage entity.
func (cc *ClassCreate) SetStage(s *Stage) *ClassCreate {
	return cc.SetStageID(s.ID)
}

// SetTeacherID sets the "teacher" edge to the User entity by ID.
func (cc *ClassCreate) SetTeacherID(id uuid.UUID) *ClassCreate {
	cc.mutation.SetTeacherID(id)
	return cc
}

// SetTeacher sets the "teacher" edge to the User entity.
func (cc *ClassCreate) SetTeacher(u *User) *ClassCreate {
	return cc.SetTeacherID(u.ID)
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (cc *ClassCreate) SetGroupID(id uuid.UUID) *ClassCreate {
	cc.mutation.SetGroupID(id)
	return cc
}

// SetGroup sets the "group" edge to the Group entity.
func (cc *ClassCreate) SetGroup(g *Group) *ClassCreate {
	return cc.SetGroupID(g.ID)
}

// AddAssignmentIDs adds the "assignments" edge to the Assignment entity by IDs.
func (cc *ClassCreate) AddAssignmentIDs(ids ...uuid.UUID) *ClassCreate {
	cc.mutation.AddAssignmentIDs(ids...)
	return cc
}

// AddAssignments adds the "assignments" edges to the Assignment entity.
func (cc *ClassCreate) AddAssignments(a ...*Assignment) *ClassCreate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cc.AddAssignmentIDs(ids...)
}

// AddAttendanceIDs adds the "attendances" edge to the Attendance entity by IDs.
func (cc *ClassCreate) AddAttendanceIDs(ids ...uuid.UUID) *ClassCreate {
	cc.mutation.AddAttendanceIDs(ids...)
	return cc
}

// AddAttendances adds the "attendances" edges to the Attendance entity.
func (cc *ClassCreate) AddAttendances(a ...*Attendance) *ClassCreate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cc.AddAttendanceIDs(ids...)
}

// AddScheduleIDs adds the "schedules" edge to the Schedule entity by IDs.
func (cc *ClassCreate) AddScheduleIDs(ids ...uuid.UUID) *ClassCreate {
	cc.mutation.AddScheduleIDs(ids...)
	return cc
}

// AddSchedules adds the "schedules" edges to the Schedule entity.
func (cc *ClassCreate) AddSchedules(s ...*Schedule) *ClassCreate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cc.AddScheduleIDs(ids...)
}

// AddCourseGradeIDs adds the "course_grades" edge to the CourseGrade entity by IDs.
func (cc *ClassCreate) AddCourseGradeIDs(ids ...uuid.UUID) *ClassCreate {
	cc.mutation.AddCourseGradeIDs(ids...)
	return cc
}

// AddCourseGrades adds the "course_grades" edges to the CourseGrade entity.
func (cc *ClassCreate) AddCourseGrades(c ...*CourseGrade) *ClassCreate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cc.AddCourseGradeIDs(ids...)
}

// Mutation returns the ClassMutation object of the builder.
func (cc *ClassCreate) Mutation() *ClassMutation {
	return cc.mutation
}

// Save creates the Class in the database.
func (cc *ClassCreate) Save(ctx context.Context) (*Class, error) {
	var (
		err  error
		node *Class
	)
	cc.defaults()
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ClassMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ClassCreate) SaveX(ctx context.Context) *Class {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ClassCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ClassCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *ClassCreate) defaults() {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := class.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		v := class.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
	if _, ok := cc.mutation.Active(); !ok {
		v := class.DefaultActive
		cc.mutation.SetActive(v)
	}
	if _, ok := cc.mutation.ID(); !ok {
		v := class.DefaultID()
		cc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *ClassCreate) check() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if v, ok := cc.mutation.Name(); ok {
		if err := class.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "name": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Active(); !ok {
		return &ValidationError{Name: "active", err: errors.New(`ent: missing required field "active"`)}
	}
	if _, ok := cc.mutation.StageID(); !ok {
		return &ValidationError{Name: "stage", err: errors.New("ent: missing required edge \"stage\"")}
	}
	if _, ok := cc.mutation.TeacherID(); !ok {
		return &ValidationError{Name: "teacher", err: errors.New("ent: missing required edge \"teacher\"")}
	}
	if _, ok := cc.mutation.GroupID(); !ok {
		return &ValidationError{Name: "group", err: errors.New("ent: missing required edge \"group\"")}
	}
	return nil
}

func (cc *ClassCreate) sqlSave(ctx context.Context) (*Class, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
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

func (cc *ClassCreate) createSpec() (*Class, *sqlgraph.CreateSpec) {
	var (
		_node = &Class{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: class.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: class.FieldID,
			},
		}
	)
	_spec.OnConflict = cc.conflict
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: class.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: class.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: class.FieldName,
		})
		_node.Name = value
	}
	if value, ok := cc.mutation.Active(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: class.FieldActive,
		})
		_node.Active = value
	}
	if value, ok := cc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: class.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if nodes := cc.mutation.StageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   class.StageTable,
			Columns: []string{class.StageColumn},
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
		_node.stage_classes = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.TeacherIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   class.TeacherTable,
			Columns: []string{class.TeacherColumn},
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
		_node.user_classes = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   class.GroupTable,
			Columns: []string{class.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.AssignmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.AssignmentsTable,
			Columns: []string{class.AssignmentsColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.AttendancesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.AttendancesTable,
			Columns: []string{class.AttendancesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: attendance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.SchedulesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.SchedulesTable,
			Columns: []string{class.SchedulesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: schedule.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.CourseGradesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.CourseGradesTable,
			Columns: []string{class.CourseGradesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: coursegrade.FieldID,
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

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Class.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ClassUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (cc *ClassCreate) OnConflict(opts ...sql.ConflictOption) *ClassUpsertOne {
	cc.conflict = opts
	return &ClassUpsertOne{
		create: cc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Class.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (cc *ClassCreate) OnConflictColumns(columns ...string) *ClassUpsertOne {
	cc.conflict = append(cc.conflict, sql.ConflictColumns(columns...))
	return &ClassUpsertOne{
		create: cc,
	}
}

type (
	// ClassUpsertOne is the builder for "upsert"-ing
	//  one Class node.
	ClassUpsertOne struct {
		create *ClassCreate
	}

	// ClassUpsert is the "OnConflict" setter.
	ClassUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *ClassUpsert) SetCreatedAt(v time.Time) *ClassUpsert {
	u.Set(class.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ClassUpsert) UpdateCreatedAt() *ClassUpsert {
	u.SetExcluded(class.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ClassUpsert) SetUpdatedAt(v time.Time) *ClassUpsert {
	u.Set(class.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ClassUpsert) UpdateUpdatedAt() *ClassUpsert {
	u.SetExcluded(class.FieldUpdatedAt)
	return u
}

// SetName sets the "name" field.
func (u *ClassUpsert) SetName(v string) *ClassUpsert {
	u.Set(class.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ClassUpsert) UpdateName() *ClassUpsert {
	u.SetExcluded(class.FieldName)
	return u
}

// SetActive sets the "active" field.
func (u *ClassUpsert) SetActive(v bool) *ClassUpsert {
	u.Set(class.FieldActive, v)
	return u
}

// UpdateActive sets the "active" field to the value that was provided on create.
func (u *ClassUpsert) UpdateActive() *ClassUpsert {
	u.SetExcluded(class.FieldActive)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ClassUpsert) SetDeletedAt(v time.Time) *ClassUpsert {
	u.Set(class.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ClassUpsert) UpdateDeletedAt() *ClassUpsert {
	u.SetExcluded(class.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *ClassUpsert) ClearDeletedAt() *ClassUpsert {
	u.SetNull(class.FieldDeletedAt)
	return u
}

// UpdateNewValues updates the fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Class.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(class.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *ClassUpsertOne) UpdateNewValues() *ClassUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(class.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Class.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *ClassUpsertOne) Ignore() *ClassUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ClassUpsertOne) DoNothing() *ClassUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ClassCreate.OnConflict
// documentation for more info.
func (u *ClassUpsertOne) Update(set func(*ClassUpsert)) *ClassUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ClassUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *ClassUpsertOne) SetCreatedAt(v time.Time) *ClassUpsertOne {
	return u.Update(func(s *ClassUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ClassUpsertOne) UpdateCreatedAt() *ClassUpsertOne {
	return u.Update(func(s *ClassUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ClassUpsertOne) SetUpdatedAt(v time.Time) *ClassUpsertOne {
	return u.Update(func(s *ClassUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ClassUpsertOne) UpdateUpdatedAt() *ClassUpsertOne {
	return u.Update(func(s *ClassUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetName sets the "name" field.
func (u *ClassUpsertOne) SetName(v string) *ClassUpsertOne {
	return u.Update(func(s *ClassUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ClassUpsertOne) UpdateName() *ClassUpsertOne {
	return u.Update(func(s *ClassUpsert) {
		s.UpdateName()
	})
}

// SetActive sets the "active" field.
func (u *ClassUpsertOne) SetActive(v bool) *ClassUpsertOne {
	return u.Update(func(s *ClassUpsert) {
		s.SetActive(v)
	})
}

// UpdateActive sets the "active" field to the value that was provided on create.
func (u *ClassUpsertOne) UpdateActive() *ClassUpsertOne {
	return u.Update(func(s *ClassUpsert) {
		s.UpdateActive()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ClassUpsertOne) SetDeletedAt(v time.Time) *ClassUpsertOne {
	return u.Update(func(s *ClassUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ClassUpsertOne) UpdateDeletedAt() *ClassUpsertOne {
	return u.Update(func(s *ClassUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *ClassUpsertOne) ClearDeletedAt() *ClassUpsertOne {
	return u.Update(func(s *ClassUpsert) {
		s.ClearDeletedAt()
	})
}

// Exec executes the query.
func (u *ClassUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ClassCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ClassUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ClassUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: ClassUpsertOne.ID is not supported by MySQL driver. Use ClassUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ClassUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ClassCreateBulk is the builder for creating many Class entities in bulk.
type ClassCreateBulk struct {
	config
	builders []*ClassCreate
	conflict []sql.ConflictOption
}

// Save creates the Class entities in the database.
func (ccb *ClassCreateBulk) Save(ctx context.Context) ([]*Class, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Class, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ClassMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ClassCreateBulk) SaveX(ctx context.Context) []*Class {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ClassCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ClassCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Class.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ClassUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (ccb *ClassCreateBulk) OnConflict(opts ...sql.ConflictOption) *ClassUpsertBulk {
	ccb.conflict = opts
	return &ClassUpsertBulk{
		create: ccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Class.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ccb *ClassCreateBulk) OnConflictColumns(columns ...string) *ClassUpsertBulk {
	ccb.conflict = append(ccb.conflict, sql.ConflictColumns(columns...))
	return &ClassUpsertBulk{
		create: ccb,
	}
}

// ClassUpsertBulk is the builder for "upsert"-ing
// a bulk of Class nodes.
type ClassUpsertBulk struct {
	create *ClassCreateBulk
}

// UpdateNewValues updates the fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Class.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(class.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *ClassUpsertBulk) UpdateNewValues() *ClassUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(class.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Class.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *ClassUpsertBulk) Ignore() *ClassUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ClassUpsertBulk) DoNothing() *ClassUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ClassCreateBulk.OnConflict
// documentation for more info.
func (u *ClassUpsertBulk) Update(set func(*ClassUpsert)) *ClassUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ClassUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *ClassUpsertBulk) SetCreatedAt(v time.Time) *ClassUpsertBulk {
	return u.Update(func(s *ClassUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ClassUpsertBulk) UpdateCreatedAt() *ClassUpsertBulk {
	return u.Update(func(s *ClassUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ClassUpsertBulk) SetUpdatedAt(v time.Time) *ClassUpsertBulk {
	return u.Update(func(s *ClassUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ClassUpsertBulk) UpdateUpdatedAt() *ClassUpsertBulk {
	return u.Update(func(s *ClassUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetName sets the "name" field.
func (u *ClassUpsertBulk) SetName(v string) *ClassUpsertBulk {
	return u.Update(func(s *ClassUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ClassUpsertBulk) UpdateName() *ClassUpsertBulk {
	return u.Update(func(s *ClassUpsert) {
		s.UpdateName()
	})
}

// SetActive sets the "active" field.
func (u *ClassUpsertBulk) SetActive(v bool) *ClassUpsertBulk {
	return u.Update(func(s *ClassUpsert) {
		s.SetActive(v)
	})
}

// UpdateActive sets the "active" field to the value that was provided on create.
func (u *ClassUpsertBulk) UpdateActive() *ClassUpsertBulk {
	return u.Update(func(s *ClassUpsert) {
		s.UpdateActive()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ClassUpsertBulk) SetDeletedAt(v time.Time) *ClassUpsertBulk {
	return u.Update(func(s *ClassUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ClassUpsertBulk) UpdateDeletedAt() *ClassUpsertBulk {
	return u.Update(func(s *ClassUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *ClassUpsertBulk) ClearDeletedAt() *ClassUpsertBulk {
	return u.Update(func(s *ClassUpsert) {
		s.ClearDeletedAt()
	})
}

// Exec executes the query.
func (u *ClassUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ClassCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ClassCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ClassUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
