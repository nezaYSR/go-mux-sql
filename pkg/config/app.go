package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var (
	db *gorm.DB
)

func Connect() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dialect := os.Getenv("SQL_DIALECT")
	sql_user := os.Getenv("SQL_USER")
	sql_password := os.Getenv("SQL_PWD")
	sql_schema := os.Getenv("SQL_SCHEMA")

	uri := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", sql_user, sql_password, sql_schema)

	d, err := gorm.Open(dialect, uri)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
