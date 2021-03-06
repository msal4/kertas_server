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

// School holds the schema definition for the School entity.
type School struct {
	ent.Schema
}

// Fields of the School.
func (School) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name").NotEmpty().Annotations(entgql.OrderField("NAME")),
		field.String("image").NotEmpty(),
		field.String("directory").NotEmpty(),
		field.Bool("active").Default(true),
		field.Time("deleted_at").Nillable().Optional(),
	}
}

func (School) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Edges of the School.
func (School) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}, entgql.Bind()),
		edge.To("stages", Stage.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}, entgql.Bind()),
	}
}

func (School) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("active"),
		index.Fields("deleted_at"),
	}
}
