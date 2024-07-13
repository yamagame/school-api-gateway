package irmodel

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yamagame/school-api-gateway/pkg/conv"
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

	desks.NewOne().Set(".id", int32(11))
	desks.NewOne().Set(".id", int32(12))
	desks.NewOne().Set(".id", int32(13))

	out := labo.ValueMap()
	snapshot.Equal(t, out, "test-labo1.json")
	// snapshot.Save(t, out, "test-labo1.json")
}

func TestLaboCSV(t *testing.T) {
	fp, _ := os.Open("./testdata/test-labo.csv")
	defer fp.Close()

	records, err := conv.ReadCSV(fp)
	assert.NoError(t, err)

	labos := NewLabo().
		NewRecords(records)
	out := labos.ValueMap()

	snapshot.Equal(t, out, "test-labo.json")
	// snapshot.Save(t, out, "test-labo.json")
}

func TestLaboHasManyCSV(t *testing.T) {
	// テストデータ準備
	fp, _ := os.Open("./testdata/test-many.csv")
	defer fp.Close()

	records, err := conv.ReadCSV(fp)
	assert.NoError(t, err)

	labos := NewLabo().
		NewRecords(records)
	assert.NoError(t, labos.Error)

	out := labos.ValueMap()

	snapshot.Equal(t, out, "test-many.json")
	// snapshot.Save(t, out, "test-many.json")
}

func TestLaboFillCSV(t *testing.T) {
	// テストデータ準備
	fp, _ := os.Open("./testdata/test-fill.csv")
	defer fp.Close()

	records, err := conv.ReadCSV(fp)
	assert.NoError(t, err)

	snapshot.Equal(t, records, "test-fill.json")
	// snapshot.Save(t, records, "test-fill.json")

	labos := NewLabo().
		NewRecords(records)
	assert.NoError(t, labos.Error)

	out := labos.ValueMap()

	snapshot.Equal(t, out, "test-labos.json")
	// snapshot.Save(t, out, "test-labos.json")
}
