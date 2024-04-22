package userexpressions

import "github.com/evsedov/GoCalculator/orchestrator/entities"

type (
	Response struct {
		UserExpressions []entities.Expression `json:"user_expressions"`
		Message         string                `json:"message"`
	}
)
