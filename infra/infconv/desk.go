package infconv

import (
	"github.com/yamagame/school-api-gateway/infra/model"
	"github.com/yamagame/school-api-gateway/irmodel"
	"github.com/yamagame/school-api-gateway/pkg/irconv"
)

type DeskConv struct{}

func (DeskConv) ToStruct(in *irconv.Record) *model.Desk {
	out := &model.Desk{}
	out.Error = in.
		ToStruct(".id", ".ID", out).
		ToStruct(".labo_id", ".LaboID", out).
		ToStruct(".name", ".Name", out).
		ToStruct(".product_code", ".ProductCode", out).
		Error
	return out
}

func (DeskConv) ToIRModel(in *model.Desk) *irconv.Record {
	out := irmodel.NewDesk()
	if in.HasError() {
		out.Error = in.Error
		return out
	}
	return out.
		FromStruct(".ID", ".id", in).
		FromStruct(".LaboID", ".labo_id", in).
		FromStruct(".Name", ".name", in).
		FromStruct(".ProductCode", ".product_code", in).
		Self()
}

func (DeskConv) NewSlice() *model.Desks {
	return &model.Desks{
		Slice: irconv.NewSlice[model.Desk](),
	}
}
