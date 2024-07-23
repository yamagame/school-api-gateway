package main

import (
	"github.com/yamagame/school-api-gateway/infra"
	"github.com/yamagame/school-api-gateway/infra/model"
	"gorm.io/gen"
)

// func genmodel(g *gen.Generator, tablename string) interface{} {
// 	switch tablename {
// 	case "labos":
// 		return g.GenerateModel(tablename,
// 			gen.FieldRelate(field.BelongsTo, "Group", g.GenerateModel("groups"), nil),
// 			gen.FieldRelate(field.BelongsTo, "Program", g.GenerateModel("programs"), nil),
// 			gen.FieldRelate(field.BelongsTo, "Building", g.GenerateModel("buildings"), nil),
// 			gen.FieldRelate(field.HasMany, "Chairs", g.GenerateModel("chairs"), &field.RelateConfig{
// 				RelateSlicePointer: true,
// 			}),
// 			gen.FieldRelate(field.HasMany, "Desks", g.GenerateModel("desks"), &field.RelateConfig{
// 				RelateSlicePointer: true,
// 			}),
// 			gen.FieldNew("Error", "error", field.Tag{"gorm": "-"}),
// 		)
// 	default:
// 		return g.GenerateModel(tablename)
// 	}
// }

func generate() {
	db := infra.DB()

	g := gen.NewGenerator(gen.Config{
		// 生成ディレクトリ、パッケージ名になる
		OutPath: "./infra/dao/query",
		// モード
		Mode:              gen.WithQueryInterface,
		FieldWithIndexTag: true,
		FieldNullable:     true,
		FieldWithTypeTag:  true,
	})

	// gorm.DBを指定する
	g.UseDB(db)

	g.ApplyBasic(model.Labo{})
	g.ApplyBasic(model.Desk{})
	g.ApplyBasic(model.Chair{})

	// // 全てのテーブルを取得
	// tableList, err := db.Migrator().GetTables()
	// if err != nil {
	// 	panic(err)
	// }

	// // 各テーブル毎にモデルを作成
	// tables := make([]interface{}, len(tableList))

	// // 残りのテーブルのモデルを作成
	// for _, tablename := range tableList {
	// 	tables = append(tables, genmodel(g, tablename))
	// }

	// // DAOを生成
	// g.ApplyBasic(
	// 	tables...,
	// )

	// コードを生成
	g.Execute()
}

func main() {
	generate()
}
