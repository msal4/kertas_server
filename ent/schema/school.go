package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// School holds the schema definition for the School entity.
type School struct {
	ent.Schema
}

// Fields of the School.
func (School) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("image").NotEmpty(),
		field.Enum("status").GoType(Status("")).Default(StatusActive.String()),
	}
}

func (School) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Edges of the School.
func (School) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("stages", Stage.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}

func (School) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("status"),
	}
}
