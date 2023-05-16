// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"entdemo/ent/tweet"
	"entdemo/ent/user"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TweetCreate is the builder for creating a Tweet entity.
type TweetCreate struct {
	config
	mutation *TweetMutation
	hooks    []Hook
}

// SetText sets the "text" field.
func (tc *TweetCreate) SetText(s string) *TweetCreate {
	tc.mutation.SetText(s)
	return tc
}

// AddLikedUserIDs adds the "liked_users" edge to the User entity by IDs.
func (tc *TweetCreate) AddLikedUserIDs(ids ...int) *TweetCreate {
	tc.mutation.AddLikedUserIDs(ids...)
	return tc
}

// AddLikedUsers adds the "liked_users" edges to the User entity.
func (tc *TweetCreate) AddLikedUsers(u ...*User) *TweetCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tc.AddLikedUserIDs(ids...)
}

// Mutation returns the TweetMutation object of the builder.
func (tc *TweetCreate) Mutation() *TweetMutation {
	return tc.mutation
}

// Save creates the Tweet in the database.
func (tc *TweetCreate) Save(ctx context.Context) (*Tweet, error) {
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TweetCreate) SaveX(ctx context.Context) *Tweet {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TweetCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TweetCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TweetCreate) check() error {
	if _, ok := tc.mutation.Text(); !ok {
		return &ValidationError{Name: "text", err: errors.New(`ent: missing required field "Tweet.text"`)}
	}
	return nil
}

func (tc *TweetCreate) sqlSave(ctx context.Context) (*Tweet, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TweetCreate) createSpec() (*Tweet, *sqlgraph.CreateSpec) {
	var (
		_node = &Tweet{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(tweet.Table, sqlgraph.NewFieldSpec(tweet.FieldID, field.TypeInt))
	)
	if value, ok := tc.mutation.Text(); ok {
		_spec.SetField(tweet.FieldText, field.TypeString, value)
		_node.Text = value
	}
	if nodes := tc.mutation.LikedUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tweet.LikedUsersTable,
			Columns: tweet.LikedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &LikeCreate{config: tc.config, mutation: newLikeMutation(tc.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TweetCreateBulk is the builder for creating many Tweet entities in bulk.
type TweetCreateBulk struct {
	config
	builders []*TweetCreate
}

// Save creates the Tweet entities in the database.
func (tcb *TweetCreateBulk) Save(ctx context.Context) ([]*Tweet, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Tweet, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TweetMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TweetCreateBulk) SaveX(ctx context.Context) []*Tweet {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TweetCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TweetCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
