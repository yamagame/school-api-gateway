package infconv

import (
	"github.com/yamagame/school-api-gateway/infra/model"
)

var Labo = LaboConv{}
var Desk = DeskConv{}
var Labos = Convs[model.Labo, LaboConv]{
	conv: Labo,
}
var Desks = Convs[model.Desk, DeskConv]{
	conv: Desk,
}
