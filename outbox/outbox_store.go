package outbox

import (
	"github.com/andrewvo148/pkg/db"
)

type OutboxStore interface {
	Save(ctx context.context, msg Message) error
	FindUnpublised(ctx context.Context, limit int) ()
	MarkPublished(ctx context.Context, ids ...string) error
}

type OutboxStoreSimple struct {
	tableName string
	db db.DB
}

func (o OutboxStoreSimple) Save(ctx context.context, msg Message) error {
	const query = "INSERT INTO %s (id, subject, data, created_at)" VALUES ($1, $2, $3, $4)"
	_, err := o.db.ExecContext(ctx, queryFormat, msg.ID(), msg.Subbject(), msg.Date, time.Now())
	return err
}

func (o OutboxStoreSimple) queryFormat(query string) string {
	return fmt.Sprintf(query, s.tableName)
}