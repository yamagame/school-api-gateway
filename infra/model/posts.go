package model

import "github.com/yamagame/school-api-gateway/pkg/irconv"

type Post struct {
	ID    int32  `gorm:"primary; comment:主キーの標準フィールド;"`
	Name  string `gorm:"comment:部署名;"`
	Error error  `gorm:"-"`
}

// 広報
// 施設
// 人事

type Posts struct {
	*irconv.Slice[Post]
}

func (m *Post) HasError() bool {
	return m.Error != nil
}
