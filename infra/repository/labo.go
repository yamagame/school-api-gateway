package repository

import (
	"context"

	"github.com/yamagame/school-api-gateway/infra/dao/query"
	"github.com/yamagame/school-api-gateway/infra/model"
	"gorm.io/gorm"
)

type LaboInterface interface {
	Create(ctx context.Context, labos []*model.Labo) error
	Update(ctx context.Context, labos []*model.Labo) error
	Find(ctx context.Context, ids []int32) ([]*model.Labo, error)
	List(ctx context.Context, limit, offset int32) ([]*model.Labo, error)
}

type Labo struct {
	db *gorm.DB
}

func NewLabo(db *gorm.DB) *Labo {
	return &Labo{
		db: db,
	}
}

func (r *Labo) Create(ctx context.Context, labos []*model.Labo) error {
	q := query.Use(r.db)
	lb := q.Labo
	return lb.WithContext(ctx).CreateInBatches(labos, 100)
}

func (r *Labo) Find(ctx context.Context, ids []int32) ([]*model.Labo, error) {
	q := query.Use(r.db)
	lb := q.Labo
	return lb.WithContext(ctx).Where(lb.ID.In(ids...)).Find()
}

func (r *Labo) Update(ctx context.Context, labos []*model.Labo) error {
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

func (r *Labo) List(ctx context.Context, limit, offset int32) ([]*model.Labo, error) {
	q := query.Use(r.db)
	lb := q.Labo
	return lb.WithContext(ctx).
		Joins(lb.Building, lb.Group, lb.Program).
		Limit(int(limit)).
		Offset(int(offset)).
		Order(lb.ID.Asc()).
		Find()
}
