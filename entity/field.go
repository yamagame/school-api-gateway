package entity

import (
	"fmt"
)

var (
	NotFoundErr = fmt.Errorf("値が見つかりません")
)

type FieldInterface interface {
	AddProp(key string, defval interface{})
	IsExist(key string) (bool, error)
	GetValue(key string) (interface{}, error)
	SetValue(key string, value interface{}) error
	SetIfExist(key string, cb func(key string, val interface{}) error) error
	ToMap() map[string]interface{}
}

type Field struct {
	values map[string]*Value
}

func NewField() Field {
	return Field{
		values: map[string]*Value{},
	}
}

func (x *Field) AddProp(key string, defval interface{}) {
	x.values[key] = &Value{
		Key:     key,
		Default: defval,
	}
}

func (x *Field) IsExist(key string) (bool, error) {
	if v, ok := x.values[key]; ok {
		return v.IsExist(), nil
	}
	return false, NotFoundErr
}

func (x *Field) GetValue(key string) (interface{}, error) {
	if v, ok := x.values[key]; ok {
		return v.Get(), nil
	}
	return nil, NotFoundErr
}

func (x *Field) SetValue(key string, value interface{}) error {
	if v, ok := x.values[key]; ok {
		v.Set(value)
		return nil
	}
	return NotFoundErr
}

func (x *Field) SetIfExist(key string, cb func(key string, val interface{}) error) error {
	if flag, _ := x.IsExist(key); flag {
		val, err := x.GetValue(key)
		if err != nil {
			return err
		}
		return cb(key, val)
	}
	return NotFoundErr
}

func (x *Field) ToMap() map[string]interface{} {
	r := map[string]interface{}{}
	for key, val := range x.values {
		r[key] = val.Get()
	}
	return r
}
