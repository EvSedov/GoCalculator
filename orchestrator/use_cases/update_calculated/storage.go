package updatecalculated

import "github.com/evsedov/GoCalculator/orchestrator/entities"

type ExpressionUpdater interface {
	UpdateCalculated(expression *entities.Expression) error
}
