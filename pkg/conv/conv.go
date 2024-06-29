package conv

func ToPtr[T any](v T) *T {
	a := v
	return &a
}

func Row[T any](v T) T {
	return v
}

func StrPtr(v interface{}) interface{} {
	return ToPtr(v.(string))
}

func PtrStr(v interface{}) interface{} {
	return *v.(*string)
}
