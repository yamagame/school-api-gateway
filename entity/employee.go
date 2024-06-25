package entity

import "github.com/yamagame/school-api-gateway/pkg/field"

func NewEmployee(id int32) *field.Field {
	v := &field.Field{}
	v.SetValue("id", id)
	v.SetValue("name", "")
	return v
}
