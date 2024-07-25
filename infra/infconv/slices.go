package infconv

import (
	"github.com/yamagame/school-api-gateway/infra/model"
)

func (DeskConv) NewSlice(models ...*model.Desk) *model.Desks {
	return &model.Desks{
		SliceContainer: Desks.NewSlice(models...),
	}
}

func (GroupConv) NewSlice(models ...*model.Group) *model.Groups {
	return &model.Groups{
		SliceContainer: Groups.NewSlice(models...),
	}
}

func (LaboConv) NewSlice(models ...*model.Labo) *model.Labos {
	return &model.Labos{
		SliceContainer: Labos.NewSlice(models...),
	}
}

func (PropertyConv) NewSlice(models ...*model.Property) *model.Properties {
	return &model.Properties{
		SliceContainer: Properties.NewSlice(models...),
	}
}

func NewDeskSlice(models ...*model.Desk) *model.Desks {
	return Desk.NewSlice(models...)
}

func NewGroupSlice(models ...*model.Group) *model.Groups {
	return Group.NewSlice(models...)
}

func NewLaboSlice(models ...*model.Labo) *model.Labos {
	return Labo.NewSlice(models...)
}

func NewPropertySlice(models ...*model.Property) *model.Properties {
	return Property.NewSlice(models...)
}
