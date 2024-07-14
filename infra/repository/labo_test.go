package repository

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yamagame/school-api-gateway/infra"
	"github.com/yamagame/school-api-gateway/infra/infconv"
	"github.com/yamagame/school-api-gateway/irmodel"
	irmodel1 "github.com/yamagame/school-api-gateway/pkg/irconv"
	"github.com/yamagame/school-api-gateway/pkg/snapshot"
	"gorm.io/gorm"
)

func TestCreateUpdate(t *testing.T) {
	ctx := context.Background()
	db := infra.DB()
	db.Transaction(func(tx *gorm.DB) error {
		repo := NewLabo(tx)

		labos := &irmodel1.Records{}
		labos.Append(
			irmodel.NewLabo().
				Set(".id", int32(1)).
				Set(".name", "サトウ1").
				Set(".url", "http://sato.com"),
			irmodel.NewLabo().
				Set(".id", int32(2)).
				Set(".name", "シミズ2"),
			irmodel.NewLabo().
				Set(".id", int32(0)).
				Set(".name", "スズキ3").
				Set(".url", "http://zuzuki.com"),
		)

		out := labos.ValueMap()

		snapshot.Equal(t, out, "create-update.json")
		// snapshot.Save(t, out, "create-update.json")

		// 中間モデルからdaoに変換
		models := infconv.Labos.ToStruct(labos)
		assert.NoError(t, models.Error)

		// Upsertする
		err := repo.Upsert(ctx, models)
		assert.NoError(t, err)

		// Find
		{
			results := repo.Find(ctx, []int32{1, 2})
			assert.NoError(t, results.Error)

			res := results.ShallowCopy()
			assert.Equal(t, 2, len(res))
			assert.Equal(t, "サトウ1", *res[0].Name)
			assert.Equal(t, "シミズ2", *res[1].Name)
		}

		{
			results := repo.FindWithName(ctx, []string{"サトウ1", "シミズ2", "スズキ3"})
			assert.NoError(t, results.Error)

			res := results.ShallowCopy()
			assert.Equal(t, 3, len(res))
			assert.Equal(t, "サトウ1", *res[0].Name)
			assert.Equal(t, int32(1), res[0].ID)
			assert.Equal(t, "シミズ2", *res[1].Name)
			assert.Equal(t, int32(2), res[1].ID)
			assert.Equal(t, "スズキ3", *res[2].Name)
			assert.NotEqual(t, int32(0), res[2].ID)
		}

		return fmt.Errorf("restore")
	})
}

func TestLabosInfraToIRModel(t *testing.T) {
	ctx := context.Background()
	db := infra.DB()
	repo := NewLabo(db)
	res := repo.List(ctx, 10, 0)
	assert.NoError(t, res.Error)

	labos := infconv.Labos.ToIRModel(res.ShallowCopy())
	assert.NoError(t, labos.Error)

	out := labos.ValueMap()

	snapshot.Equal(t, out, "test-labos.json")
	// snapshot.Save(t, out, "test-labos.json")
}

func TestLabosCSVToIRModel(t *testing.T) {
	fp, _ := os.Open("./testdata/create-labos.csv")
	defer fp.Close()

	records, err := irmodel1.ReadCSV(fp)
	assert.NoError(t, err)

	labos := irmodel.NewLabo().
		NewRecords(records)

	out := labos.ValueMap()

	snapshot.Equal(t, out, "create-labos.json")
	// snapshot.Save(t, out, "create-labos.json")
}
