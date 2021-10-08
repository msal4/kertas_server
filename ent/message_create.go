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
	"github.com/msal4/hassah_school_server/ent/group"
	"github.com/msal4/hassah_school_server/ent/message"
	"github.com/msal4/hassah_school_server/ent/user"
)

// MessageCreate is the builder for creating a Message entity.
type MessageCreate struct {
	config
	mutation *MessageMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (mc *MessageCreate) SetCreatedAt(t time.Time) *MessageCreate {
	mc.mutation.SetCreatedAt(t)
	return mc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mc *MessageCreate) SetNillableCreatedAt(t *time.Time) *MessageCreate {
	if t != nil {
		mc.SetCreatedAt(*t)
	}
	return mc
}

// SetUpdatedAt sets the "updated_at" field.
func (mc *MessageCreate) SetUpdatedAt(t time.Time) *MessageCreate {
	mc.mutation.SetUpdatedAt(t)
	return mc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mc *MessageCreate) SetNillableUpdatedAt(t *time.Time) *MessageCreate {
	if t != nil {
		mc.SetUpdatedAt(*t)
	}
	return mc
}

// SetContent sets the "content" field.
func (mc *MessageCreate) SetContent(s string) *MessageCreate {
	mc.mutation.SetContent(s)
	return mc
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (mc *MessageCreate) SetNillableContent(s *string) *MessageCreate {
	if s != nil {
		mc.SetContent(*s)
	}
	return mc
}

// SetAttachment sets the "attachment" field.
func (mc *MessageCreate) SetAttachment(s string) *MessageCreate {
	mc.mutation.SetAttachment(s)
	return mc
}

// SetNillableAttachment sets the "attachment" field if the given value is not nil.
func (mc *MessageCreate) SetNillableAttachment(s *string) *MessageCreate {
	if s != nil {
		mc.SetAttachment(*s)
	}
	return mc
}

// SetDeletedAt sets the "deleted_at" field.
func (mc *MessageCreate) SetDeletedAt(t time.Time) *MessageCreate {
	mc.mutation.SetDeletedAt(t)
	return mc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (mc *MessageCreate) SetNillableDeletedAt(t *time.Time) *MessageCreate {
	if t != nil {
		mc.SetDeletedAt(*t)
	}
	return mc
}

// SetID sets the "id" field.
func (mc *MessageCreate) SetID(u uuid.UUID) *MessageCreate {
	mc.mutation.SetID(u)
	return mc
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (mc *MessageCreate) SetGroupID(id uuid.UUID) *MessageCreate {
	mc.mutation.SetGroupID(id)
	return mc
}

// SetGroup sets the "group" edge to the Group entity.
func (mc *MessageCreate) SetGroup(g *Group) *MessageCreate {
	return mc.SetGroupID(g.ID)
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (mc *MessageCreate) SetOwnerID(id uuid.UUID) *MessageCreate {
	mc.mutation.SetOwnerID(id)
	return mc
}

// SetOwner sets the "owner" edge to the User entity.
func (mc *MessageCreate) SetOwner(u *User) *MessageCreate {
	return mc.SetOwnerID(u.ID)
}

// Mutation returns the MessageMutation object of the builder.
func (mc *MessageCreate) Mutation() *MessageMutation {
	return mc.mutation
}

// Save creates the Message in the database.
func (mc *MessageCreate) Save(ctx context.Context) (*Message, error) {
	var (
		err  error
		node *Message
	)
	mc.defaults()
	if len(mc.hooks) == 0 {
		if err = mc.check(); err != nil {
			return nil, err
		}
		node, err = mc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MessageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mc.check(); err != nil {
				return nil, err
			}
			mc.mutation = mutation
			if node, err = mc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(mc.hooks) - 1; i >= 0; i-- {
			if mc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MessageCreate) SaveX(ctx context.Context) *Message {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MessageCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MessageCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MessageCreate) defaults() {
	if _, ok := mc.mutation.CreatedAt(); !ok {
		v := message.DefaultCreatedAt()
		mc.mutation.SetCreatedAt(v)
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		v := message.DefaultUpdatedAt()
		mc.mutation.SetUpdatedAt(v)
	}
	if _, ok := mc.mutation.ID(); !ok {
		v := message.DefaultID()
		mc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MessageCreate) check() error {
	if _, ok := mc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := mc.mutation.GroupID(); !ok {
		return &ValidationError{Name: "group", err: errors.New("ent: missing required edge \"group\"")}
	}
	if _, ok := mc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner", err: errors.New("ent: missing required edge \"owner\"")}
	}
	return nil
}

func (mc *MessageCreate) sqlSave(ctx context.Context) (*Message, error) {
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
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

func (mc *MessageCreate) createSpec() (*Message, *sqlgraph.CreateSpec) {
	var (
		_node = &Message{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: message.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: message.FieldID,
			},
		}
	)
	_spec.OnConflict = mc.conflict
	if id, ok := mc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: message.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := mc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: message.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := mc.mutation.Content(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: message.FieldContent,
		})
		_node.Content = value
	}
	if value, ok := mc.mutation.Attachment(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: message.FieldAttachment,
		})
		_node.Attachment = value
	}
	if value, ok := mc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: message.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if nodes := mc.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.GroupTable,
			Columns: []string{message.GroupColumn},
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
		_node.group_messages = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.OwnerTable,
			Columns: []string{message.OwnerColumn},
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
		_node.user_messages = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Message.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MessageUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (mc *MessageCreate) OnConflict(opts ...sql.ConflictOption) *MessageUpsertOne {
	mc.conflict = opts
	return &MessageUpsertOne{
		create: mc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Message.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (mc *MessageCreate) OnConflictColumns(columns ...string) *MessageUpsertOne {
	mc.conflict = append(mc.conflict, sql.ConflictColumns(columns...))
	return &MessageUpsertOne{
		create: mc,
	}
}

type (
	// MessageUpsertOne is the builder for "upsert"-ing
	//  one Message node.
	MessageUpsertOne struct {
		create *MessageCreate
	}

	// MessageUpsert is the "OnConflict" setter.
	MessageUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *MessageUpsert) SetCreatedAt(v time.Time) *MessageUpsert {
	u.Set(message.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *MessageUpsert) UpdateCreatedAt() *MessageUpsert {
	u.SetExcluded(message.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *MessageUpsert) SetUpdatedAt(v time.Time) *MessageUpsert {
	u.Set(message.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *MessageUpsert) UpdateUpdatedAt() *MessageUpsert {
	u.SetExcluded(message.FieldUpdatedAt)
	return u
}

// SetContent sets the "content" field.
func (u *MessageUpsert) SetContent(v string) *MessageUpsert {
	u.Set(message.FieldContent, v)
	return u
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *MessageUpsert) UpdateContent() *MessageUpsert {
	u.SetExcluded(message.FieldContent)
	return u
}

// ClearContent clears the value of the "content" field.
func (u *MessageUpsert) ClearContent() *MessageUpsert {
	u.SetNull(message.FieldContent)
	return u
}

// SetAttachment sets the "attachment" field.
func (u *MessageUpsert) SetAttachment(v string) *MessageUpsert {
	u.Set(message.FieldAttachment, v)
	return u
}

// UpdateAttachment sets the "attachment" field to the value that was provided on create.
func (u *MessageUpsert) UpdateAttachment() *MessageUpsert {
	u.SetExcluded(message.FieldAttachment)
	return u
}

// ClearAttachment clears the value of the "attachment" field.
func (u *MessageUpsert) ClearAttachment() *MessageUpsert {
	u.SetNull(message.FieldAttachment)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *MessageUpsert) SetDeletedAt(v time.Time) *MessageUpsert {
	u.Set(message.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *MessageUpsert) UpdateDeletedAt() *MessageUpsert {
	u.SetExcluded(message.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *MessageUpsert) ClearDeletedAt() *MessageUpsert {
	u.SetNull(message.FieldDeletedAt)
	return u
}

// UpdateNewValues updates the fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Message.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(message.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *MessageUpsertOne) UpdateNewValues() *MessageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(message.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Message.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *MessageUpsertOne) Ignore() *MessageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MessageUpsertOne) DoNothing() *MessageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MessageCreate.OnConflict
// documentation for more info.
func (u *MessageUpsertOne) Update(set func(*MessageUpsert)) *MessageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MessageUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *MessageUpsertOne) SetCreatedAt(v time.Time) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *MessageUpsertOne) UpdateCreatedAt() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *MessageUpsertOne) SetUpdatedAt(v time.Time) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *MessageUpsertOne) UpdateUpdatedAt() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetContent sets the "content" field.
func (u *MessageUpsertOne) SetContent(v string) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.SetContent(v)
	})
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *MessageUpsertOne) UpdateContent() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateContent()
	})
}

// ClearContent clears the value of the "content" field.
func (u *MessageUpsertOne) ClearContent() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.ClearContent()
	})
}

// SetAttachment sets the "attachment" field.
func (u *MessageUpsertOne) SetAttachment(v string) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.SetAttachment(v)
	})
}

// UpdateAttachment sets the "attachment" field to the value that was provided on create.
func (u *MessageUpsertOne) UpdateAttachment() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateAttachment()
	})
}

// ClearAttachment clears the value of the "attachment" field.
func (u *MessageUpsertOne) ClearAttachment() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.ClearAttachment()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *MessageUpsertOne) SetDeletedAt(v time.Time) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *MessageUpsertOne) UpdateDeletedAt() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *MessageUpsertOne) ClearDeletedAt() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.ClearDeletedAt()
	})
}

// Exec executes the query.
func (u *MessageUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MessageCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MessageUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *MessageUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: MessageUpsertOne.ID is not supported by MySQL driver. Use MessageUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *MessageUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// MessageCreateBulk is the builder for creating many Message entities in bulk.
type MessageCreateBulk struct {
	config
	builders []*MessageCreate
	conflict []sql.ConflictOption
}

// Save creates the Message entities in the database.
func (mcb *MessageCreateBulk) Save(ctx context.Context) ([]*Message, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Message, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MessageMutation)
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
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = mcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MessageCreateBulk) SaveX(ctx context.Context) []*Message {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MessageCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MessageCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Message.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MessageUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (mcb *MessageCreateBulk) OnConflict(opts ...sql.ConflictOption) *MessageUpsertBulk {
	mcb.conflict = opts
	return &MessageUpsertBulk{
		create: mcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Message.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (mcb *MessageCreateBulk) OnConflictColumns(columns ...string) *MessageUpsertBulk {
	mcb.conflict = append(mcb.conflict, sql.ConflictColumns(columns...))
	return &MessageUpsertBulk{
		create: mcb,
	}
}

// MessageUpsertBulk is the builder for "upsert"-ing
// a bulk of Message nodes.
type MessageUpsertBulk struct {
	create *MessageCreateBulk
}

// UpdateNewValues updates the fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Message.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(message.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *MessageUpsertBulk) UpdateNewValues() *MessageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(message.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Message.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *MessageUpsertBulk) Ignore() *MessageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MessageUpsertBulk) DoNothing() *MessageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MessageCreateBulk.OnConflict
// documentation for more info.
func (u *MessageUpsertBulk) Update(set func(*MessageUpsert)) *MessageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MessageUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *MessageUpsertBulk) SetCreatedAt(v time.Time) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *MessageUpsertBulk) UpdateCreatedAt() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *MessageUpsertBulk) SetUpdatedAt(v time.Time) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *MessageUpsertBulk) UpdateUpdatedAt() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetContent sets the "content" field.
func (u *MessageUpsertBulk) SetContent(v string) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.SetContent(v)
	})
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *MessageUpsertBulk) UpdateContent() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateContent()
	})
}

// ClearContent clears the value of the "content" field.
func (u *MessageUpsertBulk) ClearContent() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.ClearContent()
	})
}

// SetAttachment sets the "attachment" field.
func (u *MessageUpsertBulk) SetAttachment(v string) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.SetAttachment(v)
	})
}

// UpdateAttachment sets the "attachment" field to the value that was provided on create.
func (u *MessageUpsertBulk) UpdateAttachment() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateAttachment()
	})
}

// ClearAttachment clears the value of the "attachment" field.
func (u *MessageUpsertBulk) ClearAttachment() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.ClearAttachment()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *MessageUpsertBulk) SetDeletedAt(v time.Time) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *MessageUpsertBulk) UpdateDeletedAt() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *MessageUpsertBulk) ClearDeletedAt() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.ClearDeletedAt()
	})
}

// Exec executes the query.
func (u *MessageUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the MessageCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MessageCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MessageUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
