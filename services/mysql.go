package services

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

// MysqlConnect function for Database connection
func MysqlConnect() *gorm.DB {

	godotenv.Load()

	mysqlName := os.Getenv("MYSQL_DB_NAME")
	mysqlUser := os.Getenv("MYSQL_DB_USER")
	mysqlPass := os.Getenv("MYSQL_DB_PASS")

	connectionString := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", mysqlUser, mysqlPass, mysqlName)

	db, err := gorm.Open("mysql", connectionString)
	// defer db.Close()

	if err != nil {
		panic(err.Error())
	}

	return db
}
