package config

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Env(key string) string {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic("Error loading .env file")
	}

	return os.Getenv(key)
}

func Connect() {
	dsn := Env("DB_USERNAME") + ":" + Env("DB_PASSWORD") + "@tcp(" + Env("DB_HOST") + ":" + Env("DB_PORT") + ")/" + Env("DB_DATABASE") + "?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
