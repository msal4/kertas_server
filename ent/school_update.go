// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/msal4/hassah_school_server/ent/predicate"
	"github.com/msal4/hassah_school_server/ent/school"
	"github.com/msal4/hassah_school_server/ent/stage"
	"github.com/msal4/hassah_school_server/ent/user"
)

// SchoolUpdate is the builder for updating School entities.
type SchoolUpdate struct {
	config
	hooks    []Hook
	mutation *SchoolMutation
}

// Where appends a list predicates to the SchoolUpdate builder.
func (su *SchoolUpdate) Where(ps ...predicate.School) *SchoolUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetName sets the "name" field.
func (su *SchoolUpdate) SetName(s string) *SchoolUpdate {
	su.mutation.SetName(s)
	return su
}

// SetImage sets the "image" field.
func (su *SchoolUpdate) SetImage(s string) *SchoolUpdate {
	su.mutation.SetImage(s)
	return su
}

// SetDirectory sets the "directory" field.
func (su *SchoolUpdate) SetDirectory(s string) *SchoolUpdate {
	su.mutation.SetDirectory(s)
	return su
}

// SetActive sets the "active" field.
func (su *SchoolUpdate) SetActive(b bool) *SchoolUpdate {
	su.mutation.SetActive(b)
	return su
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (su *SchoolUpdate) SetNillableActive(b *bool) *SchoolUpdate {
	if b != nil {
		su.SetActive(*b)
	}
	return su
}

// SetDeletedAt sets the "deleted_at" field.
func (su *SchoolUpdate) SetDeletedAt(t time.Time) *SchoolUpdate {
	su.mutation.SetDeletedAt(t)
	return su
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (su *SchoolUpdate) SetNillableDeletedAt(t *time.Time) *SchoolUpdate {
	if t != nil {
		su.SetDeletedAt(*t)
	}
	return su
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (su *SchoolUpdate) ClearDeletedAt() *SchoolUpdate {
	su.mutation.ClearDeletedAt()
	return su
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (su *SchoolUpdate) AddUserIDs(ids ...uuid.UUID) *SchoolUpdate {
	su.mutation.AddUserIDs(ids...)
	return su
}

// AddUsers adds the "users" edges to the User entity.
func (su *SchoolUpdate) AddUsers(u ...*User) *SchoolUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return su.AddUserIDs(ids...)
}

// AddStageIDs adds the "stages" edge to the Stage entity by IDs.
func (su *SchoolUpdate) AddStageIDs(ids ...uuid.UUID) *SchoolUpdate {
	su.mutation.AddStageIDs(ids...)
	return su
}

// AddStages adds the "stages" edges to the Stage entity.
func (su *SchoolUpdate) AddStages(s ...*Stage) *SchoolUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return su.AddStageIDs(ids...)
}

// Mutation returns the SchoolMutation object of the builder.
func (su *SchoolUpdate) Mutation() *SchoolMutation {
	return su.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (su *SchoolUpdate) ClearUsers() *SchoolUpdate {
	su.mutation.ClearUsers()
	return su
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (su *SchoolUpdate) RemoveUserIDs(ids ...uuid.UUID) *SchoolUpdate {
	su.mutation.RemoveUserIDs(ids...)
	return su
}

// RemoveUsers removes "users" edges to User entities.
func (su *SchoolUpdate) RemoveUsers(u ...*User) *SchoolUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return su.RemoveUserIDs(ids...)
}

// ClearStages clears all "stages" edges to the Stage entity.
func (su *SchoolUpdate) ClearStages() *SchoolUpdate {
	su.mutation.ClearStages()
	return su
}

// RemoveStageIDs removes the "stages" edge to Stage entities by IDs.
func (su *SchoolUpdate) RemoveStageIDs(ids ...uuid.UUID) *SchoolUpdate {
	su.mutation.RemoveStageIDs(ids...)
	return su
}

// RemoveStages removes "stages" edges to Stage entities.
func (su *SchoolUpdate) RemoveStages(s ...*Stage) *SchoolUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return su.RemoveStageIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SchoolUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	su.defaults()
	if len(su.hooks) == 0 {
		if err = su.check(); err != nil {
			return 0, err
		}
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SchoolMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = su.check(); err != nil {
				return 0, err
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			if su.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *SchoolUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SchoolUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SchoolUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *SchoolUpdate) defaults() {
	if _, ok := su.mutation.UpdatedAt(); !ok {
		v := school.UpdateDefaultUpdatedAt()
		su.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *SchoolUpdate) check() error {
	if v, ok := su.mutation.Name(); ok {
		if err := school.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := su.mutation.Image(); ok {
		if err := school.ImageValidator(v); err != nil {
			return &ValidationError{Name: "image", err: fmt.Errorf("ent: validator failed for field \"image\": %w", err)}
		}
	}
	if v, ok := su.mutation.Directory(); ok {
		if err := school.DirectoryValidator(v); err != nil {
			return &ValidationError{Name: "directory", err: fmt.Errorf("ent: validator failed for field \"directory\": %w", err)}
		}
	}
	return nil
}

func (su *SchoolUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   school.Table,
			Columns: school.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: school.FieldID,
			},
		},
	}
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: school.FieldUpdatedAt,
		})
	}
	if value, ok := su.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: school.FieldName,
		})
	}
	if value, ok := su.mutation.Image(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: school.FieldImage,
		})
	}
	if value, ok := su.mutation.Directory(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: school.FieldDirectory,
		})
	}
	if value, ok := su.mutation.Active(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: school.FieldActive,
		})
	}
	if value, ok := su.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: school.FieldDeletedAt,
		})
	}
	if su.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: school.FieldDeletedAt,
		})
	}
	if su.mutation.UsersCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedUsersIDs(); len(nodes) > 0 && !su.mutation.UsersCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.UsersIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.StagesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedStagesIDs(); len(nodes) > 0 && !su.mutation.StagesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.StagesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{school.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// SchoolUpdateOne is the builder for updating a single School entity.
type SchoolUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SchoolMutation
}

// SetName sets the "name" field.
func (suo *SchoolUpdateOne) SetName(s string) *SchoolUpdateOne {
	suo.mutation.SetName(s)
	return suo
}

// SetImage sets the "image" field.
func (suo *SchoolUpdateOne) SetImage(s string) *SchoolUpdateOne {
	suo.mutation.SetImage(s)
	return suo
}

// SetDirectory sets the "directory" field.
func (suo *SchoolUpdateOne) SetDirectory(s string) *SchoolUpdateOne {
	suo.mutation.SetDirectory(s)
	return suo
}

// SetActive sets the "active" field.
func (suo *SchoolUpdateOne) SetActive(b bool) *SchoolUpdateOne {
	suo.mutation.SetActive(b)
	return suo
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (suo *SchoolUpdateOne) SetNillableActive(b *bool) *SchoolUpdateOne {
	if b != nil {
		suo.SetActive(*b)
	}
	return suo
}

// SetDeletedAt sets the "deleted_at" field.
func (suo *SchoolUpdateOne) SetDeletedAt(t time.Time) *SchoolUpdateOne {
	suo.mutation.SetDeletedAt(t)
	return suo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (suo *SchoolUpdateOne) SetNillableDeletedAt(t *time.Time) *SchoolUpdateOne {
	if t != nil {
		suo.SetDeletedAt(*t)
	}
	return suo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (suo *SchoolUpdateOne) ClearDeletedAt() *SchoolUpdateOne {
	suo.mutation.ClearDeletedAt()
	return suo
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (suo *SchoolUpdateOne) AddUserIDs(ids ...uuid.UUID) *SchoolUpdateOne {
	suo.mutation.AddUserIDs(ids...)
	return suo
}

// AddUsers adds the "users" edges to the User entity.
func (suo *SchoolUpdateOne) AddUsers(u ...*User) *SchoolUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return suo.AddUserIDs(ids...)
}

// AddStageIDs adds the "stages" edge to the Stage entity by IDs.
func (suo *SchoolUpdateOne) AddStageIDs(ids ...uuid.UUID) *SchoolUpdateOne {
	suo.mutation.AddStageIDs(ids...)
	return suo
}

// AddStages adds the "stages" edges to the Stage entity.
func (suo *SchoolUpdateOne) AddStages(s ...*Stage) *SchoolUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suo.AddStageIDs(ids...)
}

// Mutation returns the SchoolMutation object of the builder.
func (suo *SchoolUpdateOne) Mutation() *SchoolMutation {
	return suo.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (suo *SchoolUpdateOne) ClearUsers() *SchoolUpdateOne {
	suo.mutation.ClearUsers()
	return suo
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (suo *SchoolUpdateOne) RemoveUserIDs(ids ...uuid.UUID) *SchoolUpdateOne {
	suo.mutation.RemoveUserIDs(ids...)
	return suo
}

// RemoveUsers removes "users" edges to User entities.
func (suo *SchoolUpdateOne) RemoveUsers(u ...*User) *SchoolUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return suo.RemoveUserIDs(ids...)
}

// ClearStages clears all "stages" edges to the Stage entity.
func (suo *SchoolUpdateOne) ClearStages() *SchoolUpdateOne {
	suo.mutation.ClearStages()
	return suo
}

// RemoveStageIDs removes the "stages" edge to Stage entities by IDs.
func (suo *SchoolUpdateOne) RemoveStageIDs(ids ...uuid.UUID) *SchoolUpdateOne {
	suo.mutation.RemoveStageIDs(ids...)
	return suo
}

// RemoveStages removes "stages" edges to Stage entities.
func (suo *SchoolUpdateOne) RemoveStages(s ...*Stage) *SchoolUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suo.RemoveStageIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SchoolUpdateOne) Select(field string, fields ...string) *SchoolUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated School entity.
func (suo *SchoolUpdateOne) Save(ctx context.Context) (*School, error) {
	var (
		err  error
		node *School
	)
	suo.defaults()
	if len(suo.hooks) == 0 {
		if err = suo.check(); err != nil {
			return nil, err
		}
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SchoolMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = suo.check(); err != nil {
				return nil, err
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			if suo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = suo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SchoolUpdateOne) SaveX(ctx context.Context) *School {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SchoolUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SchoolUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *SchoolUpdateOne) defaults() {
	if _, ok := suo.mutation.UpdatedAt(); !ok {
		v := school.UpdateDefaultUpdatedAt()
		suo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *SchoolUpdateOne) check() error {
	if v, ok := suo.mutation.Name(); ok {
		if err := school.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := suo.mutation.Image(); ok {
		if err := school.ImageValidator(v); err != nil {
			return &ValidationError{Name: "image", err: fmt.Errorf("ent: validator failed for field \"image\": %w", err)}
		}
	}
	if v, ok := suo.mutation.Directory(); ok {
		if err := school.DirectoryValidator(v); err != nil {
			return &ValidationError{Name: "directory", err: fmt.Errorf("ent: validator failed for field \"directory\": %w", err)}
		}
	}
	return nil
}

func (suo *SchoolUpdateOne) sqlSave(ctx context.Context) (_node *School, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   school.Table,
			Columns: school.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: school.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing School.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, school.FieldID)
		for _, f := range fields {
			if !school.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != school.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: school.FieldUpdatedAt,
		})
	}
	if value, ok := suo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: school.FieldName,
		})
	}
	if value, ok := suo.mutation.Image(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: school.FieldImage,
		})
	}
	if value, ok := suo.mutation.Directory(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: school.FieldDirectory,
		})
	}
	if value, ok := suo.mutation.Active(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: school.FieldActive,
		})
	}
	if value, ok := suo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: school.FieldDeletedAt,
		})
	}
	if suo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: school.FieldDeletedAt,
		})
	}
	if suo.mutation.UsersCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedUsersIDs(); len(nodes) > 0 && !suo.mutation.UsersCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.UsersIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.StagesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedStagesIDs(); len(nodes) > 0 && !suo.mutation.StagesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.StagesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &School{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{school.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
