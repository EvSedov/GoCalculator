package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/evsedov/GoCalculator/orchestrator/storage"
	"github.com/evsedov/GoCalculator/orchestrator/utils"
)

type (
	CreateExpressionRequest struct {
		RequestID  string `form:"request_id"`
		Expression string `form:"expression"`
	}

	CreateExpressionResponse struct {
		Expressions []storage.Expression `json:"expressions"`
	}

	OrchestratorHandler struct {
		Storage storage.ExpressionCreatorGetter
	}
)

func (h *OrchestratorHandler) CreateExpression(c *fiber.Ctx) error {
	var request CreateExpressionRequest

	if err := c.BodyParser(&request); err != nil {
		return fmt.Errorf("body parser: %w", err)
	}

	expression := storage.Expression{
		ExpressionId: uuid.New().String(),
		RequestID:    request.RequestID,
		Expression:   request.Expression,
	}
	exp := utils.DelSpaceFromString((expression.Expression))

	if utils.IsValidExpression(exp) {
		expression.Message = "the expression will be calculated soon"
		expression.State = "valid"
	} else {
		expression.Message = "expression parsing error"
		expression.State = "error"
		// return c.Status(400).JSON(expression)
	}

	id, err := h.Storage.CreateExpression(expression)
	if err != nil {
		return c.SendStatus(500)
	}

	newExpression, _ := h.Storage.GetExpressionById(id)

	return c.Status(200).JSON(newExpression)
}

func (h *OrchestratorHandler) GetExpressionById(c *fiber.Ctx) error {
	expression_id := c.Params("expression_id")

	expression, err := h.Storage.GetExpressionById(expression_id)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.Status(200).JSON(expression)
}

func (h *OrchestratorHandler) GetExpressions(c *fiber.Ctx) error {
	expressions := make([]storage.Expression, 0)
	ids := h.Storage.GetIds()

	for _, id := range ids {
		expression, err := h.Storage.GetExpressionById(id)
		if err != nil {
			return c.SendStatus(500)
		}

		expressions = append(expressions, expression)
	}

	return c.Status(200).JSON(expressions)
}
