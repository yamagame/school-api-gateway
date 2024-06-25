package field

import "strings"

type Key struct {
	Fields []string
}

func NewKey(key string) *Key {
	return &Key{
		Fields: strings.Split(key, "."),
	}
}

func (x *Key) HasRelation() bool {
	return len(x.Fields) > 1
}

func (x *Key) Last() string {
	return x.Fields[len(x.Fields)-1]
}
