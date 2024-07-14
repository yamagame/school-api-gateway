package infconv

import (
	"github.com/yamagame/school-api-gateway/infra/model"
	"github.com/yamagame/school-api-gateway/irmodel"
	"github.com/yamagame/school-api-gateway/pkg/conv"
)

type LaboConv struct{}

func (LaboConv) ToStruct(in *conv.Record) (*model.Labo, error) {
	out := &model.Labo{}
	if err := in.
		ToStruct(".id", ".ID", out).
		ToStruct(".name", ".Name", out, conv.StrPtr).
		ToStruct(".url", ".URL", out, conv.StrPtr).
		IfExist(".group", func(v *conv.Record) {
			v.ToStruct(".group.id", ".GroupID", out, conv.Int32Ptr).
				ToStruct(".group.name", ".Group.Name", out)
		}).
		IfExist(".program", func(v *conv.Record) {
			v.ToStruct(".program.id", ".ProgramID", out, conv.Int32Ptr).
				ToStruct(".program.name", ".Program.Name", out)
		}).
		IfExist(".building", func(v *conv.Record) {
			v.ToStruct(".building.id", ".BuildingID", out, conv.Int32Ptr).
				ToStruct(".building.name", ".Building.Name", out)
		}).
		IfExist(".desk", func(v *conv.Record) {
			if desks, err := Desks.ToStruct(v.GetHasManyRecords("desk")); err == nil {
				out.Desks = desks.ShallowCopy()
			}
		}).Error; err != nil {
		return nil, err
	}
	return out, nil
}

func (LaboConv) ToIRModel(in *model.Labo) (*conv.Record, error) {
	out := irmodel.NewLabo()
	if err := out.
		FromStruct(".ID", ".id", in).
		FromStruct(".Name", ".name", in, conv.PtrStr).
		FromStruct(".URL", ".url", in, conv.PtrStr).
		IfNotNil(".GroupID", in, func(v *conv.Record) {
			v.FromStruct(".GroupID", ".group.id", in, conv.PtrInt32).
				FromStruct(".Group.Name", ".group.name", in)
		}).
		IfNotNil(".ProgramID", in, func(v *conv.Record) {
			v.FromStruct(".ProgramID", ".program.id", in, conv.PtrInt32).
				FromStruct(".Program.Name", ".program.name", in)
		}).
		IfNotNil(".BuildingID", in, func(v *conv.Record) {
			v.FromStruct(".BuildingID", ".building.id", in, conv.PtrInt32).
				FromStruct(".Building.Name", ".building.name", in)
		}).
		IfNotNil(".Desks", in, func(v *conv.Record) {
			if values, err := Desks.ToIRModel(in.Desks); err == nil {
				v.SetHasManyRecords("desk", values)
			}
		}).Error; err != nil {
		return nil, err
	}
	return out, nil
}

func (LaboConv) ToArray(labos []*model.Labo) (*model.Labos, error) {
	return &model.Labos{
		Slice: conv.NewSlice(labos...),
	}, nil
}
