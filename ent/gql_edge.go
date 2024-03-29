// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (t *Tweet) LikedUsers(ctx context.Context) (result []*User, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = t.NamedLikedUsers(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = t.Edges.LikedUsersOrErr()
	}
	if IsNotLoaded(err) {
		result, err = t.QueryLikedUsers().All(ctx)
	}
	return result, err
}

func (u *User) LikedTweets(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int,
) (*TweetConnection, error) {
	opts := []TweetPaginateOption{}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := u.Edges.totalCount[0][alias]
	if nodes, err := u.NamedLikedTweets(alias); err == nil || hasTotalCount {
		pager, err := newTweetPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &TweetConnection{Edges: []*TweetEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return u.QueryLikedTweets().Paginate(ctx, after, first, before, last, opts...)
}
