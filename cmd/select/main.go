package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/yamagame/school-api-gateway/infra"
	"github.com/yamagame/school-api-gateway/infra/repository"
)

func main() {
	db := infra.DB()

	// labosテーブルをセレクト
	var labos []repository.Labos
	db.Find(&labos)

	// JSONに変換
	jsonBytes, err := json.MarshalIndent(labos, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	// 標準出力にJSONを出力
	fmt.Println(string(jsonBytes))
}
