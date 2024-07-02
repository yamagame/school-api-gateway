package entity

import "github.com/yamagame/school-api-gateway/pkg/conv"

func NewRoom() *conv.Record {
	v := conv.NewRecord()
	v.SetValue("id", int32(0))
	v.SetValue("name", "")
	v.SetHasOne("building", NewBuilding())
	v.SetValue("floor", 0)
	return v
}
