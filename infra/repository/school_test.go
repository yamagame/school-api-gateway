package repository

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yamagame/school-api-gateway/entity"
	"github.com/yamagame/school-api-gateway/infra"
	"github.com/yamagame/school-api-gateway/pkg/field"
)

func TestSchool(t *testing.T) {
	var err error
	ctx := context.Background()
	db := infra.DB()
	repo := NewSchool(db)
	labo1 := entity.NewLabo(0)
	err = labo1.Set("id", int32(1))
	assert.NoError(t, err)
	err = labo1.Set("name", "サトウ")
	assert.NoError(t, err)
	labo2 := entity.NewLabo(0)
	err = labo2.Set("id", int32(2))
	assert.NoError(t, err)
	err = labo2.Set("name", "シミズ")
	assert.NoError(t, err)
	labo3 := entity.NewLabo(0)
	err = labo3.Set("id", int32(0))
	assert.NoError(t, err)
	err = labo3.Set("name", "スズキ")
	assert.NoError(t, err)
	labos := []*field.Field{
		labo1,
		labo2,
		labo3,
	}
	err = repo.SaveLabos(ctx, labos)
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
