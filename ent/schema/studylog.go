package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// StudyLog holds the schema definition for the StudyLog entity.
type StudyLog struct {
	ent.Schema
}

// Fields of the StudyLog.
func (StudyLog) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Immutable().Default(uuid.New).Unique(),
		field.Time("start_at"),
		field.Time("end_at"),
		field.String("content"),
	}
}

// Edges of the StudyLog.
func (StudyLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("study_logs").
			Unique().
			Required(),

		edge.From("category", Category.Type).
			Ref("study_logs").
			Unique().
			Required(),

		edge.To("shared_group", Group.Type),
	}
}

// Indexes of the StudyLog.
func (StudyLog) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("user"),
	}
}
