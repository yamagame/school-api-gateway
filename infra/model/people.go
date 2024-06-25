package model

import "time"

type Person struct {
	ID        int32      `gorm:"primary; comment:主キーの標準フィールド;"`
	Name      string     `gorm:"type:varchar(255); comment:名前;"`
	Birthday  time.Time  `gorm:"type:date; comment:生年月日;"`
	RoleID    int32      `gorm:"comment:役割ID;"`
	Role      Role       `gorm:"comment:役割;"`
	AddressID int32      `gorm:"comment::住所ID;"`
	Address   Address    `gorm:"comment::住所;"`
	Licenses  []License  `gorm:"comment::資格;"`
	CreatedAt *time.Time `gorm:"comment:GORMによって自動的に管理される作成時間;"`
	UpdatedAt *time.Time `gorm:"comment:GORMによって自動的に管理される更新時間;"`
}
