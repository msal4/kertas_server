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

// Schedule holds the schema definition for the Schedule entity.
type Schedule struct {
	ent.Schema
}

// Fields of the Schedule.
func (Schedule) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Int("weekday").Min(0).Max(6).Annotations(entgql.OrderField("WEEKDAY")),
		field.Time("starts_at").SchemaType(map[string]string{dialect.Postgres: "time"}).Annotations(entgql.OrderField("STARTS_AT")),
		field.Int("duration").Default(60).Comment("Duration is the lecture duration in minutes").Annotations(entgql.OrderField("DURATION")),
	}
}

// Edges of the Schedule.
func (Schedule) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("class", Class.Type).Ref("schedules").Unique().Required(),
	}
}

func (Schedule) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("class").Fields("weekday", "starts_at"),
		index.Fields("weekday"),
	}
}
