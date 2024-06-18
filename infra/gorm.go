package infra

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type config struct {
	DB_USER string
	DB_PASS string
	DB_HOST string
	DB_PORT string
	DB_NAME string
	DB_LOCA string
}

func osEnv(key, defaultvalue string) string {
	val, ok := os.LookupEnv(key)
	if ok {
		return val
	}
	return defaultvalue
}

var dbconf = config{}

func init() {
	dbconf.DB_USER = osEnv("DB_USER", "root")
	dbconf.DB_PASS = osEnv("DB_PASS", "")
	dbconf.DB_HOST = osEnv("DB_HOST", "localhost")
	dbconf.DB_PORT = osEnv("DB_PORT", "3336")
	dbconf.DB_NAME = osEnv("DB_NAME", "school-database")
	dbconf.DB_LOCA = osEnv("DB_LOCA", "UTC")
}

func DB() *gorm.DB {
	dsn := dbconf.DB_USER + ":" + dbconf.DB_PASS + "@tcp(" + dbconf.DB_HOST + ":" + dbconf.DB_PORT + ")/" + dbconf.DB_NAME + "?charset=utf8mb4&parseTime=True&loc=" + dbconf.DB_LOCA
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
