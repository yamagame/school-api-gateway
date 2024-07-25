package model

import "github.com/yamagame/school-api-gateway/pkg/iconv"

type License struct {
	ID       int32  `gorm:"primary; comment:主キーの標準フィールド;"`
	Name     string `gorm:"type:varchar(255); comment:ライセンス名;"`
	PersonID int32  `gorm:"資格所有者の人物ID"`
	Person   Person `gorm:"資格所有者"`
	Error    error  `gorm:"-"`
}

type Licenses struct {
	*iconv.Slice[License]
}

func (m *License) HasError() bool {
	return m.Error != nil
}

func (m *License) GetError() string {
	if m.HasError() {
		return m.Error.Error()
	}
	return ""
}
