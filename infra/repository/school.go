package repository

import (
	"context"

	"github.com/yamagame/school-api-gateway/infra/dao/query"
	"github.com/yamagame/school-api-gateway/infra/infconv"
	"github.com/yamagame/school-api-gateway/infra/model"
	"github.com/yamagame/school-api-gateway/pkg/field"
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

func (r *School) SaveLabos(ctx context.Context, labos []*field.Field) error {
	models := []*model.Labo{}
	for _, v := range labos {
		labo, err := infconv.LaboToInfra(v)
		if err != nil {
			return err
		}
		models = append(models, labo)
	}
	q := query.Use(r.db)
	lb := q.Labo
	creates := []*model.Labo{}
	for _, labo := range models {
		if labo.ID == 0 {
			creates = append(creates, labo)
		} else {
			_, err := lb.WithContext(ctx).Where(lb.ID.Eq(labo.ID)).Updates(labo)
			if err != nil {
				return err
			}
		}
	}
	return lb.WithContext(ctx).CreateInBatches(creates, 100)
}

func (r *School) ListLabos(ctx context.Context, limit, offset int32) ([]*field.Field, error) {
	q := query.Use(r.db)
	lb := q.Labo
	labos, err := lb.WithContext(ctx).
		Joins(lb.Building, lb.Group, lb.Program).
		Limit(int(limit)).
		Offset(int(offset)).
		Order(lb.ID.Asc()).
		Find()
	if err != nil {
		return nil, err
	}

	var fields []*field.Field
	for _, v := range labos {
		labo, err := infconv.LaboToEntity(v)
		if err != nil {
			return nil, err
		}
		fields = append(fields, labo)
	}
	return fields, err
}
