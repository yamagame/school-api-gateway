package main

import (
	"github.com/yamagame/school-api-gateway/infra"
	"github.com/yamagame/school-api-gateway/infra/model"
)

func main() {
	db := infra.DB()

	// labos テーブルをマイグレート
	_ = db.AutoMigrate(&model.Role{})
	_ = db.AutoMigrate(&model.Address{})
	_ = db.AutoMigrate(&model.Person{})
	_ = db.AutoMigrate(&model.Alias{})
	_ = db.AutoMigrate(&model.Area{})
	_ = db.AutoMigrate(&model.Building{})
	_ = db.AutoMigrate(&model.Class{})
	_ = db.AutoMigrate(&model.Course{})
	_ = db.AutoMigrate(&model.Group{})
	_ = db.AutoMigrate(&model.Employee{})
	_ = db.AutoMigrate(&model.Labo{})
	_ = db.AutoMigrate(&model.Post{})
	_ = db.AutoMigrate(&model.Professor{})
	_ = db.AutoMigrate(&model.Program{})
	_ = db.AutoMigrate(&model.Room{})
	_ = db.AutoMigrate(&model.School{})
	_ = db.AutoMigrate(&model.Student{})
}
