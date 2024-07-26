package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Affiliation holds the schema definition for the Affiliation entity.
type Affiliation struct {
	ent.Schema
}

func (Affiliation) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("user_id", "group_id"),
	}
}

// Fields of the Affiliation.
func (Affiliation) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("user_id", uuid.UUID{}),
		field.UUID("group_id", uuid.UUID{}),
		field.String("nickname"),
		field.Int8("role"),
		field.Time("joined_at").Immutable().Default(time.Now),
	}
}

// Edges of the Affiliation.
func (Affiliation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Unique().
			Required().
			Field("user_id"),
		edge.To("group", Group.Type).
			Unique().
			Required().
			Field("group_id"),
	}
}
