package school

import "github.com/yamagame/school-api-gateway/pkg/iconv"

type Labos struct {
	*iconv.SliceContainer[Labo]
}

func NewLabos(labos ...*Labo) *Labos {
	return &Labos{
		SliceContainer: iconv.NewSlice(labos...),
	}
}
