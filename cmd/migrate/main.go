package main

import (
	"github.com/yamagame/school-api-gateway/infra"
	"github.com/yamagame/school-api-gateway/infra/model"
)

func main() {
	db := infra.DB()

	// labos テーブルをマイグレート
	_ = db.AutoMigrate(&model.Labo{})
}
