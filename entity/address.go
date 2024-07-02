package entity

import "github.com/yamagame/school-api-gateway/pkg/conv"

func NewAddress() *conv.Record {
	return conv.NewRecord().
		SetValue("id", int32(0)).
		SetValue("name", "")
}
