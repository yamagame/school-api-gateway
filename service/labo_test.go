package service

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yamagame/school-api-gateway/infra"
	"github.com/yamagame/school-api-gateway/infra/repository"
	"github.com/yamagame/school-api-gateway/pkg/conv"
	"github.com/yamagame/school-api-gateway/proto/school"
)

func TestLabo(t *testing.T) {
	var err error
	ctx := context.Background()

	// サービスを作成
	db := infra.DB()
	svc := NewLabo(repository.NewLabo(db))

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
}

func TestCreateWithMap(t *testing.T) {
	var err error

	fp, _ := os.Open("./testdata/create-labo.csv")
	defer fp.Close()
	records, err := conv.ReadCSV(fp)
	assert.NoError(t, err)

	ctx := context.Background()
	db := infra.DB()
	svc := NewLabo(repository.NewLabo(db))

	// []map[string]stringから作成
	id, err := svc.CreateWithMap(ctx, records)
	assert.NoError(t, err)
	assert.NotEqual(t, 0, id)
}

func TestUpdateWithMap(t *testing.T) {
	var err error

	fp, _ := os.Open("./testdata/update-labo.csv")
	defer fp.Close()
	records, err := conv.ReadCSV(fp)
	assert.NoError(t, err)

	ctx := context.Background()
	db := infra.DB()
	svc := NewLabo(repository.NewLabo(db))

	// []map[string]stringから更新
	id, err := svc.UpdateWithMap(ctx, records)
	assert.NoError(t, err)
	assert.NotEqual(t, 0, id)
}
