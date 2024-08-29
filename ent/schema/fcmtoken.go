package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// FCMToken holds the schema definition for the FCMToken entity.
type FCMToken struct {
	ent.Schema
}

// Fields of the FCMToken.
func (FCMToken) Fields() []ent.Field {
	return []ent.Field{
		field.String("push_token").Unique(),
	}
}

// Edges of the FCMToken.
func (FCMToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("session", Session.Type).
			Ref("fcm_token").
			Unique().
			Required(),
	}
}
