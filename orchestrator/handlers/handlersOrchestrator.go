package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/evsedov/GoCalculator/orchestrator/database"
	"github.com/evsedov/GoCalculator/orchestrator/models"
	"github.com/evsedov/GoCalculator/orchestrator/utils"
)

type (
	CreateExpressionRequest struct {
		RequestID  string `form:"request_id"`
		Expression string `form:"expression"`
	}

	OrchestratorHandler struct{}
)

func (h *OrchestratorHandler) CreateExpression(c *fiber.Ctx) error {
	var request CreateExpressionRequest

	if err := c.BodyParser(&request); err != nil {
		return fmt.Errorf("body parser: %w", err)
	}

	expression := new(models.Expression)

	expression.ExpressionId = uuid.New().String()
	expression.RequestID = request.RequestID
	expression.Expression = utils.DelSpaceFromString((request.Expression))

	if utils.IsValidExpression(expression.Expression) {
		expression.Message = "the expression will be calculated soon"
		expression.State = "valid"
	} else {
		expression.Message = "expression parsing error"
		expression.State = "error"
	}

	database.DB.Db.Create(&expression)

	return c.Status(200).JSON(expression)
}

func (h *OrchestratorHandler) GetExpressionById(c *fiber.Ctx) error {
	expression := models.Expression{}
	expressionId := c.Params("expression_id")
	database.DB.Db.Where("expression_id = ?", expressionId).First(&expression)

	return c.Status(200).JSON(expression)
}

func (h *OrchestratorHandler) GetExpressions(c *fiber.Ctx) error {
	expressions := []models.Expression{}
	database.DB.Db.Find(&expressions)

	return c.Status(200).JSON(expressions)
}
