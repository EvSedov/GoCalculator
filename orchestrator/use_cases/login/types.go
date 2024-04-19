package login

import "github.com/evsedov/GoCalculator/orchestrator/entities"

type (
	Response struct {
		User  entities.User `json:"user"`
		Error string        `json:"error"`
	}
)
