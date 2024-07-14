package irconv

type Slice[T any] struct {
	Error   error
	records []*T
}

func NewSlice[T any](records ...*T) *Slice[T] {
	return &Slice[T]{
		records: records,
	}
}

func (v *Slice[T]) SetError(err error) *Slice[T] {
	v.Error = err
	return v
}

func (v *Slice[T]) HasError() bool {
	return v.Error != nil
}

func (v *Slice[T]) ShallowCopy() []*T {
	r := make([]*T, len(v.records))
	for i, v := range v.records {
		r[i] = v
	}
	return r
}

func (v *Slice[T]) Append(records ...*T) *Slice[T] {
	if v.Error != nil {
		return v
	}
	v.records = append(v.records, records...)
	return v
}

func (v *Slice[T]) IndexOf(idx int) *T {
	if idx >= 0 && idx < v.Length() {
		return v.records[idx]
	}
	return nil
}

func (v *Slice[T]) NewIterator() *Iterator[T] {
	return NewIterator(v.records)
}

func (v *Slice[T]) Length() int {
	return len(v.records)
}

func (v *Slice[T]) First() *T {
	if len(v.records) > 0 {
		return v.records[0]
	}
	return nil
}

func (v *Slice[T]) Last() *T {
	if len(v.records) > 0 {
		return v.records[len(v.records)-1]
	}
	return nil
}
