package svcconv

import (
	"github.com/yamagame/school-api-gateway/infra/infconv"
	"github.com/yamagame/school-api-gateway/infra/model"
	"github.com/yamagame/school-api-gateway/irmodel"
	"github.com/yamagame/school-api-gateway/pkg/irconv"
	"github.com/yamagame/school-api-gateway/proto/school"
)

type LaboConv struct {
}

func (c LaboConv) ToStruct(in *irconv.Record) (*school.Labo, error) {
	out := &school.Labo{}
	in.
		ToStruct(".id", ".Id", out).
		ToStruct(".name", ".Name", out).
		IfExist(".group", func(v *irconv.Record) {
			out.Group = &school.Group{}
			v.ToStruct(".group.id", ".Group.Id", out).
				ToStruct(".group.name", ".Group.Name", out)
		}).
		IfExist(".program", func(v *irconv.Record) {
			out.Program = &school.Program{}
			v.ToStruct(".program.id", ".Program.Id", out).
				ToStruct(".program.name", ".Program.Name", out)
		}).
		IfExist(".building", func(v *irconv.Record) {
			out.Building = &school.Building{}
			v.ToStruct(".building.id", ".Building.Id", out).
				ToStruct(".building.name", ".Building.Name", out)
		}).
		Self()
	return out, nil
}

func (c LaboConv) ToIRModel(in *school.Labo) *irconv.Record {
	return irmodel.NewLabo().
		FromStruct(".Id", ".id", in).
		FromStruct(".Name", ".name", in).
		IfNotNil(".Group", in, func(v *irconv.Record) {
			v.FromStruct(".Group.Id", ".group.id", in).
				FromStruct(".Group.Name", ".group.name", in)
		}).
		IfNotNil(".Program", in, func(v *irconv.Record) {
			v.FromStruct(".Program.Id", ".program.id", in).
				FromStruct(".Program.Name", ".program.name", in)
		}).
		IfNotNil(".Building", in, func(v *irconv.Record) {
			v.FromStruct(".Building.Id", ".building.id", in).
				FromStruct(".Building.Name", ".building.name", in)
		}).
		Self()
}

func (LaboConv) NewSlice() *school.Labos {
	return &school.Labos{
		Slice: irconv.NewSlice[school.Labo](),
	}
}

func (c LaboConv) ProtoToInfra(labos *school.Labos) *model.Labos {
	return infconv.Labos.ToStruct(Labos.ToIRModel(labos.ShallowCopy()))
}

func (c LaboConv) InfraToProto(labos *model.Labos) *school.Labos {
	return Labos.ToStruct(infconv.Labos.ToIRModel(labos.ShallowCopy()))
}
