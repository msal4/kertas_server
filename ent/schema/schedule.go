package schema

import (
	"time"

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
		field.Int("weekday").GoType(time.Weekday(0)).Min(0).Max(6).Annotations(entgql.OrderField("WEEKDAY"), entgql.Type("Weekday")),
		field.Time("starts_at").SchemaType(map[string]string{dialect.Postgres: "time"}).Annotations(entgql.OrderField("STARTS_AT")),
		field.Int64("duration").GoType(time.Duration(0)).Annotations(entgql.OrderField("DURATION"), entgql.Type("Duration")),
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
