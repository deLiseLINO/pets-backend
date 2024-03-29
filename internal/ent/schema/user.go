package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Unique(),
		field.String("email").Unique(),
		field.String("unique_name").Unique(),
		field.String("name"),
		field.String("password"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
