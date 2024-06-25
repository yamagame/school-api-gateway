package field

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestField(t *testing.T) {
	f := Field{}
	f.SetValue("val", 100)
	f.SetField("field", &Field{})
	f.SetArray("array", &Field{}, &Field{})

	value, err := f.Value("val")
	assert.NoError(t, err)
	assert.Equal(t, 100, value.Get())

	field, err := f.Field("field")
	assert.NoError(t, err)
	assert.NotNil(t, field)

	array, err := f.Array("array")
	assert.NoError(t, err)
	assert.NotNil(t, array)
}
