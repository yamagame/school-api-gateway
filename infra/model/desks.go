package model

type Desk struct {
	ID          int32 `gorm:"primary; comment:主キーの標準フィールド;"`
	LaboID      int32
	Name        string
	ProductCode string
	Error       error `gorm:"-"`
}

func (m *Desk) HasError() bool {
	return m.Error != nil
}

func (m *Desk) GetError() string {
	if m.HasError() {
		return m.Error.Error()
	}
	return ""
}
