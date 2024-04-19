package register

import "github.com/evsedov/GoCalculator/orchestrator/entities"

type Response struct {
	User  entities.User `json:"user"`
	Error error         `json:"error"`
}
