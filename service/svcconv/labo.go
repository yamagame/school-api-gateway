package svcconv

import (
	"github.com/yamagame/school-api-gateway/pkg/field"
	"github.com/yamagame/school-api-gateway/proto/school"
)

func LaboToProto(in *field.Field) (*school.Labo, error) {
	out := &school.Labo{}
	if value, err := in.Value("id"); err == nil {
		out.Id = value.Get().(int32)
	}
	if value, err := in.Value("name"); err == nil {
		out.Name = value.Get().(string)
	}
	if group, err := in.Field("group"); err == nil {
		if value, err := group.Value("name"); err == nil {
			out.Group = value.Get().(string)
		}
	}
	if program, err := in.Field("program"); err == nil {
		if value, err := program.Value("name"); err == nil {
			out.Program = value.Get().(string)
		}
	}
	if building, err := in.Field("building"); err == nil {
		if value, err := building.Value("name"); err == nil {
			out.Building = value.Get().(string)
		}
	}
	return out, nil
}
