package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	// mysql接続
	dsn := "root:root@tcp(127.0.0.1:3306)/golang_login_test"
	_, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatal(err)
	}
}
