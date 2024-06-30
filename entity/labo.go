package entity

import (
	"github.com/yamagame/school-api-gateway/pkg/conv"
)

func NewLabo() *conv.Record {
	v := conv.NewRecord()
	v.SetValue("id", int32(0))
	v.SetValue("name", "")
	v.SetHasOne("group", NewGroup())
	v.SetHasOne("program", NewProgram())
	v.SetHasOne("building", NewBuilding())
	return v
}
