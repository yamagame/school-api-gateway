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

func (f *Records) Append(record *Record) {
	*f = append(*f, record)
}

func (f *Records) PickRecords(key, valkey string) ([]string, map[string]*Records, error) {
	keys := []string{}
	ret := map[string]*Records{}
	for _, record := range *f {
		if v, err := record.GetValue(key); err == nil {
			key := v.(string)
			if _, ok := ret[key]; !ok {
				ret[key] = NewRecords()
				keys = append(keys, key)
			}
			if f, err := record.GetHasOne(valkey); err == nil {
				ret[key].Append(f)
			}
		}
	}
	return keys, ret, nil
}

func (f *Records) MergeRecords(headkey, bodykey string, factory func() *Record) (*Records, error) {
	keys, body, err := f.PickRecords(headkey, bodykey)
	if err != nil {
		return nil, err
	}
	records := NewRecords()
	for _, key := range keys {
		record := factory()
		record.SetValue(headkey, key)
		record.SetHasManyRecords(bodykey, *body[key]...)
		records.Append(record)
	}
	return records, nil
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
