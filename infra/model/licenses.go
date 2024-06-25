package model

type License struct {
	ID       int32  `gorm:"primary; comment:主キーの標準フィールド;"`
	Name     string `gorm:"type:varchar(255); comment:ライセンス名;"`
	PersonID int32  `gorm:"資格所有者の人物ID"`
	Person   Person `gorm:"資格所有者"`
}
