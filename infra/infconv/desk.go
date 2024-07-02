package infconv

import (
	"github.com/yamagame/school-api-gateway/entity"
	"github.com/yamagame/school-api-gateway/infra/model"
	"github.com/yamagame/school-api-gateway/pkg/conv"
)

type DeskConv struct{}

func (DeskConv) ToInfra(in *conv.Record) (*model.Desk, error) {
	out := &model.Desk{}
	if err := in.
		ToStruct(".id", ".ID", out, conv.Raw).
		ToStruct(".labo_id", ".LaboID", out, conv.Raw).Error; err != nil {
		return nil, err
	}
	return out, nil
}

func (DeskConv) ToEntity(in *model.Desk) (*conv.Record, error) {
	out := entity.NewDesk()
	if err := out.
		FromStruct(".ID", ".id", in, conv.Raw).
		FromStruct(".LaboID", ".labo_id", in, conv.Raw).Error; err != nil {
		return nil, err
	}
	return out, nil
}
