package infconv

import (
	"github.com/yamagame/school-api-gateway/pkg/conv"
)

type ConvInterface[M any] interface {
	ToInfra(*conv.Record) (*M, error)
	ToEntity(*M) (*conv.Record, error)
}

type Convs[M any, B ConvInterface[M]] struct {
	conv B
}

func (c Convs[M, B]) ToInfra(in *conv.Records, err error) ([]*M, error) {
	if err != nil {
		return nil, err
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

func (c Convs[M, B]) ToEntity(in []*M) (*conv.Records, error) {
	if len(in) == 0 {
		return nil, conv.ErrEmptyArray
	}
	r := &conv.Records{}
	for _, v := range in {
		t, err := c.conv.ToEntity(v)
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

func (c Convs[M, B]) ToInfraWithMap(records []map[string]string, out valiables[M], factory func() *conv.Record) error {
	for _, record := range records {
		val, err := conv.NewRecordWithMap(record, factory)
		if err != nil {
			return err
		}
		l, err := c.conv.ToInfra(val)
		if err != nil {
			return err
		}
		out.Append(l)
	}
	return nil
}
