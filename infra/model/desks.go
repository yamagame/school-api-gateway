package model

import "github.com/yamagame/school-api-gateway/pkg/irconv"

type Desk struct {
	ID     int32 `gorm:"primary; comment:主キーの標準フィールド;"`
	LaboID int32
	Name   string
}

type Desks struct {
	*irconv.Slice[Desk]
}
