package model

import (
	"time"

	"github.com/yamagame/school-api-gateway/pkg/irconv"
)

type Role struct {
	ID        int32      `gorm:"primary; comment:主キーの標準フィールド;"`
	Name      string     `gorm:"comment:役割名;"`
	CreatedAt *time.Time `gorm:"comment:GORMによって自動的に管理される作成時間;"`
	UpdatedAt *time.Time `gorm:"comment:GORMによって自動的に管理される更新時間;"`
	Error     error      `gorm:"-"`
}

// 事務
// 教授
// 補助

type Roles struct {
	*irconv.Slice[Role]
}

func (m *Role) HasError() bool {
	return m.Error != nil
}
