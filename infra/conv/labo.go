package conv

import (
	"github.com/yamagame/school-api-gateway/entity"
	"github.com/yamagame/school-api-gateway/infra/model"
	"github.com/yamagame/school-api-gateway/proto/school"
)

func LaboToInfra(in *entity.Labo) (*model.Labo, error) {
	out := &model.Labo{}
	in.SetIfExist("id", func(key string, val interface{}) error {
		out.ID = val.(int32)
		return nil
	})
	in.SetIfExist("name", func(key string, val interface{}) error {
		out.Name = ToStrPtr(val)
		return nil
	})
	in.SetIfExist("group", func(key string, val interface{}) error {
		out.Group = ToStrPtr(val)
		return nil
	})
	in.SetIfExist("program", func(key string, val interface{}) error {
		out.Program = ToStrPtr(val)
		return nil
	})
	return out, nil
}

func LaboToProto(in *entity.Labo) (*school.Labo, error) {
	out := &school.Labo{}
	in.SetIfExist("id", func(key string, val interface{}) error {
		out.Id = val.(int32)
		return nil
	})
	in.SetIfExist("name", func(key string, val interface{}) error {
		out.Name = ToStr(val)
		return nil
	})
	in.SetIfExist("group", func(key string, val interface{}) error {
		out.Group = ToStr(val)
		return nil
	})
	in.SetIfExist("program", func(key string, val interface{}) error {
		out.Program = ToStr(val)
		return nil
	})
	return out, nil
}

func LaboToEntity(in *model.Labo) (*entity.Labo, error) {
	var err error
	out := entity.NewLabo()
	if err = out.SetValue("id", in.ID); err != nil {
		return nil, err
	}
	if err = SetIfNotNil("name", out, in.Name); err != nil {
		return nil, err
	}
	if err = SetIfNotNil("group", out, in.Group); err != nil {
		return nil, err
	}
	if err = SetIfNotNil("program", out, in.Program); err != nil {
		return nil, err
	}
	if err = SetIfNotNil("building", out, in.Building); err != nil {
		return nil, err
	}
	return out, nil
}
