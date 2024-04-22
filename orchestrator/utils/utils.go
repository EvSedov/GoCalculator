package utils

import "strings"

var Secret = []byte("CTnRkgxS4QK_bRyl_8C51swvvDos9_Xr08A5_HwPE8c")

func Calc(left, right float64, sign string) float64 {
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

func DelSpaceFromString(str string) string {
	return strings.ReplaceAll(str, " ", "")
}

func IsValidExpression(expression string) bool {
	referenceString := "0123456789+-*/()"
	isValid := true
	for _, e := range expression {
		if !strings.ContainsRune(referenceString, e) {
			isValid = false
			return isValid
		}
	}

	if strings.Contains(expression, "**") || strings.Contains(expression, "--") ||
		strings.Contains(expression, "++") || strings.Contains(expression, "//") {
		isValid = false
	}

	return isValid
}
