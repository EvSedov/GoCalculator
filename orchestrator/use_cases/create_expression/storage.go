package createexpression

import "github.com/evsedov/GoCalculator/orchestrator/entities"

type ExpressionCreater interface {
	CreateNew(expression entities.Expression) error
}
