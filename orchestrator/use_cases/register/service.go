package register

import (
	"encoding/json"

	"github.com/evsedov/GoCalculator/orchestrator/entities"
)

type (
	Service struct {
		user UserCreater
	}
)

func NewService(user UserCreater) *Service {
	return &Service{
		user: user,
	}
}

func (s *Service) CreateUser(user entities.User) Response {
	var responseUser entities.User

	u, err := s.user.Create(&user)
	if err != nil {
		return Response{
			Error: err,
		}
	}

	if err = json.Unmarshal(u, &responseUser); err != nil {
		return Response{
			Error: err,
		}
	}

	return Response{
		User:  responseUser,
		Error: nil,
	}
}
