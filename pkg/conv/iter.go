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

type Variables[T any] struct {
	Error   error
	Records []*T
}

func NewVariables[T any](records ...*T) *Variables[T] {
	return &Variables[T]{
		Records: records,
	}
}

func (v *Variables[T]) Append(records ...*T) *Variables[T] {
	if v.Error != nil {
		return v
	}
	v.Records = append(v.Records, records...)
	return v
}

func (v *Variables[T]) NewIterator() *Iterator[T] {
	return NewIterator(v.Records)
}

func (v *Variables[T]) Length() int {
	return len(v.Records)
}

func (v *Variables[T]) First() *T {
	if len(v.Records) > 0 {
		return v.Records[0]
	}
	return nil
}
