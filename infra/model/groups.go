package model

import (
	"time"

	"github.com/yamagame/school-api-gateway/pkg/irconv"
)

type Group struct {
	ID        int32      `gorm:"primary; comment:主キーの標準フィールド;"`
	Name      string     `gorm:"type:varchar(255); comment:専門領域名;"`
	CreatedAt *time.Time `gorm:"comment:GORMによって自動的に管理される作成時間;"`
	UpdatedAt *time.Time `gorm:"comment:GORMによって自動的に管理される更新時間;"`
	Error     error      `gorm:"-"`
}

// Ⅰ類（情報系）
// Ⅱ類（融合系）
// Ⅲ類（理工系）

type Groups struct {
	*irconv.Slice[Group]
}

func (m *Group) HasError() bool {
	return m.Error != nil
}
