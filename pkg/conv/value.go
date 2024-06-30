package conv

import (
	"encoding/json"
	"reflect"

	kjson "sigs.k8s.io/json"
)

type Value struct {
	key       string
	protected bool
	value     interface{}
	dirty     bool
}

type valueInternal struct {
	Key       string
	Protected bool
	Value     interface{}
	Dirty     bool
}

func NewValue(key string, val interface{}) *Value {
	return &Value{
		key:       key,
		protected: false,
		value:     val,
	}
}

func (x *Value) Copy() *Value {
	v := &Value{
		key:       x.key,
		protected: x.protected,
		value:     x.value,
		dirty:     x.dirty,
	}
	return v
}

func (x *Value) IsChanged() bool {
	return x.dirty
}

func (x *Value) IsProtected() bool {
	return x.protected
}

func (x *Value) Get() interface{} {
	return x.value
}

func (x *Value) Set(val interface{}) error {
	if reflect.TypeOf(x.value) != reflect.TypeOf(val) {
		return ErrInvalidType
	}
	x.dirty = false
	x.value = val
	return nil
}

func (x *Value) Update(val interface{}) error {
	if reflect.TypeOf(x.value) != reflect.TypeOf(val) {
		return ErrInvalidType
	}
	if x.protected {
		return ErrProtectedValue
	}
	x.dirty = true
	x.value = val
	return nil
}

func (x *Value) MarshalJSON() ([]byte, error) {
	return json.Marshal(&valueInternal{
		Key:       x.key,
		Protected: x.protected,
		Value:     x.Get(),
		Dirty:     x.dirty,
	})
}

func (x *Value) UnmarshalJSON(b []byte) error {
	aux := &valueInternal{}
	if err := kjson.UnmarshalCaseSensitivePreserveInts(b, &aux); err != nil {
		return err
	}
	x.key = aux.Key
	x.protected = aux.Protected
	x.value = aux.Value
	x.dirty = aux.Dirty
	return nil
}
