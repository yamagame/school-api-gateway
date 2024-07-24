package school

import "github.com/yamagame/school-api-gateway/pkg/irconv"

type Labos struct {
	*irconv.Slice[Labo]
}

func NewLabos(labos ...*Labo) *Labos {
	return &Labos{
		Slice: irconv.NewSlice(labos...),
	}
}
