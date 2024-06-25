package field

import (
	"encoding/csv"
	"io"
	"reflect"
)

type Field map[string]interface{}

func (m *Field) SetValue(key string, defval interface{}) *Field {
	(*m)[key] = NewValue(key, defval)
	return m
}

func (m *Field) SetField(key string, defval interface{}) *Field {
	(*m)[key] = defval
	return m
}

func (m *Field) SetArray(key string, fields ...*Field) *Field {
	(*m)[key] = fields
	return m
}

func (m *Field) Value(key string) (*Value, error) {
	if v, ok := (*m)[key]; ok {
		if reflect.TypeOf(v) == reflect.TypeOf(&Value{}) {
			value := v.(*Value)
			return value, nil
		}
	}
	return nil, ErrNotFound
}

func (m *Field) Field(key string) (*Field, error) {
	if v, ok := (*m)[key]; ok {
		if reflect.TypeOf(v) == reflect.TypeOf(&Field{}) {
			field := v.(*Field)
			return field, nil
		}
	}
	return nil, ErrNotFound
}

func (m *Field) Array(key string) ([]*Field, error) {
	if v, ok := (*m)[key]; ok {
		if reflect.TypeOf(v) == reflect.TypeOf([]*Field{}) {
			field := v.([]*Field)
			return field, nil
		}
	}
	return nil, ErrNotFound
}

func (m *Field) Set(key string, val interface{}) error {
	if value, err := m.Value(key); err == nil {
		value.Set(val)
		return nil
	}
	return ErrNotFound
}

func (m *Field) Get(key string) (interface{}, error) {
	if value, err := m.Value(key); err == nil {
		return value.Get(), nil
	}
	return nil, ErrNotFound
}

type Fields []*Field

func ReadCSV(r io.Reader, factory func(id int32) *Field) (Fields, error) {
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

func NewFieldWithMap(field map[string]interface{}, factory func(id int32) *Field) (*Field, error) {
	newone := factory(0)
	t := newone
	for column, val := range field {
		key := NewKey(column)
		if key.HasRelation() {
			for _, col := range key.Fields[:len(key.Fields)-1] {
				if v, err := t.Field(col); err == nil {
					t = v
				}
			}
		}
		if value, err := t.Value(key.Last()); err == nil {
			value.SetOriginal(val)
		}
	}
	return newone, nil
}
