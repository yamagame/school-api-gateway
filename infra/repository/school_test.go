package repository

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yamagame/school-api-gateway/entity"
	"github.com/yamagame/school-api-gateway/infra"
)

func TestSchool(t *testing.T) {
	var err error
	ctx := context.Background()
	db := infra.DB()
	repo := NewSchool(db)
	labo1 := entity.NewLabo()
	err = labo1.SetValue("id", int32(1))
	assert.NoError(t, err)
	err = labo1.SetValue("name", "サトウ")
	assert.NoError(t, err)
	labo2 := entity.NewLabo()
	err = labo2.SetValue("id", int32(2))
	assert.NoError(t, err)
	err = labo2.SetValue("name", "シミズ")
	assert.NoError(t, err)
	labo3 := entity.NewLabo()
	err = labo3.SetValue("id", int32(0))
	assert.NoError(t, err)
	err = labo3.SetValue("name", "スズキ")
	assert.NoError(t, err)
	labos := []*entity.Labo{
		labo1,
		labo2,
		labo3,
	}
	err = repo.SaveLabos(ctx, labos)
	assert.NoError(t, err)
}
