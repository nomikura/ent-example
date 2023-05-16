package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Like holds the edge schema definition for the Like edge.
type Like struct {
	ent.Schema
}

func (Like) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("user_id", "tweet_id"),
	}
}

func (Like) Fields() []ent.Field {
	return []ent.Field{
		field.Time("liked_at").
			Default(time.Now),
		field.Int("user_id"),
		field.Int("tweet_id"),
	}
}

func (Like) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Unique().
			Required().
			Field("user_id"),
		edge.To("tweet", Tweet.Type).
			Unique().
			Required().
			Field("tweet_id"),
	}
}

func (Like) Annotation() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
	}
}
