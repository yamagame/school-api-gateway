package snapshot

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Equal スナップショットと一致比較
func Equal(t *testing.T, v interface{}, fname string) {
	b, err := json.MarshalIndent(v, "", "  ")
	assert.NoError(t, err)
	fpath := filepath.Join("./testdata/", fname)
	rp, err := os.Open(fpath)
	assert.NoError(t, err)
	data, err := io.ReadAll(rp)
	assert.NoError(t, err)
	assert.Equal(t, data, b)
}

// Equal スナップショットを保存
func Save(t *testing.T, v interface{}, fname string) {
	b, err := json.MarshalIndent(v, "", "  ")
	assert.NoError(t, err)
	err = os.MkdirAll("./testdata/", 0777)
	assert.NoError(t, err)
	fpath := filepath.Join("./testdata/", fname)
	os.WriteFile(fpath, b, 0666)
}
