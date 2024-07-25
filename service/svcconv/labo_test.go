package svcconv

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yamagame/school-api-gateway/infra/infconv"
	"github.com/yamagame/school-api-gateway/infra/model"
	"github.com/yamagame/school-api-gateway/pkg/iconv"
	"github.com/yamagame/school-api-gateway/pkg/snapshot"
)

func TestLabo(t *testing.T) {
	m := &model.Labo{
		Name:    iconv.ToPtr("name"),
		GroupID: iconv.ToPtr(int32(1)),
		Group: model.Group{
			Name: "グループ名",
		},
		ProgramID: iconv.ToPtr(int32(2)),
		Program: model.Program{
			Name: "プログラム名",
		},
		BuildingID: iconv.ToPtr(int32(3)),
		Building: model.Building{
			Name: "建物名",
		},
	}

	// dao -> entity -> proto 変換
	e := infconv.Labo.ToIModel(m)
	assert.NoError(t, e.Error)
	s := laboConv.ToStruct(e)
	// assert.NoError(t, s.Error)

	snapshot.Match(t, s, "test1.json")
}
