package model

import (
	"time"

	"github.com/yamagame/school-api-gateway/pkg/irconv"
)

type Course struct {
	ID        int32      `gorm:"primary; comment:主キーの標準フィールド;"`
	Name      string     `gorm:"comment:コース名;"`
	CreatedAt *time.Time `gorm:"comment:GORMによって自動的に管理される作成時間;"`
	UpdatedAt *time.Time `gorm:"comment:GORMによって自動的に管理される更新時間;"`
	Error     error      `gorm:"-"`
}

type Courses struct {
	*irconv.Slice[Course]
}

func (m *Course) HasError() bool {
	return m.Error != nil
}
