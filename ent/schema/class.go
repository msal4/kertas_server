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

// Class holds the schema definition for the Class entity.
type Class struct {
	ent.Schema
}

// Fields of the Class.
func (Class) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name").NotEmpty().Annotations(entgql.OrderField("NAME")),
		field.Enum("status").GoType(Status("")).Default(StatusActive.String()).Annotations(entgql.OrderField("STATUS")),
		field.Time("deleted_at").Nillable().Optional(),
	}
}

func (Class) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Edges of the Class.
func (Class) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("stage", Stage.Type).Ref("classes").Unique().Required(),
		edge.From("teacher", User.Type).Ref("classes").Unique().Required(),
		edge.To("group", Group.Type).Unique().Required().Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("assignments", Assignment.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("attendances", Attendance.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("schedules", Schedule.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}

func (Class) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("stage"),
		index.Edges("teacher"),
		index.Fields("status"),
	}
}
