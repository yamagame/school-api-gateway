package entity

import (
	"time"

	"github.com/yamagame/school-api-gateway/pkg/field"
)

func NewPerson(id int32) *field.Field {
	v := &field.Field{}
	v.SetValue("id", id)
	v.SetValue("name", "")
	v.SetValue("birthday", time.Now())
	v.SetValue("address", NewAddress(0))
	v.SetValue("student", NewStudent(0))
	v.SetValue("professor", NewProfessor(0))
	v.SetArray("licenses", NewLicnese(0))
	return v
}
