package field

type Value struct {
	Key      string
	Writable bool
	Original interface{}
	Value    interface{}
}

func NewValue(key string, orgval interface{}) *Value {
	return &Value{
		Key:      key,
		Writable: false,
		Original: orgval,
	}
}

func (x *Value) IsExist() bool {
	return x.Value != nil
}

func (x *Value) IsChanged() bool {
	if x.Value != nil {
		return x.Original != x.Value
	}
	return false
}

func (x *Value) IsWritable() bool {
	return x.Writable
}

func (x *Value) Get() interface{} {
	if x.Value == nil {
		return x.Original
	}
	return x.Value
}

func (x *Value) Set(val interface{}) {
	x.Value = val
}

func (x *Value) GetOriginal() interface{} {
	return x.Original
}

func (x *Value) SetOriginal(val interface{}) {
	x.Original = val
}
