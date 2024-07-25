package infconv

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yamagame/school-api-gateway/infra/model"
	"github.com/yamagame/school-api-gateway/pkg/iconv"
	"github.com/yamagame/school-api-gateway/pkg/snapshot"
)

func TestLaboConv(t *testing.T) {
	in := &model.Labo{
		ID:      10,
		Name:    iconv.ToPtr("名前"),
		GroupID: iconv.ToPtr(int32(11)),
		Group: model.Group{
			Name: "グループ名",
		},
		ProgramID: iconv.ToPtr(int32(12)),
		Program: model.Program{
			Name: "プログラム名",
		},
		BuildingID: iconv.ToPtr(int32(12)),
		Building: model.Building{
			Name: "建物名",
		},
		Desks: []*model.Desk{
			{ID: 1, LaboID: 10},
			{ID: 2, LaboID: 11},
			{ID: 3, LaboID: 13},
		},
	}
	imodel := Labo.ToIModel(in)
	assert.NoError(t, imodel.Error)

	out := Labo.ToStruct(imodel)
	assert.NoError(t, out.Error)

	snapshot.Update(t)
	snapshot.Match(t, out, "test1.json")
}
