package model

import (
	"time"

	"github.com/yamagame/school-api-gateway/pkg/conv"
)

type Professor struct {
	ID        int32      `gorm:"primary; comment:主キーの標準フィールド;"`
	PersonID  int32      `gorm:"primary; comment:人物ID;"`
	Person    Person     `gorm:"primary; comment:人物;"`
	Field     string     `gorm:"comment:専門領域;"`
	RoomID    int32      `gorm:"comment:部屋ID;"`
	Room      Room       `gorm:"comment:部屋;"`
	CreatedAt *time.Time `gorm:"comment:GORMによって自動的に管理される作成時間;"`
	UpdatedAt *time.Time `gorm:"comment:GORMによって自動的に管理される更新時間;"`
}

type Professors struct {
	*conv.Slice[Professor]
}

func NewProfessors(variables ...*Professor) *Professors {
	return &Professors{
		Slice: conv.NewSlice(variables...),
	}
}
