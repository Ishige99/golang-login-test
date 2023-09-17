package main

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/mail"
	"regexp"
)

func encryptPassword(password string) (string, error) {
	// パスワードの文字列をハッシュ化する
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func compareHashPassword(hashedPassword, requestPassword string) error {
	// パスワードの文字列をハッシュ化して、既に登録されているハッシュ化したパスワードと比較します
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

func getAlreadyExistEmail(db *gorm.DB, email string) (bool, error) {
	var alreadyExistEmail string
	result := db.Select("email").
		Table("user").
		Where("email = ?", email).
		Scan(&alreadyExistEmail)

	if result.Error != nil {
		return false, result.Error
	}

	if result.RowsAffected >= 1 {
		return true, errors.New("email is already exist")
	}

	return false, nil
}
