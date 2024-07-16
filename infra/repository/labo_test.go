package repository

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yamagame/school-api-gateway/infra"
	"github.com/yamagame/school-api-gateway/infra/dao/query"
	"github.com/yamagame/school-api-gateway/infra/infconv"
	"github.com/yamagame/school-api-gateway/infra/model"
	"github.com/yamagame/school-api-gateway/irmodel"
	"github.com/yamagame/school-api-gateway/pkg/irconv"
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

		snapshot.Match(t, out, "create-update.json")

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

			res, err := results.ShallowCopy()
			require.NoError(t, err)
			assert.Equal(t, 2, len(res))
			assert.Equal(t, "サトウ1", *res[0].Name)
			assert.Equal(t, "シミズ2", *res[1].Name)
		}

		{
			results := repo.FindWithName(ctx, []string{"サトウ1", "シミズ2", "スズキ3"})
			assert.NoError(t, results.Error)

			res, err := results.ShallowCopy()
			require.NoError(t, err)
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

	snapshot.Match(t, out, "test-labos.json")
}

func TestLabosCSVToIRModel(t *testing.T) {
	fp, _ := os.Open("./testdata/create-labos.csv")
	defer fp.Close()

	records, err := irmodel1.ReadCSV(fp)
	assert.NoError(t, err)

	labos := irmodel.NewLabo().
		NewRecords(records)
	assert.NoError(t, labos.Error)

	out := labos.ValueMap()

	snapshot.Match(t, out, "create-labos-1.json")

	ctx := context.Background()
	db := infra.DB()
	infra.ResetAutoIncrementForTest(db)
	db.Transaction(func(tx *gorm.DB) error {
		repo := NewLabo(tx)

		// 中間モデルからdaoに変換
		models := infconv.Labos.ToStruct(labos)
		assert.NoError(t, models.Error)

		// Upsertする
		err := repo.Upsert(ctx, models)
		assert.NoError(t, err)

		{
			results := repo.Find(ctx, []int32{1, 2})
			assert.NoError(t, results.Error)
			records := results.MustShallowCopy()
			out := snapshot.Delete(t, records, "CreatedAt", "UpdatedAt")
			snapshot.Match(t, out, "create-labos-2.json")
		}

		{
			results := repo.FindWithName(ctx, []string{"スズキ"})
			assert.NoError(t, results.Error)
			records := results.MustShallowCopy()
			out := snapshot.Delete(t, records, "CreatedAt", "UpdatedAt")
			snapshot.Match(t, out, "create-labos-3.json")
		}

		return fmt.Errorf("restore")
	})
}

func TestFind(t *testing.T) {
	ctx := context.Background()
	db := infra.DB()
	infra.ResetAutoIncrementForTest(db)
	db.Transaction(func(tx *gorm.DB) error {
		repo := NewLabo(tx)

		{
			labos := repo.Find(ctx, []int32{1, 2, 3})
			values := labos.MustShallowCopy()
			values[0].Desks = append(values[0].Desks,
				&model.Desk{
					LaboID: values[0].ID,
					Name:   "机1",
				},
			)
			values[1].Desks = append(values[1].Desks,
				&model.Desk{
					LaboID: values[1].ID,
				},
				&model.Desk{
					LaboID: values[1].ID,
				})
			err := repo.Update(ctx, labos)
			require.NoError(t, err)
		}

		{
			records := repo.Find(ctx, []int32{1, 2, 3})
			out := infconv.Labos.ToIRModel(records.ShallowCopy()).ValueMap()
			snapshot.Match(t, out, "find-irmodels-1.json")
		}

		{
			labos := repo.Find(ctx, []int32{1, 2, 3})
			values := labos.MustShallowCopy()
			values[0].Desks[0].Name = "机1-A"
			values[1].Desks[0].Name = "机2"
			values[1].Desks[1].Name = "机3"
			err := repo.Update(ctx, labos)
			require.NoError(t, err)
		}

		{
			records := repo.Find(ctx, []int32{1, 2, 3})
			out := infconv.Labos.ToIRModel(records.ShallowCopy()).ValueMap()
			snapshot.Match(t, out, "find-irmodels-2.json")
		}

		labos := repo.Find(ctx, []int32{1, 2, 3}).MustShallowCopy()
		out := snapshot.Delete(t, labos, "CreatedAt", "UpdatedAt")
		snapshot.Match(t, out, "find-labos.json")
		return fmt.Errorf("restore")
	})
}

func TestDao(t *testing.T) {
	ctx := context.Background()
	db := infra.DB()
	infra.ResetAutoIncrementForTest(db)
	db.Transaction(func(tx *gorm.DB) error {
		q := query.Use(tx)
		lb := q.Labo
		labos, err := lb.WithContext(ctx).
			Joins(lb.Building, lb.Group, lb.Program).
			Find()
		require.NoError(t, err)

		last := labos[len(labos)-1]

		last.BuildingID = irconv.ToPtr(int32(5))
		lb.Building.WithContext(ctx).Model(last).Replace(&model.Building{
			ID: 5,
		})

		lb.Desks.WithContext(ctx).Model(last).Append(
			&model.Desk{
				Name:        "テーブルA-1",
				ProductCode: "1234",
			},
			&model.Desk{
				Name: "テーブルB",
			},
		)

		dk := q.Desk
		dk.WithContext(ctx).Updates(
			&model.Desk{
				ID:          1,
				Name:        "テーブルA",
				ProductCode: "1234",
			},
		)

		last.Name = irconv.ToPtr("藤田テスト")
		lb.WithContext(ctx).Updates(last)

		{
			labos, err := lb.WithContext(ctx).
				Preload(lb.Desks, lb.Chairs).
				Joins(lb.Building, lb.Group, lb.Program).
				Find()
			require.NoError(t, err)

			out := snapshot.Delete(t, labos, "CreatedAt", "UpdatedAt")
			snapshot.Match(t, out, "test-dao.json")
		}
		return fmt.Errorf("restore")
	})
}
