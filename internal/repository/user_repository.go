package repository

import (
	"QtCloudPan/config"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 注意：驱动包以 _ 前缀导入，以便初始化
)

// SaveUserToDB 将用户数据保存到数据库
func SaveUserToDB(username, password, email string) error {
	db, err := sql.Open("mysql", config.AppConfig.DatabaseDSN)
	if err != nil {
		return err
	}
	defer db.Close()

	query := "INSERT INTO user (username, password, email) VALUES (?, ?, ?)"
	_, err = db.Exec(query, username, password, email)
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}
	fmt.Println("User saved to database")
	return nil
}

func QueryUserFromDB(username, password string) error {
	db, err := sql.Open("mysql", config.AppConfig.DatabaseDSN)
	if err != nil {
		return err
	}
	defer db.Close()

	query := "SELECT * FROM user WHERE username=? AND password=?"
	rows, err := db.Query(query, username, password)
	if err != nil {
		return fmt.Errorf("failed to query user: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		fmt.Println("User found in database")
		return nil
	}
	return fmt.Errorf("user not found in database")

}
