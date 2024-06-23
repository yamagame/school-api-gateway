package entity

import (
	"encoding/csv"
	"io"
)

type Labo struct {
	Field
}

func NewLabo() *Labo {
	v := &Labo{
		NewField(),
	}
	v.AddProp("id", int32(0))
	v.AddProp("name", "")
	v.AddProp("group", "")
	v.AddProp("program", "")
	v.AddProp("building", "")
	return v
}

type Labos []*Labo

func ReadLaboCSV(r io.Reader) (Labos, error) {
	reader := csv.NewReader(r)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	labos := []*Labo{}
	header := records[0]
	for _, record := range records[1:] {
		var err error
		labo := NewLabo()
		for i, column := range header {
			if err = labo.SetValue(column, record[i]); err != nil {
				return nil, err
			}
		}
		labos = append(labos, labo)
	}
	return labos, nil
}

func (x Labos) ToMap() []map[string]interface{} {
	r := []map[string]interface{}{}
	for _, v := range x {
		r = append(r, v.ToMap())
	}
	return r
}
