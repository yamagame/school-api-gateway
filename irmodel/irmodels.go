package irmodel

import (
	"time"

	"github.com/yamagame/school-api-gateway/pkg/irconv"
	"github.com/yamagame/school-api-gateway/pkg/zero"
)

var NewAddress = irconv.NewRecord().
	SetValue("id", zero.Int32, irconv.PRIMARY).
	SetValue("name", zero.Str).
	Func()

var NewBuilding = irconv.NewRecord().
	SetValue("id", zero.Int32, irconv.PRIMARY).
	SetValue("name", zero.Str).
	Func()

var NewChair = irconv.NewRecord().
	SetValue("id", zero.Int32, irconv.PRIMARY).
	SetValue("labo_id", zero.Int32).
	Func()

var NewClass = irconv.NewRecord().
	SetValue("id", zero.Int32, irconv.PRIMARY).
	SetValue("name", zero.Str).
	Func()

var NewDesk = irconv.NewRecord().
	SetValue("id", zero.Int32, irconv.PRIMARY).
	SetValue("labo_id", zero.Int32).
	SetValue("name", zero.Str).
	SetValue("product_code", zero.Str).
	Func()

var NewEmployee = irconv.NewRecord().
	SetValue("id", zero.Int32, irconv.PRIMARY).
	SetValue("name", zero.Str).
	Func()

var NewGroup = irconv.NewRecord().
	SetValue("id", zero.Int32, irconv.PRIMARY).
	SetValue("name", zero.Str).
	Func()

var NewLabo = irconv.NewRecord().
	SetValue("id", zero.Int32, irconv.PRIMARY).
	SetValue("name", zero.Str).
	SetValue("url", zero.Str).
	SetBelongTo("group", NewGroup()).
	SetBelongTo("program", NewProgram()).
	SetBelongTo("building", NewBuilding()).
	SetHasMany("desk", irconv.NewMany(NewDesk())).
	SetHasMany("chair", irconv.NewMany(NewChair())).
	SetHasOne("property", NewProperty()).
	Func()

var NewLicnese = irconv.NewRecord().
	SetValue("id", zero.Int32, irconv.PRIMARY).
	Func()

var NewPerson = irconv.NewRecord().
	SetValue("id", zero.Int32, irconv.PRIMARY).
	SetValue("name", zero.Str).
	SetValue("birthday", time.Now()).
	SetValue("address", NewAddress()).
	SetValue("student", NewStudent()).
	SetValue("professor", NewProfessor()).
	SetHasMany("licenses", irconv.NewMany(NewLicnese())).
	Func()

var NewProfessor = irconv.NewRecord().
	SetValue("id", zero.Int32, irconv.PRIMARY).
	Func()

var NewProgram = irconv.NewRecord().
	SetValue("id", zero.Int32, irconv.PRIMARY).
	SetValue("name", zero.Str).
	Func()

var NewRoom = irconv.NewRecord().
	SetValue("id", zero.Int32, irconv.PRIMARY).
	SetValue("name", zero.Str).
	SetHasOne("building", NewBuilding()).
	SetValue("floor", zero.Int).
	Func()

var NewStudent = irconv.NewRecord().
	SetValue("id", zero.Int32, irconv.PRIMARY).
	Func()

var NewProperty = irconv.NewRecord().
	SetValue("id", zero.Int32, irconv.PRIMARY).
	SetValue("labo_id", zero.Int32).
	SetValue("name", zero.Str).
	Func()
