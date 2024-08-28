package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").NotEmpty(),
		field.Int64("tgid").Unique().Immutable(),

		// permissions
		field.Bool("customCode").Default(false).Comment("If the user can use custom code"),
		field.Bool("admin").Default(false).Comment("if the user can read, delete and modify all records"),
		field.Bool("readAll").Default(false).Comment("If the user can read all records"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("records", Record.Type),
	}
}
