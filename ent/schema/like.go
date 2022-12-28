package schema

import (
	"context"
	gen "entdemo/ent"
	"entdemo/ent/hook"
	"fmt"
	"time"

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

func (Like) Hooks() []ent.Hook {
	return []ent.Hook{
		func(next ent.Mutator) ent.Mutator {
			return hook.LikeFunc(func(ctx context.Context, m *gen.LikeMutation) (ent.Value, error) {
				// "hello" will not be output
				fmt.Println("hello")

				return next.Mutate(ctx, m)
			})
		},
	}
}
