package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLabo(t *testing.T) {
	var val string
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
