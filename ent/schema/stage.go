package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Stage holds the schema definition for the Stage entity.
type Stage struct {
	ent.Schema
}

// Fields of the Stage.
func (Stage) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name").NotEmpty().Annotations(entgql.OrderField("NAME")),
		field.Int("tuition_amount").Annotations(entgql.OrderField("TUITION_AMOUNT")),
		field.Enum("status").GoType(Status("")).Default(StatusActive.String()).Annotations(entgql.OrderField("STATUS")),
	}
}

func (Stage) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
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
