package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// AssignmentSubmission holds the schema definition for the AssignmentSubmission entity.
type AssignmentSubmission struct {
	ent.Schema
}

// Fields of the AssignmentSubmission.
func (AssignmentSubmission) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Strings("files"),
		field.Time("submitted_at").Optional().Nillable().Annotations(entgql.OrderField("SUBMITTED_AT")),
	}
}

func (AssignmentSubmission) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Edges of the AssignmentSubmission.
func (AssignmentSubmission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("student", User.Type).Ref("submissions").Unique().Required(),
		edge.From("assignment", Assignment.Type).Ref("submissions").Unique().Required(),
	}
}

func (AssignmentSubmission) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("student", "assignment").Unique(),
		index.Edges("assignment"),
	}
}
