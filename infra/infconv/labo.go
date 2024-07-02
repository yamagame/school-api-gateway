package infconv

import (
	"github.com/yamagame/school-api-gateway/entity"
	"github.com/yamagame/school-api-gateway/infra/model"
	"github.com/yamagame/school-api-gateway/pkg/conv"
)

type LaboConv struct{}

func (LaboConv) ToInfra(in *conv.Record) (*model.Labo, error) {
	out := &model.Labo{}
	in.ToStruct(".id", ".ID", out, conv.Raw)
	in.ToStruct(".name", ".Name", out, conv.StrPtr)
	in.ToStruct(".group.id", ".GroupID", out, conv.Int32Ptr)
	in.ToStruct(".group.name", ".Group.Name", out, conv.Raw)
	in.ToStruct(".program.id", ".ProgramID", out, conv.Int32Ptr)
	in.ToStruct(".program.name", ".Program.Name", out, conv.Raw)
	in.ToStruct(".building.id", ".BuildingID", out, conv.Int32Ptr)
	in.ToStruct(".building.name", ".Building.Name", out, conv.Raw)
	if values, err := in.GetHasManyRecords("desk"); err == nil {
		if desks, err := Desks.ToInfra(values); err == nil {
			out.Desks = desks
		}
	}
	return out, nil
}

func (LaboConv) ToEntity(in *model.Labo) (*conv.Record, error) {
	out := entity.NewLabo()
	out.FromStruct(".ID", ".id", in, conv.Raw)
	out.FromStruct(".Name", ".name", in, conv.PtrStr)
	out.FromStruct(".GroupID", ".group.id", in, conv.PtrInt32)
	out.FromStruct(".Group.Name", ".group.name", in, conv.Raw)
	out.FromStruct(".ProgramID", ".program.id", in, conv.PtrInt32)
	out.FromStruct(".Program.Name", ".program.name", in, conv.Raw)
	out.FromStruct(".BuildingID", ".building.id", in, conv.PtrInt32)
	out.FromStruct(".Building.Name", ".building.name", in, conv.Raw)
	if len(in.Desks) > 0 {
		if values, err := Desks.ToEntity(in.Desks); err == nil {
			out.SetHasManyRecords("desk", values...)
		}
	}
	return out, nil
}
