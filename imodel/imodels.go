package imodel

import (
	"time"

	"github.com/yamagame/school-api-gateway/pkg/iconv"
	"github.com/yamagame/school-api-gateway/pkg/zero"
)

var NewAddress = iconv.NewRecord().
	SetValue("id", zero.Int32, iconv.PRIMARY).
	SetValue("name", zero.Str).
	Func()

var NewBuilding = iconv.NewRecord().
	SetValue("id", zero.Int32, iconv.PRIMARY).
	SetValue("name", zero.Str).
	Func()

var NewChair = iconv.NewRecord().
	SetValue("id", zero.Int32, iconv.PRIMARY).
	SetValue("labo_id", zero.Int32).
	Func()

var NewClass = iconv.NewRecord().
	SetValue("id", zero.Int32, iconv.PRIMARY).
	SetValue("name", zero.Str).
	Func()

var NewDesk = iconv.NewRecord().
	SetValue("id", zero.Int32, iconv.PRIMARY).
	SetValue("labo_id", zero.Int32).
	SetValue("name", zero.Str).
	SetValue("product_code", zero.Str).
	Func()

var NewEmployee = iconv.NewRecord().
	SetValue("id", zero.Int32, iconv.PRIMARY).
	SetValue("name", zero.Str).
	Func()

var NewGroup = iconv.NewRecord().
	SetValue("id", zero.Int32, iconv.PRIMARY).
	SetValue("name", zero.Str).
	Func()

var NewLabo = iconv.NewRecord().
	SetValue("id", zero.Int32, iconv.PRIMARY).
	SetValue("name", zero.Str).
	SetValue("url", zero.Str).
	SetBelongTo("group", NewGroup()).
	SetBelongTo("program", NewProgram()).
	SetBelongTo("building", NewBuilding()).
	SetHasMany("desk", iconv.NewMany(NewDesk())).
	SetHasMany("chair", iconv.NewMany(NewChair())).
	SetHasOne("property", NewProperty()).
	Func()

var NewLicnese = iconv.NewRecord().
	SetValue("id", zero.Int32, iconv.PRIMARY).
	Func()

var NewPerson = iconv.NewRecord().
	SetValue("id", zero.Int32, iconv.PRIMARY).
	SetValue("name", zero.Str).
	SetValue("birthday", time.Now()).
	SetValue("address", NewAddress()).
	SetValue("student", NewStudent()).
	SetValue("professor", NewProfessor()).
	SetHasMany("licenses", iconv.NewMany(NewLicnese())).
	Func()

var NewProfessor = iconv.NewRecord().
	SetValue("id", zero.Int32, iconv.PRIMARY).
	Func()

var NewProgram = iconv.NewRecord().
	SetValue("id", zero.Int32, iconv.PRIMARY).
	SetValue("name", zero.Str).
	Func()

var NewRoom = iconv.NewRecord().
	SetValue("id", zero.Int32, iconv.PRIMARY).
	SetValue("name", zero.Str).
	SetHasOne("building", NewBuilding()).
	SetValue("floor", zero.Int).
	Func()

var NewStudent = iconv.NewRecord().
	SetValue("id", zero.Int32, iconv.PRIMARY).
	Func()

var NewProperty = iconv.NewRecord().
	SetValue("id", zero.Int32, iconv.PRIMARY).
	SetValue("labo_id", zero.Int32).
	SetValue("name", zero.Str).
	Func()
