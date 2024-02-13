package utils

func CalcExpression(left, right float64, sign string) float64 {
	var result float64
	switch sign {
	case "+":
		result = left + right
	case "-":
		result = left - right
	case "*":
		result = left * right
	case "/":
		result = left / right
	}

	return result
}
