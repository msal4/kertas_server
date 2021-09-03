package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// TuitionPayment holds the schema definition for the TuitionPayment entity.
type TuitionPayment struct {
	ent.Schema
}

// Fields of the TuitionPayment.
func (TuitionPayment) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("year").Validate(validateYear).Annotations(entgql.OrderField("YEAR")),
		field.Int("paid_amount").Annotations(entgql.OrderField("PAID_AMOUNT")),
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
