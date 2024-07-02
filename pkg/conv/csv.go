package conv

import (
	"encoding/csv"
	"io"
)

func ReadCSV(r io.Reader) ([]map[string]string, error) {
	reader := csv.NewReader(r)
	reader.FieldsPerRecord = -1
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	ret := []map[string]string{}
	header := records[0]
	prevalues := map[string]string{}
	for _, record := range records[1:] {
		field := map[string]string{}
		for i, column := range header {
			if record[i] != "" {
				prevalues[column] = record[i]
				field[column] = record[i]
			} else {
				field[column] = prevalues[column]
			}
		}
		ret = append(ret, field)
	}
	return ret, nil
}
