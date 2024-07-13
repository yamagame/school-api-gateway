package infconv

import (
	"github.com/yamagame/school-api-gateway/pkg/conv"
)

type ConvInterface[M any] interface {
	ToInfra(*conv.Record) (*M, error)
	ToIRModel(*M) (*conv.Record, error)
}

type Convs[M any, B ConvInterface[M]] struct {
	conv B
}

func (c Convs[M, B]) ToInfra(in *conv.Records, errs ...error) ([]*M, error) {
	if len(errs) > 0 && errs[0] != nil {
		return nil, errs[0]
	}
	r := []*M{}
	it := in.NewIterator()
	for it.HasNext() {
		v := it.Next()
		t, err := c.conv.ToInfra(v)
		if err != nil {
			return nil, err
		}
		r = append(r, t)
	}
	return r, nil
}

func (c Convs[M, B]) ToIRModel(in []*M) (*conv.Records, error) {
	if len(in) == 0 {
		return nil, conv.ErrEmptyArray
	}
	r := &conv.Records{}
	for _, v := range in {
		t, err := c.conv.ToIRModel(v)
		if err != nil {
			return nil, err
		}
		r.Append(t)
	}
	return r, nil
}

type valiables[T any] interface {
	Append(records ...*T) *conv.Variables[T]
}

func (c Convs[M, B]) ToInfraWithMap(in []map[string]interface{}, out valiables[M], factory func() *conv.Record) error {
	records := factory().NewRecords(in)
	for _, record := range records.Records() {
		l, err := c.conv.ToInfra(record)
		if err != nil {
			return err
		}
		out.Append(l)
	}
	return nil
}
