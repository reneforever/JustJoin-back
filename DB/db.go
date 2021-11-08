package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() (err error) {
	dsn := "root:123456@tcp(1227.0.0.1:3306)/justjoin"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	fmt.Println("DB connect success")
	return nil
}

func GetDB() *sql.DB {
	return db
}
