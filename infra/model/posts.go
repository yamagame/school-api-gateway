package model

type Post struct {
	ID   int32  `gorm:"primary; comment:主キーの標準フィールド;"`
	Name string `gorm:"comment:部署名;"`
}

// 広報
// 施設
// 人事
