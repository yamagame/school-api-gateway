package svcconv

import (
	"github.com/yamagame/school-api-gateway/infra/infconv"
	"github.com/yamagame/school-api-gateway/infra/model"
	"github.com/yamagame/school-api-gateway/proto/school"
)

func LaboToProto(in *model.Labo) (*school.Labo, error) {
	entity, err := infconv.LaboToEntity(in)
	if err != nil {
		return nil, err
	}
	out := &school.Labo{}
	if value, err := entity.Value("id"); err == nil {
		out.Id = value.Get().(int32)
	}
	if value, err := entity.Value("name"); err == nil {
		out.Name = value.Get().(string)
	}
	if group, err := entity.HasOne("group"); err == nil {
		if value, err := group.Value("name"); err == nil {
			out.Group = value.Get().(string)
		}
	}
	if program, err := entity.HasOne("program"); err == nil {
		if value, err := program.Value("name"); err == nil {
			out.Program = value.Get().(string)
		}
	}
	if building, err := entity.HasOne("building"); err == nil {
		if value, err := building.Value("name"); err == nil {
			out.Building = value.Get().(string)
		}
	}
	return out, nil
}
