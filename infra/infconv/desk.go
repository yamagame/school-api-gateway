package infconv

import (
	"github.com/yamagame/school-api-gateway/imodel"
	"github.com/yamagame/school-api-gateway/infra/model"
	"github.com/yamagame/school-api-gateway/pkg/iconv"
)

type DeskConv struct{}

func (DeskConv) ToStruct(in *iconv.Record) *model.Desk {
	out := &model.Desk{}
	out.Error = in.
		ToStruct(".id", ".ID", out).
		ToStruct(".labo_id", ".LaboID", out).
		ToStruct(".name", ".Name", out).
		ToStruct(".product_code", ".ProductCode", out).
		Error
	return out
}

func (DeskConv) ToIModel(in *model.Desk) *iconv.Record {
	out := imodel.NewDesk()
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

func (DeskConv) NewSlice(models ...*model.Desk) *model.Desks {
	r := &model.Desks{
		SliceWrapper: iconv.NewSlice[model.Desk](),
	}
	r.Append(models...)
	return r
}

func (v DeskConv) ToModels(desks []*iconv.Record) []*model.Desk {
	ret := []*model.Desk{}
	for _, desk := range desks {
		ret = append(ret, Desk.ToStruct(desk))
	}
	return ret
}

func (DeskConv) ToRecords(desks []*model.Desk) []*iconv.Record {
	return Desks.ToIModel(Desk.NewSlice(desks...)).Slice()
}
