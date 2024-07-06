package infconv

import (
	"github.com/yamagame/school-api-gateway/entity"
	"github.com/yamagame/school-api-gateway/infra/model"
	"github.com/yamagame/school-api-gateway/pkg/conv"
)

type LaboConv struct{}

var Labo = LaboConv{}

var Labos = Convs[model.Labo, LaboConv]{
	conv: Labo,
}

func (LaboConv) ToInfra(in *conv.Record) (*model.Labo, error) {
	out := &model.Labo{}
	if err := in.
		ToStruct(".id", ".ID", out, conv.Raw).
		ToStruct(".name", ".Name", out, conv.StrPtr).
		ToStruct(".url", ".URL", out, conv.StrPtr).
		ToStruct(".group.id", ".GroupID", out, conv.Int32Ptr).
		ToStruct(".group.name", ".Group.Name", out, conv.Raw).
		ToStruct(".program.id", ".ProgramID", out, conv.Int32Ptr).
		ToStruct(".program.name", ".Program.Name", out, conv.Raw).
		ToStruct(".building.id", ".BuildingID", out, conv.Int32Ptr).
		ToStruct(".building.name", ".Building.Name", out, conv.Raw).Error; err != nil {
		return nil, err
	}
	if values, err := in.GetHasManyRecords("desk"); err == nil {
		if desks, err := Desks.ToInfra(values); err == nil {
			out.Desks = desks
		}
	}
	return out, nil
}

func (LaboConv) ToEntity(in *model.Labo) (*conv.Record, error) {
	out := entity.NewLabo()
	if err := out.
		FromStruct(".ID", ".id", in, conv.Raw).
		FromStruct(".Name", ".name", in, conv.PtrStr).
		FromStruct(".URL", ".url", in, conv.PtrStr).
		FromStruct(".GroupID", ".group.id", in, conv.PtrInt32).
		FromStruct(".Group.Name", ".group.name", in, conv.Raw).
		FromStruct(".ProgramID", ".program.id", in, conv.PtrInt32).
		FromStruct(".Program.Name", ".program.name", in, conv.Raw).
		FromStruct(".BuildingID", ".building.id", in, conv.PtrInt32).
		FromStruct(".Building.Name", ".building.name", in, conv.Raw).Error; err != nil {
		return nil, err
	}
	if len(in.Desks) > 0 {
		if values, err := Desks.ToEntity(in.Desks); err == nil {
			out.SetHasManyRecords("desk", values)
		}
	}
	return out, nil
}
