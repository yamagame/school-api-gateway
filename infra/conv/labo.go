package conv

import (
	"github.com/yamagame/school-api-gateway/entity"
	"github.com/yamagame/school-api-gateway/infra/model"
	"github.com/yamagame/school-api-gateway/proto/school"
)

func LaboToInfra(in *entity.Labo) (*model.Labo, error) {
	out := &model.Labo{}
	if err := in.SetIfExist("id", func(_ string, val interface{}) error {
		out.ID = val.(int32)
		return nil
	}); err != nil {
		return nil, err
	}
	if err := in.SetIfExist("name", func(_ string, val interface{}) error {
		out.Name = ToStrPtr(val)
		return nil
	}); err != nil {
		return nil, err
	}
	if err := in.SetIfExist("group", func(_ string, val interface{}) error {
		out.Group = ToStrPtr(val)
		return nil
	}); err != nil {
		return nil, err
	}
	if err := in.SetIfExist("program", func(_ string, val interface{}) error {
		out.Program = ToStrPtr(val)
		return nil
	}); err != nil {
		return nil, err
	}
	return out, nil
}

func LaboToProto(in *entity.Labo) (*school.Labo, error) {
	out := &school.Labo{}
	if err := in.SetIfExist("id", func(_ string, val interface{}) error {
		out.Id = val.(int32)
		return nil
	}); err != nil {
		return nil, err
	}
	if err := in.SetIfExist("name", func(_ string, val interface{}) error {
		out.Name = ToStr(val)
		return nil
	}); err != nil {
		return nil, err
	}
	if err := in.SetIfExist("group", func(_ string, val interface{}) error {
		out.Group = ToStr(val)
		return nil
	}); err != nil {
		return nil, err
	}
	if err := in.SetIfExist("program", func(_ string, val interface{}) error {
		out.Program = ToStr(val)
		return nil
	}); err != nil {
		return nil, err
	}
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
