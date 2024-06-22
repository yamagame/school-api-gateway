package model

import "time"

type Labo struct {
	ID        uint       `gorm:"primary; comment:主キーの標準フィールド;"`
	Name      *string    `gorm:"type:varchar(255); comment:研究室の名前;"`
	Group     *string    `gorm:"type:varchar(255); comment:研究室のグループ;"`
	Program   *string    `gorm:"type:varchar(255); comment:研究室のプログラム;"`
	Building  *string    `gorm:"type:varchar(255); comment:研究室の建物;"`
	CreatedAt *time.Time `gorm:"comment:GORMによって自動的に管理される作成時間"`
	UpdatedAt *time.Time `gorm:"comment:GORMによって自動的に管理される更新時間"`
}
