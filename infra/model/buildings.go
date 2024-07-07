package model

import (
	"time"

	"github.com/yamagame/school-api-gateway/pkg/conv"
)

type Building struct {
	ID        int32      `gorm:"primary; comment:主キーの標準フィールド;"`
	Name      string     `gorm:"comment:建物名;"`
	Floors    int32      `gorm:"comment:階数;"`
	Aliases   []Alias    `gorm:"many2many:building_aliases; comment:別名;"`
	CreatedAt *time.Time `gorm:"comment:GORMによって自動的に管理される作成時間;"`
	UpdatedAt *time.Time `gorm:"comment:GORMによって自動的に管理される更新時間;"`
}

type Buildings struct {
	*conv.Variables[Building]
}

func NewBuildings(variables ...*Building) *Buildings {
	return &Buildings{
		Variables: conv.NewVariables(variables...),
	}
}
