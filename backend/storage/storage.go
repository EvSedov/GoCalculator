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
	}

	ExpressionStorage struct {
		Expressions map[string]Expression
	}

	ExpressionCreatorGetter interface {
		CreateExpression(expression Expression) (string, error)
		GetExpression(ExpressionId string) (Expression, error)
	}
)

func (e *ExpressionStorage) CreateExpression(expression Expression) (string, error) {
	e.Expressions[expression.ExpressionId] = expression

	return expression.ExpressionId, nil
}

var (
	errExpressionNotFound = errors.New("expression not found")
)

func (e *ExpressionStorage) GetExpression(ExpressionId string) (Expression, error) {
	expression, ok := e.Expressions[ExpressionId]
	if !ok {
		return Expression{}, errExpressionNotFound
	}

	return expression, nil
}
