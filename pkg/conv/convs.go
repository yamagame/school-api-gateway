package conv

type ConvInterface[M any] interface {
	ToStruct(*Record) (*M, error)
	ToIRModel(*M) (*Record, error)
}

type Convs[M any, B ConvInterface[M]] struct {
	Conv B
}

func (c Convs[M, B]) ToStruct(in *Records, errs ...error) ([]*M, error) {
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
	return r, nil
}

func (c Convs[M, B]) ToIRModel(in []*M) (*Records, error) {
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
	Append(records ...*T) *Variables[T]
}

func (c Convs[M, B]) ToStructWithMap(in []map[string]interface{}, out valiables[M], factory func() *Record) error {
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
