package storage

import (
	"errors"

	"github.com/evsedov/GoCalculator/orchestrator/entities"
	"golang.org/x/crypto/bcrypt"
)

func (s *storage) Create(user *entities.User) (err error) {
	if err = s.DB.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (s *storage) Login(user *entities.User) (err error) {
	email := user.Email
	password := user.Password
	var dbUser entities.User
	s.DB.Where("email = ?", email).First(&dbUser)
	if dbUser.Id == 0 {
		err = errors.New("user not found")
		return err
	}

	if err = bcrypt.CompareHashAndPassword(dbUser.Password, []byte(password)); err != nil {
		err2 := errors.New("incorrect login or password")
		return err2
	}

	return nil
}
