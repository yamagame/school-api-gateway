package entity

import (
	"github.com/yamagame/school-api-gateway/pkg/conv"
)

func NewLabo() *conv.Record {
	v := conv.NewRecord().
		SetValue("id", int32(0)).
		SetValue("name", "").
		SetHasOne("group", NewGroup()).
		SetHasOne("program", NewProgram()).
		SetHasOne("building", NewBuilding()).
		SetHasMany("desk", NewDesk())
	return v
}
