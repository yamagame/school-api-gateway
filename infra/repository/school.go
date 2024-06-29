package repository

import (
	"context"

	"github.com/yamagame/school-api-gateway/infra/dao/query"
	"github.com/yamagame/school-api-gateway/infra/model"
	"gorm.io/gorm"
)

type School struct {
	db *gorm.DB
}

func NewSchool(db *gorm.DB) *School {
	return &School{
		db: db,
	}
}

func (r *School) CreateLabos(ctx context.Context, labos []*model.Labo) error {
	q := query.Use(r.db)
	lb := q.Labo
	return lb.WithContext(ctx).CreateInBatches(labos, 100)
}

func (r *School) UpdateLabos(ctx context.Context, labos []*model.Labo) error {
	q := query.Use(r.db)
	lb := q.Labo
	for _, labo := range labos {
		_, err := lb.WithContext(ctx).Where(lb.ID.Eq(labo.ID)).Updates(labo)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *School) ListLabos(ctx context.Context, limit, offset int32) ([]*model.Labo, error) {
	q := query.Use(r.db)
	lb := q.Labo
	return lb.WithContext(ctx).
		Joins(lb.Building, lb.Group, lb.Program).
		Limit(int(limit)).
		Offset(int(offset)).
		Order(lb.ID.Asc()).
		Find()
}
