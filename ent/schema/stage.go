package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// Stage holds the schema definition for the Stage entity.
type Stage struct {
	ent.Schema
}

// Fields of the Stage.
func (Stage) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.Int("tuition_amount"),
		field.Enum("status").GoType(Status("")).Default(StatusActive.String()),
	}
}

func (Stage) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Edges of the Stage.
func (Stage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("school", School.Type).Ref("stages").Unique(),
		edge.To("classes", Class.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("payments", TuitionPayment.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("students", User.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}

func (Stage) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("school"),
		index.Fields("status"),
	}
}
