package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	// CSVリーダーを作成
	reader := csv.NewReader(os.Stdin)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// CSVのヘッダーを取得
	header := records[0]

	// JSONデータを格納するためのスライス
	var jsonData []map[string]string

	// 各レコードを処理
	for _, record := range records[1:] {
		item := make(map[string]string)
		for i, v := range record {
			item[header[i]] = v
		}
		jsonData = append(jsonData, item)
	}

	// JSONに変換
	jsonBytes, err := json.MarshalIndent(jsonData, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	// 標準出力にJSONを出力
	fmt.Println(string(jsonBytes))
}
