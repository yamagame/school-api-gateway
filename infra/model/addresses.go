package model

import "time"

type Address struct {
	ID         int32      `gorm:"primary; comment:主キーの標準フィールド;"`
	ZipCode    string     `gorm:"comment:郵便番号;"`
	Prefecture string     `gorm:"comment:都道府県;"`
	City       string     `gorm:"comment:市町村;"`
	Number     string     `gorm:"comment:番地;"`
	Building   string     `gorm:"comment:建物名;"`
	CreatedAt  *time.Time `gorm:"comment:GORMによって自動的に管理される作成時間;"`
	UpdatedAt  *time.Time `gorm:"comment:GORMによって自動的に管理される更新時間;"`
}
