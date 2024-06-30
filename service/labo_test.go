package service

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yamagame/school-api-gateway/infra"
	"github.com/yamagame/school-api-gateway/infra/repository"
	"github.com/yamagame/school-api-gateway/pkg/conv"
	"github.com/yamagame/school-api-gateway/proto/school"
)

func TestLabo(t *testing.T) {
	var err error
	ctx := context.Background()
	db := infra.DB()
	s := NewLabo(repository.NewLabo(db))
	id, err := s.Create(ctx)
	assert.NoError(t, err)
	assert.NotEqual(t, 0, id)
	fmt.Println(id)

	laboname := "テスト研究室"
	copyname := "テスト研究室コピー"

	id2, err := s.Update(ctx, &school.Labo{
		Id:   id,
		Name: laboname,
	})
	assert.NoError(t, err)
	assert.Equal(t, id, id2)

	labo, err := s.Find(ctx, id2)
	assert.NoError(t, err)
	assert.Equal(t, laboname, labo.Name)

	id3, err := s.Copy(ctx, id2)
	assert.NoError(t, err)

	_, err = s.Update(ctx, &school.Labo{
		Id:   id3,
		Name: copyname,
	})
	assert.NoError(t, err)
	assert.Equal(t, id, id2)

	labo3, err := s.Find(ctx, id3)
	assert.NoError(t, err)
	assert.Equal(t, copyname, labo3.Name)
}

func TestCreateWithMap(t *testing.T) {
	var err error

	fp, _ := os.Open("./testdata/test-labo.csv")
	defer fp.Close()
	records, err := conv.ReadCSV(fp)
	assert.NoError(t, err)

	ctx := context.Background()
	db := infra.DB()
	s := NewLabo(repository.NewLabo(db))
	id, err := s.CreateWithMap(ctx, records)
	assert.NoError(t, err)
	assert.NotEqual(t, 0, id)
}
