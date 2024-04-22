package updatecalculated

import (
	"github.com/evsedov/GoCalculator/orchestrator/entities"
)

type Service struct {
	expression ExpressionUpdater
}

func NewService(expression ExpressionUpdater) *Service {
	return &Service{
		expression: expression,
	}
}

func (s *Service) UpdateCalculatedExpression(expression *entities.Expression) *Response {
	err := s.expression.UpdateCalculated(expression)
	if err != nil {
		return &Response{
			Message: err.Error(),
		}
	}

	return &Response{
		Message: "ok",
	}
}
