package entity

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPerson(t *testing.T) {
	person := NewPerson(0)

	bytes, err := json.MarshalIndent(person, "", "  ")
	assert.NoError(t, err)
	fmt.Println(string(bytes))
}
