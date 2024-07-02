package conv

type Records []*Record

func NewRecords() *Records {
	return &Records{}
}

func (f *Records) Append(record *Record) {
	*f = append(*f, record)
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
