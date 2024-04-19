package login

import (
	"encoding/json"

	"github.com/evsedov/GoCalculator/orchestrator/entities"
)

type (
	Service struct {
		user UserLogin
	}
)

func NewService(user UserLogin) *Service {
	return &Service{
		user: user,
	}
}

func (s *Service) Login(user entities.User) *Response {
	var responseUser entities.User
	u, err := s.user.Login(&user)
	if err != nil {
		return &Response{
			Error: err.Error(),
		}
	}

	if err = json.Unmarshal(u, &responseUser); err != nil {
		return &Response{
			Error: err.Error(),
		}
	}

	return &Response{
		User:  responseUser,
		Error: "",
	}
}
