package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Session holds the schema definition for the Session entity.
type Session struct {
	ent.Schema
}

// Fields of the Session.
func (Session) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Immutable().Default(uuid.New).Unique(),
		field.UUID("refresh_token", uuid.UUID{}).Default(uuid.New).Unique(),
		field.Time("created_at").Immutable().Default(time.Now),
		field.Int8("device_type").Immutable(),
		field.Time("last_active").Default(time.Now),
	}
}

// Edges of the Session.
func (Session) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("sessions").
			Unique().
			Required(),
		edge.To("apns_token", APNsToken.Type).
			Unique(),
		edge.To("fcm_token", FCMToken.Type).
			Unique(),
	}
}

// Indexes of the Session.
func (Session) Indexes() []ent.Index {
	return []ent.Index{
		// unique index.
		index.Fields("refresh_token"),
	}
}
