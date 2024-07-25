package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Group holds the schema definition for the Group entity.
type Group struct {
	ent.Schema
}

// Fields of the Group.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Immutable().Default(uuid.New).Unique(),
		field.String("name"),
		field.String("description").Default(""),
		field.String("nickname_policy").Default(""),
		field.Int8("reveal_policy"),
		field.Int8("invite_policy"),
		field.Time("created_at").Immutable().Default(time.Now),
	}
}

// Edges of the Group.
func (Group) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("members", User.Type).Ref("joined_groups"),
		edge.From("owner", User.Type).
			Ref("owned_groups").
			Required().
			Unique(),
		edge.From("shared_study_logs", StudyLog.Type).
			Ref("shared_group"),
		edge.To("invite_codes", InviteCode.Type),
	}
}
