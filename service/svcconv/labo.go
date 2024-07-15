package svcconv

import (
	"github.com/yamagame/school-api-gateway/infra/infconv"
	"github.com/yamagame/school-api-gateway/infra/model"
	"github.com/yamagame/school-api-gateway/irmodel"
	irmodel1 "github.com/yamagame/school-api-gateway/pkg/irconv"
	"github.com/yamagame/school-api-gateway/proto/school"
)

type LaboConv struct {
}

func (c LaboConv) ToStruct(in *irmodel1.Record) (*school.Labo, error) {
	out := &school.Labo{}
	in.
		ToStruct(".id", ".Id", out).
		ToStruct(".name", ".Name", out).
		IfExist(".group", func(v *irmodel1.Record) {
			out.Group = &school.Group{}
			v.ToStruct(".group.id", ".Group.Id", out).
				ToStruct(".group.name", ".Group.Name", out)
		}).
		IfExist(".program", func(v *irmodel1.Record) {
			out.Program = &school.Program{}
			v.ToStruct(".program.id", ".Program.Id", out).
				ToStruct(".program.name", ".Program.Name", out)
		}).
		IfExist(".building", func(v *irmodel1.Record) {
			out.Building = &school.Building{}
			v.ToStruct(".building.id", ".Building.Id", out).
				ToStruct(".building.name", ".Building.Name", out)
		})
	return out, nil
}

func (c LaboConv) ToIRModel(in *school.Labo) *irmodel1.Record {
	return irmodel.NewLabo().
		FromStruct(".Id", ".id", in).
		FromStruct(".Name", ".name", in).
		IfNotNil(".Group", in, func(v *irmodel1.Record) {
			v.FromStruct(".Group.Id", ".group.id", in).
				FromStruct(".Group.Name", ".group.name", in)
		}).
		IfNotNil(".Program", in, func(v *irmodel1.Record) {
			v.FromStruct(".Program.Id", ".program.id", in).
				FromStruct(".Program.Name", ".program.name", in)
		}).
		IfNotNil(".Building", in, func(v *irmodel1.Record) {
			v.FromStruct(".Building.Id", ".building.id", in).
				FromStruct(".Building.Name", ".building.name", in)
		})
}

func (LaboConv) NewSlice() *school.Labos {
	return &school.Labos{
		Slice: irmodel1.NewSlice[school.Labo](),
	}
}

func (c LaboConv) ProtoToInfra(labos *school.Labos) *model.Labos {
	return infconv.Labos.ToStruct(Labos.ToIRModel(labos.ShallowCopy()))
}

func (c LaboConv) InfraToProto(labos *model.Labos) *school.Labos {
	return Labos.ToStruct(infconv.Labos.ToIRModel(labos.ShallowCopy()))
}
