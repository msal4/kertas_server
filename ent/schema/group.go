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

// Group holds the schema definition for the Group entity.
type Group struct {
	ent.Schema
}

// Fields of the Group.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name").Optional().Annotations(entgql.OrderField("NAME")),
		field.Enum("group_type").NamedValues(
			"Private", "PRIVATE",
			"Shared", "SHARED",
		).Default("SHARED").Annotations(entgql.OrderField("GROUP_TYPE")),
		field.Bool("active").Default(true),
		field.Time("deleted_at").Nillable().Optional(),
	}
}

func (Group) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Edges of the Group.
func (Group) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("class", Class.Type).Ref("group").Unique(),
		edge.From("users", User.Type).Ref("groups"),
		edge.To("messages", Message.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}

func (Group) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("class"),
		index.Fields("active"),
		index.Fields("group_type"),
		index.Fields("deleted_at"),
	}
}
