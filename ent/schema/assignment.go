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

// Assignment holds the schema definition for the Assignment entity.
type Assignment struct {
	ent.Schema
}

// Fields of the Assignment.
func (Assignment) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name").NotEmpty().Annotations(entgql.OrderField("NAME")),
		field.String("description").Optional().Annotations(entgql.OrderField("DESCRIPTION")),
		field.Bool("is_exam").Default(false).Annotations(entgql.OrderField("IS_EXAM")),
		field.Time("due_date").Annotations(entgql.OrderField("DUE_DATE")),
		field.Int("duration").Optional().Annotations(entgql.OrderField("DURATION")),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

func (Assignment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
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
