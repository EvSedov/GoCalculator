package calculateexpression

import (
	"encoding/json"

	"github.com/evsedov/GoCalculator/orchestrator/entities"
)

type Service struct {
	expression ExpressionEvaluator
}

func NewService(expression ExpressionEvaluator) *Service {
	return &Service{
		expression: expression,
	}
}

func (s *Service) GetExpressionToCalculate() *Response {
	var expression entities.Expression
	data, err := s.expression.GetToCalculate()
	if err != nil {
		return &Response{
			Message: err.Error(),
		}
	}

	err = json.Unmarshal(data, &expression)
	if err != nil {
		return &Response{
			Message: err.Error(),
		}
	}

	return &Response{
		Message:     "ok",
		Expressions: expression,
	}
}
