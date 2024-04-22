package register

import (
	"time"

	"github.com/evsedov/GoCalculator/orchestrator/entities"
	"github.com/evsedov/GoCalculator/orchestrator/utils"
	"github.com/golang-jwt/jwt/v5"
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

func (s *Service) CreateUser(user entities.User) *Response {
	_, err := s.user.Create(&user)
	if err != nil {
		return &Response{
			Error: err.Error(),
		}
	}

	hmacSampleSecret := utils.Secret
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"nbf":   now.Unix(),
		"exp":   now.Add(5 * time.Hour).Unix(),
		"iat":   now.Unix(),
	})

	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		return &Response{
			Error: err.Error(),
		}
	}

	return &Response{
		Token: tokenString,
		Error: "",
	}
}
