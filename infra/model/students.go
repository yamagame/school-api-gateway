package model

import "time"

type Student struct {
	ID        int32      `gorm:"primary; comment:主キーの標準フィールド;"`
	PersonID  int32      `gorm:"comment:人物ID;"`
	Person    Person     `gorm:"comment:人物;"`
	CourseID  int32      `gorm:"comment:コースID;"`
	Course    Course     `gorm:"comment:コース;"`
	GroupID   int32      `gorm:"comment:専門領域ID;"`
	Group     Group      `gorm:"comment:専門領域;"`
	Degree    int32      `gorm:"comment:学位;"`
	Grade     int32      `gorm:"comment:学年;"`
	CreatedAt *time.Time `gorm:"comment:GORMによって自動的に管理される作成時間;"`
	UpdatedAt *time.Time `gorm:"comment:GORMによって自動的に管理される更新時間;"`
}
