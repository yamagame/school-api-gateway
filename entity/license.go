package entity

import "github.com/yamagame/school-api-gateway/pkg/conv"

func NewLicnese(id int32) *conv.Record {
	v := conv.NewRecord()
	v.SetValue("id", id)
	return v
}
