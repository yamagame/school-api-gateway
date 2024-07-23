package model

type Chair struct {
	ID     int32 `gorm:"primary; comment:主キーの標準フィールド;"`
	LaboID int32
	Error  error `gorm:"-"`
}

func (m *Chair) HasError() bool {
	return m.Error != nil
}
