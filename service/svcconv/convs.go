package svcconv

import (
	"github.com/yamagame/school-api-gateway/pkg/irconv"
	"github.com/yamagame/school-api-gateway/proto/school"
)

var Labo = LaboConv{}

var Labos = irconv.Convs[school.Labo, school.Labos, LaboConv]{
	Conv: Labo,
}
