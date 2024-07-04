package entity

import (
	"time"

	"github.com/yamagame/school-api-gateway/pkg/conv"
)

func NewPerson() *conv.Record {
	return conv.NewRecord().
		SetValue("id", int32(0), conv.PRIMARY).
		SetValue("name", "").
		SetValue("birthday", time.Now()).
		SetValue("address", NewAddress()).
		SetValue("student", NewStudent()).
		SetValue("professor", NewProfessor()).
		SetHasMany("licenses", conv.NewMany(NewLicnese()))
}
