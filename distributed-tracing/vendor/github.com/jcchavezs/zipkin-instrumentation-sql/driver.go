package zipkinsql

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"sync"
	"time"

	zipkin "github.com/openzipkin/zipkin-go"
	zipkinmodel "github.com/openzipkin/zipkin-go/model"
)

type conn interface {
	driver.Pinger
	driver.Execer
	driver.ExecerContext
	driver.Queryer
	driver.QueryerContext
	driver.Conn
	driver.ConnPrepareContext
	driver.ConnBeginTx
}

var (
	// Type assertions
	_ driver.Driver           = &zDriver{}
	_ conn                    = &zConn{}
	_ driver.Result           = &zResult{}
	_ driver.Stmt             = &zStmt{}
	_ driver.StmtExecContext  = &zStmt{}
	_ driver.StmtQueryContext = &zStmt{}
)

var (
	regMu sync.Mutex
)

// Register initializes and registers our zipkinsql wrapped database driver
// identified by its driverName and using provided TraceOptions. On success it
// returns the generated driverName to use when calling sql.Open.
// It is possible to register multiple wrappers for the same database driver if
// needing different TraceOptions for different connections.
func Register(driverName string, tracer *zipkin.Tracer, options ...TraceOption) (string, error) {
	// retrieve the driver implementation we need to wrap with instrumentation
	db, err := sql.Open(driverName, "")
	if err != nil {
		return "", err
	}
	dri := db.Driver()
	if err = db.Close(); err != nil {
		return "", err
	}

	regMu.Lock()
	defer regMu.Unlock()
	registerName := fmt.Sprintf("%s-zipkinsql-%d", driverName, len(sql.Drivers()))
	sql.Register(registerName, Wrap(dri, tracer, options...))

	return registerName, nil
}

// Wrap takes a SQL driver and wraps it with Zipkin instrumentation.
func Wrap(d driver.Driver, t *zipkin.Tracer, options ...TraceOption) driver.Driver {
	o := TraceOptions{}
	for _, option := range options {
		option(&o)
	}
	return wrapDriver(d, t, o)
}

func (d zDriver) Open(name string) (driver.Conn, error) {
	c, err := d.parent.Open(name)
	if err != nil {
		return nil, err
	}
	return wrapConn(c, d.tracer, d.options), nil
}

// WrapConn allows an existing driver.Conn to be wrapped by zipkinsql.
func WrapConn(c driver.Conn, t *zipkin.Tracer, options ...TraceOption) driver.Conn {
	o := TraceOptions{}
	for _, option := range options {
		option(&o)
	}
	return wrapConn(c, t, o)
}

// zConn implements driver.Conn
type zConn struct {
	parent  driver.Conn
	tracer  *zipkin.Tracer
	options TraceOptions
}

func (c zConn) Ping(ctx context.Context) (err error) {
	if pinger, ok := c.parent.(driver.Pinger); ok {
		err = pinger.Ping(ctx)
	}
	return
}

func (c zConn) Exec(query string, args []driver.Value) (res driver.Result, err error) {
	if exec, ok := c.parent.(driver.Execer); ok {
		return exec.Exec(query, args)
	}

	return nil, driver.ErrSkip
}

func (c zConn) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (res driver.Result, err error) {
	if execCtx, ok := c.parent.(driver.ExecerContext); ok {
		parentSpan := zipkin.SpanFromContext(ctx)
		if parentSpan == nil && !c.options.AllowRootSpan {
			return execCtx.ExecContext(ctx, query, args)
		}

		var (
			startTime time.Time
		)

		defer func() {
			if err == nil || err != driver.ErrSkip {
				var span zipkin.Span
				span, _ = c.tracer.StartSpanFromContext(ctx, "sql/exec", zipkin.Kind(zipkinmodel.Client), zipkin.StartTime(startTime))

				if c.options.TagQuery {
					span.Tag("sql.query", query)
				}
				setSpanDefaultTags(span, c.options.DefaultTags)

				setSpanError(span, err)
				span.Finish()
			}
		}()

		startTime = time.Now()
		if res, err = execCtx.ExecContext(ctx, query, args); err != nil {
			return nil, err
		}

		return zResult{parent: res, tracer: c.tracer, ctx: ctx, options: c.options}, nil
	}

	return nil, driver.ErrSkip
}

func (c zConn) Query(query string, args []driver.Value) (rows driver.Rows, err error) {
	if queryer, ok := c.parent.(driver.Queryer); ok {
		return queryer.Query(query, args)
	}

	return nil, driver.ErrSkip
}

func (c zConn) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (rows driver.Rows, err error) {
	if queryerCtx, ok := c.parent.(driver.QueryerContext); ok {
		parentSpan := zipkin.SpanFromContext(ctx)
		if parentSpan == nil && !c.options.AllowRootSpan {
			return queryerCtx.QueryContext(ctx, query, args)
		}

		var (
			startTime time.Time
		)

		defer func() {
			if err == nil || err != driver.ErrSkip {
				var span zipkin.Span
				span, _ = c.tracer.StartSpanFromContext(ctx, "sql/query", zipkin.Kind(zipkinmodel.Client), zipkin.StartTime(startTime))

				if c.options.TagQuery {
					span.Tag("sql.query", query)
				}
				setSpanDefaultTags(span, c.options.DefaultTags)

				setSpanError(span, err)
				span.Finish()
			}
		}()

		startTime = time.Now()
		if rows, err = queryerCtx.QueryContext(ctx, query, args); err != nil {
			return nil, err
		}

		return rows, nil
	}

	return nil, driver.ErrSkip
}

func (c zConn) Prepare(query string) (stmt driver.Stmt, err error) {
	if c.options.AllowRootSpan {
		span := c.tracer.StartSpan("sql/prepare", zipkin.Kind(zipkinmodel.Client))

		if c.options.TagQuery {
			span.Tag("sql.query", query)
		}
		setSpanDefaultTags(span, c.options.DefaultTags)
		defer func() {
			setSpanError(span, err)
			span.Finish()
		}()
	}

	stmt, err = c.parent.Prepare(query)
	if err != nil {
		return nil, err
	}

	stmt = wrapStmt(stmt, query, c.tracer, c.options)
	return
}

func (c *zConn) Close() error {
	return c.parent.Close()
}

func (c *zConn) Begin() (driver.Tx, error) {
	return c.Begin()
}

func (c *zConn) PrepareContext(ctx context.Context, query string) (stmt driver.Stmt, err error) {
	var span zipkin.Span
	setSpanDefaultTags(span, c.options.DefaultTags)
	if c.options.AllowRootSpan || zipkin.SpanFromContext(ctx) != nil {
		span, ctx = c.tracer.StartSpanFromContext(ctx, "sql/prepare", zipkin.Kind(zipkinmodel.Client))
		if c.options.TagQuery {
			span.Tag("sql.query", query)
		}

		defer func() {
			setSpanError(span, err)
			span.Finish()
		}()
	}

	if prepCtx, ok := c.parent.(driver.ConnPrepareContext); ok {
		stmt, err = prepCtx.PrepareContext(ctx, query)
	} else {
		stmt, err = c.parent.Prepare(query)
	}

	if err != nil {
		return nil, err
	}

	stmt = wrapStmt(stmt, query, c.tracer, c.options)
	return
}

func (c *zConn) BeginTx(ctx context.Context, opts driver.TxOptions) (driver.Tx, error) {
	if zipkin.SpanFromContext(ctx) == nil && !c.options.AllowRootSpan {
		if connBeginTx, ok := c.parent.(driver.ConnBeginTx); ok {
			return connBeginTx.BeginTx(ctx, opts)
		}

		return c.parent.Begin()
	}

	span, _ := c.tracer.StartSpanFromContext(ctx, "sql/begin_transaction", zipkin.Kind(zipkinmodel.Client))
	defer span.Finish()

	setSpanDefaultTags(span, c.options.DefaultTags)

	if connBeginTx, ok := c.parent.(driver.ConnBeginTx); ok {
		tx, err := connBeginTx.BeginTx(ctx, opts)
		setSpanError(span, err)
		if err != nil {
			return nil, err
		}
		return zTx{parent: tx, ctx: ctx, tracer: c.tracer, options: c.options}, nil
	}

	tx, err := c.parent.Begin()
	setSpanError(span, err)
	if err != nil {
		return nil, err
	}

	return zTx{parent: tx, ctx: ctx, tracer: c.tracer, options: c.options}, nil
}

// zResult implements driver.Result
type zResult struct {
	parent  driver.Result
	ctx     context.Context
	tracer  *zipkin.Tracer
	options TraceOptions
}

func (r zResult) LastInsertId() (int64, error) {
	if !r.options.LastInsertIDSpan {
		return r.parent.LastInsertId()
	}

	span, _ := r.tracer.StartSpanFromContext(r.ctx, "sql/last_insert_id", zipkin.Kind(zipkinmodel.Client))
	defer span.Finish()

	setSpanDefaultTags(span, r.options.DefaultTags)

	id, err := r.parent.LastInsertId()
	setSpanError(span, err)

	return id, err
}

func (r zResult) RowsAffected() (cnt int64, err error) {
	zipkin.SpanFromContext(r.ctx)
	if r.options.RowsAffectedSpan && (r.options.AllowRootSpan || zipkin.SpanFromContext(r.ctx) != nil) {
		span, _ := r.tracer.StartSpanFromContext(r.ctx, "sql/rows_affected", zipkin.Kind(zipkinmodel.Client))
		setSpanDefaultTags(span, r.options.DefaultTags)
		defer func() {
			span.Tag("sql.affected_rows", fmt.Sprintf("%d", cnt))
			setSpanError(span, err)
			span.Finish()
		}()
	}

	cnt, err = r.parent.RowsAffected()
	return
}

// zStmt implements driver.Stmt
type zStmt struct {
	parent  driver.Stmt
	query   string
	tracer  *zipkin.Tracer
	options TraceOptions
}

func (s zStmt) Exec(args []driver.Value) (res driver.Result, err error) {
	if !s.options.AllowRootSpan {
		return s.parent.Exec(args)
	}

	span, ctx := s.tracer.StartSpanFromContext(context.Background(), "sql:exec", zipkin.Kind(zipkinmodel.Client))
	setSpanDefaultTags(span, s.options.DefaultTags)

	if s.options.TagQuery {
		span.Tag("sql.query", s.query)
	}

	defer func() {
		setSpanError(span, err)
		span.Finish()
	}()

	res, err = s.parent.Exec(args)
	if err != nil {
		return nil, err
	}

	res, err = zResult{parent: res, ctx: ctx, tracer: s.tracer, options: s.options}, nil
	return
}

func (s zStmt) Close() error {
	return s.parent.Close()
}

func (s zStmt) NumInput() int {
	return s.parent.NumInput()
}

func (s zStmt) Query(args []driver.Value) (rows driver.Rows, err error) {
	if !s.options.AllowRootSpan {
		return s.parent.Query(args)
	}

	span, _ := s.tracer.StartSpanFromContext(context.Background(), "sql:query", zipkin.Kind(zipkinmodel.Client))
	setSpanDefaultTags(span, s.options.DefaultTags)

	if s.options.TagQuery {
		span.Tag("sql.query", s.query)
	}

	defer func() {
		setSpanError(span, err)
		span.Finish()
	}()

	rows, err = s.parent.Query(args)
	if err != nil {
		return nil, err
	}

	return
}

func (s zStmt) ExecContext(ctx context.Context, args []driver.NamedValue) (res driver.Result, err error) {
	if zipkin.SpanFromContext(ctx) == nil && !s.options.AllowRootSpan {
		return s.parent.(driver.StmtExecContext).ExecContext(ctx, args)
	}

	span, ctx := s.tracer.StartSpanFromContext(ctx, "sql/exec", zipkin.Kind(zipkinmodel.Client))
	defer func() {
		setSpanError(span, err)
		span.Finish()
	}()

	if s.options.TagQuery {
		span.Tag("sql.query", s.query)
	}

	setSpanDefaultTags(span, s.options.DefaultTags)

	execContext := s.parent.(driver.StmtExecContext)
	res, err = execContext.ExecContext(ctx, args)
	if err != nil {
		return nil, err
	}

	if s.options.TagAffectedRows {
		if affectedRows, err := res.RowsAffected(); err != nil {
			span.Tag("sql.affected_rows", fmt.Sprintf("%d", affectedRows))
		}
	}

	res, err = zResult{parent: res, tracer: s.tracer, ctx: ctx, options: s.options}, nil
	return
}

func (s zStmt) QueryContext(ctx context.Context, args []driver.NamedValue) (rows driver.Rows, err error) {
	if zipkin.SpanFromContext(ctx) == nil && !s.options.AllowRootSpan {
		return s.parent.(driver.StmtQueryContext).QueryContext(ctx, args)
	}

	span, ctx := s.tracer.StartSpanFromContext(ctx, "sql/query", zipkin.Kind(zipkinmodel.Client))
	defer func() {
		setSpanError(span, err)
		span.Finish()
	}()

	if s.options.TagQuery {
		span.Tag("sql.query", s.query)
	}

	setSpanDefaultTags(span, s.options.DefaultTags)

	defer func() {
		setSpanError(span, err)
		span.Finish()
	}()

	// we already tested driver to implement StmtQueryContext
	queryContext := s.parent.(driver.StmtQueryContext)
	rows, err = queryContext.QueryContext(ctx, args)
	if err != nil {
		return nil, err
	}

	return
}

// zTx implemens driver.Tx
type zTx struct {
	parent  driver.Tx
	ctx     context.Context
	tracer  *zipkin.Tracer
	options TraceOptions
}

func (t zTx) Commit() (err error) {
	if zipkin.SpanFromContext(t.ctx) != nil || t.options.AllowRootSpan {
		span, _ := t.tracer.StartSpanFromContext(t.ctx, "sql/commit", zipkin.Kind(zipkinmodel.Client))
		defer func() {
			setSpanDefaultTags(span, t.options.DefaultTags)
			setSpanError(span, err)
			span.Finish()
		}()
	}
	err = t.parent.Commit()
	return
}

func (t zTx) Rollback() (err error) {
	if zipkin.SpanFromContext(t.ctx) != nil || t.options.AllowRootSpan {
		span, _ := t.tracer.StartSpanFromContext(t.ctx, "sql/rollback", zipkin.Kind(zipkinmodel.Client))
		defer func() {
			setSpanDefaultTags(span, t.options.DefaultTags)
			setSpanError(span, err)
			span.Finish()
		}()
	}
	err = t.parent.Rollback()
	return
}

func setSpanError(span zipkin.Span, err error) {
	if err != nil {
		zipkin.TagError.Set(span, err.Error())
	}
}

func setSpanDefaultTags(span zipkin.Span, tags map[string]string) {
	for key, value := range tags {
		span.Tag(key, value)
	}
}
