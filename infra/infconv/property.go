package infconv

import (
	"github.com/yamagame/school-api-gateway/infra/model"
	"github.com/yamagame/school-api-gateway/irmodel"
	"github.com/yamagame/school-api-gateway/pkg/irconv"
)

type PropertyConv struct{}

func (PropertyConv) ToStruct(in *irconv.Record) (*model.Property, error) {
	out := &model.Property{}
	out.Error = in.
		ToStruct(".id", ".ID", out).
		ToStruct(".labo_id", ".LaboID", out).
		ToStruct(".name", ".Name", out).
		Error
	return out, out.Error
}

func (PropertyConv) ToIRModel(in *model.Property) *irconv.Record {
	out := irmodel.NewProperty()
	if in.HasError() {
		out.Error = in.Error
		return out
	}
	return out.
		FromStruct(".ID", ".id", in).
		FromStruct(".LaboID", ".labo_id", in).
		FromStruct(".Name", ".name", in).
		Self()
}

func (PropertyConv) NewSlice() *model.Properties {
	return &model.Properties{
		Slice: irconv.NewSlice[model.Property](),
	}
}
