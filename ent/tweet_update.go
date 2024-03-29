// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"entdemo/ent/predicate"
	"entdemo/ent/tweet"
	"entdemo/ent/user"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TweetUpdate is the builder for updating Tweet entities.
type TweetUpdate struct {
	config
	hooks    []Hook
	mutation *TweetMutation
}

// Where appends a list predicates to the TweetUpdate builder.
func (tu *TweetUpdate) Where(ps ...predicate.Tweet) *TweetUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetText sets the "text" field.
func (tu *TweetUpdate) SetText(s string) *TweetUpdate {
	tu.mutation.SetText(s)
	return tu
}

// AddLikedUserIDs adds the "liked_users" edge to the User entity by IDs.
func (tu *TweetUpdate) AddLikedUserIDs(ids ...int) *TweetUpdate {
	tu.mutation.AddLikedUserIDs(ids...)
	return tu
}

// AddLikedUsers adds the "liked_users" edges to the User entity.
func (tu *TweetUpdate) AddLikedUsers(u ...*User) *TweetUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tu.AddLikedUserIDs(ids...)
}

// Mutation returns the TweetMutation object of the builder.
func (tu *TweetUpdate) Mutation() *TweetMutation {
	return tu.mutation
}

// ClearLikedUsers clears all "liked_users" edges to the User entity.
func (tu *TweetUpdate) ClearLikedUsers() *TweetUpdate {
	tu.mutation.ClearLikedUsers()
	return tu
}

// RemoveLikedUserIDs removes the "liked_users" edge to User entities by IDs.
func (tu *TweetUpdate) RemoveLikedUserIDs(ids ...int) *TweetUpdate {
	tu.mutation.RemoveLikedUserIDs(ids...)
	return tu
}

// RemoveLikedUsers removes "liked_users" edges to User entities.
func (tu *TweetUpdate) RemoveLikedUsers(u ...*User) *TweetUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tu.RemoveLikedUserIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TweetUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TweetUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TweetUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TweetUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tu *TweetUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(tweet.Table, tweet.Columns, sqlgraph.NewFieldSpec(tweet.FieldID, field.TypeInt))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Text(); ok {
		_spec.SetField(tweet.FieldText, field.TypeString, value)
	}
	if tu.mutation.LikedUsersCleared() {
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
		createE := &LikeCreate{config: tu.config, mutation: newLikeMutation(tu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedLikedUsersIDs(); len(nodes) > 0 && !tu.mutation.LikedUsersCleared() {
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
		createE := &LikeCreate{config: tu.config, mutation: newLikeMutation(tu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.LikedUsersIDs(); len(nodes) > 0 {
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
		createE := &LikeCreate{config: tu.config, mutation: newLikeMutation(tu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tweet.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TweetUpdateOne is the builder for updating a single Tweet entity.
type TweetUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TweetMutation
}

// SetText sets the "text" field.
func (tuo *TweetUpdateOne) SetText(s string) *TweetUpdateOne {
	tuo.mutation.SetText(s)
	return tuo
}

// AddLikedUserIDs adds the "liked_users" edge to the User entity by IDs.
func (tuo *TweetUpdateOne) AddLikedUserIDs(ids ...int) *TweetUpdateOne {
	tuo.mutation.AddLikedUserIDs(ids...)
	return tuo
}

// AddLikedUsers adds the "liked_users" edges to the User entity.
func (tuo *TweetUpdateOne) AddLikedUsers(u ...*User) *TweetUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tuo.AddLikedUserIDs(ids...)
}

// Mutation returns the TweetMutation object of the builder.
func (tuo *TweetUpdateOne) Mutation() *TweetMutation {
	return tuo.mutation
}

// ClearLikedUsers clears all "liked_users" edges to the User entity.
func (tuo *TweetUpdateOne) ClearLikedUsers() *TweetUpdateOne {
	tuo.mutation.ClearLikedUsers()
	return tuo
}

// RemoveLikedUserIDs removes the "liked_users" edge to User entities by IDs.
func (tuo *TweetUpdateOne) RemoveLikedUserIDs(ids ...int) *TweetUpdateOne {
	tuo.mutation.RemoveLikedUserIDs(ids...)
	return tuo
}

// RemoveLikedUsers removes "liked_users" edges to User entities.
func (tuo *TweetUpdateOne) RemoveLikedUsers(u ...*User) *TweetUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tuo.RemoveLikedUserIDs(ids...)
}

// Where appends a list predicates to the TweetUpdate builder.
func (tuo *TweetUpdateOne) Where(ps ...predicate.Tweet) *TweetUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TweetUpdateOne) Select(field string, fields ...string) *TweetUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Tweet entity.
func (tuo *TweetUpdateOne) Save(ctx context.Context) (*Tweet, error) {
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TweetUpdateOne) SaveX(ctx context.Context) *Tweet {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TweetUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TweetUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tuo *TweetUpdateOne) sqlSave(ctx context.Context) (_node *Tweet, err error) {
	_spec := sqlgraph.NewUpdateSpec(tweet.Table, tweet.Columns, sqlgraph.NewFieldSpec(tweet.FieldID, field.TypeInt))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Tweet.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tweet.FieldID)
		for _, f := range fields {
			if !tweet.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tweet.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Text(); ok {
		_spec.SetField(tweet.FieldText, field.TypeString, value)
	}
	if tuo.mutation.LikedUsersCleared() {
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
		createE := &LikeCreate{config: tuo.config, mutation: newLikeMutation(tuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedLikedUsersIDs(); len(nodes) > 0 && !tuo.mutation.LikedUsersCleared() {
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
		createE := &LikeCreate{config: tuo.config, mutation: newLikeMutation(tuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.LikedUsersIDs(); len(nodes) > 0 {
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
		createE := &LikeCreate{config: tuo.config, mutation: newLikeMutation(tuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Tweet{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tweet.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
