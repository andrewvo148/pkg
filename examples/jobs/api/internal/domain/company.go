package domain

import (
	"andrewvo148/pkg/examples/jobs/api/pkg/ddd"
	"errors"
)

var (
	ErrTitleCannotBeBlank = errors.New("The title cannot be blank")
	ErrDescriptionCannotBeBlank = errors.New("The description cannot be blank")

)
type Company struct {
	ddd.AggregateRoot
	ID    int64
	title int64
}

func NewCompany(id string, title string, description string) *Company {
	return &Company{}
}

