package register

import (
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
	err := s.user.Create(&user)
	if err != nil {
		return Response{
			Error: err.Error(),
		}
	}

	return Response{
		Error: "",
	}
}
