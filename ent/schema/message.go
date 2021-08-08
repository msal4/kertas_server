package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Message holds the schema definition for the Message entity.
type Message struct {
	ent.Schema
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.String("content").Optional(),
		field.String("attachment").Optional(),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

func (Message) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("group", Group.Type).Ref("messages").Unique().Required(),
		edge.From("owner", User.Type).Ref("messages").Unique().Required(),
	}
}

func (Message) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("owner"),
		index.Edges("group"),
		index.Fields("deleted_at"),
	}
}
