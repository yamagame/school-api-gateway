package svcconv

import (
	"github.com/yamagame/school-api-gateway/pkg/irconv"
	"github.com/yamagame/school-api-gateway/proto/school"
)

var (
	laboConv = LaboConv{}
	labos    = irconv.Convs[school.Labo, school.Labos, LaboConv]{
		Conv: laboConv,
	}
)
