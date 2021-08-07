package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// Assignment holds the schema definition for the Assignment entity.
type Assignment struct {
	ent.Schema
}

// Fields of the Assignment.
func (Assignment) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("description").Optional(),
		field.Bool("is_exam").Default(false),
		field.Time("due_date"),
		field.Int("duration").Optional(),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

func (Assignment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Edges of the Assignment.
func (Assignment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("class", Class.Type).Ref("assignments").Unique().Required(),
		edge.To("submissions", AssignmentSubmission.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("grades", Grade.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}

func (Assignment) Index() []ent.Index {
	return []ent.Index{
		index.Edges("class"),
		index.Fields("is_exam"),
		index.Fields("deleted_at"),
	}
}
