package model

import (
	"time"
)

type Class struct {
	ID        int32      `gorm:"primary; comment:主キーの標準フィールド;"`
	Name      string     `gorm:"comment:授業名;"`
	Rooms     []Room     `gorm:"many2many:class_rooms; comment:部屋;"`
	CreatedAt *time.Time `gorm:"comment:GORMによって自動的に管理される作成時間;"`
	UpdatedAt *time.Time `gorm:"comment:GORMによって自動的に管理される更新時間;"`
	Error     error      `gorm:"-"`
}

func (m *Class) HasError() bool {
	return m.Error != nil
}

func (m *Class) GetError() string {
	if m.HasError() {
		return m.Error.Error()
	}
	return ""
}
