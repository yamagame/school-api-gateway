package entity

import "github.com/yamagame/school-api-gateway/pkg/conv"

func NewDesk() *conv.Record {
	return conv.NewRecord().
		SetValue("id", int32(0)).
		SetValue("labo_id", int32(0))
}
