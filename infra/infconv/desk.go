package infconv

import (
	"github.com/yamagame/school-api-gateway/infra/model"
	"github.com/yamagame/school-api-gateway/irmodel"
	"github.com/yamagame/school-api-gateway/pkg/conv"
)

type DeskConv struct{}

func (DeskConv) ToStruct(in *conv.Record) (*model.Desk, error) {
	out := &model.Desk{}
	if err := in.
		ToStruct(".id", ".ID", out).
		ToStruct(".labo_id", ".LaboID", out).Error; err != nil {
		return nil, err
	}
	return out, nil
}

func (DeskConv) ToIRModel(in *model.Desk) (*conv.Record, error) {
	out := irmodel.NewDesk()
	if err := out.
		FromStruct(".ID", ".id", in).
		FromStruct(".LaboID", ".labo_id", in).Error; err != nil {
		return nil, err
	}
	return out, nil
}

func (DeskConv) ToArray(desks []*model.Desk) (*model.Desks, error) {
	return &model.Desks{
		Slice: conv.NewSlice(desks...),
	}, nil
}
