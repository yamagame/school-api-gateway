package entity

type Value struct {
	Key     string
	Default interface{}
	Value   interface{}
}

func (x *Value) IsExist() bool {
	return x.Value != nil
}

func (x *Value) Get() interface{} {
	if x.Value == nil {
		return x.Default
	}
	return x.Value
}

func (x *Value) Set(val interface{}) {
	x.Value = val
}
