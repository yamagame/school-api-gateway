package conv

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/client-go/util/jsonpath"
	kjson "sigs.k8s.io/json"
)

func TestField(t *testing.T) {
	f := NewRecord()
	f.SetValue("val", int64(100))
	field0 := NewRecord()
	field1 := NewRecord()
	field2 := NewRecord()
	f.SetHasOne("field", field0)
	f.SetHasMany("fields", field1, field2)

	field0.SetValue("val1", "hello val1")
	field1.SetValue("val2", "hello val2")
	field2.SetValue("val3", "hello val3")

	value, err := f.Value(".val")
	assert.NoError(t, err)
	assert.Equal(t, int64(100), value.Get())

	field, err := f.HasOne(".field")
	assert.NoError(t, err)
	assert.NotNil(t, field)

	array, err := f.HasMany(".fields")
	assert.NoError(t, err)
	assert.NotNil(t, array)

	bytes1, err := json.MarshalIndent(f, "", "  ")
	assert.NoError(t, err)
	fmt.Println(string(bytes1))

	g := f.Copy()
	bytes2, err := json.MarshalIndent(g, "", "  ")
	assert.NoError(t, err)
	fmt.Println(string(bytes2))

	valueg, err := g.Value(".val")
	assert.NoError(t, err)
	assert.Equal(t, int64(100), valueg.Get())

	fieldg, err := g.HasOne(".field")
	assert.NoError(t, err)
	assert.NotNil(t, fieldg)

	l := NewRecord()
	// l.SetValue("val", 100)
	// field0l := &Record{}
	// field1l := &Record{}
	// field2l := &Record{}
	// l.SetHasOne("field", field0l)
	// l.SetHasMany("fields", field1l, field2l)
	err = kjson.UnmarshalCaseSensitivePreserveInts(bytes2, &l)
	assert.NoError(t, err)
	bytes3, err := json.MarshalIndent(l, "", "  ")
	assert.NoError(t, err)
	fmt.Println(string(bytes3))

	valuel, err := l.Value(".val")
	assert.NoError(t, err)
	valuef, err := f.Value(".val")
	assert.NoError(t, err)
	assert.Equal(t, valuef.Get(), valuel.Get())

	fieldl, err := l.HasOne(".field")
	assert.NoError(t, err)
	assert.NotNil(t, fieldl)
}

func printValues(jp *jsonpath.JSONPath, values [][]reflect.Value) {
	for _, val := range values {
		jp.PrintResults(os.Stdout, val)
	}
}
