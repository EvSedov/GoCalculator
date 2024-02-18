package storage

type (
	Expression struct {
		RequestID    string `json:"request_id"`
		ExpressionId string `json:"expression_id"`
		Expression   string `json:"expression"`
		State        string `json:"state"`
		Result       string `json:"result"`
		Message      string `json:"message"`
	}

	WebExpressionStorage struct {
		Expressions []Expression `json:"expressions"`
	}

	WebExpressionGetter interface {
		GetExpressions() []Expression
		SetExpressions(expression Expression)
	}
)

func (webs *WebExpressionStorage) GetExpressions() []Expression {

	return webs.Expressions
}

func (webs *WebExpressionStorage) SetExpressions(expression Expression) {
	webs.Expressions = append(webs.Expressions, expression)
}
