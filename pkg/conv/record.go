package conv

import (
	"encoding/csv"
	"io"
	"reflect"
)

type Record struct {
	Values   map[string]*Value
	HasOnes  map[string]*Record
	HasManys map[string][]*Record
}

func NewRecord() *Record {
	return &Record{
		Values:   map[string]*Value{},
		HasOnes:  map[string]*Record{},
		HasManys: map[string][]*Record{},
	}
}

func (m *Record) SetValue(key string, val interface{}) *Record {
	m.Values[key] = NewValue(key, val)
	return m
}

func (m *Record) SetHasOne(key string, record *Record) *Record {
	m.HasOnes[key] = record
	return m
}

func (m *Record) SetHasMany(key string, records ...*Record) *Record {
	m.HasManys[key] = records
	return m
}

func (m *Record) GetValue(key string) (interface{}, error) {
	if v, ok := m.Values[key]; ok {
		if reflect.TypeOf(v) == reflect.TypeOf(&Value{}) {
			return v.Get(), nil
		}
	}
	return nil, ErrNotFound
}

func (m *Record) GetHasOne(key string) (*Record, error) {
	if v, ok := m.HasOnes[key]; ok {
		if reflect.TypeOf(v) == reflect.TypeOf(&Record{}) {
			return v, nil
		}
	}
	return nil, ErrNotFound
}

func (m *Record) GetHasMany(key string) ([]*Record, error) {
	if v, ok := m.HasManys[key]; ok {
		if reflect.TypeOf(v) == reflect.TypeOf([]*Record{}) {
			return v, nil
		}
	}
	return nil, ErrNotFound
}

func (m *Record) Value(jsonpath string) (*Value, error) {
	values := m.allValues()
	if v, err := GetVal(values, jsonpath); err == nil {
		if reflect.TypeOf(v) == reflect.TypeOf(&Value{}) {
			value := v.(*Value)
			return value, nil
		}
	}
	return nil, ErrNotFound
}

func (m *Record) HasOne(jsonpath string) (*Record, error) {
	values := m.allHasOne()
	if v, err := GetVal(values, jsonpath); err == nil {
		if reflect.TypeOf(v) == reflect.TypeOf(&Record{}) {
			field := v.(*Record)
			return field, nil
		}
	}
	return nil, ErrNotFound
}

func (m *Record) HasMany(jsonpath string) ([]*Record, error) {
	values := m.allHasMany()
	if v, err := GetVal(values, jsonpath); err == nil {
		if reflect.TypeOf(v) == reflect.TypeOf([]*Record{}) {
			field := v.([]*Record)
			return field, nil
		}
	}
	return nil, ErrNotFound
}

func (m *Record) Set(jsonpath string, val interface{}) error {
	if v, err := m.getVal(jsonpath); err == nil {
		v.(*Value).Set(val)
		return nil
	}
	return ErrNotFound
}

func (m *Record) Get(jsonpath string) (interface{}, error) {
	if val, err := m.getVal(jsonpath); err == nil {
		return val.(*Value).Get(), nil
	}
	return nil, ErrNotFound
}

func (m *Record) ValueMap() map[string]interface{} {
	ret := map[string]interface{}{}
	for key, v := range m.Values {
		ret[key] = v.Get()
	}
	for key, v := range m.HasOnes {
		ret[key] = v.ValueMap()
	}
	for key, v := range m.HasManys {
		m := []map[string]interface{}{}
		for _, t := range v {
			m = append(m, t.ValueMap())
		}
		ret[key] = m
	}
	return ret
}

func (m *Record) allValues() map[string]interface{} {
	ret := map[string]interface{}{}
	for key, v := range m.Values {
		ret[key] = v
	}
	for key, v := range m.HasOnes {
		ret[key] = v.allValues()
	}
	for key, v := range m.HasManys {
		m := []map[string]interface{}{}
		for _, t := range v {
			m = append(m, t.allValues())
		}
		ret[key] = m
	}
	return ret
}

func (m *Record) allHasOne() map[string]interface{} {
	ret := map[string]interface{}{}
	for key, v := range m.HasOnes {
		ret[key] = v
	}
	return ret
}

func (m *Record) allHasMany() map[string]interface{} {
	ret := map[string]interface{}{}
	for key, v := range m.HasManys {
		m := []map[string]interface{}{}
		for _, t := range v {
			m = append(m, t.allValues())
		}
		ret[key] = m
	}
	return ret
}

func (m *Record) Copy() *Record {
	r := NewRecord()
	values := m.allValues()
	for key, val := range values {
		switch reflect.TypeOf(val) {
		case reflect.TypeOf(&Value{}):
			r.Values[key] = val.(*Value).Copy()
		case reflect.TypeOf(&Record{}):
			field := val.(*Record).Copy()
			r.SetHasOne(key, field)
		case reflect.TypeOf([]*Record{}):
			fields := []*Record{}
			for _, v := range val.([]*Record) {
				fields = append(fields, v.Copy())
			}
			r.SetHasMany(key, fields...)
		}
	}
	return r
}

func (m *Record) getVal(template string) (interface{}, error) {
	data := m.allValues()
	return GetVal(data, template)
}

func (m *Record) setVal(template string, val interface{}) error {
	data := m.allValues()
	return SetVal(data, template, val)
}

func (m *Record) ToStruct(src, dst string, data interface{}, conv func(v interface{}) interface{}) error {
	if value, err := m.Value(src); err == nil {
		return SetVal(data, dst, conv(value.Get()))
	}
	return ErrNotFound
}

func (m *Record) FromStruct(src, dst string, data interface{}, conv func(v interface{}) interface{}) error {
	if v, err := GetVal(data, src); err == nil {
		return m.Set(dst, conv(v))
	}
	return ErrNotFound
}

type Fields []*Record

func (f *Fields) ValueMap() []map[string]interface{} {
	r := []map[string]interface{}{}
	for _, v := range *f {
		r = append(r, v.ValueMap())
	}
	return r
}

func ReadCSV(r io.Reader, factory func(id int32) *Record) (Fields, error) {
	reader := csv.NewReader(r)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	fields := Fields{}
	header := records[0]
	for _, record := range records[1:] {
		field := map[string]interface{}{}
		for i, column := range header {
			field[column] = record[i]
		}
		newone, err := NewFieldWithMap(field, factory)
		if err != nil {
			return nil, err
		}
		fields = append(fields, newone)
	}
	return fields, nil
}

func NewFieldWithMap(field map[string]interface{}, factory func(id int32) *Record) (*Record, error) {
	newone := factory(0)
	for key, val := range field {
		newone.Set(key, val)
	}
	return newone, nil
}