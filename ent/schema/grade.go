package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Grade holds the schema definition for the Grade entity.
type Grade struct {
	ent.Schema
}

// Fields of the Grade.
func (Grade) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Int("exam_grade").Range(0, 100).Annotations(entgql.OrderField("EXAM_GRADE")),
	}
}

func (Grade) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
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
