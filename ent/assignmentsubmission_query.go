// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/msal4/hassah_school_server/ent/assignment"
	"github.com/msal4/hassah_school_server/ent/assignmentsubmission"
	"github.com/msal4/hassah_school_server/ent/predicate"
	"github.com/msal4/hassah_school_server/ent/user"
)

// AssignmentSubmissionQuery is the builder for querying AssignmentSubmission entities.
type AssignmentSubmissionQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AssignmentSubmission
	// eager-loading edges.
	withStudent    *UserQuery
	withAssignment *AssignmentQuery
	withFKs        bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AssignmentSubmissionQuery builder.
func (asq *AssignmentSubmissionQuery) Where(ps ...predicate.AssignmentSubmission) *AssignmentSubmissionQuery {
	asq.predicates = append(asq.predicates, ps...)
	return asq
}

// Limit adds a limit step to the query.
func (asq *AssignmentSubmissionQuery) Limit(limit int) *AssignmentSubmissionQuery {
	asq.limit = &limit
	return asq
}

// Offset adds an offset step to the query.
func (asq *AssignmentSubmissionQuery) Offset(offset int) *AssignmentSubmissionQuery {
	asq.offset = &offset
	return asq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (asq *AssignmentSubmissionQuery) Unique(unique bool) *AssignmentSubmissionQuery {
	asq.unique = &unique
	return asq
}

// Order adds an order step to the query.
func (asq *AssignmentSubmissionQuery) Order(o ...OrderFunc) *AssignmentSubmissionQuery {
	asq.order = append(asq.order, o...)
	return asq
}

// QueryStudent chains the current query on the "student" edge.
func (asq *AssignmentSubmissionQuery) QueryStudent() *UserQuery {
	query := &UserQuery{config: asq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := asq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := asq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(assignmentsubmission.Table, assignmentsubmission.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, assignmentsubmission.StudentTable, assignmentsubmission.StudentColumn),
		)
		fromU = sqlgraph.SetNeighbors(asq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAssignment chains the current query on the "assignment" edge.
func (asq *AssignmentSubmissionQuery) QueryAssignment() *AssignmentQuery {
	query := &AssignmentQuery{config: asq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := asq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := asq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(assignmentsubmission.Table, assignmentsubmission.FieldID, selector),
			sqlgraph.To(assignment.Table, assignment.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, assignmentsubmission.AssignmentTable, assignmentsubmission.AssignmentColumn),
		)
		fromU = sqlgraph.SetNeighbors(asq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AssignmentSubmission entity from the query.
// Returns a *NotFoundError when no AssignmentSubmission was found.
func (asq *AssignmentSubmissionQuery) First(ctx context.Context) (*AssignmentSubmission, error) {
	nodes, err := asq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{assignmentsubmission.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (asq *AssignmentSubmissionQuery) FirstX(ctx context.Context) *AssignmentSubmission {
	node, err := asq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AssignmentSubmission ID from the query.
// Returns a *NotFoundError when no AssignmentSubmission ID was found.
func (asq *AssignmentSubmissionQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = asq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{assignmentsubmission.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (asq *AssignmentSubmissionQuery) FirstIDX(ctx context.Context) int {
	id, err := asq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AssignmentSubmission entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one AssignmentSubmission entity is not found.
// Returns a *NotFoundError when no AssignmentSubmission entities are found.
func (asq *AssignmentSubmissionQuery) Only(ctx context.Context) (*AssignmentSubmission, error) {
	nodes, err := asq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{assignmentsubmission.Label}
	default:
		return nil, &NotSingularError{assignmentsubmission.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (asq *AssignmentSubmissionQuery) OnlyX(ctx context.Context) *AssignmentSubmission {
	node, err := asq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AssignmentSubmission ID in the query.
// Returns a *NotSingularError when exactly one AssignmentSubmission ID is not found.
// Returns a *NotFoundError when no entities are found.
func (asq *AssignmentSubmissionQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = asq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{assignmentsubmission.Label}
	default:
		err = &NotSingularError{assignmentsubmission.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (asq *AssignmentSubmissionQuery) OnlyIDX(ctx context.Context) int {
	id, err := asq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AssignmentSubmissions.
func (asq *AssignmentSubmissionQuery) All(ctx context.Context) ([]*AssignmentSubmission, error) {
	if err := asq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return asq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (asq *AssignmentSubmissionQuery) AllX(ctx context.Context) []*AssignmentSubmission {
	nodes, err := asq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AssignmentSubmission IDs.
func (asq *AssignmentSubmissionQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := asq.Select(assignmentsubmission.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (asq *AssignmentSubmissionQuery) IDsX(ctx context.Context) []int {
	ids, err := asq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (asq *AssignmentSubmissionQuery) Count(ctx context.Context) (int, error) {
	if err := asq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return asq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (asq *AssignmentSubmissionQuery) CountX(ctx context.Context) int {
	count, err := asq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (asq *AssignmentSubmissionQuery) Exist(ctx context.Context) (bool, error) {
	if err := asq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return asq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (asq *AssignmentSubmissionQuery) ExistX(ctx context.Context) bool {
	exist, err := asq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AssignmentSubmissionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (asq *AssignmentSubmissionQuery) Clone() *AssignmentSubmissionQuery {
	if asq == nil {
		return nil
	}
	return &AssignmentSubmissionQuery{
		config:         asq.config,
		limit:          asq.limit,
		offset:         asq.offset,
		order:          append([]OrderFunc{}, asq.order...),
		predicates:     append([]predicate.AssignmentSubmission{}, asq.predicates...),
		withStudent:    asq.withStudent.Clone(),
		withAssignment: asq.withAssignment.Clone(),
		// clone intermediate query.
		sql:  asq.sql.Clone(),
		path: asq.path,
	}
}

// WithStudent tells the query-builder to eager-load the nodes that are connected to
// the "student" edge. The optional arguments are used to configure the query builder of the edge.
func (asq *AssignmentSubmissionQuery) WithStudent(opts ...func(*UserQuery)) *AssignmentSubmissionQuery {
	query := &UserQuery{config: asq.config}
	for _, opt := range opts {
		opt(query)
	}
	asq.withStudent = query
	return asq
}

// WithAssignment tells the query-builder to eager-load the nodes that are connected to
// the "assignment" edge. The optional arguments are used to configure the query builder of the edge.
func (asq *AssignmentSubmissionQuery) WithAssignment(opts ...func(*AssignmentQuery)) *AssignmentSubmissionQuery {
	query := &AssignmentQuery{config: asq.config}
	for _, opt := range opts {
		opt(query)
	}
	asq.withAssignment = query
	return asq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AssignmentSubmission.Query().
//		GroupBy(assignmentsubmission.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (asq *AssignmentSubmissionQuery) GroupBy(field string, fields ...string) *AssignmentSubmissionGroupBy {
	group := &AssignmentSubmissionGroupBy{config: asq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := asq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return asq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.AssignmentSubmission.Query().
//		Select(assignmentsubmission.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (asq *AssignmentSubmissionQuery) Select(fields ...string) *AssignmentSubmissionSelect {
	asq.fields = append(asq.fields, fields...)
	return &AssignmentSubmissionSelect{AssignmentSubmissionQuery: asq}
}

func (asq *AssignmentSubmissionQuery) prepareQuery(ctx context.Context) error {
	for _, f := range asq.fields {
		if !assignmentsubmission.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if asq.path != nil {
		prev, err := asq.path(ctx)
		if err != nil {
			return err
		}
		asq.sql = prev
	}
	return nil
}

func (asq *AssignmentSubmissionQuery) sqlAll(ctx context.Context) ([]*AssignmentSubmission, error) {
	var (
		nodes       = []*AssignmentSubmission{}
		withFKs     = asq.withFKs
		_spec       = asq.querySpec()
		loadedTypes = [2]bool{
			asq.withStudent != nil,
			asq.withAssignment != nil,
		}
	)
	if asq.withStudent != nil || asq.withAssignment != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, assignmentsubmission.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &AssignmentSubmission{config: asq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, asq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := asq.withStudent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*AssignmentSubmission)
		for i := range nodes {
			if nodes[i].user_submissions == nil {
				continue
			}
			fk := *nodes[i].user_submissions
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_submissions" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Student = n
			}
		}
	}

	if query := asq.withAssignment; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*AssignmentSubmission)
		for i := range nodes {
			if nodes[i].assignment_submissions == nil {
				continue
			}
			fk := *nodes[i].assignment_submissions
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(assignment.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "assignment_submissions" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Assignment = n
			}
		}
	}

	return nodes, nil
}

func (asq *AssignmentSubmissionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := asq.querySpec()
	return sqlgraph.CountNodes(ctx, asq.driver, _spec)
}

func (asq *AssignmentSubmissionQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := asq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (asq *AssignmentSubmissionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   assignmentsubmission.Table,
			Columns: assignmentsubmission.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: assignmentsubmission.FieldID,
			},
		},
		From:   asq.sql,
		Unique: true,
	}
	if unique := asq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := asq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, assignmentsubmission.FieldID)
		for i := range fields {
			if fields[i] != assignmentsubmission.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := asq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := asq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := asq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := asq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (asq *AssignmentSubmissionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(asq.driver.Dialect())
	t1 := builder.Table(assignmentsubmission.Table)
	columns := asq.fields
	if len(columns) == 0 {
		columns = assignmentsubmission.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if asq.sql != nil {
		selector = asq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range asq.predicates {
		p(selector)
	}
	for _, p := range asq.order {
		p(selector)
	}
	if offset := asq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := asq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AssignmentSubmissionGroupBy is the group-by builder for AssignmentSubmission entities.
type AssignmentSubmissionGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (asgb *AssignmentSubmissionGroupBy) Aggregate(fns ...AggregateFunc) *AssignmentSubmissionGroupBy {
	asgb.fns = append(asgb.fns, fns...)
	return asgb
}

// Scan applies the group-by query and scans the result into the given value.
func (asgb *AssignmentSubmissionGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := asgb.path(ctx)
	if err != nil {
		return err
	}
	asgb.sql = query
	return asgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (asgb *AssignmentSubmissionGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := asgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (asgb *AssignmentSubmissionGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(asgb.fields) > 1 {
		return nil, errors.New("ent: AssignmentSubmissionGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := asgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (asgb *AssignmentSubmissionGroupBy) StringsX(ctx context.Context) []string {
	v, err := asgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (asgb *AssignmentSubmissionGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = asgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{assignmentsubmission.Label}
	default:
		err = fmt.Errorf("ent: AssignmentSubmissionGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (asgb *AssignmentSubmissionGroupBy) StringX(ctx context.Context) string {
	v, err := asgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (asgb *AssignmentSubmissionGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(asgb.fields) > 1 {
		return nil, errors.New("ent: AssignmentSubmissionGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := asgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (asgb *AssignmentSubmissionGroupBy) IntsX(ctx context.Context) []int {
	v, err := asgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (asgb *AssignmentSubmissionGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = asgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{assignmentsubmission.Label}
	default:
		err = fmt.Errorf("ent: AssignmentSubmissionGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (asgb *AssignmentSubmissionGroupBy) IntX(ctx context.Context) int {
	v, err := asgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (asgb *AssignmentSubmissionGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(asgb.fields) > 1 {
		return nil, errors.New("ent: AssignmentSubmissionGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := asgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (asgb *AssignmentSubmissionGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := asgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (asgb *AssignmentSubmissionGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = asgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{assignmentsubmission.Label}
	default:
		err = fmt.Errorf("ent: AssignmentSubmissionGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (asgb *AssignmentSubmissionGroupBy) Float64X(ctx context.Context) float64 {
	v, err := asgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (asgb *AssignmentSubmissionGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(asgb.fields) > 1 {
		return nil, errors.New("ent: AssignmentSubmissionGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := asgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (asgb *AssignmentSubmissionGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := asgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (asgb *AssignmentSubmissionGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = asgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{assignmentsubmission.Label}
	default:
		err = fmt.Errorf("ent: AssignmentSubmissionGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (asgb *AssignmentSubmissionGroupBy) BoolX(ctx context.Context) bool {
	v, err := asgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (asgb *AssignmentSubmissionGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range asgb.fields {
		if !assignmentsubmission.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := asgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := asgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (asgb *AssignmentSubmissionGroupBy) sqlQuery() *sql.Selector {
	selector := asgb.sql.Select()
	aggregation := make([]string, 0, len(asgb.fns))
	for _, fn := range asgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(asgb.fields)+len(asgb.fns))
		for _, f := range asgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(asgb.fields...)...)
}

// AssignmentSubmissionSelect is the builder for selecting fields of AssignmentSubmission entities.
type AssignmentSubmissionSelect struct {
	*AssignmentSubmissionQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ass *AssignmentSubmissionSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ass.prepareQuery(ctx); err != nil {
		return err
	}
	ass.sql = ass.AssignmentSubmissionQuery.sqlQuery(ctx)
	return ass.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ass *AssignmentSubmissionSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ass.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ass *AssignmentSubmissionSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ass.fields) > 1 {
		return nil, errors.New("ent: AssignmentSubmissionSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ass.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ass *AssignmentSubmissionSelect) StringsX(ctx context.Context) []string {
	v, err := ass.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ass *AssignmentSubmissionSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ass.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{assignmentsubmission.Label}
	default:
		err = fmt.Errorf("ent: AssignmentSubmissionSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ass *AssignmentSubmissionSelect) StringX(ctx context.Context) string {
	v, err := ass.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ass *AssignmentSubmissionSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ass.fields) > 1 {
		return nil, errors.New("ent: AssignmentSubmissionSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ass.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ass *AssignmentSubmissionSelect) IntsX(ctx context.Context) []int {
	v, err := ass.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ass *AssignmentSubmissionSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ass.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{assignmentsubmission.Label}
	default:
		err = fmt.Errorf("ent: AssignmentSubmissionSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ass *AssignmentSubmissionSelect) IntX(ctx context.Context) int {
	v, err := ass.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ass *AssignmentSubmissionSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ass.fields) > 1 {
		return nil, errors.New("ent: AssignmentSubmissionSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ass.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ass *AssignmentSubmissionSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ass.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ass *AssignmentSubmissionSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ass.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{assignmentsubmission.Label}
	default:
		err = fmt.Errorf("ent: AssignmentSubmissionSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ass *AssignmentSubmissionSelect) Float64X(ctx context.Context) float64 {
	v, err := ass.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ass *AssignmentSubmissionSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ass.fields) > 1 {
		return nil, errors.New("ent: AssignmentSubmissionSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ass.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ass *AssignmentSubmissionSelect) BoolsX(ctx context.Context) []bool {
	v, err := ass.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ass *AssignmentSubmissionSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ass.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{assignmentsubmission.Label}
	default:
		err = fmt.Errorf("ent: AssignmentSubmissionSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ass *AssignmentSubmissionSelect) BoolX(ctx context.Context) bool {
	v, err := ass.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ass *AssignmentSubmissionSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ass.sql.Query()
	if err := ass.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
