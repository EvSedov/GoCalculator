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

func (s *storage) Login(user *entities.User) (data []byte, err error) {
	email := user.Email
	password := user.Password
	var dbUser entities.User
	s.DB.Where("email = ?", email).First(&dbUser)
	if dbUser.Id == 0 {
		err = errors.New("user not found")
		return nil, err
	}

	dbHash := dbUser.Password

	if err = bcrypt.CompareHashAndPassword(dbHash, []byte(password)); err != nil {
		err2 := errors.New("incorrect login or password")
		return nil, err2
	}

	return dbHash, nil
}
