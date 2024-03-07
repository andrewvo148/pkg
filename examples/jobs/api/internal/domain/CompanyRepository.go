package domain

import (
	"context"

	"github.com/google/uuid"
)

type CompanyRepo interface {
	GetAll(context.Context) ([]*Company, error)
	GetByID(context.Context, uuid.UUID)
	Create(context.Context, Company) error
	Update(context.Context, Company) (*Company, error)
}
