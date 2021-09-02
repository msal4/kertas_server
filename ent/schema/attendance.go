package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Attendance holds the schema definition for the Attendance entity.
type Attendance struct {
	ent.Schema
}

// Fields of the Attendance.
func (Attendance) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Time("date").SchemaType(map[string]string{dialect.Postgres: "date"}).Annotations(entgql.OrderField("DATE")),
		field.Enum("state").NamedValues(
			"Present", "PRESENT",
			"Absent", "ABSENT",
			"ExcusedAbsence", "EXCUSED_ABSENCE",
			"Sick", "SICK",
		).Default("PRESENT").Annotations(entgql.OrderField("STATE")),
	}
}

func (Attendance) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
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
