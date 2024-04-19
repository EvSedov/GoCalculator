package storage

import (
	"encoding/json"
	"errors"

	"github.com/evsedov/GoCalculator/orchestrator/entities"
	"golang.org/x/crypto/bcrypt"
)

func (s *storage) Create(user *entities.User) (data []byte, err error) {
	if err = s.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	data, err = json.Marshal(user)
	if err != nil {
		return nil, err
	}

	return data, nil
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

	if err = bcrypt.CompareHashAndPassword(dbUser.Password, []byte(password)); err != nil {
		err2 := errors.New("incorrect password")
		return nil, errors.Join(err, err2)
	}

	data, err = json.Marshal(dbUser)
	if err != nil {
		return nil, err
	}

	return data, nil
}
