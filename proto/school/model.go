package school

import "github.com/yamagame/school-api-gateway/pkg/iconv"

type Labos struct {
	*iconv.SliceWrapper[Labo]
}

func NewLabos(labos ...*Labo) *Labos {
	return &Labos{
		SliceWrapper: iconv.NewSlice(labos...),
	}
}
