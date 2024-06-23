package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func record2column(record []string) string {
	return "`" + strings.Join(record, "`, `") + "`"
}

func record2value(record []string) string {
	return "(\"" + strings.Join(record, "\", \"") + "\", NOW(), NOW())"
}

func main() {
	// CSVリーダーを作成
	reader := csv.NewReader(os.Stdin)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// JSONデータを格納するためのスライス
	sql := "INSERT INTO labos (" + record2column(records[0]) + ", `created_at`, `updated_at`) VALUES\n"

	// 各レコードを処理
	sql += record2value(records[1])
	for _, record := range records[2:] {
		sql += ",\n"
		sql += record2value(record)
	}

	sql += ";"

	// 標準出力にJSONを出力
	fmt.Println(sql)
}
