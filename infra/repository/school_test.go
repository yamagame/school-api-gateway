package repository

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yamagame/school-api-gateway/entity"
	"github.com/yamagame/school-api-gateway/infra"
	"github.com/yamagame/school-api-gateway/infra/infconv"
	"github.com/yamagame/school-api-gateway/infra/model"
	"github.com/yamagame/school-api-gateway/pkg/conv"
)

func TestSchool(t *testing.T) {
	var err error
	ctx := context.Background()
	db := infra.DB()
	repo := NewSchool(db)
	labo1 := entity.NewLabo(0)
	err = labo1.Set(".id", int32(1))
	assert.NoError(t, err)
	err = labo1.Set(".name", "サトウ")
	assert.NoError(t, err)
	labo2 := entity.NewLabo(0)
	err = labo2.Set(".id", int32(2))
	assert.NoError(t, err)
	err = labo2.Set(".name", "シミズ")
	assert.NoError(t, err)
	labo3 := entity.NewLabo(0)
	err = labo3.Set(".id", int32(0))
	assert.NoError(t, err)
	err = labo3.Set(".name", "スズキ")
	assert.NoError(t, err)
	labos := []*conv.Record{
		labo1,
		labo2,
		labo3,
	}

	models := []*model.Labo{}
	for _, entity := range labos {
		labo, err := infconv.LaboToInfra(entity)
		assert.NoError(t, err)
		models = append(models, labo)
	}

	creates := []*model.Labo{}
	updates := []*model.Labo{}
	for _, labo := range models {
		if labo.ID == 0 {
			creates = append(creates, labo)
		} else {
			updates = append(updates, labo)
		}
	}

	err = repo.UpdateLabos(ctx, updates)
	assert.NoError(t, err)
	err = repo.CreateLabos(ctx, creates)
	assert.NoError(t, err)
}

func TestLabos(t *testing.T) {
	var err error
	ctx := context.Background()
	db := infra.DB()
	repo := NewSchool(db)
	res, err := repo.ListLabos(ctx, 10, 0)
	assert.NoError(t, err)
	fmt.Println(res)
}
