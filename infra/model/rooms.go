package model

import (
	"time"

	"github.com/yamagame/school-api-gateway/pkg/conv"
)

type Room struct {
	ID         int32      `gorm:"primary; comment:主キーの標準フィールド;"`
	Name       string     `gorm:"comment:部屋名;"`
	BuildingID int32      `gorm:"comment:建物ID;"`
	Building   Building   `gorm:"comment:建物;"`
	Floor      int32      `gorm:"comment:階数;"`
	CreatedAt  *time.Time `gorm:"comment:GORMによって自動的に管理される作成時間;"`
	UpdatedAt  *time.Time `gorm:"comment:GORMによって自動的に管理される更新時間;"`
}

type Rooms struct {
	*conv.Slice[Room]
}

func NewRooms(variables ...*Room) *Rooms {
	return &Rooms{
		Slice: conv.NewSlice(variables...),
	}
}
