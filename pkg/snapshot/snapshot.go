package snapshot

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func Update(t *testing.T) {
	t.Setenv("UPDATE_SNAPS", "true")
}

func Match(t *testing.T, v interface{}, fname string) {
	update := os.Getenv("UPDATE_SNAPS")
	if update == "true" {
		Save(t, v, fname)
	} else {
		Equal(t, v, fname)
	}
}

func Delete[T any](t *testing.T, in []*T, keys ...string) []interface{} {
	var _delete func(i interface{}, keys ...string)
	_delete = func(i interface{}, keys ...string) {
		if m, ok := i.(map[string]interface{}); ok {
			for _, key := range keys {
				delete(m, key)
			}
			for _, v := range m {
				if n, ok := v.(map[string]interface{}); ok {
					_delete(n, keys...)
				}
			}
		}
	}
	ret := []interface{}{}
	for _, v := range in {
		input, err := json.Marshal(v)
		require.NoError(t, err)
		var i interface{}
		err = json.Unmarshal([]byte(input), &i)
		require.NoError(t, err)
		_delete(i, keys...)
		ret = append(ret, i)
	}
	return ret
}
