package iconv

type HasMany struct {
	Error   error
	Model   *Record
	records *Records
}

func NewMany(model *Record) *HasMany {
	return &HasMany{
		Model:   model,
		records: NewRecords(model),
	}
}

func (m *HasMany) Records() *Records {
	return m.records
}

func (m *HasMany) NewIterator() *Iterator[Record] {
	return NewIterator(m.records.Slice())
}

func (m *HasMany) Take(jsonp string, val interface{}) *Record {
	itr := m.NewIterator()
	for itr.HasNext() {
		field := itr.Next()
		if v, err := field.Get(jsonp); err == nil {
			if v == val {
				return field
			}
		}
	}
	r := m.Model.Copy()
	r.Error = ErrNotFound
	return r
}

func (m *HasMany) ValueMap() []map[string]interface{} {
	r := []map[string]interface{}{}
	for _, v := range m.records.Slice() {
		r = append(r, v.ValueMap())
	}
	return r
}

func (m *HasMany) IsExist() bool {
	for _, v := range m.records.Slice() {
		if v.IsExist() {
			return true
		}
	}
	return false
}

func (m *HasMany) Append(records ...*Record) *HasMany {
	if m.HasError() {
		return m
	}
	m.records.Append(records...)
	return m
}

func (m *HasMany) Clear() *HasMany {
	if m.HasError() {
		return m
	}
	m.records = NewRecords(m.Model)
	return m
}

func (m *HasMany) Copy() *HasMany {
	r := &HasMany{}
	r.Error = m.Error
	r.Model = m.Model.Copy()
	r.records = NewRecords(m.Model)
	for _, v := range m.records.Slice() {
		r.records.Append(v.Copy())
	}
	return r
}

func (m *HasMany) NewOne() *Record {
	r := m.Model.Copy()
	m.records.Append(r)
	return r
}

func (m *HasMany) SetError(err error) *HasMany {
	m.Error = err
	return m
}

func (m *HasMany) HasError() bool {
	return m.Error != nil
}

func (m *HasMany) GetError() string {
	if m.HasError() {
		return m.Error.Error()
	}
	return ""
}
