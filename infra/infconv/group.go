package infconv

import (
	"github.com/yamagame/school-api-gateway/infra/model"
	"github.com/yamagame/school-api-gateway/irmodel"
	"github.com/yamagame/school-api-gateway/pkg/irconv"
)

type GroupConv struct{}

func (GroupConv) ToStruct(in *irconv.Record) (*model.Group, error) {
	out := &model.Group{}
	out.Error = in.
		ToStruct(".id", ".ID", out, irconv.Int32Ptr).
		ToStruct(".name", ".Name", out).
		Error
	return out, out.Error
}

func (GroupConv) ToIRModel(in *model.Group) *irconv.Record {
	out := irmodel.NewGroup()
	if in.HasError() {
		out.Error = in.Error
		return out
	}
	return out.
		FromStruct(".ID", ".id", in, irconv.PtrInt32).
		FromStruct(".Name", ".name", in).
		Self()
}

func (GroupConv) NewSlice() *model.Groups {
	return &model.Groups{
		Slice: irconv.NewSlice[model.Group](),
	}
}
