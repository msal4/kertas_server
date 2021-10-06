// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/msal4/hassah_school_server/ent/class"
	"github.com/msal4/hassah_school_server/ent/notification"
	"github.com/msal4/hassah_school_server/ent/predicate"
	"github.com/msal4/hassah_school_server/ent/school"
	"github.com/msal4/hassah_school_server/ent/stage"
	"github.com/msal4/hassah_school_server/ent/tuitionpayment"
	"github.com/msal4/hassah_school_server/ent/user"
)

// StageQuery is the builder for querying Stage entities.
type StageQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Stage
	// eager-loading edges.
	withSchool        *SchoolQuery
	withClasses       *ClassQuery
	withPayments      *TuitionPaymentQuery
	withStudents      *UserQuery
	withNotifications *NotificationQuery
	withFKs           bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the StageQuery builder.
func (sq *StageQuery) Where(ps ...predicate.Stage) *StageQuery {
	sq.predicates = append(sq.predicates, ps...)
	return sq
}

// Limit adds a limit step to the query.
func (sq *StageQuery) Limit(limit int) *StageQuery {
	sq.limit = &limit
	return sq
}

// Offset adds an offset step to the query.
func (sq *StageQuery) Offset(offset int) *StageQuery {
	sq.offset = &offset
	return sq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sq *StageQuery) Unique(unique bool) *StageQuery {
	sq.unique = &unique
	return sq
}

// Order adds an order step to the query.
func (sq *StageQuery) Order(o ...OrderFunc) *StageQuery {
	sq.order = append(sq.order, o...)
	return sq
}

// QuerySchool chains the current query on the "school" edge.
func (sq *StageQuery) QuerySchool() *SchoolQuery {
	query := &SchoolQuery{config: sq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(stage.Table, stage.FieldID, selector),
			sqlgraph.To(school.Table, school.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, stage.SchoolTable, stage.SchoolColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryClasses chains the current query on the "classes" edge.
func (sq *StageQuery) QueryClasses() *ClassQuery {
	query := &ClassQuery{config: sq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(stage.Table, stage.FieldID, selector),
			sqlgraph.To(class.Table, class.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, stage.ClassesTable, stage.ClassesColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPayments chains the current query on the "payments" edge.
func (sq *StageQuery) QueryPayments() *TuitionPaymentQuery {
	query := &TuitionPaymentQuery{config: sq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(stage.Table, stage.FieldID, selector),
			sqlgraph.To(tuitionpayment.Table, tuitionpayment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, stage.PaymentsTable, stage.PaymentsColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryStudents chains the current query on the "students" edge.
func (sq *StageQuery) QueryStudents() *UserQuery {
	query := &UserQuery{config: sq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(stage.Table, stage.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, stage.StudentsTable, stage.StudentsColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNotifications chains the current query on the "notifications" edge.
func (sq *StageQuery) QueryNotifications() *NotificationQuery {
	query := &NotificationQuery{config: sq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(stage.Table, stage.FieldID, selector),
			sqlgraph.To(notification.Table, notification.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, stage.NotificationsTable, stage.NotificationsColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Stage entity from the query.
// Returns a *NotFoundError when no Stage was found.
func (sq *StageQuery) First(ctx context.Context) (*Stage, error) {
	nodes, err := sq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{stage.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sq *StageQuery) FirstX(ctx context.Context) *Stage {
	node, err := sq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Stage ID from the query.
// Returns a *NotFoundError when no Stage ID was found.
func (sq *StageQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = sq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{stage.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sq *StageQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := sq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Stage entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one Stage entity is not found.
// Returns a *NotFoundError when no Stage entities are found.
func (sq *StageQuery) Only(ctx context.Context) (*Stage, error) {
	nodes, err := sq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{stage.Label}
	default:
		return nil, &NotSingularError{stage.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sq *StageQuery) OnlyX(ctx context.Context) *Stage {
	node, err := sq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Stage ID in the query.
// Returns a *NotSingularError when exactly one Stage ID is not found.
// Returns a *NotFoundError when no entities are found.
func (sq *StageQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = sq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{stage.Label}
	default:
		err = &NotSingularError{stage.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sq *StageQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := sq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Stages.
func (sq *StageQuery) All(ctx context.Context) ([]*Stage, error) {
	if err := sq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return sq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (sq *StageQuery) AllX(ctx context.Context) []*Stage {
	nodes, err := sq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Stage IDs.
func (sq *StageQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := sq.Select(stage.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sq *StageQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := sq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sq *StageQuery) Count(ctx context.Context) (int, error) {
	if err := sq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return sq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (sq *StageQuery) CountX(ctx context.Context) int {
	count, err := sq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sq *StageQuery) Exist(ctx context.Context) (bool, error) {
	if err := sq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return sq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (sq *StageQuery) ExistX(ctx context.Context) bool {
	exist, err := sq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the StageQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sq *StageQuery) Clone() *StageQuery {
	if sq == nil {
		return nil
	}
	return &StageQuery{
		config:            sq.config,
		limit:             sq.limit,
		offset:            sq.offset,
		order:             append([]OrderFunc{}, sq.order...),
		predicates:        append([]predicate.Stage{}, sq.predicates...),
		withSchool:        sq.withSchool.Clone(),
		withClasses:       sq.withClasses.Clone(),
		withPayments:      sq.withPayments.Clone(),
		withStudents:      sq.withStudents.Clone(),
		withNotifications: sq.withNotifications.Clone(),
		// clone intermediate query.
		sql:  sq.sql.Clone(),
		path: sq.path,
	}
}

// WithSchool tells the query-builder to eager-load the nodes that are connected to
// the "school" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *StageQuery) WithSchool(opts ...func(*SchoolQuery)) *StageQuery {
	query := &SchoolQuery{config: sq.config}
	for _, opt := range opts {
		opt(query)
	}
	sq.withSchool = query
	return sq
}

// WithClasses tells the query-builder to eager-load the nodes that are connected to
// the "classes" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *StageQuery) WithClasses(opts ...func(*ClassQuery)) *StageQuery {
	query := &ClassQuery{config: sq.config}
	for _, opt := range opts {
		opt(query)
	}
	sq.withClasses = query
	return sq
}

// WithPayments tells the query-builder to eager-load the nodes that are connected to
// the "payments" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *StageQuery) WithPayments(opts ...func(*TuitionPaymentQuery)) *StageQuery {
	query := &TuitionPaymentQuery{config: sq.config}
	for _, opt := range opts {
		opt(query)
	}
	sq.withPayments = query
	return sq
}

// WithStudents tells the query-builder to eager-load the nodes that are connected to
// the "students" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *StageQuery) WithStudents(opts ...func(*UserQuery)) *StageQuery {
	query := &UserQuery{config: sq.config}
	for _, opt := range opts {
		opt(query)
	}
	sq.withStudents = query
	return sq
}

// WithNotifications tells the query-builder to eager-load the nodes that are connected to
// the "notifications" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *StageQuery) WithNotifications(opts ...func(*NotificationQuery)) *StageQuery {
	query := &NotificationQuery{config: sq.config}
	for _, opt := range opts {
		opt(query)
	}
	sq.withNotifications = query
	return sq
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
//	client.Stage.Query().
//		GroupBy(stage.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (sq *StageQuery) GroupBy(field string, fields ...string) *StageGroupBy {
	group := &StageGroupBy{config: sq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return sq.sqlQuery(ctx), nil
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
//	client.Stage.Query().
//		Select(stage.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (sq *StageQuery) Select(fields ...string) *StageSelect {
	sq.fields = append(sq.fields, fields...)
	return &StageSelect{StageQuery: sq}
}

func (sq *StageQuery) prepareQuery(ctx context.Context) error {
	for _, f := range sq.fields {
		if !stage.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sq.path != nil {
		prev, err := sq.path(ctx)
		if err != nil {
			return err
		}
		sq.sql = prev
	}
	return nil
}

func (sq *StageQuery) sqlAll(ctx context.Context) ([]*Stage, error) {
	var (
		nodes       = []*Stage{}
		withFKs     = sq.withFKs
		_spec       = sq.querySpec()
		loadedTypes = [5]bool{
			sq.withSchool != nil,
			sq.withClasses != nil,
			sq.withPayments != nil,
			sq.withStudents != nil,
			sq.withNotifications != nil,
		}
	)
	if sq.withSchool != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, stage.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Stage{config: sq.config}
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
	if err := sqlgraph.QueryNodes(ctx, sq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := sq.withSchool; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*Stage)
		for i := range nodes {
			if nodes[i].school_stages == nil {
				continue
			}
			fk := *nodes[i].school_stages
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(school.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "school_stages" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.School = n
			}
		}
	}

	if query := sq.withClasses; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[uuid.UUID]*Stage)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Classes = []*Class{}
		}
		query.withFKs = true
		query.Where(predicate.Class(func(s *sql.Selector) {
			s.Where(sql.InValues(stage.ClassesColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.stage_classes
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "stage_classes" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "stage_classes" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Classes = append(node.Edges.Classes, n)
		}
	}

	if query := sq.withPayments; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[uuid.UUID]*Stage)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Payments = []*TuitionPayment{}
		}
		query.withFKs = true
		query.Where(predicate.TuitionPayment(func(s *sql.Selector) {
			s.Where(sql.InValues(stage.PaymentsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.stage_payments
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "stage_payments" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "stage_payments" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Payments = append(node.Edges.Payments, n)
		}
	}

	if query := sq.withStudents; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[uuid.UUID]*Stage)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Students = []*User{}
		}
		query.withFKs = true
		query.Where(predicate.User(func(s *sql.Selector) {
			s.Where(sql.InValues(stage.StudentsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.stage_students
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "stage_students" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "stage_students" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Students = append(node.Edges.Students, n)
		}
	}

	if query := sq.withNotifications; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[uuid.UUID]*Stage)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Notifications = []*Notification{}
		}
		query.withFKs = true
		query.Where(predicate.Notification(func(s *sql.Selector) {
			s.Where(sql.InValues(stage.NotificationsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.stage_notifications
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "stage_notifications" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "stage_notifications" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Notifications = append(node.Edges.Notifications, n)
		}
	}

	return nodes, nil
}

func (sq *StageQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sq.querySpec()
	return sqlgraph.CountNodes(ctx, sq.driver, _spec)
}

func (sq *StageQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := sq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (sq *StageQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   stage.Table,
			Columns: stage.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: stage.FieldID,
			},
		},
		From:   sq.sql,
		Unique: true,
	}
	if unique := sq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := sq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, stage.FieldID)
		for i := range fields {
			if fields[i] != stage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sq *StageQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sq.driver.Dialect())
	t1 := builder.Table(stage.Table)
	columns := sq.fields
	if len(columns) == 0 {
		columns = stage.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sq.sql != nil {
		selector = sq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range sq.predicates {
		p(selector)
	}
	for _, p := range sq.order {
		p(selector)
	}
	if offset := sq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// StageGroupBy is the group-by builder for Stage entities.
type StageGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *StageGroupBy) Aggregate(fns ...AggregateFunc) *StageGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the group-by query and scans the result into the given value.
func (sgb *StageGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := sgb.path(ctx)
	if err != nil {
		return err
	}
	sgb.sql = query
	return sgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sgb *StageGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := sgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (sgb *StageGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(sgb.fields) > 1 {
		return nil, errors.New("ent: StageGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := sgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sgb *StageGroupBy) StringsX(ctx context.Context) []string {
	v, err := sgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sgb *StageGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = sgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{stage.Label}
	default:
		err = fmt.Errorf("ent: StageGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (sgb *StageGroupBy) StringX(ctx context.Context) string {
	v, err := sgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (sgb *StageGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(sgb.fields) > 1 {
		return nil, errors.New("ent: StageGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := sgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sgb *StageGroupBy) IntsX(ctx context.Context) []int {
	v, err := sgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sgb *StageGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = sgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{stage.Label}
	default:
		err = fmt.Errorf("ent: StageGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (sgb *StageGroupBy) IntX(ctx context.Context) int {
	v, err := sgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (sgb *StageGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(sgb.fields) > 1 {
		return nil, errors.New("ent: StageGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := sgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sgb *StageGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := sgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sgb *StageGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = sgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{stage.Label}
	default:
		err = fmt.Errorf("ent: StageGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (sgb *StageGroupBy) Float64X(ctx context.Context) float64 {
	v, err := sgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (sgb *StageGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(sgb.fields) > 1 {
		return nil, errors.New("ent: StageGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := sgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sgb *StageGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := sgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sgb *StageGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = sgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{stage.Label}
	default:
		err = fmt.Errorf("ent: StageGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (sgb *StageGroupBy) BoolX(ctx context.Context) bool {
	v, err := sgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sgb *StageGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range sgb.fields {
		if !stage.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := sgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sgb *StageGroupBy) sqlQuery() *sql.Selector {
	selector := sgb.sql.Select()
	aggregation := make([]string, 0, len(sgb.fns))
	for _, fn := range sgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(sgb.fields)+len(sgb.fns))
		for _, f := range sgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(sgb.fields...)...)
}

// StageSelect is the builder for selecting fields of Stage entities.
type StageSelect struct {
	*StageQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ss *StageSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ss.prepareQuery(ctx); err != nil {
		return err
	}
	ss.sql = ss.StageQuery.sqlQuery(ctx)
	return ss.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ss *StageSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ss.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ss *StageSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ss.fields) > 1 {
		return nil, errors.New("ent: StageSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ss *StageSelect) StringsX(ctx context.Context) []string {
	v, err := ss.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ss *StageSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ss.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{stage.Label}
	default:
		err = fmt.Errorf("ent: StageSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ss *StageSelect) StringX(ctx context.Context) string {
	v, err := ss.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ss *StageSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ss.fields) > 1 {
		return nil, errors.New("ent: StageSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ss *StageSelect) IntsX(ctx context.Context) []int {
	v, err := ss.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ss *StageSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ss.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{stage.Label}
	default:
		err = fmt.Errorf("ent: StageSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ss *StageSelect) IntX(ctx context.Context) int {
	v, err := ss.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ss *StageSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ss.fields) > 1 {
		return nil, errors.New("ent: StageSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ss *StageSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ss.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ss *StageSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ss.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{stage.Label}
	default:
		err = fmt.Errorf("ent: StageSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ss *StageSelect) Float64X(ctx context.Context) float64 {
	v, err := ss.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ss *StageSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ss.fields) > 1 {
		return nil, errors.New("ent: StageSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ss *StageSelect) BoolsX(ctx context.Context) []bool {
	v, err := ss.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ss *StageSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ss.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{stage.Label}
	default:
		err = fmt.Errorf("ent: StageSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ss *StageSelect) BoolX(ctx context.Context) bool {
	v, err := ss.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ss *StageSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ss.sql.Query()
	if err := ss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
