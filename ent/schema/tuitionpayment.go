package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// TuitionPayment holds the schema definition for the TuitionPayment entity.
type TuitionPayment struct {
	ent.Schema
}

// Fields of the TuitionPayment.
func (TuitionPayment) Fields() []ent.Field {
	return []ent.Field{
		field.Int("paid_amount"),
	}
}

func (TuitionPayment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Edges of the TuitionPayment.
func (TuitionPayment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("student", User.Type).Ref("payments").Unique().Required(),
		edge.From("stage", Stage.Type).Ref("payments").Unique().Required(),
	}
}

func (TuitionPayment) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("student"),
		index.Edges("stage"),
	}
}
