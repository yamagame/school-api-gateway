package entity

import (
	"time"

	"github.com/yamagame/school-api-gateway/pkg/conv"
)

func NewPerson(id int32) *conv.Record {
	v := conv.NewRecord()
	v.SetValue("id", id)
	v.SetValue("name", "")
	v.SetValue("birthday", time.Now())
	v.SetValue("address", NewAddress(0))
	v.SetValue("student", NewStudent(0))
	v.SetValue("professor", NewProfessor(0))
	v.SetHasMany("licenses", NewLicnese(0))
	return v
}
