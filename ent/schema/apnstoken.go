package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// APNsToken holds the schema definition for the APNsToken entity.
type APNsToken struct {
	ent.Schema
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "apns_tokens"},
	}
}

// Fields of the APNsToken.
func (APNsToken) Fields() []ent.Field {
	return []ent.Field{
		field.String("start_token").Unique(),
		field.String("update_token").Unique(),
	}
}

// Edges of the APNsToken.
func (APNsToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("session", Session.Type).
			Ref("apns_token").
			Unique().
			Required(),
	}
}
