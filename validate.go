package main

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"net/mail"
	"regexp"
)

func encryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func compareHashPassword(hashedPassword, requestPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(requestPassword)); err != nil {
		return err
	}
	return nil
}

func validateParameter(email, password string) error {
	// validate email
	if _, err := mail.ParseAddress(email); err != nil {
		return err
	}

	// validate password
	if err := validatePassword(password); err != nil {
		return err
	}

	return nil
}

func validatePassword(password string) error {
	// 正規表現でのバリデーション
	passwordPattern := "^[a-zA-Z0-9]{4,12}$"
	match, err := regexp.MatchString(passwordPattern, password)
	if err != nil {
		return err
	}
	if !match {
		return errors.New("password is invalid. please enter alphanumeric characters between 4 and 12 characters")
	}
	return nil
}
