package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/evsedov/GoCalculator/backend/storage"
)

type (
	CreateExpressionRequest struct {
		RequestID  string `form:"request_id"`
		Expression string `form:"expression"`
	}

	CreateExpressionResponse struct {
		RequestID    string `json:"request_id"`
		ExpressionId string `json:"expression_id"`
		Expression   string `json:"expression"`
		Message      string `json:"message"`
	}

	ExpressionHandler struct {
		Storage storage.ExpressionCreatorGetter
	}
)

func (h *ExpressionHandler) CreateExpression(c *fiber.Ctx) error {
	var request CreateExpressionRequest
	if err := c.BodyParser(&request); err != nil {
		return fmt.Errorf("body parser: %w", err)
	}

	expression := storage.Expression{
		ExpressionId: uuid.New().String(),
		RequestID:    request.RequestID,
		Expression:   request.Expression,
		State:        "in_progress",
	}

	id, err := h.Storage.CreateExpression(expression)
	if err != nil {
		return fmt.Errorf("create order: %w", err)
	}

	return c.Status(200).JSON(CreateExpressionResponse{
		ExpressionId: id,
		RequestID:    request.RequestID,
		Expression:   request.Expression,
		Message:      "the expression will be calculated soon",
	})
}

func (h *ExpressionHandler) GetExpressionById(c *fiber.Ctx) error {
	expression_id := c.Params("expression_id")

	expression, err := h.Storage.GetExpression(expression_id)
	if err != nil {
		return fmt.Errorf("get order: %w", err)
	}

	return c.JSON(expression)
}
