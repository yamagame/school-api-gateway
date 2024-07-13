package conv

type Many struct {
	Error   error
	Model   *Record
	Records []*Record
}

func NewMany(model *Record) *Many {
	return &Many{
		Model:   model,
		Records: []*Record{},
	}
}

func (m *Many) ValueMap() []map[string]interface{} {
	r := []map[string]interface{}{}
	for _, v := range m.Records {
		r = append(r, v.ValueMap())
	}
	return r
}

func (m *Many) IsExist() bool {
	for _, v := range m.Records {
		if v.IsExist() {
			return true
		}
	}
	return false
}

func (m *Many) Append(records ...*Record) *Many {
	if m.Error != nil {
		return m
	}
	m.Records = append(m.Records, records...)
	return m
}

func (m *Many) Clear() *Many {
	if m.Error != nil {
		return m
	}
	m.Records = []*Record{}
	return m
}

func (m *Many) Copy() *Many {
	r := &Many{}
	r.Error = m.Error
	r.Model = m.Model.Copy()
	r.Records = []*Record{}
	for _, v := range m.Records {
		r.Records = append(r.Records, v.Copy())
	}
	return r
}

func (m *Many) NewOne() *Record {
	r := m.Model.Copy()
	m.Records = append(m.Records, r)
	return r
}
