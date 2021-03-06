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
	"github.com/msal4/hassah_school_server/ent/assignmentsubmission"
	"github.com/msal4/hassah_school_server/ent/user"
)

// AssignmentSubmissionCreate is the builder for creating a AssignmentSubmission entity.
type AssignmentSubmissionCreate struct {
	config
	mutation *AssignmentSubmissionMutation
	hooks    []Hook
	conflict []sql.ConflictOption
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
	if _spec.ID.Value != nil {
		_node.ID = _spec.ID.Value.(uuid.UUID)
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
	_spec.OnConflict = asc.conflict
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

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AssignmentSubmission.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AssignmentSubmissionUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (asc *AssignmentSubmissionCreate) OnConflict(opts ...sql.ConflictOption) *AssignmentSubmissionUpsertOne {
	asc.conflict = opts
	return &AssignmentSubmissionUpsertOne{
		create: asc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AssignmentSubmission.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (asc *AssignmentSubmissionCreate) OnConflictColumns(columns ...string) *AssignmentSubmissionUpsertOne {
	asc.conflict = append(asc.conflict, sql.ConflictColumns(columns...))
	return &AssignmentSubmissionUpsertOne{
		create: asc,
	}
}

type (
	// AssignmentSubmissionUpsertOne is the builder for "upsert"-ing
	//  one AssignmentSubmission node.
	AssignmentSubmissionUpsertOne struct {
		create *AssignmentSubmissionCreate
	}

	// AssignmentSubmissionUpsert is the "OnConflict" setter.
	AssignmentSubmissionUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *AssignmentSubmissionUpsert) SetCreatedAt(v time.Time) *AssignmentSubmissionUpsert {
	u.Set(assignmentsubmission.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AssignmentSubmissionUpsert) UpdateCreatedAt() *AssignmentSubmissionUpsert {
	u.SetExcluded(assignmentsubmission.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AssignmentSubmissionUpsert) SetUpdatedAt(v time.Time) *AssignmentSubmissionUpsert {
	u.Set(assignmentsubmission.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AssignmentSubmissionUpsert) UpdateUpdatedAt() *AssignmentSubmissionUpsert {
	u.SetExcluded(assignmentsubmission.FieldUpdatedAt)
	return u
}

// SetFiles sets the "files" field.
func (u *AssignmentSubmissionUpsert) SetFiles(v []string) *AssignmentSubmissionUpsert {
	u.Set(assignmentsubmission.FieldFiles, v)
	return u
}

// UpdateFiles sets the "files" field to the value that was provided on create.
func (u *AssignmentSubmissionUpsert) UpdateFiles() *AssignmentSubmissionUpsert {
	u.SetExcluded(assignmentsubmission.FieldFiles)
	return u
}

// SetSubmittedAt sets the "submitted_at" field.
func (u *AssignmentSubmissionUpsert) SetSubmittedAt(v time.Time) *AssignmentSubmissionUpsert {
	u.Set(assignmentsubmission.FieldSubmittedAt, v)
	return u
}

// UpdateSubmittedAt sets the "submitted_at" field to the value that was provided on create.
func (u *AssignmentSubmissionUpsert) UpdateSubmittedAt() *AssignmentSubmissionUpsert {
	u.SetExcluded(assignmentsubmission.FieldSubmittedAt)
	return u
}

// ClearSubmittedAt clears the value of the "submitted_at" field.
func (u *AssignmentSubmissionUpsert) ClearSubmittedAt() *AssignmentSubmissionUpsert {
	u.SetNull(assignmentsubmission.FieldSubmittedAt)
	return u
}

// UpdateNewValues updates the fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.AssignmentSubmission.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(assignmentsubmission.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AssignmentSubmissionUpsertOne) UpdateNewValues() *AssignmentSubmissionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(assignmentsubmission.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.AssignmentSubmission.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *AssignmentSubmissionUpsertOne) Ignore() *AssignmentSubmissionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AssignmentSubmissionUpsertOne) DoNothing() *AssignmentSubmissionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AssignmentSubmissionCreate.OnConflict
// documentation for more info.
func (u *AssignmentSubmissionUpsertOne) Update(set func(*AssignmentSubmissionUpsert)) *AssignmentSubmissionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AssignmentSubmissionUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *AssignmentSubmissionUpsertOne) SetCreatedAt(v time.Time) *AssignmentSubmissionUpsertOne {
	return u.Update(func(s *AssignmentSubmissionUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AssignmentSubmissionUpsertOne) UpdateCreatedAt() *AssignmentSubmissionUpsertOne {
	return u.Update(func(s *AssignmentSubmissionUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AssignmentSubmissionUpsertOne) SetUpdatedAt(v time.Time) *AssignmentSubmissionUpsertOne {
	return u.Update(func(s *AssignmentSubmissionUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AssignmentSubmissionUpsertOne) UpdateUpdatedAt() *AssignmentSubmissionUpsertOne {
	return u.Update(func(s *AssignmentSubmissionUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetFiles sets the "files" field.
func (u *AssignmentSubmissionUpsertOne) SetFiles(v []string) *AssignmentSubmissionUpsertOne {
	return u.Update(func(s *AssignmentSubmissionUpsert) {
		s.SetFiles(v)
	})
}

// UpdateFiles sets the "files" field to the value that was provided on create.
func (u *AssignmentSubmissionUpsertOne) UpdateFiles() *AssignmentSubmissionUpsertOne {
	return u.Update(func(s *AssignmentSubmissionUpsert) {
		s.UpdateFiles()
	})
}

// SetSubmittedAt sets the "submitted_at" field.
func (u *AssignmentSubmissionUpsertOne) SetSubmittedAt(v time.Time) *AssignmentSubmissionUpsertOne {
	return u.Update(func(s *AssignmentSubmissionUpsert) {
		s.SetSubmittedAt(v)
	})
}

// UpdateSubmittedAt sets the "submitted_at" field to the value that was provided on create.
func (u *AssignmentSubmissionUpsertOne) UpdateSubmittedAt() *AssignmentSubmissionUpsertOne {
	return u.Update(func(s *AssignmentSubmissionUpsert) {
		s.UpdateSubmittedAt()
	})
}

// ClearSubmittedAt clears the value of the "submitted_at" field.
func (u *AssignmentSubmissionUpsertOne) ClearSubmittedAt() *AssignmentSubmissionUpsertOne {
	return u.Update(func(s *AssignmentSubmissionUpsert) {
		s.ClearSubmittedAt()
	})
}

// Exec executes the query.
func (u *AssignmentSubmissionUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AssignmentSubmissionCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AssignmentSubmissionUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AssignmentSubmissionUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: AssignmentSubmissionUpsertOne.ID is not supported by MySQL driver. Use AssignmentSubmissionUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AssignmentSubmissionUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AssignmentSubmissionCreateBulk is the builder for creating many AssignmentSubmission entities in bulk.
type AssignmentSubmissionCreateBulk struct {
	config
	builders []*AssignmentSubmissionCreate
	conflict []sql.ConflictOption
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
					spec.OnConflict = ascb.conflict
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

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AssignmentSubmission.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AssignmentSubmissionUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (ascb *AssignmentSubmissionCreateBulk) OnConflict(opts ...sql.ConflictOption) *AssignmentSubmissionUpsertBulk {
	ascb.conflict = opts
	return &AssignmentSubmissionUpsertBulk{
		create: ascb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AssignmentSubmission.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ascb *AssignmentSubmissionCreateBulk) OnConflictColumns(columns ...string) *AssignmentSubmissionUpsertBulk {
	ascb.conflict = append(ascb.conflict, sql.ConflictColumns(columns...))
	return &AssignmentSubmissionUpsertBulk{
		create: ascb,
	}
}

// AssignmentSubmissionUpsertBulk is the builder for "upsert"-ing
// a bulk of AssignmentSubmission nodes.
type AssignmentSubmissionUpsertBulk struct {
	create *AssignmentSubmissionCreateBulk
}

// UpdateNewValues updates the fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.AssignmentSubmission.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(assignmentsubmission.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AssignmentSubmissionUpsertBulk) UpdateNewValues() *AssignmentSubmissionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(assignmentsubmission.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AssignmentSubmission.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *AssignmentSubmissionUpsertBulk) Ignore() *AssignmentSubmissionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AssignmentSubmissionUpsertBulk) DoNothing() *AssignmentSubmissionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AssignmentSubmissionCreateBulk.OnConflict
// documentation for more info.
func (u *AssignmentSubmissionUpsertBulk) Update(set func(*AssignmentSubmissionUpsert)) *AssignmentSubmissionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AssignmentSubmissionUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *AssignmentSubmissionUpsertBulk) SetCreatedAt(v time.Time) *AssignmentSubmissionUpsertBulk {
	return u.Update(func(s *AssignmentSubmissionUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AssignmentSubmissionUpsertBulk) UpdateCreatedAt() *AssignmentSubmissionUpsertBulk {
	return u.Update(func(s *AssignmentSubmissionUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AssignmentSubmissionUpsertBulk) SetUpdatedAt(v time.Time) *AssignmentSubmissionUpsertBulk {
	return u.Update(func(s *AssignmentSubmissionUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AssignmentSubmissionUpsertBulk) UpdateUpdatedAt() *AssignmentSubmissionUpsertBulk {
	return u.Update(func(s *AssignmentSubmissionUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetFiles sets the "files" field.
func (u *AssignmentSubmissionUpsertBulk) SetFiles(v []string) *AssignmentSubmissionUpsertBulk {
	return u.Update(func(s *AssignmentSubmissionUpsert) {
		s.SetFiles(v)
	})
}

// UpdateFiles sets the "files" field to the value that was provided on create.
func (u *AssignmentSubmissionUpsertBulk) UpdateFiles() *AssignmentSubmissionUpsertBulk {
	return u.Update(func(s *AssignmentSubmissionUpsert) {
		s.UpdateFiles()
	})
}

// SetSubmittedAt sets the "submitted_at" field.
func (u *AssignmentSubmissionUpsertBulk) SetSubmittedAt(v time.Time) *AssignmentSubmissionUpsertBulk {
	return u.Update(func(s *AssignmentSubmissionUpsert) {
		s.SetSubmittedAt(v)
	})
}

// UpdateSubmittedAt sets the "submitted_at" field to the value that was provided on create.
func (u *AssignmentSubmissionUpsertBulk) UpdateSubmittedAt() *AssignmentSubmissionUpsertBulk {
	return u.Update(func(s *AssignmentSubmissionUpsert) {
		s.UpdateSubmittedAt()
	})
}

// ClearSubmittedAt clears the value of the "submitted_at" field.
func (u *AssignmentSubmissionUpsertBulk) ClearSubmittedAt() *AssignmentSubmissionUpsertBulk {
	return u.Update(func(s *AssignmentSubmissionUpsert) {
		s.ClearSubmittedAt()
	})
}

// Exec executes the query.
func (u *AssignmentSubmissionUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AssignmentSubmissionCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AssignmentSubmissionCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AssignmentSubmissionUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
