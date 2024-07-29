package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Category holds the schema definition for the Category entity.
type Category struct {
	ent.Schema
}

// Fields of the Category.
func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.String("name"),
	}
}

// Edges of the Category.
func (Category) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("owned_categories").
			Unique().
			Required(),
		edge.To("subjects", Subject.Type),
	}
}

// Indexes of the Category.
func (Category) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("user"),
	}
}
