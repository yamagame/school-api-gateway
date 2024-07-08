package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExec(t *testing.T) {
	jsondata, err := os.ReadFile(filepath.Join("testdata", "test.json"))
	assert.NoError(t, err)
	tmpldata, err := os.ReadFile(filepath.Join("testdata", "test.tmpl"))
	assert.NoError(t, err)
	exec(jsondata, tmpldata, "")
}
