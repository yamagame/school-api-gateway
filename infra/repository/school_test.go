package repository

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yamagame/school-api-gateway/entity"
	"github.com/yamagame/school-api-gateway/infra"
)

func TestSchool(t *testing.T) {
	ctx := context.Background()
	db := infra.DB()
	repo := NewSchool(db)
	labo1 := entity.NewLabo()
	labo1.SetValue("id", int32(1))
	labo1.SetValue("name", "サトウ")
	labo2 := entity.NewLabo()
	labo2.SetValue("id", int32(2))
	labo2.SetValue("name", "シミズ")
	labo3 := entity.NewLabo()
	labo3.SetValue("id", int32(0))
	labo3.SetValue("name", "スズキ")
	labos := []*entity.Labo{
		labo1,
		labo2,
		labo3,
	}
	err := repo.SaveLabos(ctx, labos)
	assert.NoError(t, err)
}
