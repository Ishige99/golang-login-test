package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type Handlers struct {
	db *gorm.DB
}

type User struct {
	Id        int       `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

func (h *Handlers) createUserHandler(c *gin.Context) {
	var request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&request); err != nil {
		returnError(c, err, http.StatusBadRequest)
	}

	// todo: validate mail and password

	user := User{
		Email:    request.Email,
		Password: request.Password,
	}

	// create user
	if err := h.db.Table("user").Create(&user).Error; err != nil {
		returnError(c, err, http.StatusInternalServerError)
	}

	c.JSONP(http.StatusOK, user)
}

func helloWorldHandler(c *gin.Context) {
	c.JSONP(http.StatusOK, "hello world!")
}

func returnError(c *gin.Context, err error, status int) {
	apiError := struct {
		Message string `json:"message"`
		Status  int    `json:"status"`
	}{
		Message: err.Error(),
		Status:  status,
	}
	c.JSONP(status, apiError)
}
