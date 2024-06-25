package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type UniqMap struct {
	values []string
	mapval map[string]int
}

func newUniqMap() *UniqMap {
	return &UniqMap{
		values: []string{},
		mapval: map[string]int{},
	}
}

func (x *UniqMap) Add(val string) {
	x.values = append(x.values, val)
}

func (x *UniqMap) Sort() {
	slices.Sort(x.values)
	values := slices.Compact(x.values)
	slices.Sort(values)
	x.values = values
	for i, key := range values {
		x.mapval[key] = i + 1
	}
}

func (x UniqMap) GetIndex(key string) string {
	return fmt.Sprintf("%d", x.mapval[key])
}

func (x UniqMap) SQL() string {
	sql := ""
	sql += fmt.Sprintf("(\"%s\", NOW(), NOW())", x.values[0])
	for _, val := range x.values[1:] {
		sql += fmt.Sprintf(",\n(\"%s\", NOW(), NOW())", val)
	}
	sql += ";\n"
	return sql
}

type Generate struct {
	groups    *UniqMap
	programs  *UniqMap
	buildings *UniqMap
}

func (x Generate) record2column(record []string) string {
	return "`" + strings.Join(record, "`, `") + "`"
}

func (x Generate) record2value(record []string) string {
	return "(\"" + record[0] + "\", " + x.groups.GetIndex(record[1]) + ", " + x.programs.GetIndex(record[2]) + ", " + x.buildings.GetIndex(record[3]) + ", NOW(), NOW())"
}

func (x Generate) groupsSQL() string {
	sql := "INSERT INTO `groups` (`name`, `created_at`, `updated_at`) VALUES\n"
	sql += x.groups.SQL()
	return sql
}

func (x Generate) programsSQL() string {
	sql := "INSERT INTO `programs` (`name`, `created_at`, `updated_at`) VALUES\n"
	sql += x.programs.SQL()
	return sql
}

func (x Generate) buildingsSQL() string {
	sql := "INSERT INTO `buildings` (`name`, `created_at`, `updated_at`) VALUES\n"
	sql += x.buildings.SQL()
	return sql
}

func (x Generate) labosSQL(records [][]string) string {
	sql := "INSERT INTO `labos` (`name`, `group_id`, `program_id`, `building_id`, `created_at`, `updated_at`) VALUES\n"

	// 各レコードを処理
	sql += x.record2value(records[1])
	for _, record := range records[2:] {
		sql += ",\n"
		sql += x.record2value(record)
	}

	sql += ";"
	return sql
}

func main() {
	// CSVリーダーを作成
	reader := csv.NewReader(os.Stdin)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	groups := newUniqMap()
	programs := newUniqMap()
	buildings := newUniqMap()
	for _, record := range records[2:] {
		groups.Add(record[1])
		programs.Add(record[2])
		buildings.Add(record[3])
	}
	groups.Sort()
	programs.Sort()
	buildings.Sort()

	gen := Generate{
		groups:    groups,
		programs:  programs,
		buildings: buildings,
	}

	// 標準出力にJSONを出力
	fmt.Println(gen.groupsSQL())
	fmt.Println(gen.programsSQL())
	fmt.Println(gen.buildingsSQL())
	fmt.Println(gen.labosSQL(records))
}
