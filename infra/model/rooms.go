package model

import "time"

type Room struct {
	ID         int32      `gorm:"primary; comment:主キーの標準フィールド;"`
	Name       string     `gorm:"comment:部屋名;"`
	BuildingID int32      `gorm:"comment:建物ID;"`
	Building   Building   `gorm:"comment:建物;"`
	Floor      int32      `gorm:"comment:階数;"`
	CreatedAt  *time.Time `gorm:"comment:GORMによって自動的に管理される作成時間;"`
	UpdatedAt  *time.Time `gorm:"comment:GORMによって自動的に管理される更新時間;"`
}
