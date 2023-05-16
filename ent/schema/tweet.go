package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Tweet struct {
	ent.Schema
}

// Fields of the Tweet.
func (Tweet) Fields() []ent.Field {
	return []ent.Field{
		field.Text("text"),
	}
}

// Edges of the Tweet.
func (Tweet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("liked_users", User.Type).
			Ref("liked_tweets").
			Through("likes", Like.Type),
	}
}

func (Tweet) Annotation() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
	}
}
