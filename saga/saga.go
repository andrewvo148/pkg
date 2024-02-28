package saga

import "context"

type Sage interface {
	Execute(ctx context.Context) error
	Rollback(ctx context.Context) error
}
