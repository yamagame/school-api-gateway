package irmodel

import (
	"time"

	"github.com/yamagame/school-api-gateway/pkg/irconv"
)

var NewAddress = irconv.NewRecord().
	SetValue("id", int32(0), irconv.PRIMARY).
	SetValue("name", "").
	Func()

var NewBuilding = irconv.NewRecord().
	SetValue("id", int32(0), irconv.PRIMARY).
	SetValue("name", "").
	Func()

var NewChair = irconv.NewRecord().
	SetValue("id", int32(0), irconv.PRIMARY).
	SetValue("labo_id", int32(0)).
	Func()

var NewClass = irconv.NewRecord().
	SetValue("id", int32(0), irconv.PRIMARY).
	SetValue("name", "").
	Func()

var NewDesk = irconv.NewRecord().
	SetValue("id", int32(0), irconv.PRIMARY).
	SetValue("labo_id", int32(0)).
	SetValue("name", "").
	Func()

var NewEmployee = irconv.NewRecord().
	SetValue("id", int32(0), irconv.PRIMARY).
	SetValue("name", "").
	Func()

var NewGroup = irconv.NewRecord().
	SetValue("id", int32(0), irconv.PRIMARY).
	SetValue("name", "").
	Func()

var NewLabo = irconv.NewRecord().
	SetValue("id", int32(0), irconv.PRIMARY).
	SetValue("name", "").
	SetValue("url", "").
	SetBelongTo("group", NewGroup()).
	SetBelongTo("program", NewProgram()).
	SetBelongTo("building", NewBuilding()).
	SetHasMany("desk", irconv.NewMany(NewDesk())).
	SetHasMany("chair", irconv.NewMany(NewChair())).
	Func()

var NewLicnese = irconv.NewRecord().
	SetValue("id", int32(0), irconv.PRIMARY).
	Func()

var NewPerson = irconv.NewRecord().
	SetValue("id", int32(0), irconv.PRIMARY).
	SetValue("name", "").
	SetValue("birthday", time.Now()).
	SetValue("address", NewAddress()).
	SetValue("student", NewStudent()).
	SetValue("professor", NewProfessor()).
	SetHasMany("licenses", irconv.NewMany(NewLicnese())).
	Func()

var NewProfessor = irconv.NewRecord().
	SetValue("id", int32(0), irconv.PRIMARY).
	Func()

var NewProgram = irconv.NewRecord().
	SetValue("id", int32(0), irconv.PRIMARY).
	SetValue("name", "").
	Func()

var NewRoom = irconv.NewRecord().
	SetValue("id", int32(0), irconv.PRIMARY).
	SetValue("name", "").
	SetHasOne("building", NewBuilding()).
	SetValue("floor", 0).
	Func()

var NewStudent = irconv.NewRecord().
	SetValue("id", int32(0), irconv.PRIMARY).
	Func()
