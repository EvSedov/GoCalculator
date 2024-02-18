package storage

import (
	"errors"
)

type (
	Expression struct {
		RequestID    string `json:"request_id"`
		ExpressionId string `json:"expression_id"`
		Expression   string `json:"expression"`
		State        string `json:"state"`
		Result       string `json:"result"`
		Message      string `json:"message"`
	}

	ExpressionStorage struct {
		Expressions map[string]Expression
		Ids         []string
	}

	ExpressionCreatorGetter interface {
		CreateExpression(expression Expression) (string, error)
		GetExpressionById(ExpressionId string) (Expression, error)
		GetIds() []string
	}
)

func (e *ExpressionStorage) CreateExpression(expression Expression) (string, error) {
	e.Expressions[expression.ExpressionId] = expression
	e.Ids = append(e.Ids, expression.ExpressionId)

	return expression.ExpressionId, nil
}

var (
	errExpressionNotFound = errors.New("expression not found")
)

func (e *ExpressionStorage) GetExpressionById(ExpressionId string) (Expression, error) {
	expression, ok := e.Expressions[ExpressionId]
	if !ok {
		return Expression{}, errExpressionNotFound
	}

	return expression, nil
}

func (e *ExpressionStorage) GetIds() []string {

	return e.Ids
}
