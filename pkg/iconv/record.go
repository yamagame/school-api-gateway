package iconv

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

type Record struct {
	Error     error
	Values    map[string]*Value
	BelongTos map[string]*Record
	HasOnes   map[string]*Record
	HasManys  map[string]*HasMany
}

func NewRecord() *Record {
	return &Record{
		Values:    map[string]*Value{},
		BelongTos: map[string]*Record{},
		HasOnes:   map[string]*Record{},
		HasManys:  map[string]*HasMany{},
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

func NewRecordWithMap(field map[string]interface{}, factory func() *Record) (*Record, error) {
	record := factory()
	for key, val := range field {
		v, err := record.Get(key)
		if err != nil {
			return nil, err
		}
		if val == "-" {
			continue
		}
		if s, ok := val.(string); ok {
			if err := record.Set(key, strTo(v, s)).Error; err != nil {
				return nil, err
			}
		} else if a, ok := val.([]map[string]interface{}); ok {
			many, err := record.HasMany(key)
			if err != nil {
				return nil, err
			}
			factory := many.Model.Func()
			for _, v := range a {
				r, err := NewRecordWithMap(v, factory)
				if err != nil {
					return nil, err
				}
				many.Append(r)
			}
		}
	}
	return record, nil
}

func (m *Record) PrimaryKeys() []string {
	r := []string{}
	for k, v := range m.Values {
		if v.IsPrimary() {
			r = append(r, k)
		}
	}
	return r
}

func (m *Record) SetValue(key string, val interface{}, options ...string) *Record {
	if m.Error != nil {
		return m
	}
	m.Values[key] = NewValue(key, val, options...)
	return m
}

func (m *Record) SetBelongTo(key string, record *Record) *Record {
	if m.Error != nil {
		return m
	}
	m.BelongTos[key] = record
	return m
}

func (m *Record) SetHasOne(key string, record *Record) *Record {
	if m.Error != nil {
		return m
	}
	m.HasOnes[key] = record
	return m
}

func (m *Record) SetHasMany(key string, many *HasMany) *Record {
	if m.Error != nil {
		return m
	}
	m.HasManys[key] = many.Copy()
	return m
}

// func (m *Record) SetHasManyRecords(key string, records *Records) *Record {
// 	if m.Error != nil {
// 		return m
// 	}
// 	if v, ok := m.HasManys[key]; ok {
// 		v.Append(records.Records()...)
// 	}
// 	return m
// }

func (m *Record) GetValue(key string) *Value {
	if v, ok := m.Values[key]; ok {
		if reflect.TypeOf(v) == reflect.TypeOf(&Value{}) {
			return v
		}
	}
	return &Value{
		Error: errors.Join(ErrNotFound, fmt.Errorf("GetValue key: %v", key)),
	}
}

func (m *Record) GetBelongTo(key string) *Record {
	if m.HasError() {
		return m
	}
	if v, ok := m.BelongTos[key]; ok {
		if reflect.TypeOf(v) == reflect.TypeOf(&Record{}) {
			return v
		}
	}
	m.Error = errors.Join(ErrNotFound, fmt.Errorf("GetBelongTo key: %v", key))
	return m
}

func (m *Record) GetHasOne(key string) *Record {
	if m.HasError() {
		return m
	}
	if v, ok := m.HasOnes[key]; ok {
		if reflect.TypeOf(v) == reflect.TypeOf(&Record{}) {
			return v
		}
	}
	m.Error = errors.Join(ErrNotFound, fmt.Errorf("GetHasOne key: %v", key))
	return m
}

func (m *Record) GetHasMany(key string) *HasMany {
	if m.HasError() {
		return &HasMany{
			Error: m.Error,
		}
	}
	if v, ok := m.HasManys[key]; ok {
		if reflect.TypeOf(v) == reflect.TypeOf(&HasMany{}) {
			return v
		}
	}
	return &HasMany{
		Error: errors.Join(ErrNotFound, fmt.Errorf("GetHasMany key: %v", key)),
	}
}

// func (m *Record) GetHasManyRecords(key string) *Records {
// 	if m.HasError() {
// 		return &Records{
// 			Error: m.Error,
// 		}
// 	}
// 	if v, ok := m.HasManys[key]; ok {
// 		if reflect.TypeOf(v) == reflect.TypeOf(&HasMany{}) {
// 			r := NewRecords(v.Model.Copy())
// 			r.Append(v.records...)
// 			return r
// 		}
// 	}
// 	return &Records{
// 		Error: errors.Join(ErrNotFound, fmt.Errorf("GetHasManyRecords key: %v", key)),
// 	}
// }

func (m *Record) Value(jsonpath string) (*Value, error) {
	values := m.allValues()
	if v, err := GetVal(values, jsonpath); err == nil {
		if reflect.TypeOf(v) == reflect.TypeOf(&Value{}) {
			value := v.(*Value)
			return value, nil
		}
	}
	return nil, errors.Join(ErrNotFound, fmt.Errorf("Value jsonpath: %v", jsonpath))
}

func (m *Record) HasOne(jsonpath string) (*Record, error) {
	values := m.allHasOne()
	if v, err := GetVal(values, jsonpath); err == nil {
		if reflect.TypeOf(v) == reflect.TypeOf(&Record{}) {
			field := v.(*Record)
			return field, nil
		}
	}
	return nil, errors.Join(ErrNotFound, fmt.Errorf("HasOne jsonpath: %v", jsonpath))
}

func (m *Record) HasMany(jsonpath string) (*HasMany, error) {
	values := m.allMany()
	if v, err := GetVal(values, jsonpath); err == nil {
		if reflect.TypeOf(v) == reflect.TypeOf(&HasMany{}) {
			field := v.(*HasMany)
			return field, nil
		}
	}
	return nil, errors.Join(ErrNotFound, fmt.Errorf("HasMany jsonpath: %v", jsonpath))
}

func (m *Record) Set(jsonpath string, val interface{}) *Record {
	if v, err := m.getVal(jsonpath); err == nil {
		m.Error = v.(*Value).Set(val)
	} else {
		m.Error = errors.Join(ErrNotFound, fmt.Errorf("Set jsonpath: %v", jsonpath))
	}
	return m
}

func (m *Record) Update(jsonpath string, val interface{}) *Record {
	if v, err := m.getVal(jsonpath); err == nil {
		m.Error = v.(*Value).Update(val)
	} else {
		m.Error = errors.Join(ErrNotFound, fmt.Errorf("Update jsonpath: %v", jsonpath))
	}
	return m
}

func (m *Record) Get(jsonpath string) (interface{}, error) {
	if v, err := m.getVal(jsonpath); err == nil {
		if reflect.TypeOf(v) == reflect.TypeOf(&Value{}) {
			return v.(*Value).Get(), nil
		}
		return v, nil
	}
	return nil, errors.Join(ErrNotFound, fmt.Errorf("Get jsonpath: %v", jsonpath))
}

func (m *Record) ValueMap() map[string]interface{} {
	ret := map[string]interface{}{}
	for key, v := range m.Values {
		ret[key] = v.Get()
	}
	for key, v := range m.BelongTos {
		ret[key] = v.ValueMap()
	}
	for key, v := range m.HasOnes {
		ret[key] = v.ValueMap()
	}
	for key, v := range m.HasManys {
		m := []map[string]interface{}{}
		it := v.NewIterator()
		for it.HasNext() {
			t := it.Next()
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
	for key, v := range m.BelongTos {
		ret[key] = v.allValues()
	}
	for key, v := range m.HasOnes {
		ret[key] = v.allValues()
	}
	for key, v := range m.HasManys {
		m := []map[string]interface{}{}
		it := v.NewIterator()
		for it.HasNext() {
			t := it.Next()
			m = append(m, t.allValues())
		}
		ret[key] = m
	}
	return ret
}

func (m *Record) allHasOne() map[string]interface{} {
	ret := map[string]interface{}{}
	for key, v := range m.Values {
		ret[key] = v
	}
	for key, v := range m.BelongTos {
		ret[key] = v
	}
	for key, v := range m.HasOnes {
		ret[key] = v
	}
	for key, v := range m.HasManys {
		m := []map[string]interface{}{}
		it := v.NewIterator()
		for it.HasNext() {
			t := it.Next()
			m = append(m, t.allHasOne())
		}
		ret[key] = m
	}
	return ret
}

func (m *Record) allMany() map[string]interface{} {
	ret := map[string]interface{}{}
	for key, v := range m.Values {
		ret[key] = v
	}
	for key, v := range m.BelongTos {
		ret[key] = v.allMany()
	}
	for key, v := range m.HasOnes {
		ret[key] = v.allMany()
	}
	for key, v := range m.HasManys {
		ret[key] = v
	}
	return ret
}

func (m *Record) Copy() *Record {
	r := NewRecord()
	if m.Error != nil {
		r.Error = m.Error
		return r
	}
	for key, v := range m.Values {
		r.Values[key] = v.Copy()
	}
	for key, v := range m.BelongTos {
		r.BelongTos[key] = v.Copy()
	}
	for key, v := range m.HasOnes {
		r.HasOnes[key] = v.Copy()
	}
	for key, v := range m.HasManys {
		r.HasManys[key] = v.Copy()
	}
	return r
}

func (m *Record) NewRecords(records []map[string]interface{}) *Records {
	work := NewRecords(m.Copy())
	if m.Error != nil {
		work.Error = m.Error
		return work
	}
	for _, record := range records {
		r, err := NewRecordWithMap(record, func() *Record {
			return m.Copy()
		})
		if err != nil {
			work.Error = err
			return work
		}
		work.Append(r)
	}
	return work
}

func (m *Record) Func() func() *Record {
	return func() *Record {
		return m.Copy()
	}
}

func (m *Record) Self() *Record {
	return m
}

func (m *Record) getVal(template string) (interface{}, error) {
	data := m.allValues()
	return GetVal(data, template)
}

func (m *Record) getMany(template string) (interface{}, error) {
	data := m.allMany()
	return GetVal(data, template)
}

func (m *Record) setVal(template string, val interface{}) error {
	data := m.allValues()
	return SetVal(data, template, val)
}

func (m *Record) ToStruct(src, dst string, out interface{}, convs ...func(v interface{}, in ...interface{}) interface{}) *Record {
	if m.Error != nil {
		return m
	}
	if value, err := m.Value(src); err == nil {
		if value.IsExist() {
			v := value.Get()
			for _, conv := range convs {
				v = conv(v, m)
			}
			m.Error = SetVal(out, dst, v)
		}
		return m
	}
	m.Error = errors.Join(ErrNotFound, fmt.Errorf("ToStruct src: %s, dst: %s", src, dst))
	return m
}

func (m *Record) FromStruct(src, dst string, in interface{}, convs ...func(v interface{}, in ...interface{}) interface{}) *Record {
	if m.Error != nil {
		return m
	}
	if v, err := GetVal(in, src); err == nil {
		value := reflect.ValueOf(v)
		if value.Kind() == reflect.Ptr && value.IsNil() {
			return m
		}
		for _, conv := range convs {
			v = conv(v, in)
		}
		m.Error = m.Set(dst, v).Error
		return m
	}
	m.Error = errors.Join(ErrNotFound, fmt.Errorf("FromStruct src: %s, dst: %s", src, dst))
	return m
}

func (m *Record) IfNotNil(jpath string, in interface{}, cb func(v *Record)) *Record {
	var err error
	if v, err := GetVal(in, jpath); err == nil {
		value := reflect.ValueOf(v)
		if value.Kind() == reflect.Ptr && value.IsNil() {
			return m
		}
		cb(m)
		return m
	}
	m.Error = err
	return m
}

func (m *Record) IfExist(jpath string, cb func(v *Record)) *Record {
	var isExist func(v interface{}) bool
	isExist = func(v interface{}) bool {
		if reflect.TypeOf(v) == reflect.TypeOf(&Record{}) {
			value := v.(*Record)
			if value.IsExist() {
				return true
			}
		} else if reflect.TypeOf(v) == reflect.TypeOf(&HasMany{}) {
			value := v.(*HasMany)
			if value.IsExist() {
				return true
			}
		} else if reflect.TypeOf(v) == reflect.TypeOf(&Value{}) {
			value := v.(*Value)
			if value.IsExist() {
				return true
			}
		} else if reflect.TypeOf(v) == reflect.TypeOf([]map[string]interface{}{}) {
			values := v.([]map[string]interface{})
			for _, q := range values {
				if isExist(q) {
					return true
				}
			}
		} else if reflect.TypeOf(v) == reflect.TypeOf(map[string]interface{}{}) {
			values := v.(map[string]interface{})
			for _, q := range values {
				if isExist(q) {
					return true
				}
			}
		}
		return false
	}
	data := m.allValues()
	if v, err := GetVal(data, jpath); err == nil {
		if isExist(v) {
			cb(m)
		}
	}
	return m
}

func (m *Record) Updates() map[string]interface{} {
	ret1 := m.UpdateValues()
	ret2 := m.UpdateBelongTos()
	ret3 := m.UpdateHasOnes()
	ret4 := m.UpdateHasManyes()
	return MergeMap(ret1, ret2, ret3, ret4)
}

func (m *Record) UpdateValues() map[string]interface{} {
	ret := map[string]interface{}{}
	for key, v := range m.Values {
		if !v.IsSynced() && v.IsExist() {
			ret[key] = v.Get()
		}
	}
	return ret
}

func (m *Record) UpdateBelongTos() map[string]interface{} {
	ret := map[string]interface{}{}
	for key, v := range m.BelongTos {
		t := v.Updates()
		if len(t) > 0 {
			ret[key] = t
		}
	}
	return ret
}

func (m *Record) UpdateHasOnes() map[string]interface{} {
	ret := map[string]interface{}{}
	for key, v := range m.HasOnes {
		t := v.Updates()
		if len(t) > 0 {
			ret[key] = t
		}
	}
	return ret
}

func (m *Record) UpdateHasManyes() map[string]interface{} {
	ret := map[string]interface{}{}
	for key, v := range m.HasManys {
		m := []map[string]interface{}{}
		it := v.NewIterator()
		for it.HasNext() {
			v := it.Next()
			t := v.Updates()
			if len(t) > 0 {
				m = append(m, t)
			}
		}
		if len(m) > 0 {
			ret[key] = m
		}
	}
	return ret
}

func (m *Record) IsExist() bool {
	for _, v := range m.Values {
		if v.IsExist() {
			return true
		}
	}
	for _, v := range m.BelongTos {
		if v.IsExist() {
			return true
		}
	}
	for _, v := range m.HasOnes {
		if v.IsExist() {
			return true
		}
	}
	for _, m := range m.HasManys {
		it := m.NewIterator()
		for it.HasNext() {
			v := it.Next()
			if v.IsExist() {
				return true
			}
		}
	}
	return false
}

func (m *Record) SetError(err error) *Record {
	m.Error = err
	return m
}

func (m *Record) HasError() bool {
	return m.Error != nil
}

func (m *Record) GetError() string {
	if m.HasError() {
		return m.Error.Error()
	}
	return ""
}
