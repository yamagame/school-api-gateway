package infra

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type config struct {
	DBUser string
	DBPass string
	DBHost string
	DBPort string
	DBName string
	DBLoca string
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
	dbconf.DBUser = osEnv("DB_USER", "root")
	dbconf.DBPass = osEnv("DB_PASS", "")
	dbconf.DBHost = osEnv("DB_HOST", "localhost")
	dbconf.DBPort = osEnv("DB_PORT", "3336")
	dbconf.DBName = osEnv("DB_NAME", "school-database")
	dbconf.DBLoca = osEnv("DB_LOCA", "UTC")
}

func DB() *gorm.DB {
	dsn := dbconf.DBUser + ":" + dbconf.DBPass + "@tcp(" + dbconf.DBHost + ":" + dbconf.DBPort + ")/" + dbconf.DBName + "?charset=utf8mb4&parseTime=True&loc=" + dbconf.DBLoca
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
