package infconv

import (
	"github.com/yamagame/school-api-gateway/infra/model"
	"github.com/yamagame/school-api-gateway/pkg/irconv"
)

var Labo = LaboConv{}

var Labos = irconv.Convs[model.Labo, model.Labos, LaboConv]{
	Conv: Labo,
}

var Desk = DeskConv{}

var Desks = irconv.Convs[model.Desk, model.Desks, DeskConv]{
	Conv: Desk,
}

var Property = PropertyConv{}

var Properties = irconv.Convs[model.Property, model.Properties, PropertyConv]{
	Conv: Property,
}
