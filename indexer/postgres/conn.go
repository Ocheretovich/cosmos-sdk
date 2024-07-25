package postgres

import (
	"context"
	"database/sql"
)

// DBConn is an interface that abstracts the *sql.DB, *sql.Tx and *sql.Conn types.
type DBConn interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}
