package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// OtpCodes holds the schema definition for the OtpCodes entity.
type OtpCodes struct {
	ent.Schema
}

// Fields of the OtpCodes.
func (OtpCodes) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Unique(),
		field.String("code"),
		field.String("email"),
		field.Time("next_send_time"),
		field.Time("exparation_time"),
	}
}

// Edges of the OtpCodes.
func (OtpCodes) Edges() []ent.Edge {
	return nil
}
