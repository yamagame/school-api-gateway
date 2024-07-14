package infconv

import (
	"github.com/yamagame/school-api-gateway/infra/model"
	"github.com/yamagame/school-api-gateway/pkg/conv"
)

var Labo = LaboConv{}

var Labos = conv.Convs[model.Labo, model.Labos, LaboConv]{
	Conv: Labo,
}

var Desk = DeskConv{}

var Desks = conv.Convs[model.Desk, model.Desks, DeskConv]{
	Conv: Desk,
}
