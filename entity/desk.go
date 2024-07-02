package entity

import "github.com/yamagame/school-api-gateway/pkg/conv"

func NewDesk() *conv.Record {
	v := conv.NewRecord()
	v.SetValue("id", int32(0))
	v.SetValue("labo_id", int32(0))
	return v
}
