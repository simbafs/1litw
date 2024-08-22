package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Record holds the schema definition for the Record entity.
type Record struct {
	ent.Schema
}

// Fields of the Record.
func (Record) Fields() []ent.Field {
	return []ent.Field{
		field.String("code").NotEmpty(),
		field.String("target").NotEmpty(),
		field.Time("created_at").Immutable().Default(func() time.Time {
			return time.Now()
		}),
	}
}

// Edges of the Record.
func (Record) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("records").
			Unique(),
	}
}
