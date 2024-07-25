package infconv

import (
	"github.com/yamagame/school-api-gateway/imodel"
	"github.com/yamagame/school-api-gateway/infra/model"
	"github.com/yamagame/school-api-gateway/pkg/iconv"
)

type PropertyConv struct{}

func (PropertyConv) ToStruct(in *iconv.Record) *model.Property {
	out := &model.Property{}
	out.Error = in.
		ToStruct(".id", ".ID", out).
		ToStruct(".labo_id", ".LaboID", out).
		ToStruct(".name", ".Name", out).
		Error
	return out
}

func (PropertyConv) ToIModel(in *model.Property) *iconv.Record {
	out := imodel.NewProperty()
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

func (PropertyConv) NewSlice(models ...*model.Property) *model.Properties {
	r := &model.Properties{
		SliceWrapper: iconv.NewSlice[model.Property](),
	}
	r.Append(models...)
	return r
}
