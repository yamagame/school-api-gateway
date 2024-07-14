package svcconv

import (
	"github.com/yamagame/school-api-gateway/pkg/conv"
	"github.com/yamagame/school-api-gateway/proto/school"
)

var Labo = LaboConv{}

var Labos = conv.Convs[school.Labo, school.Labos, LaboConv]{
	Conv: Labo,
}
