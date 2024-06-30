package conv

import (
	"encoding/csv"
	"io"
	"reflect"
	"strconv"
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

func strTo(v interface{}, s string) interface{} {
	switch reflect.ValueOf(v).Kind() {
	case reflect.Bool:
		v, _ := strconv.ParseBool(s)
		return v
	case reflect.Int:
		v, _ := strconv.Atoi(s)
		return v
	case reflect.Int64:
		v, _ := strconv.ParseInt(s, 10, 64)
		return int64(v)
	case reflect.Int32:
		v, _ := strconv.ParseInt(s, 10, 32)
		return int32(v)
	case reflect.Int16:
		v, _ := strconv.ParseInt(s, 10, 16)
		return int16(v)
	case reflect.Int8:
		v, _ := strconv.ParseInt(s, 10, 8)
		return int8(v)
	case reflect.Uint:
		v, _ := strconv.ParseUint(s, 10, 32)
		return uint(v)
	case reflect.Uint64:
		v, _ := strconv.ParseUint(s, 10, 64)
		return uint64(v)
	case reflect.Uint32:
		v, _ := strconv.ParseUint(s, 10, 32)
		return uint32(v)
	case reflect.Uint16:
		v, _ := strconv.ParseUint(s, 10, 16)
		return uint16(v)
	case reflect.Uint8:
		v, _ := strconv.ParseUint(s, 10, 8)
		return uint8(v)
	case reflect.Float32:
		v, _ := strconv.ParseFloat(s, 32)
		return float32(v)
	case reflect.Float64:
		v, _ := strconv.ParseFloat(s, 64)
		return float64(v)
	case reflect.String:
		return s
	default:
	}
	return v
}

func NewRecordWithMap(field map[string]string, factory func() *Record) (*Record, error) {
	newone := factory()
	for key, val := range field {
		v, err := newone.Get(key)
		if err != nil {
			return nil, err
		}
		if err := newone.Set(key, strTo(v, val)); err != nil {
			return nil, err
		}
	}
	return newone, nil
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
		return v.(*Value).Set(val)
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
		value := reflect.ValueOf(v)
		if value.Kind() == reflect.Ptr && value.IsNil() {
			return nil
		}
		t := conv(v)
		return m.Set(dst, t)
	}
	return ErrNotFound
}

type Records []*Record

func (f *Records) ValueMap() []map[string]interface{} {
	r := []map[string]interface{}{}
	for _, v := range *f {
		r = append(r, v.ValueMap())
	}
	return r
}

func ReadCSV(r io.Reader) ([]map[string]string, error) {
	reader := csv.NewReader(r)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	ret := []map[string]string{}
	header := records[0]
	for _, record := range records[1:] {
		field := map[string]string{}
		for i, column := range header {
			field[column] = record[i]
		}
		ret = append(ret, field)
	}
	return ret, nil
}

// func ReadCSV(r io.Reader, factory func() *Record) (Fields, error) {
// 	reader := csv.NewReader(r)
// 	records, err := reader.ReadAll()
// 	if err != nil {
// 		return nil, err
// 	}
// 	fields := Fields{}
// 	header := records[0]
// 	for _, record := range records[1:] {
// 		field := map[string]interface{}{}
// 		for i, column := range header {
// 			field[column] = record[i]
// 		}
// 		newone, err := NewRecordWithMap(field, factory)
// 		if err != nil {
// 			return nil, err
// 		}
// 		fields = append(fields, newone)
// 	}
// 	return fields, nil
// }
