package entity

import "github.com/yamagame/school-api-gateway/pkg/conv"

func NewLicnese() *conv.Record {
	return conv.NewRecord().
		SetValue("id", int32(0))
}
