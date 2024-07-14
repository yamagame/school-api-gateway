package model

import "github.com/yamagame/school-api-gateway/pkg/conv"

type Post struct {
	ID   int32  `gorm:"primary; comment:主キーの標準フィールド;"`
	Name string `gorm:"comment:部署名;"`
}

// 広報
// 施設
// 人事

type Posts struct {
	*conv.Slice[Post]
}

func NewPosts(variables ...*Post) *Posts {
	return &Posts{
		Slice: conv.NewSlice(variables...),
	}
}
