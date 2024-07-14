package model

import (
	"time"

	"github.com/yamagame/school-api-gateway/pkg/irconv"
)

type Labo struct {
	ID         int32        `gorm:"primary; comment:主キーの標準フィールド;"`
	Name       *string      `gorm:"type:varchar(255); comment:研究室の名前;"`
	URL        *string      `gorm:"type:varchar(255); comment:研究室のURL;"`
	GroupID    *int32       `gorm:"comment:専門領域ID;"`
	Group      Group        `gorm:"comment:専門領域;"`
	ProgramID  *int32       `gorm:"comment:プログラムID;"`
	Program    Program      `gorm:"comment:プログラム;"`
	BuildingID *int32       `gorm:"comment:建物ID;"`
	Building   Building     `gorm:"comment:建物;"`
	Professors []*Professor `gorm:"many2many:labo_professors; comment:担当教授;"`
	Students   []*Student   `gorm:"many2many:labo_students; comment:受講生徒;"`
	Rooms      []*Room      `gorm:"many2many:labo_rooms; comment:教室;"`
	Desks      []*Desk      `gorm:"comment:研究室のテーブル;"`
	Chairs     []*Chair     `gorm:"comment:研究室の椅子;"`
	CreatedAt  *time.Time   `gorm:"comment:GORMによって自動的に管理される作成時間;"`
	UpdatedAt  *time.Time   `gorm:"comment:GORMによって自動的に管理される更新時間;"`
}

type Labos struct {
	*irconv.Slice[Labo]
}
