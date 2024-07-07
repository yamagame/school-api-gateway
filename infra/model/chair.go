package model

import "github.com/yamagame/school-api-gateway/pkg/conv"

type Chair struct {
	ID     int32 `gorm:"primary; comment:主キーの標準フィールド;"`
	LaboID int32
}

type Chairs struct {
	*conv.Variables[Chair]
}

func NewChairs(variables ...*Chair) *Chairs {
	return &Chairs{
		Variables: conv.NewVariables(variables...),
	}
}
