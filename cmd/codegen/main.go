package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

func exec(jsondata, tempfile []byte, outdir string) {
	var err error
	files := []map[string]interface{}{}
	json.Unmarshal(jsondata, &files)
	tmpl := template.Must(template.New("codegen").Funcs(sprig.FuncMap()).Parse(string(tempfile)))
	for _, dat := range files {
		filename, ok := dat["filename"].(string)
		if !ok {
			continue
		}
		fp := os.Stdout
		if outdir != "" {
			fname := filepath.Join(outdir, filename)
			fp, err = os.Create(fname)
			if err != nil {
				panic(err)
			}
		}
		err := tmpl.Execute(fp, dat)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	args := os.Args
	if len(args) == 0 {
		fmt.Println("json ファイルを指定してください")
		return
	}
	jsonfile := args[1]
	outdir := ""
	if len(args) > 2 {
		outdir = args[2]
	}
	jsondata, err := os.ReadFile(jsonfile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	tempfile, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	exec(jsondata, tempfile, outdir)
}
