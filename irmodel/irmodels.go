package irmodel

import (
	"time"

	"github.com/yamagame/school-api-gateway/pkg/conv"
)

var NewAddress = conv.NewRecord().
	SetValue("id", int32(0), conv.PRIMARY).
	SetValue("name", "").
	Func()

var NewBuilding = conv.NewRecord().
	SetValue("id", int32(0), conv.PRIMARY).
	SetValue("name", "").
	Func()

var NewChair = conv.NewRecord().
	SetValue("id", int32(0), conv.PRIMARY).
	SetValue("labo_id", int32(0)).
	Func()

var NewClass = conv.NewRecord().
	SetValue("id", int32(0), conv.PRIMARY).
	SetValue("name", "").
	Func()

var NewDesk = conv.NewRecord().
	SetValue("id", int32(0), conv.PRIMARY).
	SetValue("labo_id", int32(0)).
	Func()

var NewEmployee = conv.NewRecord().
	SetValue("id", int32(0), conv.PRIMARY).
	SetValue("name", "").
	Func()

var NewGroup = conv.NewRecord().
	SetValue("id", int32(0), conv.PRIMARY).
	SetValue("name", "").
	Func()

var NewLabo = conv.NewRecord().
	SetValue("id", int32(0), conv.PRIMARY).
	SetValue("name", "").
	SetValue("url", "").
	SetBelongTo("group", NewGroup()).
	SetBelongTo("program", NewProgram()).
	SetBelongTo("building", NewBuilding()).
	SetHasMany("desk", conv.NewMany(NewDesk())).
	SetHasMany("chair", conv.NewMany(NewChair())).
	Func()

var NewLicnese = conv.NewRecord().
	SetValue("id", int32(0), conv.PRIMARY).
	Func()

var NewPerson = conv.NewRecord().
	SetValue("id", int32(0), conv.PRIMARY).
	SetValue("name", "").
	SetValue("birthday", time.Now()).
	SetValue("address", NewAddress()).
	SetValue("student", NewStudent()).
	SetValue("professor", NewProfessor()).
	SetHasMany("licenses", conv.NewMany(NewLicnese())).
	Func()

var NewProfessor = conv.NewRecord().
	SetValue("id", int32(0), conv.PRIMARY).
	Func()

var NewProgram = conv.NewRecord().
	SetValue("id", int32(0), conv.PRIMARY).
	SetValue("name", "").
	Func()

var NewRoom = conv.NewRecord().
	SetValue("id", int32(0), conv.PRIMARY).
	SetValue("name", "").
	SetHasOne("building", NewBuilding()).
	SetValue("floor", 0).
	Func()

var NewStudent = conv.NewRecord().
	SetValue("id", int32(0), conv.PRIMARY).
	Func()
