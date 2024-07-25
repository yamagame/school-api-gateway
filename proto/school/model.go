package school

import "github.com/yamagame/school-api-gateway/pkg/iconv"

type Labos struct {
	*iconv.Slice[Labo]
}

func NewLabos(labos ...*Labo) *Labos {
	return &Labos{
		Slice: iconv.NewSlice(labos...),
	}
}
