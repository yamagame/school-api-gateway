package conv

type Iterator[T any] struct {
	records []*T
	index   int
}

func NewIterator[T any](records []*T) *Iterator[T] {
	return &Iterator[T]{
		records: records,
		index:   0,
	}
}

func (i *Iterator[T]) HasNext() bool {
	if i.index < len(i.records) {
		return true
	}
	return false
}

func (i *Iterator[T]) Next() *T {
	r := i.records[i.index]
	i.index++
	return r
}

type Slice[T any] struct {
	Error   error
	records []*T
}

func NewSlice[T any](records ...*T) *Slice[T] {
	return &Slice[T]{
		records: records,
	}
}

func (v *Slice[T]) Copy() []*T {
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
