package conv

import (
	"github.com/yamagame/school-api-gateway/entity"
	"github.com/yamagame/school-api-gateway/infra/model"
	"github.com/yamagame/school-api-gateway/pkg/field"
	"github.com/yamagame/school-api-gateway/proto/school"
)

func LaboToInfra(in *field.Field) (*model.Labo, error) {
	out := &model.Labo{}
	if value, err := in.Value("id"); err == nil {
		out.ID = value.Get().(int32)
	}
	if value, err := in.Value("name"); err == nil {
		out.Name = ToPtr(value.Get().(string))
	}
	if group, err := in.Field("group"); err == nil {
		if value, err := group.Value("name"); err == nil {
			out.Group.Name = value.Get().(string)
		}
	}
	if program, err := in.Field("program"); err == nil {
		if value, err := program.Value("name"); err == nil {
			out.Program.Name = value.Get().(string)
		}
	}
	if building, err := in.Field("building"); err == nil {
		if value, err := building.Value("name"); err == nil {
			out.Building.Name = value.Get().(string)
		}
	}
	return out, nil
}

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

func LaboToEntity(in *model.Labo) (*field.Field, error) {
	var err error
	out := entity.NewLabo(0)
	if err = out.Set("id", in.ID); err != nil {
		return nil, err
	}
	if in.Name != nil {
		if err = out.Set("name", *in.Name); err != nil {
			return nil, err
		}
	}
	if group, err := out.Field("group"); err == nil {
		if err = group.Set("id", in.GroupID); err != nil {
			return nil, err
		}
		if err = group.Set("name", in.Group.Name); err != nil {
			return nil, err
		}
	}
	if program, err := out.Field("program"); err == nil {
		if err = program.Set("id", in.ProgramID); err != nil {
			return nil, err
		}
		if err = program.Set("name", in.Program.Name); err != nil {
			return nil, err
		}
	}
	if building, err := out.Field("building"); err == nil {
		if err = building.Set("id", in.BuildingID); err != nil {
			return nil, err
		}
		if err = building.Set("name", in.Building.Name); err != nil {
			return nil, err
		}
	}
	return out, nil
}
