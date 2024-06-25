package repository

import (
	"context"

	"github.com/yamagame/school-api-gateway/infra/conv"
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
		labo, err := conv.LaboToInfra(v)
		if err != nil {
			return err
		}
		models = append(models, labo)
	}
	creates := []*model.Labo{}
	for _, labo := range models {
		if labo.ID == 0 {
			creates = append(creates, labo)
		} else {
			err := r.db.WithContext(ctx).Debug().Where("id", labo.ID).Updates(labo).Error
			if err != nil {
				return err
			}
		}
	}
	return r.db.WithContext(ctx).Debug().CreateInBatches(creates, 100).Error
}

func (r *School) ListLabos(ctx context.Context, limit, offset int32) ([]*field.Field, error) {
	var models []*model.Labo
	err := r.db.WithContext(ctx).
		Joins("Group").
		Joins("Program").
		Joins("Building").
		Limit(int(limit)).
		Offset(int(offset)).
		Order("id").
		Find(&models).Error
	if err != nil {
		return nil, err
	}
	labos := []*field.Field{}
	for _, v := range models {
		labo, err := conv.LaboToEntity(v)
		if err != nil {
			return nil, err
		}
		labos = append(labos, labo)
	}
	return labos, err
}
