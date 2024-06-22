package entity

import (
	"fmt"
)

var (
	NotFoundErr = fmt.Errorf("値が見つかりません")
)

type FieldInterface interface {
	AddProp(key, defval string)
	GetValue(key string) (string, error)
	SetValue(key, value string) error
	AddUintProp(key, defval uint)
	GetUintValue(key string) (uint, error)
	SetUintValue(key string, value uint) error
}

type Field struct {
	values map[string]*Value
}

func NewField() Field {
	return Field{
		values: map[string]*Value{},
	}
}

func (x *Field) AddProp(key, defval string) {
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

func (x *Field) GetValue(key string) (string, error) {
	if v, ok := x.values[key]; ok {
		return v.Get().(string), nil
	}
	return "", NotFoundErr
}

func (x *Field) SetValue(key, value string) error {
	if v, ok := x.values[key]; ok {
		v.Set(value)
		return nil
	}
	return NotFoundErr
}

func (x *Field) AddUintProp(key string, defval uint) {
	x.values[key] = &Value{
		Key:     key,
		Default: defval,
	}
}

func (x *Field) GetUintValue(key string) (uint, error) {
	if v, ok := x.values[key]; ok {
		return v.Get().(uint), nil
	}
	return 0, NotFoundErr
}

func (x *Field) SetUintValue(key string, value uint) error {
	if v, ok := x.values[key]; ok {
		v.Set(value)
		return nil
	}
	return NotFoundErr
}
