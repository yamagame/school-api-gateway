package irconv

type ConvsInterface[M any] interface {
	Append(...*M) *Slice[M]
	SetError(err error) *Slice[M]
}

type ConvInterface[M any, N any] interface {
	ToStruct(*Record) (*M, error)
	ToIRModel(*M) (*Record, error)
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
		t, err := c.Conv.ToStruct(v)
		if err != nil {
			(*r).SetError(err)
			return r
		}
		(*r).Append(t)
	}
	return r
}

func (c Convs[M, N, B]) ToIRModel(in []*M) *Records {
	r := &Records{}
	for _, v := range in {
		t, err := c.Conv.ToIRModel(v)
		if err != nil {
			r.Error = err
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
		l, err := c.Conv.ToStruct(record)
		if err != nil {
			return err
		}
		out.Append(l)
	}
	return nil
}
