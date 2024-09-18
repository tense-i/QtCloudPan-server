package utils

import (
	"QtCloudPan/config"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB // 全局变量，用于存储数据库连接
func InitDB() {
	fmt.Println("Init")
	fmt.Println(config.AppConfig.DatabaseDSN)

	DB, err := sql.Open("mysql", config.AppConfig.DatabaseDSN)
	if err != nil {
		panic(err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	fmt.Println("Database connected")

}

func Close() {
	DB.Close()
}

func QueryRow(sql string, args ...interface{}) *sql.Row {
	return DB.QueryRow(sql, args...)
}

func Query(sql string, args ...interface{}) (*sql.Rows, error) {
	return DB.Query(sql, args...)
}

func Exec(sql string, args ...interface{}) (sql.Result, error) {
	return DB.Exec(sql, args...)
}
