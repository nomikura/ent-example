package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Default("Unknown"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("liked_tweets", Tweet.Type).
			Through("likes", Like.Type).
			Annotations(entgql.RelayConnection()),
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
	}
}
