package irconv

import (
	"fmt"

	"golang.org/x/exp/maps"
)

type Records struct {
	Error   error
	model   *Record
	records []*Record
}

func NewRecords(model *Record) *Records {
	return &Records{
		model:   model,
		records: []*Record{},
	}
}

func (f *Records) Records() []*Record {
	return f.records
}

func (f *Records) Take(jsonpath string, val interface{}) (*Record, error) {
	for _, record := range f.records {
		got, err := record.Get(jsonpath)
		if err != nil {
			return nil, err
		}
		if got == val {
			return record, nil
		}
	}
	return nil, ErrNotFound
}

func (f *Records) Find(jsonpath string, val interface{}) ([]*Record, error) {
	ret := []*Record{}
	for _, record := range f.records {
		got, err := record.Get(jsonpath)
		if err != nil {
			return nil, err
		}
		if got == val {
			ret = append(ret, record)
		}
	}
	return ret, nil
}

func (f *Records) Uniq(jsonpath string) ([]interface{}, error) {
	isExist := func(ret []interface{}, val interface{}) bool {
		for _, r := range ret {
			if r == val {
				return true
			}
		}
		return false
	}
	ret := []interface{}{}
	for _, record := range f.records {
		got, err := record.Get(jsonpath)
		if err != nil {
			return nil, err
		}
		if !isExist(ret, got) {
			ret = append(ret, got)
		}
	}
	return ret, nil
}

func (f *Records) Append(record ...*Record) *Records {
	if f.Error != nil {
		return f
	}
	f.records = append(f.records, record...)
	return f
}

func (f *Records) ValueMap() []map[string]interface{} {
	r := []map[string]interface{}{}
	for _, v := range f.records {
		r = append(r, v.ValueMap())
	}
	return r
}

func (f *Records) Updates() []map[string]interface{} {
	r := []map[string]interface{}{}
	for _, v := range f.records {
		r = append(r, v.Updates())
	}
	return r
}

func (f *Records) UpdateValues() []map[string]interface{} {
	r := []map[string]interface{}{}
	for _, v := range f.records {
		r = append(r, v.UpdateValues())
	}
	return r
}

func (f *Records) UpdateHasOnes() []map[string]interface{} {
	r := []map[string]interface{}{}
	for _, v := range f.records {
		r = append(r, v.UpdateHasOnes())
	}
	return r
}

func (f *Records) UpdateHasManyes() []map[string]interface{} {
	r := []map[string]interface{}{}
	for _, v := range f.records {
		r = append(r, v.UpdateHasManyes())
	}
	return r
}

func (f *Records) MergeValue(records *Records, pkpath, valpath string) *Records {
	for _, record := range records.records {
		pk, err := record.Get(pkpath)
		if err != nil {
			f.Error = err
			return f
		}
		target, err := f.Take(pkpath, pk)
		if err != nil {
			f.Error = err
			return f
		}
		val, err := record.Get(valpath)
		if err != nil {
			f.Error = err
			return f
		}
		target.Set(valpath, val)
	}
	return f
}

func (f *Records) MergeHasMany(records *Records, pkpath, valpath string) *Records {
	for _, record := range records.records {
		pk, err := record.Get(pkpath)
		if err != nil {
			f.Error = err
			return f
		}
		record, err := record.HasOne(valpath)
		if err != nil {
			f.Error = err
			return f
		}
		target, err := f.Take(pkpath, pk)
		if err != nil {
			f.Error = err
			return f
		}
		records, err := target.HasMany(valpath)
		if err != nil {
			f.Error = err
			return f
		}
		if record.IsExist() {
			records.Append(record)
		}
	}
	return f
}

func (f *Records) Convert(model *Record) *Records {
	if f.Error != nil {
		return f
	}

	newrecords := NewRecords(model.Copy())

	pkeys := f.model.PrimaryKeys()
	if len(pkeys) == 0 {
		newrecords.Error = ErrNotFoundPrimaryKey
		return newrecords
	}
	primkey := fmt.Sprintf(".%s", pkeys[0])

	valuekeys := maps.Keys(f.model.Values)
	hasonekeys := maps.Keys(f.model.HasOnes)

	values, err := f.Uniq(primkey)
	if err != nil {
		f.Error = err
		return f
	}
	for _, pkey := range values {
		newrecords.Append(model.Copy().Set(primkey, pkey))
	}
	// マージ
	for _, vkey := range valuekeys {
		if err = newrecords.MergeValue(f, primkey, fmt.Sprintf(".%s", vkey)).Error; err != nil {
			return newrecords
		}
	}
	for _, mkey := range hasonekeys {
		if err = newrecords.MergeHasMany(f, primkey, fmt.Sprintf(".%s", mkey)).Error; err != nil {
			return newrecords
		}
	}
	return newrecords
}

func (f *Records) NewIterator() *Iterator[Record] {
	return NewIterator(f.records)
}

func (f *Records) SetError(err error) *Records {
	f.Error = err
	return f
}

func (f *Records) HasError() bool {
	return f.Error != nil
}
