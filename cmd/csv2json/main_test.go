package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yamagame/school-api-gateway/pkg/snapshot"
)

func TestCSV2JSON(t *testing.T) {
	testdata := filepath.Join("testdata", "test-data.csv")
	fp, err := os.Open(testdata)
	assert.NoError(t, err)
	v, err := dumpCSV(fp, true)
	assert.NoError(t, err)

	out := []map[string]interface{}{}
	err = json.Unmarshal(v, &out)
	assert.NoError(t, err)

	snapshot.Match(t, out, "test-data.json")
}
