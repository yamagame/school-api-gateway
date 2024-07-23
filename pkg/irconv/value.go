package irconv

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	kjson "sigs.k8s.io/json"
)

const (
	PRIMARY = "PRIMARY"
)

type Value struct {
	key       string
	protected bool
	value     interface{}
	exist     bool
	synced    bool
	primary   bool
	Error     error
}

type valueInternal struct {
	Key       string
	Protected bool
	Value     interface{}
	Exist     bool
	Synced    bool
	Primary   bool
}

func NewValue(key string, val interface{}, options ...string) *Value {
	return &Value{
		key:       key,
		protected: false,
		value:     val,
		primary:   Exist(options, PRIMARY),
	}
}

func (x *Value) Copy() *Value {
	v := &Value{
		key:       x.key,
		protected: x.protected,
		value:     x.value,
		exist:     x.exist,
		synced:    x.synced,
		primary:   x.primary,
	}
	return v
}

func (x *Value) IsPrimary() bool {
	return x.primary
}

func (x *Value) IsExist() bool {
	return x.exist
}

func (x *Value) IsSynced() bool {
	return x.synced
}

func (x *Value) IsProtected() bool {
	return x.protected
}

func (x *Value) Get() interface{} {
	return x.value
}

func (x *Value) Set(val interface{}) error {
	if reflect.TypeOf(x.value) != reflect.TypeOf(val) {
		fmt.Printf("%v\n", x)
		return errors.Join(ErrInvalidType, fmt.Errorf("Set val: %v", val))
	}
	x.exist = true
	x.synced = true
	x.value = val
	return nil
}

func (x *Value) Update(val interface{}) error {
	if reflect.TypeOf(x.value) != reflect.TypeOf(val) {
		fmt.Printf("%v\n", x)
		return errors.Join(ErrInvalidType, fmt.Errorf("Update val: %v", val))
	}
	if x.protected {
		return ErrProtectedValue
	}
	x.exist = true
	x.synced = false
	x.value = val
	return nil
}

func (x *Value) MarshalJSON() ([]byte, error) {
	return json.Marshal(&valueInternal{
		Key:       x.key,
		Protected: x.protected,
		Value:     x.Get(),
		Exist:     x.exist,
		Synced:    x.synced,
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
	x.exist = aux.Exist
	x.synced = aux.Synced
	return nil
}

func (x *Value) HasError() bool {
	return x.Error != nil
}
