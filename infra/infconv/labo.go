package infconv

import (
	"github.com/yamagame/school-api-gateway/infra/model"
	"github.com/yamagame/school-api-gateway/irmodel"
	irmodel1 "github.com/yamagame/school-api-gateway/pkg/irconv"
)

type LaboConv struct{}

func (LaboConv) ToStruct(in *irmodel1.Record) (*model.Labo, error) {
	out := &model.Labo{}
	if err := in.
		ToStruct(".id", ".ID", out).
		ToStruct(".name", ".Name", out, irmodel1.StrPtr).
		ToStruct(".url", ".URL", out, irmodel1.StrPtr).
		IfExist(".group", func(v *irmodel1.Record) {
			v.ToStruct(".group.id", ".GroupID", out, irmodel1.Int32Ptr).
				ToStruct(".group.name", ".Group.Name", out)
		}).
		IfExist(".program", func(v *irmodel1.Record) {
			v.ToStruct(".program.id", ".ProgramID", out, irmodel1.Int32Ptr).
				ToStruct(".program.name", ".Program.Name", out)
		}).
		IfExist(".building", func(v *irmodel1.Record) {
			v.ToStruct(".building.id", ".BuildingID", out, irmodel1.Int32Ptr).
				ToStruct(".building.name", ".Building.Name", out)
		}).
		IfExist(".desk", func(v *irmodel1.Record) {
			if desks, err := Desks.ToStruct(v.GetHasManyRecords("desk")).ShallowCopy(); err == nil {
				out.Desks = desks
			}
		}).Error; err != nil {
		return nil, err
	}
	return out, nil
}

func (LaboConv) ToIRModel(in *model.Labo) (*irmodel1.Record, error) {
	out := irmodel.NewLabo()
	if err := out.
		FromStruct(".ID", ".id", in).
		FromStruct(".Name", ".name", in, irmodel1.PtrStr).
		FromStruct(".URL", ".url", in, irmodel1.PtrStr).
		IfNotNil(".GroupID", in, func(v *irmodel1.Record) {
			v.FromStruct(".GroupID", ".group.id", in, irmodel1.PtrInt32).
				FromStruct(".Group.Name", ".group.name", in)
		}).
		IfNotNil(".ProgramID", in, func(v *irmodel1.Record) {
			v.FromStruct(".ProgramID", ".program.id", in, irmodel1.PtrInt32).
				FromStruct(".Program.Name", ".program.name", in)
		}).
		IfNotNil(".BuildingID", in, func(v *irmodel1.Record) {
			v.FromStruct(".BuildingID", ".building.id", in, irmodel1.PtrInt32).
				FromStruct(".Building.Name", ".building.name", in)
		}).
		IfNotNil(".Desks", in, func(v *irmodel1.Record) {
			v.SetHasManyRecords("desk", Desks.ToIRModel(in.Desks))
		}).Error; err != nil {
		return nil, err
	}
	return out, nil
}

func (LaboConv) NewSlice() *model.Labos {
	return &model.Labos{
		Slice: irmodel1.NewSlice[model.Labo](),
	}
}
