package model

import "github.com/yamagame/school-api-gateway/pkg/conv"

type Desk struct {
	ID     int32 `gorm:"primary; comment:主キーの標準フィールド;"`
	LaboID int32
}

type Desks struct {
	*conv.Variables[Desk]
}

func NewDesks(variables ...*Desk) *Desks {
	return &Desks{
		Variables: conv.NewVariables(variables...),
	}
}
