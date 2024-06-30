package entity

import "github.com/yamagame/school-api-gateway/pkg/conv"

func NewStudent() *conv.Record {
	v := conv.NewRecord()
	v.SetValue("id", int32(0))
	return v
}
