package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	// connect MySQL
	db, err := connectDatabase()
	if err != nil {
		log.Fatal(err)
	}

	h := &Handlers{
		db: db,
	}

	// router
	router := gin.Default()
	router.GET("/hello", helloWorldHandler)
	router.POST("/user", h.createUserHandler)
	router.Run()
}

func connectDatabase() (*gorm.DB, error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/golang_login_test"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return nil, err
	}
	return db, nil
}
