// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/organization"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/predicate"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/workflow"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/workflowcontract"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/workflowcontractversion"
	"github.com/google/uuid"
)

// WorkflowContractQuery is the builder for querying WorkflowContract entities.
type WorkflowContractQuery struct {
	config
	ctx              *QueryContext
	order            []workflowcontract.OrderOption
	inters           []Interceptor
	predicates       []predicate.WorkflowContract
	withVersions     *WorkflowContractVersionQuery
	withOrganization *OrganizationQuery
	withWorkflows    *WorkflowQuery
	withFKs          bool
	modifiers        []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WorkflowContractQuery builder.
func (wcq *WorkflowContractQuery) Where(ps ...predicate.WorkflowContract) *WorkflowContractQuery {
	wcq.predicates = append(wcq.predicates, ps...)
	return wcq
}

// Limit the number of records to be returned by this query.
func (wcq *WorkflowContractQuery) Limit(limit int) *WorkflowContractQuery {
	wcq.ctx.Limit = &limit
	return wcq
}

// Offset to start from.
func (wcq *WorkflowContractQuery) Offset(offset int) *WorkflowContractQuery {
	wcq.ctx.Offset = &offset
	return wcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wcq *WorkflowContractQuery) Unique(unique bool) *WorkflowContractQuery {
	wcq.ctx.Unique = &unique
	return wcq
}

// Order specifies how the records should be ordered.
func (wcq *WorkflowContractQuery) Order(o ...workflowcontract.OrderOption) *WorkflowContractQuery {
	wcq.order = append(wcq.order, o...)
	return wcq
}

// QueryVersions chains the current query on the "versions" edge.
func (wcq *WorkflowContractQuery) QueryVersions() *WorkflowContractVersionQuery {
	query := (&WorkflowContractVersionClient{config: wcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workflowcontract.Table, workflowcontract.FieldID, selector),
			sqlgraph.To(workflowcontractversion.Table, workflowcontractversion.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, workflowcontract.VersionsTable, workflowcontract.VersionsColumn),
		)
		fromU = sqlgraph.SetNeighbors(wcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrganization chains the current query on the "organization" edge.
func (wcq *WorkflowContractQuery) QueryOrganization() *OrganizationQuery {
	query := (&OrganizationClient{config: wcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workflowcontract.Table, workflowcontract.FieldID, selector),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, workflowcontract.OrganizationTable, workflowcontract.OrganizationColumn),
		)
		fromU = sqlgraph.SetNeighbors(wcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryWorkflows chains the current query on the "workflows" edge.
func (wcq *WorkflowContractQuery) QueryWorkflows() *WorkflowQuery {
	query := (&WorkflowClient{config: wcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workflowcontract.Table, workflowcontract.FieldID, selector),
			sqlgraph.To(workflow.Table, workflow.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, workflowcontract.WorkflowsTable, workflowcontract.WorkflowsColumn),
		)
		fromU = sqlgraph.SetNeighbors(wcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first WorkflowContract entity from the query.
// Returns a *NotFoundError when no WorkflowContract was found.
func (wcq *WorkflowContractQuery) First(ctx context.Context) (*WorkflowContract, error) {
	nodes, err := wcq.Limit(1).All(setContextOp(ctx, wcq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{workflowcontract.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wcq *WorkflowContractQuery) FirstX(ctx context.Context) *WorkflowContract {
	node, err := wcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first WorkflowContract ID from the query.
// Returns a *NotFoundError when no WorkflowContract ID was found.
func (wcq *WorkflowContractQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = wcq.Limit(1).IDs(setContextOp(ctx, wcq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{workflowcontract.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wcq *WorkflowContractQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := wcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single WorkflowContract entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one WorkflowContract entity is found.
// Returns a *NotFoundError when no WorkflowContract entities are found.
func (wcq *WorkflowContractQuery) Only(ctx context.Context) (*WorkflowContract, error) {
	nodes, err := wcq.Limit(2).All(setContextOp(ctx, wcq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{workflowcontract.Label}
	default:
		return nil, &NotSingularError{workflowcontract.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wcq *WorkflowContractQuery) OnlyX(ctx context.Context) *WorkflowContract {
	node, err := wcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only WorkflowContract ID in the query.
// Returns a *NotSingularError when more than one WorkflowContract ID is found.
// Returns a *NotFoundError when no entities are found.
func (wcq *WorkflowContractQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = wcq.Limit(2).IDs(setContextOp(ctx, wcq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{workflowcontract.Label}
	default:
		err = &NotSingularError{workflowcontract.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wcq *WorkflowContractQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := wcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of WorkflowContracts.
func (wcq *WorkflowContractQuery) All(ctx context.Context) ([]*WorkflowContract, error) {
	ctx = setContextOp(ctx, wcq.ctx, "All")
	if err := wcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*WorkflowContract, *WorkflowContractQuery]()
	return withInterceptors[[]*WorkflowContract](ctx, wcq, qr, wcq.inters)
}

// AllX is like All, but panics if an error occurs.
func (wcq *WorkflowContractQuery) AllX(ctx context.Context) []*WorkflowContract {
	nodes, err := wcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of WorkflowContract IDs.
func (wcq *WorkflowContractQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if wcq.ctx.Unique == nil && wcq.path != nil {
		wcq.Unique(true)
	}
	ctx = setContextOp(ctx, wcq.ctx, "IDs")
	if err = wcq.Select(workflowcontract.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wcq *WorkflowContractQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := wcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wcq *WorkflowContractQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, wcq.ctx, "Count")
	if err := wcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, wcq, querierCount[*WorkflowContractQuery](), wcq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (wcq *WorkflowContractQuery) CountX(ctx context.Context) int {
	count, err := wcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wcq *WorkflowContractQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, wcq.ctx, "Exist")
	switch _, err := wcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (wcq *WorkflowContractQuery) ExistX(ctx context.Context) bool {
	exist, err := wcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WorkflowContractQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wcq *WorkflowContractQuery) Clone() *WorkflowContractQuery {
	if wcq == nil {
		return nil
	}
	return &WorkflowContractQuery{
		config:           wcq.config,
		ctx:              wcq.ctx.Clone(),
		order:            append([]workflowcontract.OrderOption{}, wcq.order...),
		inters:           append([]Interceptor{}, wcq.inters...),
		predicates:       append([]predicate.WorkflowContract{}, wcq.predicates...),
		withVersions:     wcq.withVersions.Clone(),
		withOrganization: wcq.withOrganization.Clone(),
		withWorkflows:    wcq.withWorkflows.Clone(),
		// clone intermediate query.
		sql:  wcq.sql.Clone(),
		path: wcq.path,
	}
}

// WithVersions tells the query-builder to eager-load the nodes that are connected to
// the "versions" edge. The optional arguments are used to configure the query builder of the edge.
func (wcq *WorkflowContractQuery) WithVersions(opts ...func(*WorkflowContractVersionQuery)) *WorkflowContractQuery {
	query := (&WorkflowContractVersionClient{config: wcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wcq.withVersions = query
	return wcq
}

// WithOrganization tells the query-builder to eager-load the nodes that are connected to
// the "organization" edge. The optional arguments are used to configure the query builder of the edge.
func (wcq *WorkflowContractQuery) WithOrganization(opts ...func(*OrganizationQuery)) *WorkflowContractQuery {
	query := (&OrganizationClient{config: wcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wcq.withOrganization = query
	return wcq
}

// WithWorkflows tells the query-builder to eager-load the nodes that are connected to
// the "workflows" edge. The optional arguments are used to configure the query builder of the edge.
func (wcq *WorkflowContractQuery) WithWorkflows(opts ...func(*WorkflowQuery)) *WorkflowContractQuery {
	query := (&WorkflowClient{config: wcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wcq.withWorkflows = query
	return wcq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.WorkflowContract.Query().
//		GroupBy(workflowcontract.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (wcq *WorkflowContractQuery) GroupBy(field string, fields ...string) *WorkflowContractGroupBy {
	wcq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &WorkflowContractGroupBy{build: wcq}
	grbuild.flds = &wcq.ctx.Fields
	grbuild.label = workflowcontract.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.WorkflowContract.Query().
//		Select(workflowcontract.FieldName).
//		Scan(ctx, &v)
func (wcq *WorkflowContractQuery) Select(fields ...string) *WorkflowContractSelect {
	wcq.ctx.Fields = append(wcq.ctx.Fields, fields...)
	sbuild := &WorkflowContractSelect{WorkflowContractQuery: wcq}
	sbuild.label = workflowcontract.Label
	sbuild.flds, sbuild.scan = &wcq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a WorkflowContractSelect configured with the given aggregations.
func (wcq *WorkflowContractQuery) Aggregate(fns ...AggregateFunc) *WorkflowContractSelect {
	return wcq.Select().Aggregate(fns...)
}

func (wcq *WorkflowContractQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range wcq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, wcq); err != nil {
				return err
			}
		}
	}
	for _, f := range wcq.ctx.Fields {
		if !workflowcontract.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if wcq.path != nil {
		prev, err := wcq.path(ctx)
		if err != nil {
			return err
		}
		wcq.sql = prev
	}
	return nil
}

func (wcq *WorkflowContractQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*WorkflowContract, error) {
	var (
		nodes       = []*WorkflowContract{}
		withFKs     = wcq.withFKs
		_spec       = wcq.querySpec()
		loadedTypes = [3]bool{
			wcq.withVersions != nil,
			wcq.withOrganization != nil,
			wcq.withWorkflows != nil,
		}
	)
	if wcq.withOrganization != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, workflowcontract.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*WorkflowContract).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &WorkflowContract{config: wcq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(wcq.modifiers) > 0 {
		_spec.Modifiers = wcq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, wcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := wcq.withVersions; query != nil {
		if err := wcq.loadVersions(ctx, query, nodes,
			func(n *WorkflowContract) { n.Edges.Versions = []*WorkflowContractVersion{} },
			func(n *WorkflowContract, e *WorkflowContractVersion) { n.Edges.Versions = append(n.Edges.Versions, e) }); err != nil {
			return nil, err
		}
	}
	if query := wcq.withOrganization; query != nil {
		if err := wcq.loadOrganization(ctx, query, nodes, nil,
			func(n *WorkflowContract, e *Organization) { n.Edges.Organization = e }); err != nil {
			return nil, err
		}
	}
	if query := wcq.withWorkflows; query != nil {
		if err := wcq.loadWorkflows(ctx, query, nodes,
			func(n *WorkflowContract) { n.Edges.Workflows = []*Workflow{} },
			func(n *WorkflowContract, e *Workflow) { n.Edges.Workflows = append(n.Edges.Workflows, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (wcq *WorkflowContractQuery) loadVersions(ctx context.Context, query *WorkflowContractVersionQuery, nodes []*WorkflowContract, init func(*WorkflowContract), assign func(*WorkflowContract, *WorkflowContractVersion)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*WorkflowContract)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.WorkflowContractVersion(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(workflowcontract.VersionsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.workflow_contract_versions
		if fk == nil {
			return fmt.Errorf(`foreign-key "workflow_contract_versions" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "workflow_contract_versions" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (wcq *WorkflowContractQuery) loadOrganization(ctx context.Context, query *OrganizationQuery, nodes []*WorkflowContract, init func(*WorkflowContract), assign func(*WorkflowContract, *Organization)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*WorkflowContract)
	for i := range nodes {
		if nodes[i].organization_workflow_contracts == nil {
			continue
		}
		fk := *nodes[i].organization_workflow_contracts
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(organization.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "organization_workflow_contracts" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (wcq *WorkflowContractQuery) loadWorkflows(ctx context.Context, query *WorkflowQuery, nodes []*WorkflowContract, init func(*WorkflowContract), assign func(*WorkflowContract, *Workflow)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*WorkflowContract)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(workflowcontract.WorkflowsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.workflow_contract
		if fk == nil {
			return fmt.Errorf(`foreign-key "workflow_contract" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "workflow_contract" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (wcq *WorkflowContractQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wcq.querySpec()
	if len(wcq.modifiers) > 0 {
		_spec.Modifiers = wcq.modifiers
	}
	_spec.Node.Columns = wcq.ctx.Fields
	if len(wcq.ctx.Fields) > 0 {
		_spec.Unique = wcq.ctx.Unique != nil && *wcq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, wcq.driver, _spec)
}

func (wcq *WorkflowContractQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(workflowcontract.Table, workflowcontract.Columns, sqlgraph.NewFieldSpec(workflowcontract.FieldID, field.TypeUUID))
	_spec.From = wcq.sql
	if unique := wcq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if wcq.path != nil {
		_spec.Unique = true
	}
	if fields := wcq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, workflowcontract.FieldID)
		for i := range fields {
			if fields[i] != workflowcontract.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := wcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wcq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wcq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wcq *WorkflowContractQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wcq.driver.Dialect())
	t1 := builder.Table(workflowcontract.Table)
	columns := wcq.ctx.Fields
	if len(columns) == 0 {
		columns = workflowcontract.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wcq.sql != nil {
		selector = wcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if wcq.ctx.Unique != nil && *wcq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range wcq.modifiers {
		m(selector)
	}
	for _, p := range wcq.predicates {
		p(selector)
	}
	for _, p := range wcq.order {
		p(selector)
	}
	if offset := wcq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wcq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (wcq *WorkflowContractQuery) Modify(modifiers ...func(s *sql.Selector)) *WorkflowContractSelect {
	wcq.modifiers = append(wcq.modifiers, modifiers...)
	return wcq.Select()
}

// WorkflowContractGroupBy is the group-by builder for WorkflowContract entities.
type WorkflowContractGroupBy struct {
	selector
	build *WorkflowContractQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wcgb *WorkflowContractGroupBy) Aggregate(fns ...AggregateFunc) *WorkflowContractGroupBy {
	wcgb.fns = append(wcgb.fns, fns...)
	return wcgb
}

// Scan applies the selector query and scans the result into the given value.
func (wcgb *WorkflowContractGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wcgb.build.ctx, "GroupBy")
	if err := wcgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WorkflowContractQuery, *WorkflowContractGroupBy](ctx, wcgb.build, wcgb, wcgb.build.inters, v)
}

func (wcgb *WorkflowContractGroupBy) sqlScan(ctx context.Context, root *WorkflowContractQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(wcgb.fns))
	for _, fn := range wcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*wcgb.flds)+len(wcgb.fns))
		for _, f := range *wcgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*wcgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wcgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// WorkflowContractSelect is the builder for selecting fields of WorkflowContract entities.
type WorkflowContractSelect struct {
	*WorkflowContractQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (wcs *WorkflowContractSelect) Aggregate(fns ...AggregateFunc) *WorkflowContractSelect {
	wcs.fns = append(wcs.fns, fns...)
	return wcs
}

// Scan applies the selector query and scans the result into the given value.
func (wcs *WorkflowContractSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wcs.ctx, "Select")
	if err := wcs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WorkflowContractQuery, *WorkflowContractSelect](ctx, wcs.WorkflowContractQuery, wcs, wcs.inters, v)
}

func (wcs *WorkflowContractSelect) sqlScan(ctx context.Context, root *WorkflowContractQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(wcs.fns))
	for _, fn := range wcs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*wcs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (wcs *WorkflowContractSelect) Modify(modifiers ...func(s *sql.Selector)) *WorkflowContractSelect {
	wcs.modifiers = append(wcs.modifiers, modifiers...)
	return wcs
}