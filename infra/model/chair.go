package model

import "github.com/yamagame/school-api-gateway/pkg/conv"

type Chair struct {
	ID     int32 `gorm:"primary; comment:主キーの標準フィールド;"`
	LaboID int32
}

type Chairs struct {
	*conv.Slice[Chair]
}

func NewChairs(variables ...*Chair) *Chairs {
	return &Chairs{
		Slice: conv.NewSlice(variables...),
	}
}
