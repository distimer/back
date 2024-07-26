package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Immutable().Unique(),
		field.String("name").Default("유저"),
		field.String("oauth_id"),
		field.Int8("oauth_provider"),
		field.Time("created_at").Immutable().Default(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("joined_groups", Group.Type).Through("affiliations", Affiliation.Type),
		edge.To("owned_groups", Group.Type),
		edge.To("study_logs", StudyLog.Type),
		edge.To("refresh_tokens", RefreshToken.Type),
		edge.To("owned_categories", Category.Type),
	}
}

// Indexes of the User.
func (User) Indexes() []ent.Index {
	return []ent.Index{
		// unique index
		index.Fields("oauth_id", "oauth_provider").
			Unique(),
	}
}
