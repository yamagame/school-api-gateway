package model

import (
	"time"
)

type Alias struct {
	ID        int32      `gorm:"primary; comment:主キーの標準フィールド;"`
	Name      string     `gorm:"comment:別名;"`
	CreatedAt *time.Time `gorm:"comment:GORMによって自動的に管理される作成時間;"`
	UpdatedAt *time.Time `gorm:"comment:GORMによって自動的に管理される更新時間;"`
	Error     error      `gorm:"-"`
}

func (m *Alias) HasError() bool {
	return m.Error != nil
}

func (m *Alias) GetError() string {
	if m.HasError() {
		return m.Error.Error()
	}
	return ""
}
