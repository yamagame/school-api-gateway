package infconv

import (
	"github.com/yamagame/school-api-gateway/infra/model"
	"github.com/yamagame/school-api-gateway/pkg/iconv"
)

var Labo = LaboConv{}

var Labos = iconv.Convs[model.Labo, model.Labos, LaboConv]{
	Conv: Labo,
}

var Desk = DeskConv{}

var Desks = iconv.Convs[model.Desk, model.Desks, DeskConv]{
	Conv: Desk,
}

var Property = PropertyConv{}

var Properties = iconv.Convs[model.Property, model.Properties, PropertyConv]{
	Conv: Property,
}
