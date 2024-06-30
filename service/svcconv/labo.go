package svcconv

import (
	"github.com/yamagame/school-api-gateway/entity"
	"github.com/yamagame/school-api-gateway/pkg/conv"
	"github.com/yamagame/school-api-gateway/proto/school"
)

func LaboToProto(in *conv.Record) (*school.Labo, error) {
	out := &school.Labo{
		Group:    &school.Group{},
		Program:  &school.Program{},
		Building: &school.Building{},
	}
	in.ToStruct(".id", ".Id", out, conv.Raw)
	in.ToStruct(".name", ".Name", out, conv.Raw)
	in.ToStruct(".group.id", ".Group.Id", out, conv.Raw)
	in.ToStruct(".group.name", ".Group.Name", out, conv.Raw)
	in.ToStruct(".program.id", ".Program.Id", out, conv.Raw)
	in.ToStruct(".program.name", ".Program.Name", out, conv.Raw)
	in.ToStruct(".building.id", ".Building.Id", out, conv.Raw)
	in.ToStruct(".building.name", ".Building.Name", out, conv.Raw)
	return out, nil
}

func LaboToEntity(in *school.Labo) (*conv.Record, error) {
	out := entity.NewLabo()
	out.FromStruct(".Id", ".id", in, conv.Raw)
	out.FromStruct(".Name", ".name", in, conv.Raw)
	out.FromStruct(".Group.Id", ".group.id", in, conv.Raw)
	out.FromStruct(".Group.Name", ".group.name", in, conv.Raw)
	out.FromStruct(".Program.Id", ".program.id", in, conv.Raw)
	out.FromStruct(".Program.Name", ".program.name", in, conv.Raw)
	out.FromStruct(".Building.Id", ".building.id", in, conv.Raw)
	out.FromStruct(".Building.Name", ".building.name", in, conv.Raw)
	return out, nil
}
