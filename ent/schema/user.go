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

const (
	passwordMinLength = 6
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name").NotEmpty().Annotations(entgql.OrderField("NAME")),
		field.String("username").Unique().NotEmpty().Annotations(entgql.OrderField("USERNAME")),
		field.String("password").MinLen(passwordMinLength),
		field.String("phone").NotEmpty().Annotations(entgql.OrderField("PHONE")),
		field.String("image").Optional(),
		field.String("directory").NotEmpty(),
		field.Int("token_version").Default(0),
		field.Enum("role").NamedValues(
			"SuperAdmin", "SUPER_ADMIN",
			"SchoolAdmin", "SCHOOL_ADMIN",
			"Teacher", "TEACHER",
			"Student", "STUDENT",
		).Default("STUDENT").Annotations(entgql.OrderField("ROLE")),
		field.Enum("status").GoType(Status("")).Default(StatusActive.String()).Annotations(entgql.OrderField("STATUS")),
		field.Time("deleted_at").Nillable().Optional(),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("stage", Stage.Type).Ref("students").Unique(),
		edge.From("school", School.Type).Ref("users").Unique(),
		edge.To("classes", Class.Type),
		edge.To("messages", Message.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("submissions", AssignmentSubmission.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("attendances", Attendance.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("payments", TuitionPayment.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("grades", Grade.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("stage"),
		index.Edges("school"),
		index.Fields("status"),
		index.Fields("role"),
	}
}
