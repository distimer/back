package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Subject holds the schema definition for the Subject entity.
type Subject struct {
	ent.Schema
}

// Fields of the Subject.
func (Subject) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.String("name"),
		field.String("color"),
		field.Int8("order").Default(-1),
	}
}

// Edges of the Subject.
func (Subject) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("category", Category.Type).
			Ref("subjects").
			Unique().
			Required(),
		edge.To("study_logs", StudyLog.Type),
		edge.To("timers", Timer.Type),
	}
}

// Indexes of the Subject.
func (Subject) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("category"),
	}
}
