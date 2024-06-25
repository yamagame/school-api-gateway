package infconv

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yamagame/school-api-gateway/infra/model"
	"github.com/yamagame/school-api-gateway/pkg/conv"
)

func TestLaboToEntity(t *testing.T) {
	in := &model.Labo{
		ID:      10,
		Name:    conv.ToPtr("名前"),
		GroupID: conv.ToPtr(int32(11)),
		Group: model.Group{
			Name: "グループ名",
		},
		ProgramID: conv.ToPtr(int32(12)),
		Program: model.Program{
			Name: "プログラム名",
		},
		BuildingID: conv.ToPtr(int32(12)),
		Building: model.Building{
			Name: "建物名",
		},
	}
	out, err := LaboToEntity(in)
	assert.NoError(t, err)

	bytes, err := json.MarshalIndent(out, "", "  ")
	assert.NoError(t, err)
	fmt.Println(string(bytes))
}