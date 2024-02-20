package utils

import (
	"fmt"

	"github.com/Knetic/govaluate"
)

func CalcExpression(expression string) (string, error) {
	e, err := govaluate.NewEvaluableExpression(expression)
	if err != nil {
		return "", err
	}

	result, err := e.Evaluate(nil)
	if err != nil {
		return "", err
	}

	return string(fmt.Sprint(result)), nil
}
