package login

import "github.com/evsedov/GoCalculator/orchestrator/entities"

type UserLogin interface {
	Login(user *entities.User) ([]byte, error)
}
