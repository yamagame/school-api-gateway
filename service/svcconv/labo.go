package svcconv

import (
	"github.com/yamagame/school-api-gateway/entity"
	"github.com/yamagame/school-api-gateway/pkg/conv"
	"github.com/yamagame/school-api-gateway/proto/school"
)

type LaboConv struct {
}

func (LaboConv) ToProto(in *conv.Record) (*school.Labo, error) {
	out := &school.Labo{
		Group:    &school.Group{},
		Program:  &school.Program{},
		Building: &school.Building{},
	}
	in.
		ToStruct(".id", ".Id", out).
		ToStruct(".name", ".Name", out).
		ToStruct(".group.id", ".Group.Id", out).
		ToStruct(".group.name", ".Group.Name", out).
		ToStruct(".program.id", ".Program.Id", out).
		ToStruct(".program.name", ".Program.Name", out).
		ToStruct(".building.id", ".Building.Id", out).
		ToStruct(".building.name", ".Building.Name", out)
	return out, nil
}

func (LaboConv) ToEntity(in *school.Labo) (*conv.Record, error) {
	out := entity.NewLabo().
		FromStruct(".Id", ".id", in).
		FromStruct(".Name", ".name", in).
		FromStruct(".Group.Id", ".group.id", in).
		FromStruct(".Group.Name", ".group.name", in).
		FromStruct(".Program.Id", ".program.id", in).
		FromStruct(".Program.Name", ".program.name", in).
		FromStruct(".Building.Id", ".building.id", in).
		FromStruct(".Building.Name", ".building.name", in)
	return out, nil
}
