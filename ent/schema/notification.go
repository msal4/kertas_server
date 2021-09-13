package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Notification holds the schema definition for the Notification entity.
type Notification struct {
	ent.Schema
}

// Fields of the Notification.
func (Notification) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("title").Annotations(entgql.OrderField("TITLE")),
		field.String("body").Optional().Annotations(entgql.OrderField("BODY")),
		field.String("image").Optional(),
		field.String("route").Optional().MaxLen(9),
		field.String("color").Optional(),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

func (Notification) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Edges of the Notification.
func (Notification) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("stage", Stage.Type).Ref("notifications").Unique().Required(),
	}
}

func (Notification) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("stage"),
		index.Fields("deleted_at"),
	}
}
