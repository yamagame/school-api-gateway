package iconv

import "fmt"

type ConvsInterface[M any] interface {
	Append(...*M) *SliceWrapper[M]
	SetError(err error) *SliceWrapper[M]
	GetSlice() []*M
	HasError() bool
	GetError() string
}

type ConvInterface[M any, N any] interface {
	ToStruct(*Record) *M
	ToIModel(*M) *Record
	NewSlice(...*M) *N
}

type Convs[M any, N ConvsInterface[M], B ConvInterface[M, N]] struct {
	Conv B
}

func (c Convs[M, N, B]) ToStruct(in *Records) *N {
	r := c.Conv.NewSlice()
	if in.Error != nil {
		(*r).SetError(in.Error)
		return r
	}
	it := in.NewIterator()
	for it.HasNext() {
		v := it.Next()
		t := c.Conv.ToStruct(v)
		(*r).Append(t)
	}
	return r
}

func (c Convs[M, N, B]) ToIModel(in *N) *Records {
	r := &Records{}
	if (*in).HasError() {
		r.Error = fmt.Errorf((*in).GetError())
		return r
	}
	slice := (*in).GetSlice()
	for _, v := range slice {
		t := c.Conv.ToIModel(v)
		if t.HasError() {
			r.Error = t.Error
			return r
		}
		r.Append(t)
	}
	return r
}

type valiables[T any] interface {
	Append(records ...*T) *SliceWrapper[T]
}

func (c Convs[M, N, B]) ToStructWithMap(in []map[string]interface{}, out valiables[M], factory func() *Record) error {
	records := factory().NewRecords(in)
	for _, record := range records.Slice() {
		l := c.Conv.ToStruct(record)
		out.Append(l)
	}
	return nil
}
