package svcconv

import (
	"fmt"

	"github.com/yamagame/school-api-gateway/imodel"
	"github.com/yamagame/school-api-gateway/pkg/iconv"
	"github.com/yamagame/school-api-gateway/proto/school"
)

type LaboConv struct {
}

func (LaboConv) NewSlice(models ...*school.Labo) *school.Labos {
	ret := &school.Labos{
		SliceWrapper: iconv.NewSlice[school.Labo](),
	}
	ret.Append(models...)
	return ret
}

func (c LaboConv) ToStruct(in *iconv.Record) *school.Labo {
	out := &school.Labo{}
	in.
		ToStruct(".id", ".Id", out, iconv.Keep).
		ToStruct(".name", ".Name", out, iconv.Keep).
		IfExist(".group", func(v *iconv.Record) {
			out.Group = &school.Group{}
			v.ToStruct(".group.id", ".Group.Id", out, iconv.Keep).
				ToStruct(".group.name", ".Group.Name", out, iconv.Keep)
		}).
		IfExist(".program", func(v *iconv.Record) {
			out.Program = &school.Program{}
			v.ToStruct(".program.id", ".Program.Id", out, iconv.Keep).
				ToStruct(".program.name", ".Program.Name", out, iconv.Keep)
		}).
		IfExist(".building", func(v *iconv.Record) {
			out.Building = &school.Building{}
			v.ToStruct(".building.id", ".Building.Id", out, iconv.Keep).
				ToStruct(".building.name", ".Building.Name", out, iconv.Keep)
		}).
		Self()
	out.Error = in.GetError()
	return out
}

func (c LaboConv) ToIModel(in *school.Labo) *iconv.Record {
	if in.Error != "" {
		return &iconv.Record{
			Error: fmt.Errorf(in.Error),
		}
	}
	return imodel.NewLabo().
		FromStruct(".Id", ".id", in, iconv.Keep).
		FromStruct(".Name", ".name", in, iconv.Keep).
		IfNotNil(".Group", in, func(v *iconv.Record) {
			v.FromStruct(".Group.Id", ".group.id", in, iconv.Keep).
				FromStruct(".Group.Name", ".group.name", in, iconv.Keep)
		}).
		IfNotNil(".Program", in, func(v *iconv.Record) {
			v.FromStruct(".Program.Id", ".program.id", in).
				FromStruct(".Program.Name", ".program.name", in, iconv.Keep)
		}).
		IfNotNil(".Building", in, func(v *iconv.Record) {
			v.FromStruct(".Building.Id", ".building.id", in, iconv.Keep).
				FromStruct(".Building.Name", ".building.name", in, iconv.Keep)
		}).
		Self()
}
