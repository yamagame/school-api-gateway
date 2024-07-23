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
		ToStruct(".id", ".Id", out, irconv.Keep).
		ToStruct(".name", ".Name", out, irconv.Keep).
		IfExist(".group", func(v *irconv.Record) {
			out.Group = &school.Group{}
			v.ToStruct(".group.id", ".Group.Id", out, irconv.Keep).
				ToStruct(".group.name", ".Group.Name", out, irconv.Keep)
		}).
		IfExist(".program", func(v *irconv.Record) {
			out.Program = &school.Program{}
			v.ToStruct(".program.id", ".Program.Id", out, irconv.Keep).
				ToStruct(".program.name", ".Program.Name", out, irconv.Keep)
		}).
		IfExist(".building", func(v *irconv.Record) {
			out.Building = &school.Building{}
			v.ToStruct(".building.id", ".Building.Id", out, irconv.Keep).
				ToStruct(".building.name", ".Building.Name", out, irconv.Keep)
		}).
		Self()
	return out, nil
}

func (c LaboConv) ToIRModel(in *school.Labo) *irconv.Record {
	return irmodel.NewLabo().
		FromStruct(".Id", ".id", in, irconv.Keep).
		FromStruct(".Name", ".name", in, irconv.Keep).
		IfNotNil(".Group", in, func(v *irconv.Record) {
			v.FromStruct(".Group.Id", ".group.id", in, irconv.Keep).
				FromStruct(".Group.Name", ".group.name", in, irconv.Keep)
		}).
		IfNotNil(".Program", in, func(v *irconv.Record) {
			v.FromStruct(".Program.Id", ".program.id", in).
				FromStruct(".Program.Name", ".program.name", in, irconv.Keep)
		}).
		IfNotNil(".Building", in, func(v *irconv.Record) {
			v.FromStruct(".Building.Id", ".building.id", in, irconv.Keep).
				FromStruct(".Building.Name", ".building.name", in, irconv.Keep)
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
