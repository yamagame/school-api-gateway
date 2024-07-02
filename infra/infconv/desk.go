package infconv

import (
	"github.com/yamagame/school-api-gateway/entity"
	"github.com/yamagame/school-api-gateway/infra/model"
	"github.com/yamagame/school-api-gateway/pkg/conv"
)

type DeskConv struct{}

func (DeskConv) ToInfra(in *conv.Record) (*model.Desk, error) {
	out := &model.Desk{}
	in.ToStruct(".id", ".ID", out, conv.Raw)
	in.ToStruct(".labo_id", ".LaboID", out, conv.Raw)
	return out, nil
}

func (DeskConv) ToEntity(in *model.Desk) (*conv.Record, error) {
	out := entity.NewDesk()
	out.FromStruct(".ID", ".id", in, conv.Raw)
	out.FromStruct(".LaboID", ".labo_id", in, conv.Raw)
	return out, nil
}
