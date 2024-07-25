package svcconv

import (
	"github.com/yamagame/school-api-gateway/infra/infconv"
	"github.com/yamagame/school-api-gateway/infra/model"
	"github.com/yamagame/school-api-gateway/proto/school"
)

type LabosConv struct {
}

func (c LabosConv) ToInfra(in *school.Labos) *model.Labos {
	return infconv.Labos.ToStruct(labos.ToIModel(in))
}

func (c LabosConv) ToProto(in *model.Labos) *school.Labos {
	return labos.ToStruct(infconv.Labos.ToIModel(in))
}
