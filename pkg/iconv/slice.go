package iconv

type SliceContainer[T any] struct {
	Error   error
	records []*T
}

func NewSlice[T any](records ...*T) *SliceContainer[T] {
	return &SliceContainer[T]{
		records: records,
	}
}

func (v *SliceContainer[T]) SetError(err error) *SliceContainer[T] {
	v.Error = err
	return v
}

func (v *SliceContainer[T]) HasError() bool {
	return v.Error != nil
}

func (v *SliceContainer[T]) GetError() string {
	if v.HasError() {
		return v.Error.Error()
	}
	return ""
}

func (v *SliceContainer[T]) Slice() []*T {
	return v.MustShallowCopy()
}

func (v *SliceContainer[T]) ShallowCopy() ([]*T, error) {
	if v.HasError() {
		return []*T{}, v.Error
	}
	r := make([]*T, len(v.records))
	for i, v := range v.records {
		r[i] = v
	}
	return r, nil
}

func (v *SliceContainer[T]) MustShallowCopy() []*T {
	r, err := v.ShallowCopy()
	if err != nil {
		panic(err)
	}
	return r
}

func (v *SliceContainer[T]) Append(records ...*T) *SliceContainer[T] {
	if v.HasError() {
		return v
	}
	v.records = append(v.records, records...)
	return v
}

func (v *SliceContainer[T]) IndexOf(idx int) *T {
	if idx >= 0 && idx < v.Length() {
		return v.records[idx]
	}
	return nil
}

func (v *SliceContainer[T]) NewIterator() *Iterator[T] {
	return NewIterator(v.records)
}

func (v *SliceContainer[T]) Length() int {
	return len(v.records)
}

func (v *SliceContainer[T]) First() *T {
	if len(v.records) > 0 {
		return v.records[0]
	}
	return nil
}

func (v *SliceContainer[T]) Last() *T {
	if len(v.records) > 0 {
		return v.records[len(v.records)-1]
	}
	return nil
}
