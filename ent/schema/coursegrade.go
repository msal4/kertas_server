package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// CourseGrade holds the schema definition for the CourseGrade entity.
type CourseGrade struct {
	ent.Schema
}

func (CourseGrade) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Int("activity_first").Positive().Optional().Annotations(entgql.OrderField("ACTIVITY_FIRST")),
		field.Int("activity_second").Positive().Optional().Annotations(entgql.OrderField("ACTIVITY_SECOND")),
		field.Int("written_first").Positive().Annotations(entgql.OrderField("WRITTEN_FIRST")),
		field.Int("written_second").Positive().Annotations(entgql.OrderField("WRITTEN_SECOND")),
		field.Int("course_final").Positive().Annotations(entgql.OrderField("COURSE_FINAL")),
	}
}

func (CourseGrade) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (CourseGrade) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("student", User.Type).Ref("course_grades").Unique().Required(),
		edge.From("class", Class.Type).Ref("course_grades").Unique().Required(),
		edge.From("stage", Stage.Type).Ref("course_grades").Unique().Required(),
	}
}

func (CourseGrade) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("student", "class", "stage").Unique(),
	}
}