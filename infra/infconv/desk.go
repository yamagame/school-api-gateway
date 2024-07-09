package infconv

import (
	"github.com/yamagame/school-api-gateway/entity"
	"github.com/yamagame/school-api-gateway/infra/model"
	"github.com/yamagame/school-api-gateway/pkg/conv"
)

type DeskType struct{}

var Desk = DeskType{}

var Desks = Convs[model.Desk, DeskType]{
	conv: Desk,
}

func (DeskType) ToInfra(in *conv.Record) (*model.Desk, error) {
	out := &model.Desk{}
	if err := in.
		ToStruct(".id", ".ID", out).
		ToStruct(".labo_id", ".LaboID", out).Error; err != nil {
		return nil, err
	}
	return out, nil
}

func (DeskType) ToEntity(in *model.Desk) (*conv.Record, error) {
	out := entity.NewDesk()
	if err := out.
		FromStruct(".ID", ".id", in).
		FromStruct(".LaboID", ".labo_id", in).Error; err != nil {
		return nil, err
	}
	return out, nil
}
