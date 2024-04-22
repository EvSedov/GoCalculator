package calculateexpression

import "github.com/evsedov/GoCalculator/orchestrator/entities"

type (
	Response struct {
		Expressions entities.Expression `json:"expressions"`
		Message     string              `json:"message"`
	}
)
