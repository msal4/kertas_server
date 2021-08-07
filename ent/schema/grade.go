package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// Grade holds the schema definition for the Grade entity.
type Grade struct {
	ent.Schema
}

// Fields of the Grade.
func (Grade) Fields() []ent.Field {
	return []ent.Field{
		field.Float("exam_grade").Range(0, 100),
	}
}

func (Grade) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Edges of the Grade.
func (Grade) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("student", User.Type).Ref("grades").Unique().Required(),
		edge.From("exam", Assignment.Type).Ref("grades").Unique().Required(),
	}
}

func (Grade) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("student", "exam").Unique(),
	}
}
