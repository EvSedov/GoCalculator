package calculateexpression

import "github.com/evsedov/GoCalculator/orchestrator/entities"

type (
	Response struct {
		Expression entities.Expression `json:"expression"`
		Message    string              `json:"message"`
	}
)
