package svcconv

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yamagame/school-api-gateway/infra/infconv"
	"github.com/yamagame/school-api-gateway/infra/model"
	"github.com/yamagame/school-api-gateway/pkg/conv"
	"github.com/yamagame/school-api-gateway/pkg/snapshot"
)

func TestLabo(t *testing.T) {
	m := &model.Labo{
		Name:    conv.ToPtr("name"),
		GroupID: conv.ToPtr(int32(1)),
		Group: model.Group{
			Name: "グループ名",
		},
		ProgramID: conv.ToPtr(int32(2)),
		Program: model.Program{
			Name: "プログラム名",
		},
		BuildingID: conv.ToPtr(int32(3)),
		Building: model.Building{
			Name: "建物名",
		},
	}
	e, err := infconv.Labo.ToEntity(m)
	assert.NoError(t, err)
	s, err := Labo.ToProto(e)
	assert.NoError(t, err)

	snapshot.Equal(t, s, "test1.json")
	// snapshot.Save(t, s, "test1.json")
}
