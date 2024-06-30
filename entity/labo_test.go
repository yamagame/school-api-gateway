package entity

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yamagame/school-api-gateway/pkg/conv"
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

	out := conv.Records{}
	for _, record := range records {
		r, err := conv.NewRecordWithMap(record, NewLabo)
		assert.NoError(t, err)
		out = append(out, r)
	}

	bytes, err := json.MarshalIndent(out.ValueMap(), "", "  ")
	assert.NoError(t, err)
	fmt.Println(string(bytes))
}
