package irconv

type HasMany struct {
	Error   error
	Model   *Record
	records []*Record
}

func NewMany(model *Record) *HasMany {
	return &HasMany{
		Model:   model,
		records: []*Record{},
	}
}

func (m *HasMany) NewIterator() *Iterator[Record] {
	return NewIterator(m.records)
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
	for _, v := range m.records {
		r = append(r, v.ValueMap())
	}
	return r
}

func (m *HasMany) IsExist() bool {
	for _, v := range m.records {
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
	m.records = append(m.records, records...)
	return m
}

func (m *HasMany) Clear() *HasMany {
	if m.HasError() {
		return m
	}
	m.records = []*Record{}
	return m
}

func (m *HasMany) Copy() *HasMany {
	r := &HasMany{}
	r.Error = m.Error
	r.Model = m.Model.Copy()
	r.records = []*Record{}
	for _, v := range m.records {
		r.records = append(r.records, v.Copy())
	}
	return r
}

func (m *HasMany) NewOne() *Record {
	r := m.Model.Copy()
	m.records = append(m.records, r)
	return r
}

func (m *HasMany) SetError(err error) *HasMany {
	m.Error = err
	return m
}

func (m *HasMany) HasError() bool {
	return m.Error != nil
}
