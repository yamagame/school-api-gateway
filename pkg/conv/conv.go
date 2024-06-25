package conv

func ToPtr[T any](v T) *T {
	a := v
	return &a
}
