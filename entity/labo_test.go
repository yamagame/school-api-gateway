package entity

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
	err = labo.Set(".name", "名前")
	assert.NoError(t, err)
	val, err = labo.Get(".name")
	assert.NoError(t, err)
	assert.Equal(t, "名前", val)
}

func TestLaboCSV(t *testing.T) {
	fp, _ := os.Open("./testdata/test-labo.csv")
	defer fp.Close()

	records, err := conv.ReadCSV(fp)
	assert.NoError(t, err)

	o := conv.Records{}
	for _, record := range records {
		r, err := conv.NewRecordWithMap(record, NewLabo)
		assert.NoError(t, err)
		o = append(o, r)
	}
	out := o.ValueMap()

	snapshot.Equal(t, out, "test-labo.json")
	// snapshot.Save(t, out, "test-labo.json")
}

func TestLaboHasManyCSV(t *testing.T) {
	fp, _ := os.Open("./testdata/test-many.csv")
	defer fp.Close()

	records, err := conv.ReadCSV(fp)
	assert.NoError(t, err)

	newMany := func() *conv.Record {
		v := conv.NewRecord()
		v.SetValue("name", "")
		v.SetHasOne("desk", NewDesk())
		return v
	}

	work, err := conv.NewRecordsWithMap(records, newMany)
	assert.NoError(t, err)

	desks, err := work.MergeRecords("name", "desk", NewLabo)
	assert.NoError(t, err)
	out := desks.ValueMap()

	snapshot.Equal(t, out, "test-many.json")
	// snapshot.Save(t, out, "test-many.json")
}
