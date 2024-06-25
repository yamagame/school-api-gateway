package main

import (
	"github.com/yamagame/school-api-gateway/infra"
	"github.com/yamagame/school-api-gateway/infra/model"
	"gorm.io/gen"
)

func main() {
	db := infra.DB()

	g := gen.NewGenerator(gen.Config{
		// 生成ディレクトリ、パッケージ名になる
		OutPath: "./infra/dao/query",
		// モード
		Mode:              gen.WithoutContext | gen.WithQueryInterface,
		FieldWithIndexTag: true,
		FieldNullable:     true,
		FieldWithTypeTag:  true,
	})

	// gorm.DBを指定する
	g.UseDB(db)

	g.ApplyBasic(model.Labo{})

	// // 全てのテーブルを取得
	// tableList, err := db.Migrator().GetTables()
	// if err != nil {
	// 	panic(err)
	// }

	// // 各テーブル毎にモデルを作成
	// tables := make([]interface{}, len(tableList))

	// // 残りのテーブルのモデルを作成
	// for _, tableName := range tableList {
	// 	tables = append(tables, g.GenerateModel(tableName))
	// }

	// // DAOを生成
	// g.ApplyBasic(
	// 	tables...,
	// )

	// コードを生成
	g.Execute()
}
