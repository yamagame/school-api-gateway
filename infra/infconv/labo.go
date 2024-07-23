package infconv

import (
	"github.com/yamagame/school-api-gateway/infra/model"
	"github.com/yamagame/school-api-gateway/irmodel"
	"github.com/yamagame/school-api-gateway/pkg/irconv"
)

type LaboConv struct{}

func (LaboConv) ToStruct(in *irconv.Record) (*model.Labo, error) {
	out := &model.Labo{}
	out.Error = in.
		ToStruct(".id", ".ID", out, irconv.Keep).
		ToStruct(".name", ".Name", out, irconv.StrPtr).
		ToStruct(".url", ".URL", out, irconv.StrPtr).
		IfExist(".group", func(v *irconv.Record) {
			v.ToStruct(".group.id", ".GroupID", out, irconv.Int32Ptr).
				ToStruct(".group.name", ".Group.Name", out, irconv.Keep)
		}).
		IfExist(".program", func(v *irconv.Record) {
			v.ToStruct(".program.id", ".ProgramID", out, irconv.Int32Ptr).
				ToStruct(".program.name", ".Program.Name", out, irconv.Keep)
		}).
		IfExist(".building", func(v *irconv.Record) {
			v.ToStruct(".building.id", ".BuildingID", out, irconv.Int32Ptr).
				ToStruct(".building.name", ".Building.Name", out, irconv.Keep)
		}).
		IfExist(".property", func(v *irconv.Record) {
			val, err := Property.ToStruct(v.GetHasOne("property"))
			if err != nil {
				out.Error = err
				return
			}
			out.Property = val
		}).
		IfExist(".desk", func(v *irconv.Record) {
			if desks, err := Desks.ToStruct(v.GetHasMany("desk").Records()).ShallowCopy(); err == nil {
				out.Desks = desks
			}
		}).
		Error
	return out, out.Error
}

func (LaboConv) ToIRModel(in *model.Labo) *irconv.Record {
	out := irmodel.NewLabo()
	if in.HasError() {
		out.Error = in.Error
		return out
	}
	return out.
		FromStruct(".ID", ".id", in, irconv.Keep).
		FromStruct(".Name", ".name", in, irconv.PtrStr).
		FromStruct(".URL", ".url", in, irconv.PtrStr).
		IfNotNil(".GroupID", in, func(v *irconv.Record) {
			v.FromStruct(".GroupID", ".group.id", in, irconv.PtrInt32).
				FromStruct(".Group.Name", ".group.name", in, irconv.Keep)
		}).
		IfNotNil(".ProgramID", in, func(v *irconv.Record) {
			v.FromStruct(".ProgramID", ".program.id", in, irconv.PtrInt32).
				FromStruct(".Program.Name", ".program.name", in, irconv.Keep)
		}).
		IfNotNil(".BuildingID", in, func(v *irconv.Record) {
			v.FromStruct(".BuildingID", ".building.id", in, irconv.PtrInt32).
				FromStruct(".Building.Name", ".building.name", in, irconv.Keep)
		}).
		IfNotNil(".Property", in, func(v *irconv.Record) {
			v.SetHasOne("property", Property.ToIRModel(in.Property))
		}).
		IfNotNil(".Desks", in, func(v *irconv.Record) {
			v.SetHasMany("desk", out.GetHasMany("desk").Append(Desks.ToIRModel(in.Desks).Records()...))
		}).
		Self()
}

func (LaboConv) NewSlice() *model.Labos {
	return &model.Labos{
		Slice: irconv.NewSlice[model.Labo](),
	}
}
