package svcconv

import (
	"github.com/yamagame/school-api-gateway/infra/infconv"
	"github.com/yamagame/school-api-gateway/infra/model"
	"github.com/yamagame/school-api-gateway/irmodel"
	"github.com/yamagame/school-api-gateway/pkg/conv"
	"github.com/yamagame/school-api-gateway/proto/school"
)

type LaboConv struct {
}

func (c LaboConv) ToProto(in *conv.Record) (*school.Labo, error) {
	out := &school.Labo{}
	in.
		ToStruct(".id", ".Id", out).
		ToStruct(".name", ".Name", out).
		IfExist(".group", func(v *conv.Record) {
			out.Group = &school.Group{}
			v.ToStruct(".group.id", ".Group.Id", out).
				ToStruct(".group.name", ".Group.Name", out)
		}).
		IfExist(".program", func(v *conv.Record) {
			out.Program = &school.Program{}
			v.ToStruct(".program.id", ".Program.Id", out).
				ToStruct(".program.name", ".Program.Name", out)
		}).
		IfExist(".building", func(v *conv.Record) {
			out.Building = &school.Building{}
			v.ToStruct(".building.id", ".Building.Id", out).
				ToStruct(".building.name", ".Building.Name", out)
		})
	return out, nil
}

func (c LaboConv) ToIRModel(in *school.Labo) (*conv.Record, error) {
	out := irmodel.NewLabo().
		FromStruct(".Id", ".id", in).
		FromStruct(".Name", ".name", in).
		IfNotNil(".Group", in, func(v *conv.Record) {
			v.FromStruct(".Group.Id", ".group.id", in).
				FromStruct(".Group.Name", ".group.name", in)
		}).
		IfNotNil(".Program", in, func(v *conv.Record) {
			v.FromStruct(".Program.Id", ".program.id", in).
				FromStruct(".Program.Name", ".program.name", in)
		}).
		IfNotNil(".Building", in, func(v *conv.Record) {
			v.FromStruct(".Building.Id", ".building.id", in).
				FromStruct(".Building.Name", ".building.name", in)
		})
	return out, nil
}

func (c LaboConv) ProtoToInfra(labos []*school.Labo) (*model.Labos, error) {
	res := model.NewLabos()
	for _, labo := range labos {
		t, err := c.ToIRModel(labo)
		if err != nil {
			return nil, err
		}
		l, err := infconv.Labo.ToInfra(t)
		if err != nil {
			return nil, err
		}
		res.Append(l)
	}
	return res, nil
}

func (c LaboConv) InfraToProto(labos *model.Labos) ([]*school.Labo, error) {
	res := []*school.Labo{}
	it := labos.NewIterator()
	for it.HasNext() {
		labo := it.Next()
		t, err := infconv.Labo.ToIRModel(labo)
		if err != nil {
			return nil, err
		}
		l, err := c.ToProto(t)
		if err != nil {
			return nil, err
		}
		res = append(res, l)
	}
	return res, nil
}
