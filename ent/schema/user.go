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
		field.Int64("userid").Unique().Immutable(),

		// permissions
		field.Bool("superAdmin").Default(false).Comment("If the user can manage users' permissions"),
		field.Bool("admin").Default(false).Comment("If the user can manipulate users' records"),
		field.Bool("create").Default(true).Comment("If the user can create records"),
		field.Bool("customCode").Default(false).Comment("If the user can use custom code"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("records", Record.Type),
	}
}
