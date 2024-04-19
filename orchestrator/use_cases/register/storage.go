package register

import "github.com/evsedov/GoCalculator/orchestrator/entities"

type UserCreater interface {
	Create(user *entities.User) ([]byte, error)
}
