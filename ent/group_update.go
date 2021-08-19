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
	"github.com/msal4/hassah_school_server/ent/class"
	"github.com/msal4/hassah_school_server/ent/group"
	"github.com/msal4/hassah_school_server/ent/message"
	"github.com/msal4/hassah_school_server/ent/predicate"
	"github.com/msal4/hassah_school_server/ent/user"
)

// GroupUpdate is the builder for updating Group entities.
type GroupUpdate struct {
	config
	hooks    []Hook
	mutation *GroupMutation
}

// Where appends a list predicates to the GroupUpdate builder.
func (gu *GroupUpdate) Where(ps ...predicate.Group) *GroupUpdate {
	gu.mutation.Where(ps...)
	return gu
}

// SetName sets the "name" field.
func (gu *GroupUpdate) SetName(s string) *GroupUpdate {
	gu.mutation.SetName(s)
	return gu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (gu *GroupUpdate) SetNillableName(s *string) *GroupUpdate {
	if s != nil {
		gu.SetName(*s)
	}
	return gu
}

// ClearName clears the value of the "name" field.
func (gu *GroupUpdate) ClearName() *GroupUpdate {
	gu.mutation.ClearName()
	return gu
}

// SetGroupType sets the "group_type" field.
func (gu *GroupUpdate) SetGroupType(gt group.GroupType) *GroupUpdate {
	gu.mutation.SetGroupType(gt)
	return gu
}

// SetNillableGroupType sets the "group_type" field if the given value is not nil.
func (gu *GroupUpdate) SetNillableGroupType(gt *group.GroupType) *GroupUpdate {
	if gt != nil {
		gu.SetGroupType(*gt)
	}
	return gu
}

// SetActive sets the "active" field.
func (gu *GroupUpdate) SetActive(b bool) *GroupUpdate {
	gu.mutation.SetActive(b)
	return gu
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (gu *GroupUpdate) SetNillableActive(b *bool) *GroupUpdate {
	if b != nil {
		gu.SetActive(*b)
	}
	return gu
}

// SetDeletedAt sets the "deleted_at" field.
func (gu *GroupUpdate) SetDeletedAt(t time.Time) *GroupUpdate {
	gu.mutation.SetDeletedAt(t)
	return gu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gu *GroupUpdate) SetNillableDeletedAt(t *time.Time) *GroupUpdate {
	if t != nil {
		gu.SetDeletedAt(*t)
	}
	return gu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (gu *GroupUpdate) ClearDeletedAt() *GroupUpdate {
	gu.mutation.ClearDeletedAt()
	return gu
}

// SetClassID sets the "class" edge to the Class entity by ID.
func (gu *GroupUpdate) SetClassID(id uuid.UUID) *GroupUpdate {
	gu.mutation.SetClassID(id)
	return gu
}

// SetNillableClassID sets the "class" edge to the Class entity by ID if the given value is not nil.
func (gu *GroupUpdate) SetNillableClassID(id *uuid.UUID) *GroupUpdate {
	if id != nil {
		gu = gu.SetClassID(*id)
	}
	return gu
}

// SetClass sets the "class" edge to the Class entity.
func (gu *GroupUpdate) SetClass(c *Class) *GroupUpdate {
	return gu.SetClassID(c.ID)
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (gu *GroupUpdate) AddUserIDs(ids ...uuid.UUID) *GroupUpdate {
	gu.mutation.AddUserIDs(ids...)
	return gu
}

// AddUsers adds the "users" edges to the User entity.
func (gu *GroupUpdate) AddUsers(u ...*User) *GroupUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return gu.AddUserIDs(ids...)
}

// AddMessageIDs adds the "messages" edge to the Message entity by IDs.
func (gu *GroupUpdate) AddMessageIDs(ids ...uuid.UUID) *GroupUpdate {
	gu.mutation.AddMessageIDs(ids...)
	return gu
}

// AddMessages adds the "messages" edges to the Message entity.
func (gu *GroupUpdate) AddMessages(m ...*Message) *GroupUpdate {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return gu.AddMessageIDs(ids...)
}

// Mutation returns the GroupMutation object of the builder.
func (gu *GroupUpdate) Mutation() *GroupMutation {
	return gu.mutation
}

// ClearClass clears the "class" edge to the Class entity.
func (gu *GroupUpdate) ClearClass() *GroupUpdate {
	gu.mutation.ClearClass()
	return gu
}

// ClearUsers clears all "users" edges to the User entity.
func (gu *GroupUpdate) ClearUsers() *GroupUpdate {
	gu.mutation.ClearUsers()
	return gu
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (gu *GroupUpdate) RemoveUserIDs(ids ...uuid.UUID) *GroupUpdate {
	gu.mutation.RemoveUserIDs(ids...)
	return gu
}

// RemoveUsers removes "users" edges to User entities.
func (gu *GroupUpdate) RemoveUsers(u ...*User) *GroupUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return gu.RemoveUserIDs(ids...)
}

// ClearMessages clears all "messages" edges to the Message entity.
func (gu *GroupUpdate) ClearMessages() *GroupUpdate {
	gu.mutation.ClearMessages()
	return gu
}

// RemoveMessageIDs removes the "messages" edge to Message entities by IDs.
func (gu *GroupUpdate) RemoveMessageIDs(ids ...uuid.UUID) *GroupUpdate {
	gu.mutation.RemoveMessageIDs(ids...)
	return gu
}

// RemoveMessages removes "messages" edges to Message entities.
func (gu *GroupUpdate) RemoveMessages(m ...*Message) *GroupUpdate {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return gu.RemoveMessageIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gu *GroupUpdate) Save(ctx context.Context) (int, error) {
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
			mutation, ok := m.(*GroupMutation)
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
func (gu *GroupUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GroupUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GroupUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gu *GroupUpdate) defaults() {
	if _, ok := gu.mutation.UpdatedAt(); !ok {
		v := group.UpdateDefaultUpdatedAt()
		gu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gu *GroupUpdate) check() error {
	if v, ok := gu.mutation.GroupType(); ok {
		if err := group.GroupTypeValidator(v); err != nil {
			return &ValidationError{Name: "group_type", err: fmt.Errorf("ent: validator failed for field \"group_type\": %w", err)}
		}
	}
	return nil
}

func (gu *GroupUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   group.Table,
			Columns: group.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: group.FieldID,
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
	if value, ok := gu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: group.FieldUpdatedAt,
		})
	}
	if value, ok := gu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: group.FieldName,
		})
	}
	if gu.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: group.FieldName,
		})
	}
	if value, ok := gu.mutation.GroupType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: group.FieldGroupType,
		})
	}
	if value, ok := gu.mutation.Active(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: group.FieldActive,
		})
	}
	if value, ok := gu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: group.FieldDeletedAt,
		})
	}
	if gu.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: group.FieldDeletedAt,
		})
	}
	if gu.mutation.ClassCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   group.ClassTable,
			Columns: []string{group.ClassColumn},
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
	if nodes := gu.mutation.ClassIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   group.ClassTable,
			Columns: []string{group.ClassColumn},
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
	if gu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   group.UsersTable,
			Columns: group.UsersPrimaryKey,
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
	if nodes := gu.mutation.RemovedUsersIDs(); len(nodes) > 0 && !gu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   group.UsersTable,
			Columns: group.UsersPrimaryKey,
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
	if nodes := gu.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   group.UsersTable,
			Columns: group.UsersPrimaryKey,
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
	if gu.mutation.MessagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.MessagesTable,
			Columns: []string{group.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: message.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.RemovedMessagesIDs(); len(nodes) > 0 && !gu.mutation.MessagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.MessagesTable,
			Columns: []string{group.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: message.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.MessagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.MessagesTable,
			Columns: []string{group.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: message.FieldID,
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
			err = &NotFoundError{group.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// GroupUpdateOne is the builder for updating a single Group entity.
type GroupUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GroupMutation
}

// SetName sets the "name" field.
func (guo *GroupUpdateOne) SetName(s string) *GroupUpdateOne {
	guo.mutation.SetName(s)
	return guo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (guo *GroupUpdateOne) SetNillableName(s *string) *GroupUpdateOne {
	if s != nil {
		guo.SetName(*s)
	}
	return guo
}

// ClearName clears the value of the "name" field.
func (guo *GroupUpdateOne) ClearName() *GroupUpdateOne {
	guo.mutation.ClearName()
	return guo
}

// SetGroupType sets the "group_type" field.
func (guo *GroupUpdateOne) SetGroupType(gt group.GroupType) *GroupUpdateOne {
	guo.mutation.SetGroupType(gt)
	return guo
}

// SetNillableGroupType sets the "group_type" field if the given value is not nil.
func (guo *GroupUpdateOne) SetNillableGroupType(gt *group.GroupType) *GroupUpdateOne {
	if gt != nil {
		guo.SetGroupType(*gt)
	}
	return guo
}

// SetActive sets the "active" field.
func (guo *GroupUpdateOne) SetActive(b bool) *GroupUpdateOne {
	guo.mutation.SetActive(b)
	return guo
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (guo *GroupUpdateOne) SetNillableActive(b *bool) *GroupUpdateOne {
	if b != nil {
		guo.SetActive(*b)
	}
	return guo
}

// SetDeletedAt sets the "deleted_at" field.
func (guo *GroupUpdateOne) SetDeletedAt(t time.Time) *GroupUpdateOne {
	guo.mutation.SetDeletedAt(t)
	return guo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (guo *GroupUpdateOne) SetNillableDeletedAt(t *time.Time) *GroupUpdateOne {
	if t != nil {
		guo.SetDeletedAt(*t)
	}
	return guo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (guo *GroupUpdateOne) ClearDeletedAt() *GroupUpdateOne {
	guo.mutation.ClearDeletedAt()
	return guo
}

// SetClassID sets the "class" edge to the Class entity by ID.
func (guo *GroupUpdateOne) SetClassID(id uuid.UUID) *GroupUpdateOne {
	guo.mutation.SetClassID(id)
	return guo
}

// SetNillableClassID sets the "class" edge to the Class entity by ID if the given value is not nil.
func (guo *GroupUpdateOne) SetNillableClassID(id *uuid.UUID) *GroupUpdateOne {
	if id != nil {
		guo = guo.SetClassID(*id)
	}
	return guo
}

// SetClass sets the "class" edge to the Class entity.
func (guo *GroupUpdateOne) SetClass(c *Class) *GroupUpdateOne {
	return guo.SetClassID(c.ID)
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (guo *GroupUpdateOne) AddUserIDs(ids ...uuid.UUID) *GroupUpdateOne {
	guo.mutation.AddUserIDs(ids...)
	return guo
}

// AddUsers adds the "users" edges to the User entity.
func (guo *GroupUpdateOne) AddUsers(u ...*User) *GroupUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return guo.AddUserIDs(ids...)
}

// AddMessageIDs adds the "messages" edge to the Message entity by IDs.
func (guo *GroupUpdateOne) AddMessageIDs(ids ...uuid.UUID) *GroupUpdateOne {
	guo.mutation.AddMessageIDs(ids...)
	return guo
}

// AddMessages adds the "messages" edges to the Message entity.
func (guo *GroupUpdateOne) AddMessages(m ...*Message) *GroupUpdateOne {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return guo.AddMessageIDs(ids...)
}

// Mutation returns the GroupMutation object of the builder.
func (guo *GroupUpdateOne) Mutation() *GroupMutation {
	return guo.mutation
}

// ClearClass clears the "class" edge to the Class entity.
func (guo *GroupUpdateOne) ClearClass() *GroupUpdateOne {
	guo.mutation.ClearClass()
	return guo
}

// ClearUsers clears all "users" edges to the User entity.
func (guo *GroupUpdateOne) ClearUsers() *GroupUpdateOne {
	guo.mutation.ClearUsers()
	return guo
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (guo *GroupUpdateOne) RemoveUserIDs(ids ...uuid.UUID) *GroupUpdateOne {
	guo.mutation.RemoveUserIDs(ids...)
	return guo
}

// RemoveUsers removes "users" edges to User entities.
func (guo *GroupUpdateOne) RemoveUsers(u ...*User) *GroupUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return guo.RemoveUserIDs(ids...)
}

// ClearMessages clears all "messages" edges to the Message entity.
func (guo *GroupUpdateOne) ClearMessages() *GroupUpdateOne {
	guo.mutation.ClearMessages()
	return guo
}

// RemoveMessageIDs removes the "messages" edge to Message entities by IDs.
func (guo *GroupUpdateOne) RemoveMessageIDs(ids ...uuid.UUID) *GroupUpdateOne {
	guo.mutation.RemoveMessageIDs(ids...)
	return guo
}

// RemoveMessages removes "messages" edges to Message entities.
func (guo *GroupUpdateOne) RemoveMessages(m ...*Message) *GroupUpdateOne {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return guo.RemoveMessageIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (guo *GroupUpdateOne) Select(field string, fields ...string) *GroupUpdateOne {
	guo.fields = append([]string{field}, fields...)
	return guo
}

// Save executes the query and returns the updated Group entity.
func (guo *GroupUpdateOne) Save(ctx context.Context) (*Group, error) {
	var (
		err  error
		node *Group
	)
	guo.defaults()
	if len(guo.hooks) == 0 {
		if err = guo.check(); err != nil {
			return nil, err
		}
		node, err = guo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GroupMutation)
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
func (guo *GroupUpdateOne) SaveX(ctx context.Context) *Group {
	node, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (guo *GroupUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GroupUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (guo *GroupUpdateOne) defaults() {
	if _, ok := guo.mutation.UpdatedAt(); !ok {
		v := group.UpdateDefaultUpdatedAt()
		guo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (guo *GroupUpdateOne) check() error {
	if v, ok := guo.mutation.GroupType(); ok {
		if err := group.GroupTypeValidator(v); err != nil {
			return &ValidationError{Name: "group_type", err: fmt.Errorf("ent: validator failed for field \"group_type\": %w", err)}
		}
	}
	return nil
}

func (guo *GroupUpdateOne) sqlSave(ctx context.Context) (_node *Group, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   group.Table,
			Columns: group.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: group.FieldID,
			},
		},
	}
	id, ok := guo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Group.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := guo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, group.FieldID)
		for _, f := range fields {
			if !group.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != group.FieldID {
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
	if value, ok := guo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: group.FieldUpdatedAt,
		})
	}
	if value, ok := guo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: group.FieldName,
		})
	}
	if guo.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: group.FieldName,
		})
	}
	if value, ok := guo.mutation.GroupType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: group.FieldGroupType,
		})
	}
	if value, ok := guo.mutation.Active(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: group.FieldActive,
		})
	}
	if value, ok := guo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: group.FieldDeletedAt,
		})
	}
	if guo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: group.FieldDeletedAt,
		})
	}
	if guo.mutation.ClassCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   group.ClassTable,
			Columns: []string{group.ClassColumn},
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
	if nodes := guo.mutation.ClassIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   group.ClassTable,
			Columns: []string{group.ClassColumn},
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
	if guo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   group.UsersTable,
			Columns: group.UsersPrimaryKey,
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
	if nodes := guo.mutation.RemovedUsersIDs(); len(nodes) > 0 && !guo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   group.UsersTable,
			Columns: group.UsersPrimaryKey,
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
	if nodes := guo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   group.UsersTable,
			Columns: group.UsersPrimaryKey,
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
	if guo.mutation.MessagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.MessagesTable,
			Columns: []string{group.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: message.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.RemovedMessagesIDs(); len(nodes) > 0 && !guo.mutation.MessagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.MessagesTable,
			Columns: []string{group.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: message.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.MessagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.MessagesTable,
			Columns: []string{group.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: message.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Group{config: guo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, guo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{group.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
