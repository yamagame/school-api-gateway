package model

import "github.com/yamagame/school-api-gateway/pkg/irconv"

type Chair struct {
	ID     int32 `gorm:"primary; comment:主キーの標準フィールド;"`
	LaboID int32
	Error  error `gorm:"-"`
}

type Chairs struct {
	*irconv.Slice[Chair]
}

func (m *Chair) HasError() bool {
	return m.Error != nil
}
