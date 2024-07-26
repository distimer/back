package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// InviteCode holds the schema definition for the InviteCode entity.
type InviteCode struct {
	ent.Schema
}

// Fields of the InviteCode.
func (InviteCode) Fields() []ent.Field {
	return []ent.Field{
		field.String("code").Unique(),
		field.Int32("used").Default(0),
	}
}

// Edges of the InviteCode.
func (InviteCode) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("group", Group.Type).Ref("invite_codes").Unique().Required(),
	}
}
