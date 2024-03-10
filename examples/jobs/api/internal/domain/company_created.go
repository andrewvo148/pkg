package domain

import (
	"andrewvo148/pkg/examples/jobs/api/pkg/ddd"

	"github.com/docker/distribution/uuid"
	"github.com/google/uuid"
)


type CompanyCreated struct {
	ddd.DomainEvent
	Id uuid.UUID
	title string
	description string
}

