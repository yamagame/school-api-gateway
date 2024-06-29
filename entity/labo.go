package entity

import (
	"github.com/yamagame/school-api-gateway/pkg/conv"
)

func NewLabo(id int32) *conv.Record {
	v := conv.NewRecord()
	v.SetValue("id", id)
	v.SetValue("name", "")
	v.SetHasOne("group", NewGroup(0))
	v.SetHasOne("program", NewProgram(0))
	v.SetHasOne("building", NewBuilding(0))
	return v
}
