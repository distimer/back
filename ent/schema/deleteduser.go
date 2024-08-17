package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// DeletedUser holds the schema definition for the DeletedUser entity.
type DeletedUser struct {
	ent.Schema
}

// Fields of the DeletedUser.
func (DeletedUser) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Immutable().Unique(),
		field.String("name"),
		field.String("oauth_id"),
		field.Int8("oauth_provider"),
		field.Time("created_at").Immutable(),
		field.Time("deleted_at").Immutable().Default(time.Now),
	}
}

// Edges of the DeletedUser.
func (DeletedUser) Edges() []ent.Edge {
	return nil
}
