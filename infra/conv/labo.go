package conv

import (
	"github.com/yamagame/school-api-gateway/entity"
	"github.com/yamagame/school-api-gateway/infra/model"
	"github.com/yamagame/school-api-gateway/proto/school"
)

func LaboToInfra(in *entity.Labo) (*model.Labo, error) {
	out := &model.Labo{}
	if flag, _ := in.IsExist("id"); flag {
		id, err := in.GetUintValue("id")
		if err != nil {
			return nil, err
		}
		out.ID = id
	}
	if flag, _ := in.IsExist("name"); flag {
		name, err := in.GetValue("name")
		if err != nil {
			return nil, err
		}
		out.Name = &name
	}
	if flag, _ := in.IsExist("group"); flag {
		group, err := in.GetValue("group")
		if err != nil {
			return nil, err
		}
		out.Group = &group
	}
	if flag, _ := in.IsExist("program"); flag {
		program, err := in.GetValue("program")
		if err != nil {
			return nil, err
		}
		out.Program = &program
	}
	if flag, _ := in.IsExist("building"); flag {
		building, err := in.GetValue("building")
		if err != nil {
			return nil, err
		}
		out.Building = &building
	}
	return out, nil
}

func LaboToProto(in *entity.Labo) (*school.Labo, error) {
	out := &school.Labo{}
	if flag, _ := in.IsExist("id"); flag {
		id, err := in.GetUintValue("id")
		if err != nil {
			return nil, err
		}
		out.Id = uint32(id)
	}
	if flag, _ := in.IsExist("name"); flag {
		name, err := in.GetValue("name")
		if err != nil {
			return nil, err
		}
		out.Name = name
	}
	if flag, _ := in.IsExist("group"); flag {
		group, err := in.GetValue("group")
		if err != nil {
			return nil, err
		}
		out.Group = group
	}
	if flag, _ := in.IsExist("program"); flag {
		program, err := in.GetValue("program")
		if err != nil {
			return nil, err
		}
		out.Program = program
	}
	if flag, _ := in.IsExist("building"); flag {
		building, err := in.GetValue("building")
		if err != nil {
			return nil, err
		}
		out.Building = building
	}
	return out, nil
}

func LaboToEntity(in *model.Labo) (*entity.Labo, error) {
	var err error
	out := entity.NewLabo()
	err = out.SetUintValue("id", in.ID)
	if err != nil {
		return nil, err
	}
	if in.Name != nil {
		err = out.SetValue("name", *in.Name)
		if err != nil {
			return nil, err
		}
	}
	if in.Group != nil {
		err = out.SetValue("group", *in.Group)
		if err != nil {
			return nil, err
		}
	}
	if in.Program != nil {
		err = out.SetValue("program", *in.Program)
		if err != nil {
			return nil, err
		}
	}
	if in.Building != nil {
		err = out.SetValue("building", *in.Building)
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}
