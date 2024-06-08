package repository

import "time"

type Labos struct {
	ID        uint      // 主キーの標準フィールド
	Name      string    // 研究室の名前
	CreatedAt time.Time // GORMによって自動的に管理される作成時間
	UpdatedAt time.Time // GORMによって自動的に管理される更新時間
}
