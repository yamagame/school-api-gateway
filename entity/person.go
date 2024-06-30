package entity

import (
	"time"

	"github.com/yamagame/school-api-gateway/pkg/conv"
)

func NewPerson() *conv.Record {
	v := conv.NewRecord()
	v.SetValue("id", int32(0))
	v.SetValue("name", "")
	v.SetValue("birthday", time.Now())
	v.SetValue("address", NewAddress())
	v.SetValue("student", NewStudent())
	v.SetValue("professor", NewProfessor())
	v.SetHasMany("licenses", NewLicnese())
	return v
}
