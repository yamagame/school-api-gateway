package model

import (
	"time"

	"github.com/yamagame/school-api-gateway/pkg/conv"
)

type Alias struct {
	ID        int32      `gorm:"primary; comment:主キーの標準フィールド;"`
	Name      string     `gorm:"comment:別名;"`
	CreatedAt *time.Time `gorm:"comment:GORMによって自動的に管理される作成時間;"`
	UpdatedAt *time.Time `gorm:"comment:GORMによって自動的に管理される更新時間;"`
}

type Aliases struct {
	*conv.Slice[Alias]
}

func NewAliases(variables ...*Alias) *Aliases {
	return &Aliases{
		Slice: conv.NewSlice(variables...),
	}
}
