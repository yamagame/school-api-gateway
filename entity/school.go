package entity

import (
	"context"
)

type SchoolRepositoryInterface interface {
}

type SchoolInterface interface {
}

type School struct {
	ctx  context.Context
	repo SchoolRepositoryInterface
}

func NewSchool(ctx context.Context, repo SchoolRepositoryInterface) *School {
	return &School{
		ctx:  ctx,
		repo: repo,
	}
}
