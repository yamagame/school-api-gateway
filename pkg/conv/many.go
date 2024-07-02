package conv

type Many struct {
	Model   *Record
	Records []*Record
}

func (m *Many) ValueMap() []map[string]interface{} {
	r := []map[string]interface{}{}
	for _, v := range m.Records {
		r = append(r, v.ValueMap())
	}
	return r
}

func (m *Many) Append(record *Record) error {
	m.Records = append(m.Records, record)
	return nil
}

func (m *Many) Clear() error {
	m.Records = []*Record{}
	return nil
}

func (m *Many) NewOne() *Record {
	r := m.Model.Copy()
	m.Records = append(m.Records, r)
	return r
}
