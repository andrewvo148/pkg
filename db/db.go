package db


type DB interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}