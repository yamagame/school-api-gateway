package infconv

import (
	"github.com/yamagame/school-api-gateway/infra/model"
	"github.com/yamagame/school-api-gateway/irmodel"
	irmodel1 "github.com/yamagame/school-api-gateway/pkg/irconv"
)

type DeskConv struct{}

func (DeskConv) ToStruct(in *irmodel1.Record) (*model.Desk, error) {
	out := &model.Desk{}
	if err := in.
		ToStruct(".id", ".ID", out).
		ToStruct(".labo_id", ".LaboID", out).
		ToStruct(".name", ".Name", out).
		Error; err != nil {
		return nil, err
	}
	return out, nil
}

func (DeskConv) ToIRModel(in *model.Desk) *irmodel1.Record {
	out := irmodel.NewDesk()
	return out.
		FromStruct(".ID", ".id", in).
		FromStruct(".LaboID", ".labo_id", in).
		FromStruct(".Name", ".name", in)
}

func (DeskConv) NewSlice() *model.Desks {
	return &model.Desks{
		Slice: irmodel1.NewSlice[model.Desk](),
	}
}
