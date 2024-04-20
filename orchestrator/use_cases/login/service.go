package login

import (
	"time"

	"github.com/evsedov/GoCalculator/orchestrator/entities"
	"github.com/golang-jwt/jwt/v5"
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
	err := s.user.Login(&user)
	if err != nil {
		return &Response{
			Error: err.Error(),
		}
	}

	const hmacSampleSecret = "qwerty_secret_357"
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"nbf":   now.Unix(),
		"exp":   now.Add(5 * time.Hour).Unix(),
		"iat":   now.Unix(),
	})

	tokenString, err := token.SignedString([]byte(hmacSampleSecret))
	if err != nil {
		panic(err)
	}

	return &Response{
		Token: tokenString,
		Error: "",
	}
}
