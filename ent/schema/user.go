package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
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
		field.String("name").NotEmpty(),
		field.String("username").Unique().NotEmpty(),
		field.String("password").MinLen(passwordMinLength),
		field.String("phone").NotEmpty(),
		field.String("image").Optional(),
		field.Int("token_version").Default(0),
		field.Enum("role").Values("super_admin", "school_admin", "teacher", "student").Default("student"),
		field.Enum("status").GoType(Status("")).Default("active"),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
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
