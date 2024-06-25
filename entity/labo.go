package entity

import (
	"github.com/yamagame/school-api-gateway/pkg/field"
)

func NewLabo(id int32) *field.Field {
	v := &field.Field{}
	v.SetValue("id", id)
	v.SetValue("name", "")
	v.SetField("group", NewGroup(0))
	v.SetField("program", NewProgram(0))
	v.SetField("building", NewBuilding(0))
	return v
}
