package util

import (
	"github.com/pkg/errors"
	"go_gin_app/models"
	"golang.org/x/crypto/bcrypt"
)

func PasswordHashing(user *models.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.WithStack(err)
	}
	user.Password = string(hash)
	return nil
}

func ComparePasswords(user *models.User, password string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, nil
		}
		return false, errors.WithStack(err)
	}
	return true, nil
}
