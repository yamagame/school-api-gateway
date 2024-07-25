package model

import "time"

type Property struct {
	ID        int32      `gorm:"primary; comment:主キーの標準フィールド;"`
	LaboID    int32      `gorm:"comment:参照キー;"`
	Name      string     `gorm:"type:varchar(255); comment:プログラム名;"`
	CreatedAt *time.Time `gorm:"comment:GORMによって自動的に管理される作成時間;"`
	UpdatedAt *time.Time `gorm:"comment:GORMによって自動的に管理される更新時間;"`
	Error     error      `gorm:"-"`
}

func (m *Property) HasError() bool {
	return m.Error != nil
}

func (m *Property) GetError() string {
	if m.HasError() {
		return m.Error.Error()
	}
	return ""
}
