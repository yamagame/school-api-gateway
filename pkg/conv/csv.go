package conv

import (
	"encoding/csv"
	"io"
)

func ReadCSV(r io.Reader) ([]map[string]string, error) {
	reader := csv.NewReader(r)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	ret := []map[string]string{}
	header := records[0]
	for _, record := range records[1:] {
		field := map[string]string{}
		for i, column := range header {
			field[column] = record[i]
		}
		ret = append(ret, field)
	}
	return ret, nil
}
