package domain

import (
	"context"

	"github.com/google/uuid"
)

type JobRepo interface {
	GetAll(context.Context) ([]*Job, error)
	GetByID(context.Context, uuid.UUID)
	Create(context.Context, *Job) error
	Update(context.Context, *Job) (*Job, error)
}
