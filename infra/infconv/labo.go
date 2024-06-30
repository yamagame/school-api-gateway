package infconv

import (
	"github.com/yamagame/school-api-gateway/entity"
	"github.com/yamagame/school-api-gateway/infra/model"
	"github.com/yamagame/school-api-gateway/pkg/conv"
)

func LaboToInfra(in *conv.Record) (*model.Labo, error) {
	out := &model.Labo{}
	in.ToStruct(".id", ".ID", out, conv.Raw)
	in.ToStruct(".name", ".Name", out, conv.StrPtr)
	in.ToStruct(".group.name", ".Group.Name", out, conv.Raw)
	in.ToStruct(".program.name", ".Program.Name", out, conv.Raw)
	in.ToStruct(".building.name", ".Building.Name", out, conv.Raw)
	return out, nil
}

func LaboToEntity(in *model.Labo) (*conv.Record, error) {
	out := entity.NewLabo(0)
	out.FromStruct(".ID", ".id", in, conv.Raw)
	out.FromStruct(".Name", ".name", in, conv.PtrStr)
	out.FromStruct(".GroupID", ".group.id", in, conv.Raw)
	out.FromStruct(".Group.Name", ".group.name", in, conv.Raw)
	out.FromStruct(".ProgramID", ".program.id", in, conv.Raw)
	out.FromStruct(".Program.Name", ".program.name", in, conv.Raw)
	out.FromStruct(".BuildingID", ".building.id", in, conv.Raw)
	out.FromStruct(".Building.Name", ".building.name", in, conv.Raw)
	return out, nil
}
