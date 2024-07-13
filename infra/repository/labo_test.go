package repository

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yamagame/school-api-gateway/entity"
	"github.com/yamagame/school-api-gateway/infra"
	"github.com/yamagame/school-api-gateway/infra/infconv"
	"github.com/yamagame/school-api-gateway/infra/model"
	"github.com/yamagame/school-api-gateway/pkg/conv"
	"github.com/yamagame/school-api-gateway/pkg/snapshot"
)

func TestCreateUpdate(t *testing.T) {
	var err error
	ctx := context.Background()
	db := infra.DB()
	repo := NewLabo(db)

	labos := &conv.Records{}
	labos.Append(
		entity.NewLabo().
			Set(".id", int32(1)).
			Set(".name", "サトウ1").
			Set(".url", "http://sato.com"),
		entity.NewLabo().
			Set(".id", int32(2)).
			Set(".name", "シミズ2"),
		entity.NewLabo().
			Set(".id", int32(0)).
			Set(".name", "スズキ3").
			Set(".url", "http://zuzuki.com"),
	)

	out := labos.ValueMap()

	snapshot.Equal(t, out, "create-update.json")
	// snapshot.Save(t, out, "create-update.json")

	models, err := infconv.Labos.ToInfra(labos, nil)
	assert.NoError(t, err)

	err = repo.Upsert(ctx, model.NewLabos(models...))
	assert.NoError(t, err)
}

func TestLabosInfraToEntity(t *testing.T) {
	var err error
	ctx := context.Background()
	db := infra.DB()
	repo := NewLabo(db)
	res := repo.List(ctx, 10, 0)
	assert.NoError(t, res.Error)

	labos, err := infconv.Labos.ToEntity(res.Records)
	assert.NoError(t, err)

	out := labos.ValueMap()

	snapshot.Equal(t, out, "test-labos.json")
	// snapshot.Save(t, out, "test-labos.json")
}

func TestCreateFromCSV(t *testing.T) {
	fp, _ := os.Open("./testdata/create-labos.csv")
	defer fp.Close()

	records, err := conv.ReadCSV(fp)
	assert.NoError(t, err)

	labos := entity.NewLabo().
		NewRecords(records)

	out := labos.ValueMap()

	snapshot.Equal(t, out, "create-labos.json")
	// snapshot.Save(t, out, "create-labos.json")
}
