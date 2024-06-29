package infconv

import (
	"github.com/yamagame/school-api-gateway/entity"
	"github.com/yamagame/school-api-gateway/infra/model"
	"github.com/yamagame/school-api-gateway/pkg/conv"
)

func LaboToInfra(in *conv.Record) (*model.Labo, error) {
	out := &model.Labo{}
	in.ToStruct(".id", ".ID", out, conv.Row)
	in.ToStruct(".name", ".Name", out, conv.StrPtr)
	in.ToStruct(".group.name", ".Group.Name", out, conv.Row)
	in.ToStruct(".program.name", ".Program.Name", out, conv.Row)
	in.ToStruct(".building.name", ".Building.Name", out, conv.Row)
	return out, nil
}

func LaboToEntity(in *model.Labo) (*conv.Record, error) {
	out := entity.NewLabo(0)
	out.FromStruct(".ID", ".id", in, conv.Row)
	out.FromStruct(".Name", ".name", in, conv.PtrStr)
	out.FromStruct(".GroupID", ".group.id", in, conv.Row)
	out.FromStruct(".Group.Name", ".group.name", in, conv.Row)
	out.FromStruct(".ProgramID", ".program.id", in, conv.Row)
	out.FromStruct(".Program.Name", ".program.name", in, conv.Row)
	out.FromStruct(".BuildingID", ".building.id", in, conv.Row)
	out.FromStruct(".Building.Name", ".building.name", in, conv.Row)
	return out, nil
}
