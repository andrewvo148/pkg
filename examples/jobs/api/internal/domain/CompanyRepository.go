package domain

import "context"

type CompanyRepo interface {
	Save(ctx context.Context, Job *Job) error
	Find(ctx context.Context, JobID uint64) (*Job, error)
}
