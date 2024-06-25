package model

import "time"

type School struct {
	ID        int32      `gorm:"primary; comment:主キーの標準フィールド;"`
	Name      string     `gorm:"comment:学校名;"`
	Buildings []Building `gorm:"many2many:school_buildings;"`
	Persons   []Person   `gorm:"many2many:school_persons;"`
	Classes   []Class    `gorm:"many2many:school_classes;"`
	CreatedAt *time.Time `gorm:"comment:GORMによって自動的に管理される作成時間;"`
	UpdatedAt *time.Time `gorm:"comment:GORMによって自動的に管理される更新時間;"`
}
