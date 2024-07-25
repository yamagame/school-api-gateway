package imodel

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yamagame/school-api-gateway/pkg/iconv"
	"github.com/yamagame/school-api-gateway/pkg/snapshot"
)

func TestLabo(t *testing.T) {
	var val interface{}
	var err error
	labo := NewLabo()
	val, err = labo.Get(".name")
	assert.NoError(t, err)
	assert.Equal(t, "", val)

	err = labo.Set(".name", "名前").Error
	assert.NoError(t, err)

	val, err = labo.Get(".name")
	assert.NoError(t, err)
	assert.Equal(t, "名前", val)

	desks, err := labo.HasMany(".desk")
	assert.NoError(t, err)

	desks.NewOne().Set(".id", int32(11)).Set(".labo_id", int32(1))
	desks.NewOne().Set(".id", int32(12)).Set(".labo_id", int32(2))
	desks.NewOne().Set(".id", int32(13)).Set(".labo_id", int32(3))

	desks.Take(".id", int32(12)).Set(".labo_id", int32(5))

	chairs, err := labo.HasMany(".chair")
	assert.NoError(t, err)

	chairs.NewOne().Set(".id", int32(21)).Set(".labo_id", int32(10))
	chairs.NewOne().Set(".id", int32(22)).Set(".labo_id", int32(20))

	out := labo.ValueMap()
	snapshot.Match(t, out, "test-labo1.json")
}

func TestLaboCSV(t *testing.T) {
	fp, _ := os.Open("./testdata/test-labo.csv")
	defer fp.Close()

	records, err := iconv.ReadCSV(fp)
	assert.NoError(t, err)

	labos := NewLabo().
		NewRecords(records)
	out := labos.ValueMap()

	snapshot.Match(t, out, "test-labo.json")
}

func TestLaboHasManyCSV(t *testing.T) {
	// テストデータ準備
	fp, _ := os.Open("./testdata/test-many.csv")
	defer fp.Close()

	records, err := iconv.ReadCSV(fp)
	assert.NoError(t, err)

	labos := NewLabo().
		NewRecords(records)
	assert.NoError(t, labos.Error)

	out := labos.ValueMap()

	snapshot.Match(t, out, "test-many.json")
}

func TestLaboFillCSV(t *testing.T) {
	// テストデータ準備
	fp, _ := os.Open("./testdata/test-fill.csv")
	defer fp.Close()

	records, err := iconv.ReadCSV(fp)
	assert.NoError(t, err)

	snapshot.Match(t, records, "test-fill.json")

	labos := NewLabo().
		NewRecords(records)
	assert.NoError(t, labos.Error)

	out := labos.ValueMap()

	snapshot.Match(t, out, "test-labos.json")
}
