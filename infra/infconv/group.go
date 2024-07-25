package infconv

import (
	"github.com/yamagame/school-api-gateway/imodel"
	"github.com/yamagame/school-api-gateway/infra/model"
	"github.com/yamagame/school-api-gateway/pkg/iconv"
)

type GroupConv struct{}

func (GroupConv) ToStruct(in *iconv.Record) (*model.Group, error) {
	out := &model.Group{}
	out.Error = in.
		ToStruct(".id", ".ID", out, iconv.Int32Ptr).
		ToStruct(".name", ".Name", out).
		Error
	return out, out.Error
}

func (GroupConv) ToIModel(in *model.Group) *iconv.Record {
	out := imodel.NewGroup()
	if in.HasError() {
		out.Error = in.Error
		return out
	}
	return out.
		FromStruct(".ID", ".id", in, iconv.PtrInt32).
		FromStruct(".Name", ".name", in).
		Self()
}

func (GroupConv) NewSlice() *model.Groups {
	return &model.Groups{
		Slice: iconv.NewSlice[model.Group](),
	}
}
