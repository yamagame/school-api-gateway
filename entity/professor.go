package entity

import "github.com/yamagame/school-api-gateway/pkg/field"

func NewProfessor(id int32) *field.Field {
	v := &field.Field{}
	v.SetValue("id", id)
	return v
}
