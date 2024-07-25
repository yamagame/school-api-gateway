package iconv

import "errors"

type ConvsInterface[M any] interface {
	Append(...*M) *Slice[M]
	SetError(err error) *Slice[M]
}

type ConvInterface[M any, N any] interface {
	ToStruct(*Record) *M
	ToIModel(*M) *Record
	NewSlice() *N
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

func (c Convs[M, N, B]) ToIModel(in []*M, err ...error) *Records {
	r := &Records{}
	if len(err) > 0 && err[0] != nil {
		r.Error = errors.Join(err...)
		return r
	}
	for _, v := range in {
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
	Append(records ...*T) *Slice[T]
}

func (c Convs[M, N, B]) ToStructWithMap(in []map[string]interface{}, out valiables[M], factory func() *Record) error {
	records := factory().NewRecords(in)
	for _, record := range records.Records() {
		l := c.Conv.ToStruct(record)
		out.Append(l)
	}
	return nil
}
