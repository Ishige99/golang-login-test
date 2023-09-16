package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func connectDatabase() (*gorm.DB, error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/golang_login_test"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	// mysql接続
	_, err := connectDatabase()
	if err != nil {
		log.Fatal(err)
	}
}
