package model

import (
	"time"

	"github.com/yamagame/school-api-gateway/pkg/irconv"
)

type Area struct {
	ID        int32      `gorm:"primary; comment:主キーの標準フィールド;"`
	Name      string     `gorm:"comment:地区名;"`
	CreatedAt *time.Time `gorm:"comment:GORMによって自動的に管理される作成時間;"`
	UpdatedAt *time.Time `gorm:"comment:GORMによって自動的に管理される更新時間;"`
}

// 東地区
// 西地区

type Areas struct {
	*irconv.Slice[Area]
}
