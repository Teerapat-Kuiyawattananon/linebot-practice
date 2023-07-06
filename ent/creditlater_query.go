// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"entdemo/ent/creditlater"
	"entdemo/ent/lineuser"
	"entdemo/ent/predicate"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CreditLaterQuery is the builder for querying CreditLater entities.
type CreditLaterQuery struct {
	config
	ctx        *QueryContext
	order      []creditlater.OrderOption
	inters     []Interceptor
	predicates []predicate.CreditLater
	withOwner  *LineUserQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CreditLaterQuery builder.
func (clq *CreditLaterQuery) Where(ps ...predicate.CreditLater) *CreditLaterQuery {
	clq.predicates = append(clq.predicates, ps...)
	return clq
}

// Limit the number of records to be returned by this query.
func (clq *CreditLaterQuery) Limit(limit int) *CreditLaterQuery {
	clq.ctx.Limit = &limit
	return clq
}

// Offset to start from.
func (clq *CreditLaterQuery) Offset(offset int) *CreditLaterQuery {
	clq.ctx.Offset = &offset
	return clq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (clq *CreditLaterQuery) Unique(unique bool) *CreditLaterQuery {
	clq.ctx.Unique = &unique
	return clq
}

// Order specifies how the records should be ordered.
func (clq *CreditLaterQuery) Order(o ...creditlater.OrderOption) *CreditLaterQuery {
	clq.order = append(clq.order, o...)
	return clq
}

// QueryOwner chains the current query on the "owner" edge.
func (clq *CreditLaterQuery) QueryOwner() *LineUserQuery {
	query := (&LineUserClient{config: clq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := clq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := clq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(creditlater.Table, creditlater.FieldID, selector),
			sqlgraph.To(lineuser.Table, lineuser.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, creditlater.OwnerTable, creditlater.OwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(clq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CreditLater entity from the query.
// Returns a *NotFoundError when no CreditLater was found.
func (clq *CreditLaterQuery) First(ctx context.Context) (*CreditLater, error) {
	nodes, err := clq.Limit(1).All(setContextOp(ctx, clq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{creditlater.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (clq *CreditLaterQuery) FirstX(ctx context.Context) *CreditLater {
	node, err := clq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CreditLater ID from the query.
// Returns a *NotFoundError when no CreditLater ID was found.
func (clq *CreditLaterQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = clq.Limit(1).IDs(setContextOp(ctx, clq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{creditlater.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (clq *CreditLaterQuery) FirstIDX(ctx context.Context) int {
	id, err := clq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CreditLater entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CreditLater entity is found.
// Returns a *NotFoundError when no CreditLater entities are found.
func (clq *CreditLaterQuery) Only(ctx context.Context) (*CreditLater, error) {
	nodes, err := clq.Limit(2).All(setContextOp(ctx, clq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{creditlater.Label}
	default:
		return nil, &NotSingularError{creditlater.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (clq *CreditLaterQuery) OnlyX(ctx context.Context) *CreditLater {
	node, err := clq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CreditLater ID in the query.
// Returns a *NotSingularError when more than one CreditLater ID is found.
// Returns a *NotFoundError when no entities are found.
func (clq *CreditLaterQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = clq.Limit(2).IDs(setContextOp(ctx, clq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{creditlater.Label}
	default:
		err = &NotSingularError{creditlater.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (clq *CreditLaterQuery) OnlyIDX(ctx context.Context) int {
	id, err := clq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CreditLaters.
func (clq *CreditLaterQuery) All(ctx context.Context) ([]*CreditLater, error) {
	ctx = setContextOp(ctx, clq.ctx, "All")
	if err := clq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*CreditLater, *CreditLaterQuery]()
	return withInterceptors[[]*CreditLater](ctx, clq, qr, clq.inters)
}

// AllX is like All, but panics if an error occurs.
func (clq *CreditLaterQuery) AllX(ctx context.Context) []*CreditLater {
	nodes, err := clq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CreditLater IDs.
func (clq *CreditLaterQuery) IDs(ctx context.Context) (ids []int, err error) {
	if clq.ctx.Unique == nil && clq.path != nil {
		clq.Unique(true)
	}
	ctx = setContextOp(ctx, clq.ctx, "IDs")
	if err = clq.Select(creditlater.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (clq *CreditLaterQuery) IDsX(ctx context.Context) []int {
	ids, err := clq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (clq *CreditLaterQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, clq.ctx, "Count")
	if err := clq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, clq, querierCount[*CreditLaterQuery](), clq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (clq *CreditLaterQuery) CountX(ctx context.Context) int {
	count, err := clq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (clq *CreditLaterQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, clq.ctx, "Exist")
	switch _, err := clq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (clq *CreditLaterQuery) ExistX(ctx context.Context) bool {
	exist, err := clq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CreditLaterQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (clq *CreditLaterQuery) Clone() *CreditLaterQuery {
	if clq == nil {
		return nil
	}
	return &CreditLaterQuery{
		config:     clq.config,
		ctx:        clq.ctx.Clone(),
		order:      append([]creditlater.OrderOption{}, clq.order...),
		inters:     append([]Interceptor{}, clq.inters...),
		predicates: append([]predicate.CreditLater{}, clq.predicates...),
		withOwner:  clq.withOwner.Clone(),
		// clone intermediate query.
		sql:  clq.sql.Clone(),
		path: clq.path,
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (clq *CreditLaterQuery) WithOwner(opts ...func(*LineUserQuery)) *CreditLaterQuery {
	query := (&LineUserClient{config: clq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	clq.withOwner = query
	return clq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		TransactionRef string `json:"transaction_ref,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CreditLater.Query().
//		GroupBy(creditlater.FieldTransactionRef).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (clq *CreditLaterQuery) GroupBy(field string, fields ...string) *CreditLaterGroupBy {
	clq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CreditLaterGroupBy{build: clq}
	grbuild.flds = &clq.ctx.Fields
	grbuild.label = creditlater.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		TransactionRef string `json:"transaction_ref,omitempty"`
//	}
//
//	client.CreditLater.Query().
//		Select(creditlater.FieldTransactionRef).
//		Scan(ctx, &v)
func (clq *CreditLaterQuery) Select(fields ...string) *CreditLaterSelect {
	clq.ctx.Fields = append(clq.ctx.Fields, fields...)
	sbuild := &CreditLaterSelect{CreditLaterQuery: clq}
	sbuild.label = creditlater.Label
	sbuild.flds, sbuild.scan = &clq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CreditLaterSelect configured with the given aggregations.
func (clq *CreditLaterQuery) Aggregate(fns ...AggregateFunc) *CreditLaterSelect {
	return clq.Select().Aggregate(fns...)
}

func (clq *CreditLaterQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range clq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, clq); err != nil {
				return err
			}
		}
	}
	for _, f := range clq.ctx.Fields {
		if !creditlater.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if clq.path != nil {
		prev, err := clq.path(ctx)
		if err != nil {
			return err
		}
		clq.sql = prev
	}
	return nil
}

func (clq *CreditLaterQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CreditLater, error) {
	var (
		nodes       = []*CreditLater{}
		withFKs     = clq.withFKs
		_spec       = clq.querySpec()
		loadedTypes = [1]bool{
			clq.withOwner != nil,
		}
	)
	if clq.withOwner != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, creditlater.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*CreditLater).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &CreditLater{config: clq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, clq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := clq.withOwner; query != nil {
		if err := clq.loadOwner(ctx, query, nodes, nil,
			func(n *CreditLater, e *LineUser) { n.Edges.Owner = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (clq *CreditLaterQuery) loadOwner(ctx context.Context, query *LineUserQuery, nodes []*CreditLater, init func(*CreditLater), assign func(*CreditLater, *LineUser)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*CreditLater)
	for i := range nodes {
		if nodes[i].line_user_creditlaters == nil {
			continue
		}
		fk := *nodes[i].line_user_creditlaters
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(lineuser.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "line_user_creditlaters" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (clq *CreditLaterQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := clq.querySpec()
	_spec.Node.Columns = clq.ctx.Fields
	if len(clq.ctx.Fields) > 0 {
		_spec.Unique = clq.ctx.Unique != nil && *clq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, clq.driver, _spec)
}

func (clq *CreditLaterQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(creditlater.Table, creditlater.Columns, sqlgraph.NewFieldSpec(creditlater.FieldID, field.TypeInt))
	_spec.From = clq.sql
	if unique := clq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if clq.path != nil {
		_spec.Unique = true
	}
	if fields := clq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, creditlater.FieldID)
		for i := range fields {
			if fields[i] != creditlater.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := clq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := clq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := clq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := clq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (clq *CreditLaterQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(clq.driver.Dialect())
	t1 := builder.Table(creditlater.Table)
	columns := clq.ctx.Fields
	if len(columns) == 0 {
		columns = creditlater.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if clq.sql != nil {
		selector = clq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if clq.ctx.Unique != nil && *clq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range clq.predicates {
		p(selector)
	}
	for _, p := range clq.order {
		p(selector)
	}
	if offset := clq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := clq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CreditLaterGroupBy is the group-by builder for CreditLater entities.
type CreditLaterGroupBy struct {
	selector
	build *CreditLaterQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (clgb *CreditLaterGroupBy) Aggregate(fns ...AggregateFunc) *CreditLaterGroupBy {
	clgb.fns = append(clgb.fns, fns...)
	return clgb
}

// Scan applies the selector query and scans the result into the given value.
func (clgb *CreditLaterGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, clgb.build.ctx, "GroupBy")
	if err := clgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CreditLaterQuery, *CreditLaterGroupBy](ctx, clgb.build, clgb, clgb.build.inters, v)
}

func (clgb *CreditLaterGroupBy) sqlScan(ctx context.Context, root *CreditLaterQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(clgb.fns))
	for _, fn := range clgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*clgb.flds)+len(clgb.fns))
		for _, f := range *clgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*clgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := clgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CreditLaterSelect is the builder for selecting fields of CreditLater entities.
type CreditLaterSelect struct {
	*CreditLaterQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cls *CreditLaterSelect) Aggregate(fns ...AggregateFunc) *CreditLaterSelect {
	cls.fns = append(cls.fns, fns...)
	return cls
}

// Scan applies the selector query and scans the result into the given value.
func (cls *CreditLaterSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cls.ctx, "Select")
	if err := cls.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CreditLaterQuery, *CreditLaterSelect](ctx, cls.CreditLaterQuery, cls, cls.inters, v)
}

func (cls *CreditLaterSelect) sqlScan(ctx context.Context, root *CreditLaterQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cls.fns))
	for _, fn := range cls.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cls.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
