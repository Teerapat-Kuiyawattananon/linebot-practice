// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"entdemo/ent/creditlater"
	"entdemo/ent/linelog"
	"entdemo/ent/lineuser"
	"entdemo/ent/predicate"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LineUserQuery is the builder for querying LineUser entities.
type LineUserQuery struct {
	config
	ctx              *QueryContext
	order            []lineuser.OrderOption
	inters           []Interceptor
	predicates       []predicate.LineUser
	withLinelogs     *LineLogQuery
	withCreditlaters *CreditLaterQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LineUserQuery builder.
func (luq *LineUserQuery) Where(ps ...predicate.LineUser) *LineUserQuery {
	luq.predicates = append(luq.predicates, ps...)
	return luq
}

// Limit the number of records to be returned by this query.
func (luq *LineUserQuery) Limit(limit int) *LineUserQuery {
	luq.ctx.Limit = &limit
	return luq
}

// Offset to start from.
func (luq *LineUserQuery) Offset(offset int) *LineUserQuery {
	luq.ctx.Offset = &offset
	return luq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (luq *LineUserQuery) Unique(unique bool) *LineUserQuery {
	luq.ctx.Unique = &unique
	return luq
}

// Order specifies how the records should be ordered.
func (luq *LineUserQuery) Order(o ...lineuser.OrderOption) *LineUserQuery {
	luq.order = append(luq.order, o...)
	return luq
}

// QueryLinelogs chains the current query on the "linelogs" edge.
func (luq *LineUserQuery) QueryLinelogs() *LineLogQuery {
	query := (&LineLogClient{config: luq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := luq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := luq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(lineuser.Table, lineuser.FieldID, selector),
			sqlgraph.To(linelog.Table, linelog.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, lineuser.LinelogsTable, lineuser.LinelogsColumn),
		)
		fromU = sqlgraph.SetNeighbors(luq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCreditlaters chains the current query on the "creditlaters" edge.
func (luq *LineUserQuery) QueryCreditlaters() *CreditLaterQuery {
	query := (&CreditLaterClient{config: luq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := luq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := luq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(lineuser.Table, lineuser.FieldID, selector),
			sqlgraph.To(creditlater.Table, creditlater.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, lineuser.CreditlatersTable, lineuser.CreditlatersColumn),
		)
		fromU = sqlgraph.SetNeighbors(luq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first LineUser entity from the query.
// Returns a *NotFoundError when no LineUser was found.
func (luq *LineUserQuery) First(ctx context.Context) (*LineUser, error) {
	nodes, err := luq.Limit(1).All(setContextOp(ctx, luq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{lineuser.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (luq *LineUserQuery) FirstX(ctx context.Context) *LineUser {
	node, err := luq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first LineUser ID from the query.
// Returns a *NotFoundError when no LineUser ID was found.
func (luq *LineUserQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = luq.Limit(1).IDs(setContextOp(ctx, luq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{lineuser.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (luq *LineUserQuery) FirstIDX(ctx context.Context) int {
	id, err := luq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single LineUser entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one LineUser entity is found.
// Returns a *NotFoundError when no LineUser entities are found.
func (luq *LineUserQuery) Only(ctx context.Context) (*LineUser, error) {
	nodes, err := luq.Limit(2).All(setContextOp(ctx, luq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{lineuser.Label}
	default:
		return nil, &NotSingularError{lineuser.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (luq *LineUserQuery) OnlyX(ctx context.Context) *LineUser {
	node, err := luq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only LineUser ID in the query.
// Returns a *NotSingularError when more than one LineUser ID is found.
// Returns a *NotFoundError when no entities are found.
func (luq *LineUserQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = luq.Limit(2).IDs(setContextOp(ctx, luq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{lineuser.Label}
	default:
		err = &NotSingularError{lineuser.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (luq *LineUserQuery) OnlyIDX(ctx context.Context) int {
	id, err := luq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of LineUsers.
func (luq *LineUserQuery) All(ctx context.Context) ([]*LineUser, error) {
	ctx = setContextOp(ctx, luq.ctx, "All")
	if err := luq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*LineUser, *LineUserQuery]()
	return withInterceptors[[]*LineUser](ctx, luq, qr, luq.inters)
}

// AllX is like All, but panics if an error occurs.
func (luq *LineUserQuery) AllX(ctx context.Context) []*LineUser {
	nodes, err := luq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of LineUser IDs.
func (luq *LineUserQuery) IDs(ctx context.Context) (ids []int, err error) {
	if luq.ctx.Unique == nil && luq.path != nil {
		luq.Unique(true)
	}
	ctx = setContextOp(ctx, luq.ctx, "IDs")
	if err = luq.Select(lineuser.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (luq *LineUserQuery) IDsX(ctx context.Context) []int {
	ids, err := luq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (luq *LineUserQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, luq.ctx, "Count")
	if err := luq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, luq, querierCount[*LineUserQuery](), luq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (luq *LineUserQuery) CountX(ctx context.Context) int {
	count, err := luq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (luq *LineUserQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, luq.ctx, "Exist")
	switch _, err := luq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (luq *LineUserQuery) ExistX(ctx context.Context) bool {
	exist, err := luq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LineUserQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (luq *LineUserQuery) Clone() *LineUserQuery {
	if luq == nil {
		return nil
	}
	return &LineUserQuery{
		config:           luq.config,
		ctx:              luq.ctx.Clone(),
		order:            append([]lineuser.OrderOption{}, luq.order...),
		inters:           append([]Interceptor{}, luq.inters...),
		predicates:       append([]predicate.LineUser{}, luq.predicates...),
		withLinelogs:     luq.withLinelogs.Clone(),
		withCreditlaters: luq.withCreditlaters.Clone(),
		// clone intermediate query.
		sql:  luq.sql.Clone(),
		path: luq.path,
	}
}

// WithLinelogs tells the query-builder to eager-load the nodes that are connected to
// the "linelogs" edge. The optional arguments are used to configure the query builder of the edge.
func (luq *LineUserQuery) WithLinelogs(opts ...func(*LineLogQuery)) *LineUserQuery {
	query := (&LineLogClient{config: luq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	luq.withLinelogs = query
	return luq
}

// WithCreditlaters tells the query-builder to eager-load the nodes that are connected to
// the "creditlaters" edge. The optional arguments are used to configure the query builder of the edge.
func (luq *LineUserQuery) WithCreditlaters(opts ...func(*CreditLaterQuery)) *LineUserQuery {
	query := (&CreditLaterClient{config: luq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	luq.withCreditlaters = query
	return luq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserId string `json:"userId,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.LineUser.Query().
//		GroupBy(lineuser.FieldUserId).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (luq *LineUserQuery) GroupBy(field string, fields ...string) *LineUserGroupBy {
	luq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &LineUserGroupBy{build: luq}
	grbuild.flds = &luq.ctx.Fields
	grbuild.label = lineuser.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UserId string `json:"userId,omitempty"`
//	}
//
//	client.LineUser.Query().
//		Select(lineuser.FieldUserId).
//		Scan(ctx, &v)
func (luq *LineUserQuery) Select(fields ...string) *LineUserSelect {
	luq.ctx.Fields = append(luq.ctx.Fields, fields...)
	sbuild := &LineUserSelect{LineUserQuery: luq}
	sbuild.label = lineuser.Label
	sbuild.flds, sbuild.scan = &luq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a LineUserSelect configured with the given aggregations.
func (luq *LineUserQuery) Aggregate(fns ...AggregateFunc) *LineUserSelect {
	return luq.Select().Aggregate(fns...)
}

func (luq *LineUserQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range luq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, luq); err != nil {
				return err
			}
		}
	}
	for _, f := range luq.ctx.Fields {
		if !lineuser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if luq.path != nil {
		prev, err := luq.path(ctx)
		if err != nil {
			return err
		}
		luq.sql = prev
	}
	return nil
}

func (luq *LineUserQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*LineUser, error) {
	var (
		nodes       = []*LineUser{}
		_spec       = luq.querySpec()
		loadedTypes = [2]bool{
			luq.withLinelogs != nil,
			luq.withCreditlaters != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*LineUser).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &LineUser{config: luq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, luq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := luq.withLinelogs; query != nil {
		if err := luq.loadLinelogs(ctx, query, nodes,
			func(n *LineUser) { n.Edges.Linelogs = []*LineLog{} },
			func(n *LineUser, e *LineLog) { n.Edges.Linelogs = append(n.Edges.Linelogs, e) }); err != nil {
			return nil, err
		}
	}
	if query := luq.withCreditlaters; query != nil {
		if err := luq.loadCreditlaters(ctx, query, nodes, nil,
			func(n *LineUser, e *CreditLater) { n.Edges.Creditlaters = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (luq *LineUserQuery) loadLinelogs(ctx context.Context, query *LineLogQuery, nodes []*LineUser, init func(*LineUser), assign func(*LineUser, *LineLog)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*LineUser)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.LineLog(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(lineuser.LinelogsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.line_user_linelogs
		if fk == nil {
			return fmt.Errorf(`foreign-key "line_user_linelogs" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "line_user_linelogs" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (luq *LineUserQuery) loadCreditlaters(ctx context.Context, query *CreditLaterQuery, nodes []*LineUser, init func(*LineUser), assign func(*LineUser, *CreditLater)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*LineUser)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	query.withFKs = true
	query.Where(predicate.CreditLater(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(lineuser.CreditlatersColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.line_user_creditlaters
		if fk == nil {
			return fmt.Errorf(`foreign-key "line_user_creditlaters" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "line_user_creditlaters" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (luq *LineUserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := luq.querySpec()
	_spec.Node.Columns = luq.ctx.Fields
	if len(luq.ctx.Fields) > 0 {
		_spec.Unique = luq.ctx.Unique != nil && *luq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, luq.driver, _spec)
}

func (luq *LineUserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(lineuser.Table, lineuser.Columns, sqlgraph.NewFieldSpec(lineuser.FieldID, field.TypeInt))
	_spec.From = luq.sql
	if unique := luq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if luq.path != nil {
		_spec.Unique = true
	}
	if fields := luq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, lineuser.FieldID)
		for i := range fields {
			if fields[i] != lineuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := luq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := luq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := luq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := luq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (luq *LineUserQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(luq.driver.Dialect())
	t1 := builder.Table(lineuser.Table)
	columns := luq.ctx.Fields
	if len(columns) == 0 {
		columns = lineuser.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if luq.sql != nil {
		selector = luq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if luq.ctx.Unique != nil && *luq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range luq.predicates {
		p(selector)
	}
	for _, p := range luq.order {
		p(selector)
	}
	if offset := luq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := luq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// LineUserGroupBy is the group-by builder for LineUser entities.
type LineUserGroupBy struct {
	selector
	build *LineUserQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lugb *LineUserGroupBy) Aggregate(fns ...AggregateFunc) *LineUserGroupBy {
	lugb.fns = append(lugb.fns, fns...)
	return lugb
}

// Scan applies the selector query and scans the result into the given value.
func (lugb *LineUserGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lugb.build.ctx, "GroupBy")
	if err := lugb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LineUserQuery, *LineUserGroupBy](ctx, lugb.build, lugb, lugb.build.inters, v)
}

func (lugb *LineUserGroupBy) sqlScan(ctx context.Context, root *LineUserQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(lugb.fns))
	for _, fn := range lugb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*lugb.flds)+len(lugb.fns))
		for _, f := range *lugb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*lugb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lugb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// LineUserSelect is the builder for selecting fields of LineUser entities.
type LineUserSelect struct {
	*LineUserQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (lus *LineUserSelect) Aggregate(fns ...AggregateFunc) *LineUserSelect {
	lus.fns = append(lus.fns, fns...)
	return lus
}

// Scan applies the selector query and scans the result into the given value.
func (lus *LineUserSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lus.ctx, "Select")
	if err := lus.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LineUserQuery, *LineUserSelect](ctx, lus.LineUserQuery, lus, lus.inters, v)
}

func (lus *LineUserSelect) sqlScan(ctx context.Context, root *LineUserQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(lus.fns))
	for _, fn := range lus.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*lus.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lus.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
