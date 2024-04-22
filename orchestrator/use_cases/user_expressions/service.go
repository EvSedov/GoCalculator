package userexpressions

import (
	"encoding/json"

	"github.com/evsedov/GoCalculator/orchestrator/entities"
)

type Service struct {
	storage UserExpressionsGetter
}

func NewService(storage UserExpressionsGetter) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) GetUserExpressions(email string) *Response {
	expressions := []entities.Expression{}
	data, err := s.storage.GetExpressionsByEmail(email)
	if err != nil {
		return &Response{
			Message: err.Error(),
		}
	}

	err = json.Unmarshal(data, &expressions)
	if err != nil {
		return &Response{
			Message: err.Error(),
		}
	}

	return &Response{
		Message:         "ok",
		UserExpressions: expressions,
	}
}
