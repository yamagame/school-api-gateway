package conv

type Records []*Record

func NewRecords() *Records {
	return &Records{}
}

func NewRecordsWithMap(records []map[string]string, factory func() *Record) (*Records, error) {
	work := NewRecords()
	for _, record := range records {
		desk, err := NewRecordWithMap(record, factory)
		if err != nil {
			return nil, err
		}
		work.Append(desk)
	}
	return work, nil
}

func (f *Records) Take(jsonpath string, val interface{}) (*Record, error) {
	for _, record := range *f {
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
	for _, record := range *f {
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
	for _, record := range *f {
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

func (f *Records) Append(record *Record) {
	*f = append(*f, record)
}

func (f *Records) MergeRecords(records *Records, pkpath, valpath string) error {
	for _, record := range *records {
		pk, err := record.Get(pkpath)
		if err != nil {
			return err
		}
		record, err := record.HasOne(valpath)
		if err != nil {
			return err
		}
		target, err := f.Take(pkpath, pk)
		if err != nil {
			return err
		}
		records, err := target.HasMany(valpath)
		if err != nil {
			return err
		}
		records.Append(record)
	}
	return nil
}

func (f *Records) ValueMap() []map[string]interface{} {
	r := []map[string]interface{}{}
	for _, v := range *f {
		r = append(r, v.ValueMap())
	}
	return r
}

func (f *Records) Updates() []map[string]interface{} {
	r := []map[string]interface{}{}
	for _, v := range *f {
		r = append(r, v.Updates())
	}
	return r
}

func (f *Records) UpdateValues() []map[string]interface{} {
	r := []map[string]interface{}{}
	for _, v := range *f {
		r = append(r, v.UpdateValues())
	}
	return r
}

func (f *Records) UpdateHasOnes() []map[string]interface{} {
	r := []map[string]interface{}{}
	for _, v := range *f {
		r = append(r, v.UpdateHasOnes())
	}
	return r
}

func (f *Records) UpdateHasManyes() []map[string]interface{} {
	r := []map[string]interface{}{}
	for _, v := range *f {
		r = append(r, v.UpdateHasManyes())
	}
	return r
}
