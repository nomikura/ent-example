// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"entdemo/ent/like"
	"entdemo/ent/predicate"
	"entdemo/ent/tweet"
	"entdemo/ent/user"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// LikeQuery is the builder for querying Like entities.
type LikeQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	inters     []Interceptor
	predicates []predicate.Like
	withUser   *UserQuery
	withTweet  *TweetQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LikeQuery builder.
func (lq *LikeQuery) Where(ps ...predicate.Like) *LikeQuery {
	lq.predicates = append(lq.predicates, ps...)
	return lq
}

// Limit the number of records to be returned by this query.
func (lq *LikeQuery) Limit(limit int) *LikeQuery {
	lq.limit = &limit
	return lq
}

// Offset to start from.
func (lq *LikeQuery) Offset(offset int) *LikeQuery {
	lq.offset = &offset
	return lq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (lq *LikeQuery) Unique(unique bool) *LikeQuery {
	lq.unique = &unique
	return lq
}

// Order specifies how the records should be ordered.
func (lq *LikeQuery) Order(o ...OrderFunc) *LikeQuery {
	lq.order = append(lq.order, o...)
	return lq
}

// QueryUser chains the current query on the "user" edge.
func (lq *LikeQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: lq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(like.Table, like.UserColumn, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, like.UserTable, like.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(lq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTweet chains the current query on the "tweet" edge.
func (lq *LikeQuery) QueryTweet() *TweetQuery {
	query := (&TweetClient{config: lq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(like.Table, like.TweetColumn, selector),
			sqlgraph.To(tweet.Table, tweet.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, like.TweetTable, like.TweetColumn),
		)
		fromU = sqlgraph.SetNeighbors(lq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Like entity from the query.
// Returns a *NotFoundError when no Like was found.
func (lq *LikeQuery) First(ctx context.Context) (*Like, error) {
	nodes, err := lq.Limit(1).All(newQueryContext(ctx, TypeLike, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{like.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (lq *LikeQuery) FirstX(ctx context.Context) *Like {
	node, err := lq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// Only returns a single Like entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Like entity is found.
// Returns a *NotFoundError when no Like entities are found.
func (lq *LikeQuery) Only(ctx context.Context) (*Like, error) {
	nodes, err := lq.Limit(2).All(newQueryContext(ctx, TypeLike, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{like.Label}
	default:
		return nil, &NotSingularError{like.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (lq *LikeQuery) OnlyX(ctx context.Context) *Like {
	node, err := lq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// All executes the query and returns a list of Likes.
func (lq *LikeQuery) All(ctx context.Context) ([]*Like, error) {
	ctx = newQueryContext(ctx, TypeLike, "All")
	if err := lq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Like, *LikeQuery]()
	return withInterceptors[[]*Like](ctx, lq, qr, lq.inters)
}

// AllX is like All, but panics if an error occurs.
func (lq *LikeQuery) AllX(ctx context.Context) []*Like {
	nodes, err := lq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// Count returns the count of the given query.
func (lq *LikeQuery) Count(ctx context.Context) (int, error) {
	ctx = newQueryContext(ctx, TypeLike, "Count")
	if err := lq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, lq, querierCount[*LikeQuery](), lq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (lq *LikeQuery) CountX(ctx context.Context) int {
	count, err := lq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (lq *LikeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = newQueryContext(ctx, TypeLike, "Exist")
	switch _, err := lq.First(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (lq *LikeQuery) ExistX(ctx context.Context) bool {
	exist, err := lq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LikeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (lq *LikeQuery) Clone() *LikeQuery {
	if lq == nil {
		return nil
	}
	return &LikeQuery{
		config:     lq.config,
		limit:      lq.limit,
		offset:     lq.offset,
		order:      append([]OrderFunc{}, lq.order...),
		predicates: append([]predicate.Like{}, lq.predicates...),
		withUser:   lq.withUser.Clone(),
		withTweet:  lq.withTweet.Clone(),
		// clone intermediate query.
		sql:    lq.sql.Clone(),
		path:   lq.path,
		unique: lq.unique,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (lq *LikeQuery) WithUser(opts ...func(*UserQuery)) *LikeQuery {
	query := (&UserClient{config: lq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	lq.withUser = query
	return lq
}

// WithTweet tells the query-builder to eager-load the nodes that are connected to
// the "tweet" edge. The optional arguments are used to configure the query builder of the edge.
func (lq *LikeQuery) WithTweet(opts ...func(*TweetQuery)) *LikeQuery {
	query := (&TweetClient{config: lq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	lq.withTweet = query
	return lq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		LikedAt time.Time `json:"liked_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Like.Query().
//		GroupBy(like.FieldLikedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (lq *LikeQuery) GroupBy(field string, fields ...string) *LikeGroupBy {
	lq.fields = append([]string{field}, fields...)
	grbuild := &LikeGroupBy{build: lq}
	grbuild.flds = &lq.fields
	grbuild.label = like.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		LikedAt time.Time `json:"liked_at,omitempty"`
//	}
//
//	client.Like.Query().
//		Select(like.FieldLikedAt).
//		Scan(ctx, &v)
func (lq *LikeQuery) Select(fields ...string) *LikeSelect {
	lq.fields = append(lq.fields, fields...)
	sbuild := &LikeSelect{LikeQuery: lq}
	sbuild.label = like.Label
	sbuild.flds, sbuild.scan = &lq.fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a LikeSelect configured with the given aggregations.
func (lq *LikeQuery) Aggregate(fns ...AggregateFunc) *LikeSelect {
	return lq.Select().Aggregate(fns...)
}

func (lq *LikeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range lq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, lq); err != nil {
				return err
			}
		}
	}
	for _, f := range lq.fields {
		if !like.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if lq.path != nil {
		prev, err := lq.path(ctx)
		if err != nil {
			return err
		}
		lq.sql = prev
	}
	return nil
}

func (lq *LikeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Like, error) {
	var (
		nodes       = []*Like{}
		_spec       = lq.querySpec()
		loadedTypes = [2]bool{
			lq.withUser != nil,
			lq.withTweet != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Like).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Like{config: lq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, lq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := lq.withUser; query != nil {
		if err := lq.loadUser(ctx, query, nodes, nil,
			func(n *Like, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := lq.withTweet; query != nil {
		if err := lq.loadTweet(ctx, query, nodes, nil,
			func(n *Like, e *Tweet) { n.Edges.Tweet = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (lq *LikeQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Like, init func(*Like), assign func(*Like, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Like)
	for i := range nodes {
		fk := nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (lq *LikeQuery) loadTweet(ctx context.Context, query *TweetQuery, nodes []*Like, init func(*Like), assign func(*Like, *Tweet)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Like)
	for i := range nodes {
		fk := nodes[i].TweetID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(tweet.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "tweet_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (lq *LikeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := lq.querySpec()
	_spec.Unique = false
	_spec.Node.Columns = nil
	return sqlgraph.CountNodes(ctx, lq.driver, _spec)
}

func (lq *LikeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   like.Table,
			Columns: like.Columns,
		},
		From:   lq.sql,
		Unique: true,
	}
	if unique := lq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := lq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		for i := range fields {
			_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
		}
	}
	if ps := lq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := lq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := lq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := lq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (lq *LikeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(lq.driver.Dialect())
	t1 := builder.Table(like.Table)
	columns := lq.fields
	if len(columns) == 0 {
		columns = like.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if lq.sql != nil {
		selector = lq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if lq.unique != nil && *lq.unique {
		selector.Distinct()
	}
	for _, p := range lq.predicates {
		p(selector)
	}
	for _, p := range lq.order {
		p(selector)
	}
	if offset := lq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := lq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// LikeGroupBy is the group-by builder for Like entities.
type LikeGroupBy struct {
	selector
	build *LikeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lgb *LikeGroupBy) Aggregate(fns ...AggregateFunc) *LikeGroupBy {
	lgb.fns = append(lgb.fns, fns...)
	return lgb
}

// Scan applies the selector query and scans the result into the given value.
func (lgb *LikeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeLike, "GroupBy")
	if err := lgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LikeQuery, *LikeGroupBy](ctx, lgb.build, lgb, lgb.build.inters, v)
}

func (lgb *LikeGroupBy) sqlScan(ctx context.Context, root *LikeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(lgb.fns))
	for _, fn := range lgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*lgb.flds)+len(lgb.fns))
		for _, f := range *lgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*lgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// LikeSelect is the builder for selecting fields of Like entities.
type LikeSelect struct {
	*LikeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ls *LikeSelect) Aggregate(fns ...AggregateFunc) *LikeSelect {
	ls.fns = append(ls.fns, fns...)
	return ls
}

// Scan applies the selector query and scans the result into the given value.
func (ls *LikeSelect) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeLike, "Select")
	if err := ls.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LikeQuery, *LikeSelect](ctx, ls.LikeQuery, ls, ls.inters, v)
}

func (ls *LikeSelect) sqlScan(ctx context.Context, root *LikeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ls.fns))
	for _, fn := range ls.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ls.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}