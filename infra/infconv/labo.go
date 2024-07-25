package infconv

import (
	"github.com/yamagame/school-api-gateway/imodel"
	"github.com/yamagame/school-api-gateway/infra/model"
	"github.com/yamagame/school-api-gateway/pkg/iconv"
)

type LaboConv struct{}

func (LaboConv) ToStruct(in *iconv.Record) *model.Labo {
	out := &model.Labo{}
	out.Error = in.
		ToStruct(".id", ".ID", out, iconv.Keep).
		ToStruct(".name", ".Name", out, iconv.StrPtr).
		ToStruct(".url", ".URL", out, iconv.StrPtr).
		IfExist(".group", func(v *iconv.Record) {
			v.ToStruct(".group.id", ".GroupID", out, iconv.Int32Ptr).
				ToStruct(".group.name", ".Group.Name", out, iconv.Keep)
		}).
		IfExist(".program", func(v *iconv.Record) {
			v.ToStruct(".program.id", ".ProgramID", out, iconv.Int32Ptr).
				ToStruct(".program.name", ".Program.Name", out, iconv.Keep)
		}).
		IfExist(".building", func(v *iconv.Record) {
			v.ToStruct(".building.id", ".BuildingID", out, iconv.Int32Ptr).
				ToStruct(".building.name", ".Building.Name", out, iconv.Keep)
		}).
		IfExist(".property", func(v *iconv.Record) {
			out.Property = Property.ToStruct(v.GetHasOne("property"))
		}).
		IfExist(".desk", func(v *iconv.Record) {
			if desks, err := Desks.ToStruct(v.GetHasMany("desk").Records()).ShallowCopy(); err == nil {
				out.Desks = desks
			}
		}).
		Error
	return out
}

func (LaboConv) ToIModel(in *model.Labo) *iconv.Record {
	out := imodel.NewLabo()
	if in.HasError() {
		out.Error = in.Error
		return out
	}
	return out.
		FromStruct(".ID", ".id", in, iconv.Keep).
		FromStruct(".Name", ".name", in, iconv.PtrStr).
		FromStruct(".URL", ".url", in, iconv.PtrStr).
		IfNotNil(".GroupID", in, func(v *iconv.Record) {
			v.FromStruct(".GroupID", ".group.id", in, iconv.PtrInt32).
				FromStruct(".Group.Name", ".group.name", in, iconv.Keep)
		}).
		IfNotNil(".ProgramID", in, func(v *iconv.Record) {
			v.FromStruct(".ProgramID", ".program.id", in, iconv.PtrInt32).
				FromStruct(".Program.Name", ".program.name", in, iconv.Keep)
		}).
		IfNotNil(".BuildingID", in, func(v *iconv.Record) {
			v.FromStruct(".BuildingID", ".building.id", in, iconv.PtrInt32).
				FromStruct(".Building.Name", ".building.name", in, iconv.Keep)
		}).
		IfNotNil(".Property", in, func(v *iconv.Record) {
			v.SetHasOne("property", Property.ToIModel(in.Property))
		}).
		IfNotNil(".Desks", in, func(v *iconv.Record) {
			v.SetHasMany("desk", out.GetHasMany("desk").Append(Desks.ToIModel(in.Desks).Records()...))
		}).
		Self()
}

func (LaboConv) NewSlice() *model.Labos {
	return &model.Labos{
		Slice: iconv.NewSlice[model.Labo](),
	}
}
