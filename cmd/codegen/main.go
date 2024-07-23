package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

func toLower(in []string) []string {
	ret := []string{}
	for _, s := range in {
		ret = append(ret, strings.ToLower(s))
	}
	return ret
}

func exec(jsondata, tempfile []byte, outdir string, tarclasses ...string) {
	var err error
	files := []map[string]interface{}{}
	json.Unmarshal(jsondata, &files)
	tmpl := template.Must(template.New("codegen").Funcs(sprig.FuncMap()).Parse(string(tempfile)))
	tarclasses = toLower(tarclasses)
	for _, dat := range files {
		if len(tarclasses) > 0 && tarclasses[0] != "" && !slices.Contains(tarclasses, strings.ToLower(dat["classes"].(string))) {
			continue
		}
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
	if len(args) == 1 {
		fmt.Println("json ファイルを指定してください")
		return
	}
	jsonfile := args[1]
	var tarclass string
	if len(args) > 2 {
		tarclass = args[2]
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
	exec(jsondata, tempfile, "", tarclass)
}
