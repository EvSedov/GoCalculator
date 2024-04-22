package calculateexpression

type ExpressionEvaluator interface {
	GetToCalculate() ([]byte, error)
}
