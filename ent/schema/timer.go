package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Timer holds the schema definition for the Timer entity.
type Timer struct {
	ent.Schema
}

// Fields of the Timer.
func (Timer) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Immutable().Default(uuid.New).Unique(),
		field.Time("start_at").Default(time.Now),
		field.String("content"),
		field.UUID("user_id", uuid.UUID{}).Unique(),
	}
}

// Edges of the Timer.
func (Timer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("timers").
			Unique().
			Field("user_id").
			Required(),

		edge.From("subject", Subject.Type).
			Ref("timers").
			Unique().
			Required(),

		edge.To("shared_group", Group.Type),
	}
}
