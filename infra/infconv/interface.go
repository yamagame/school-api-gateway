package infconv

import "github.com/yamagame/school-api-gateway/pkg/conv"

type ConvInterface[M any] interface {
	ToInfra(*conv.Record) (*M, error)
	ToEntity(*M) (*conv.Record, error)
}

type Convs[M any, B ConvInterface[M]] struct {
	conv B
}

func (c Convs[M, B]) ToInfra(in *conv.Records) ([]*M, error) {
	r := []*M{}
	for _, v := range in.Records() {
		t, err := c.conv.ToInfra(v)
		if err != nil {
			return nil, err
		}
		r = append(r, t)
	}
	return r, nil
}

func (c Convs[M, B]) ToEntity(in []*M) (*conv.Records, error) {
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
