package main

import (
	"errors"
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
		return
	}

	if err := validateParameter(request.Email, request.Password); err != nil {
		returnError(c, err, http.StatusBadRequest)
		return
	}

	// create hash password
	hashPassword, err := encryptPassword(request.Password)
	if err != nil {
		returnError(c, err, http.StatusInternalServerError)
		return
	}

	user := User{
		Email:    request.Email,
		Password: hashPassword,
	}

	// create user
	if err := h.db.Table("user").Create(&user).Error; err != nil {
		returnError(c, err, http.StatusInternalServerError)
		return
	}

	c.JSONP(http.StatusOK, user)
}

func (h *Handlers) loginUserHandler(c *gin.Context) {
	var request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&request); err != nil {
		returnError(c, err, http.StatusBadRequest)
		return
	}

	if err := validateParameter(request.Email, request.Password); err != nil {
		returnError(c, err, http.StatusBadRequest)
		return
	}

	// get user
	var user User
	if err := h.db.Table("user").
		Where("email = ?", request.Email).
		Scan(&user).Error; err != nil {
		returnError(c, err, http.StatusInternalServerError)
		return
	}

	// compare hash password
	if err := compareHashPassword(user.Password, request.Password); err != nil {
		returnError(c, errors.New("password is invalid"), http.StatusBadRequest)
		return
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
