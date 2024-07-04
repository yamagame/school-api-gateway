package entity

import (
	"github.com/yamagame/school-api-gateway/pkg/conv"
)

func NewLabo() *conv.Record {
	v := conv.NewRecord().
		SetValue("id", int32(0), conv.PRIMARY).
		SetValue("name", "").
		SetValue("url", "").
		SetBelongTo("group", NewGroup()).
		SetBelongTo("program", NewProgram()).
		SetBelongTo("building", NewBuilding()).
		SetHasMany("desk", conv.NewMany(NewDesk()))
	return v
}
