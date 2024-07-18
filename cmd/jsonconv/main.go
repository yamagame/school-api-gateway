package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/huandu/xstrings"
	"github.com/jinzhu/inflection"
)

type TABLE struct {
	Filename   string    `json:"filename"`
	Class      string    `json:"class"`
	Classes    string    `json:"classes"`
	Attributes []*COLUMN `json:"attributes"`
}

type COLUMN struct {
	Comment   string `json:"COLUMN_COMMENT"`
	Key       string `json:"COLUMN_KEY"`
	Name      string `json:"COLUMN_NAME"`
	Type      string `json:"COLUMN_TYPE"`
	Nullable  string `json:"IS_NULLABLE"`
	TableName string `json:"TABLE_NAME"`
}

func exec(jsondata []byte) {
	columns := []COLUMN{}
	json.Unmarshal(jsondata, &columns)
	tmap := map[string]*TABLE{}
	for _, col := range columns {
		if _, ok := tmap[col.TableName]; !ok {
			tmap[col.TableName] = &TABLE{
				Filename:   inflection.Singular(col.TableName) + ".go",
				Class:      xstrings.ToPascalCase(inflection.Singular(col.TableName)),
				Classes:    xstrings.ToPascalCase(col.TableName),
				Attributes: []*COLUMN{},
			}
		}
		tmap[col.TableName].Attributes = append(tmap[col.TableName].Attributes, &col)
	}

	ret := []*TABLE{}
	for _, v := range tmap {
		ret = append(ret, v)
	}

	b, _ := json.MarshalIndent(ret, "", "  ")
	fmt.Println(string(b))
}

func main() {
	jsondata, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	exec(jsondata)
}
