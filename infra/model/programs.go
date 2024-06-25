package model

import "time"

type Program struct {
	ID        int32      `gorm:"primary; comment:主キーの標準フィールド;"`
	Name      string     `gorm:"type:varchar(255); comment:プログラム名;"`
	CreatedAt *time.Time `gorm:"comment:GORMによって自動的に管理される作成時間;"`
	UpdatedAt *time.Time `gorm:"comment:GORMによって自動的に管理される更新時間;"`
}

// 01.メディア情報学プログラム
// 02.経営・社会情報学プログラム
// 03.情報数理工学プログラム
// 04.コンピュータサイエンスプログラム
// 05.デザイン思考・データサイエンスプログラム
// 06.セキュリティ情報学プログラム
// 07.情報通信工学プログラム
// 08.電子情報学プログラム
// 09.計測・制御システムプログラム
// 10.先端ロボティクスプログラム
// 11.機械システムプログラム
// 12.電子工学プログラム
// 13.光工学プログラム
// 14.物理工学プログラム
// 15.化学生命工学プログラム
