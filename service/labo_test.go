package service

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yamagame/school-api-gateway/infra"
	"github.com/yamagame/school-api-gateway/infra/repository"
	"github.com/yamagame/school-api-gateway/pkg/irconv"
	"github.com/yamagame/school-api-gateway/pkg/snapshot"
	"github.com/yamagame/school-api-gateway/proto/school"
	"gorm.io/gorm"
)

func TestLabo(t *testing.T) {
	ctx := context.Background()

	// サービスを作成
	db := infra.DB()
	db.Transaction(func(tx *gorm.DB) error {
		svc := NewLabo(repository.NewLabo(tx))

		// 1レコード作成
		id, err := svc.Create(ctx)
		require.NoError(t, err)
		require.NotEqual(t, 0, id)

		laboname := "テスト研究室"
		copyname := "テスト研究室コピー"

		// プライマリキーでカラムを更新
		id2, err := svc.Update(ctx, &school.Labo{
			Id:   id,
			Name: laboname,
		})
		require.NoError(t, err)
		require.Equal(t, id, id2)

		// プライマリキーで検索
		labo, err := svc.Find(ctx, id2)
		require.NoError(t, err)
		require.Equal(t, laboname, labo.Name)

		// プライマリキーでコピー
		id3, err := svc.Copy(ctx, id2)
		require.NoError(t, err)
		require.NotEqual(t, 0, id3)

		// プライマリキーでカラムを更新
		_, err = svc.Update(ctx, &school.Labo{
			Id:   id3,
			Name: copyname,
		})
		require.NoError(t, err)

		// プライマリキーで検索
		labo3, err := svc.Find(ctx, id3)
		require.NoError(t, err)
		require.Equal(t, copyname, labo3.Name)

		return fmt.Errorf("restore")
	})
}

func TestCreateWithMap(t *testing.T) {
	var err error

	fp, _ := os.Open("./testdata/create-labo.csv")
	defer fp.Close()
	records, err := irconv.ReadCSV(fp)
	assert.NoError(t, err)

	ctx := context.Background()
	db := infra.DB()
	db.Transaction(func(tx *gorm.DB) error {
		svc := NewLabo(repository.NewLabo(tx))

		// []map[string]stringから作成
		id, err := svc.CreateWithMap(ctx, records)
		assert.NoError(t, err)
		assert.NotEqual(t, 0, id)
		return fmt.Errorf("restore")
	})
}

func TestUpdateWithMap(t *testing.T) {
	var err error

	fp, _ := os.Open("./testdata/update-labo.csv")
	defer fp.Close()
	records, err := irconv.ReadCSV(fp)
	assert.NoError(t, err)

	ctx := context.Background()
	db := infra.DB()
	db.Transaction(func(tx *gorm.DB) error {
		svc := NewLabo(repository.NewLabo(tx))

		// []map[string]stringから更新
		id, err := svc.UpdateWithMap(ctx, records)
		assert.NoError(t, err)
		assert.NotEqual(t, 0, id)
		return fmt.Errorf("restore")
	})
}

func TestList(t *testing.T) {
	ctx := context.Background()
	db := infra.DB()
	svc := NewLabo(repository.NewLabo(db))

	tests := []struct {
		pageSize int
		offset   int
		wants    int
	}{
		{0, 0, 0},
		{10, 0, 10},
	}
	results := []interface{}{}
	for _, tt := range tests {
		labos, err := svc.List(ctx, int32(tt.pageSize), int32(tt.offset))
		assert.NoError(t, err)
		assert.Equal(t, tt.wants, len(labos))
		results = append(results, labos)
	}

	snapshot.Match(t, results, "test-list.json")
}
