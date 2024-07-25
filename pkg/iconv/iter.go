package iconv

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
