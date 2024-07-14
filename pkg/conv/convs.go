package conv

type ConvInterface[M any, N any] interface {
	ToStruct(*Record) (*M, error)
	ToIRModel(*M) (*Record, error)
	ToArray([]*M) (*N, error)
}

type Convs[M any, N any, B ConvInterface[M, N]] struct {
	Conv B
}

func (c Convs[M, N, B]) ToStruct(in *Records, errs ...error) (*N, error) {
	if len(errs) > 0 && errs[0] != nil {
		return nil, errs[0]
	}
	r := []*M{}
	it := in.NewIterator()
	for it.HasNext() {
		v := it.Next()
		t, err := c.Conv.ToStruct(v)
		if err != nil {
			return nil, err
		}
		r = append(r, t)
	}
	return c.Conv.ToArray(r)
}

func (c Convs[M, N, B]) ToIRModel(in []*M) (*Records, error) {
	if len(in) == 0 {
		return nil, ErrEmptyArray
	}
	r := &Records{}
	for _, v := range in {
		t, err := c.Conv.ToIRModel(v)
		if err != nil {
			return nil, err
		}
		r.Append(t)
	}
	return r, nil
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
