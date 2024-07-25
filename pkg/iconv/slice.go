package iconv

type SliceWrapper[T any] struct {
	Error   error
	records []*T
}

func NewSlice[T any](records ...*T) *SliceWrapper[T] {
	return &SliceWrapper[T]{
		records: records,
	}
}

func (v *SliceWrapper[T]) SetError(err error) *SliceWrapper[T] {
	v.Error = err
	return v
}

func (v *SliceWrapper[T]) HasError() bool {
	return v.Error != nil
}

func (v *SliceWrapper[T]) GetError() string {
	if v.HasError() {
		return v.Error.Error()
	}
	return ""
}

func (v *SliceWrapper[T]) Slice() []*T {
	return v.MustShallowCopy()
}

func (v *SliceWrapper[T]) ShallowCopy() ([]*T, error) {
	if v.HasError() {
		return []*T{}, v.Error
	}
	r := make([]*T, len(v.records))
	for i, v := range v.records {
		r[i] = v
	}
	return r, nil
}

func (v *SliceWrapper[T]) MustShallowCopy() []*T {
	r, err := v.ShallowCopy()
	if err != nil {
		panic(err)
	}
	return r
}

func (v *SliceWrapper[T]) Append(records ...*T) *SliceWrapper[T] {
	if v.HasError() {
		return v
	}
	v.records = append(v.records, records...)
	return v
}

func (v *SliceWrapper[T]) IndexOf(idx int) *T {
	if idx >= 0 && idx < v.Length() {
		return v.records[idx]
	}
	return nil
}

func (v *SliceWrapper[T]) NewIterator() *Iterator[T] {
	return NewIterator(v.records)
}

func (v *SliceWrapper[T]) Length() int {
	return len(v.records)
}

func (v *SliceWrapper[T]) First() *T {
	if len(v.records) > 0 {
		return v.records[0]
	}
	return nil
}

func (v *SliceWrapper[T]) Last() *T {
	if len(v.records) > 0 {
		return v.records[len(v.records)-1]
	}
	return nil
}
