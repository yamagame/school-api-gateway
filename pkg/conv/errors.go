package conv

import "fmt"

var (
	ErrNotFound    = fmt.Errorf("値が見つかりません")
	ErrProtected   = fmt.Errorf("値の変更はできません")
	ErrInvalidType = fmt.Errorf("値の型が異なります")
)
