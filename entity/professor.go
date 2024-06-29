package entity

import "github.com/yamagame/school-api-gateway/pkg/conv"

func NewProfessor(id int32) *conv.Record {
	v := conv.NewRecord()
	v.SetValue("id", id)
	return v
}
