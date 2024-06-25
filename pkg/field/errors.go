package field

import "fmt"

var (
	ErrNotFound  = fmt.Errorf("値が見つかりません")
	ErrProtected = fmt.Errorf("値の変更はできません")
)
