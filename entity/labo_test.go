package entity

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLabo(t *testing.T) {
	var val interface{}
	var err error
	labo := NewLabo()
	val, err = labo.GetValue("name")
	assert.NoError(t, err)
	assert.Equal(t, "", val)
	err = labo.SetValue("name", "名前")
	assert.NoError(t, err)
	val, err = labo.GetValue("name")
	assert.NoError(t, err)
	assert.Equal(t, "名前", val)
}

func TestLaboCSV(t *testing.T) {
	fp, _ := os.Open("./testdata/test-labo.csv")
	defer fp.Close()

	labos, err := ReadLaboCSV(fp)
	assert.NoError(t, err)

	bytes, err := json.MarshalIndent(labos.ToMap(), "", "  ")
	assert.NoError(t, err)
	fmt.Println(string(bytes))
}
