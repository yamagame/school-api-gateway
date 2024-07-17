package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"slices"
	"strings"
)

func fillCSV(records [][]string) []map[string]string {
	ret := []map[string]string{}
	header := records[0]
	prevalues := map[string]string{}
	for _, record := range records[1:] {
		field := map[string]string{}
		for i, column := range header {
			r := strings.TrimSpace(record[i])
			if r == "-" {
				continue
			} else if r != "" {
				prevalues[column] = r
				field[column] = r
			} else {
				field[column] = prevalues[column]
			}
		}
		ret = append(ret, field)
	}
	return ret
}

func isEqualStrings(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func mergeRecords(records []map[string]interface{}) []map[string]interface{} {
	ret := []map[string]interface{}{}

	re1, _ := regexp.Compile(`(\..+\[\])(.*)`)
	re2, _ := regexp.Compile(`(\..+)\[\](.*)`)

	pkeys := []string{}
	var r map[string]interface{}
	for _, record := range records {
		heads := []string{}

		for k, v := range record {
			matches := re1.FindAllStringSubmatch(k, -1)
			if len(matches) > 0 {
			} else {
				heads = append(heads, v.(string))
			}
		}

		slices.Sort(heads)
		if !isEqualStrings(pkeys, heads) {
			pkeys = heads
			r = map[string]interface{}{}
			ret = append(ret, r)
		}

		for k, v := range record {
			matches := re2.FindAllStringSubmatch(k, -1)
			if len(matches) > 0 {
				head := matches[0][1]
				var h []map[string]string
				if _, ok := r[head]; !ok {
					h = []map[string]string{}
				} else {
					h = r[head].([]map[string]string)
				}
				h = append(h, v.(map[string]string))
				r[head] = h
			} else {
				r[k] = v
			}
		}
	}

	return ret
}

func dumpCSV(in io.Reader) ([]byte, error) {
	// CSVリーダーを作成
	reader := csv.NewReader(in)
	csvdata, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	records := fillCSV(csvdata)

	// JSONデータを格納するためのスライス
	var jsondata []map[string]interface{}

	re, _ := regexp.Compile(`(\..+\[\])(.*)`)

	// 各レコードを処理
	for _, record := range records {
		item := make(map[string]interface{})
		for head, v := range record {
			matches := re.FindAllStringSubmatch(head, -1)
			if len(matches) > 0 {
				head := matches[0][1]
				val := matches[0][2]
				if _, ok := item[head]; !ok {
					item[head] = map[string]string{}
				}
				t := item[head].(map[string]string)
				t[val] = v
			} else {
				item[head] = v
			}
		}
		jsondata = append(jsondata, item)
	}

	jsondata = mergeRecords(jsondata)

	// JSONに変換
	return json.MarshalIndent(jsondata, "", "  ")
}

func main() {
	v, err := dumpCSV(os.Stdin)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(v))
}
