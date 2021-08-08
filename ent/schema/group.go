package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Group holds the schema definition for the Group entity.
type Group struct {
	ent.Schema
}

// Fields of the Group.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional(),
		field.Enum("type").Values("private", "shared").Default("shared"),
		field.Enum("status").GoType(Status("")).Default(StatusActive.String()),
	}
}

func (Group) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Edges of the Group.
func (Group) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("class", Class.Type).Ref("group").Unique(),
		edge.To("messages", Message.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}

func (Group) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("class"),
		index.Fields("status"),
		index.Fields("type"),
	}
}
