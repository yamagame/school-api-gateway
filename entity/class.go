package entity

import "github.com/yamagame/school-api-gateway/pkg/conv"

func NewClass() *conv.Record {
	v := conv.NewRecord()
	v.SetValue("id", int32(0))
	v.SetValue("name", "")
	return v
}
