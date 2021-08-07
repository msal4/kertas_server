package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// Attendance holds the schema definition for the Attendance entity.
type Attendance struct {
	ent.Schema
}

// Fields of the Attendance.
func (Attendance) Fields() []ent.Field {
	return []ent.Field{
		field.Time("date").SchemaType(map[string]string{dialect.Postgres: "date"}),
		field.Enum("state").Values("present", "absent", "excused_absence", "sick"),
	}
}

func (Attendance) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Edges of the Attendance.
func (Attendance) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("class", Class.Type).Ref("attendances").Unique().Required(),
		edge.From("student", User.Type).Ref("attendances").Unique().Required(),
	}
}

func (Attendance) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("class", "student").Fields("date"),
		index.Fields("state"),
		index.Fields("date"),
	}
}
