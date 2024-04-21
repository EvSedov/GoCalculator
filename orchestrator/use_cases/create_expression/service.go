package createexpression

import "github.com/evsedov/GoCalculator/orchestrator/entities"

type Service struct {
	expression ExpressionCreater
}

func NewService(expression ExpressionCreater) *Service {
	return &Service{
		expression: expression,
	}
}

func (s *Service) CreateExpression(expression entities.Expression) *Response {

	if err := s.expression.CreateNew(&expression); err != nil {
		return &Response{
			Message: err.Error(),
		}
	}

	return &Response{
		Message: "expression has been created",
	}
}
