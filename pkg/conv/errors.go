package conv

import "fmt"

var (
	ErrNotFound        = fmt.Errorf("値が見つかりません")
	ErrProtectedValue  = fmt.Errorf("値の変更はできません")
	ErrInvalidType     = fmt.Errorf("値の型が異なります")
	ErrNotStruct       = fmt.Errorf("構造体ではありません")
	ErrDifferentStruct = fmt.Errorf("構造体が異なります")
)
