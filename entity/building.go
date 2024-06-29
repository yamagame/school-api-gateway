package entity

import (
	"github.com/yamagame/school-api-gateway/pkg/conv"
)

func NewBuilding(id int32) *conv.Record {
	v := conv.NewRecord()
	v.SetValue("id", id)
	v.SetValue("name", "")
	return v
}
