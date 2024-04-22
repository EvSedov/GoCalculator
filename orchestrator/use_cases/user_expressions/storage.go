package userexpressions

type UserExpressionsGetter interface {
	GetExpressionsByEmail(email string) ([]byte, error)
}
