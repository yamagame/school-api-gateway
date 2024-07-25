package svcconv

import (
	"github.com/yamagame/school-api-gateway/pkg/iconv"
	"github.com/yamagame/school-api-gateway/proto/school"
)

var (
	laboConv = LaboConv{}
	labos    = iconv.Convs[school.Labo, school.Labos, LaboConv]{
		Conv: laboConv,
	}
)
