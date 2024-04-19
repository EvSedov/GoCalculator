package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/evsedov/GoCalculator/orchestrator/entities"
	// "github.com/evsedov/GoCalculator/orchestrator/storage"
	"github.com/evsedov/GoCalculator/orchestrator/utils"
)

type (
	CreateExpressionRequest struct {
		RequestID  string `form:"request_id"`
		Expression string `form:"expression"`
	}

	CreateExpressionWork struct {
		ExpressionId string `json:"expression_id"`
		Expression   string `json:"expression"`
		Result       string `json:"result"`
		State        string `json:"state"`
		Message      string `json:"message"`
	}

	Handler struct{}
)

func (h *Handler) CreateExpression(c *fiber.Ctx) error {
	var request CreateExpressionRequest

	if err := c.BodyParser(&request); err != nil {
		return fmt.Errorf("body parser: %w", err)
	}

	expression := new(entities.Expression)

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

	// h.Storage.DB.Create(&expression)

	return c.Status(200).JSON(expression)
}

func (h *Handler) GetExpressionById(c *fiber.Ctx) error {
	expression := entities.Expression{}
	// expressionId := c.Params("expression_id")
	// h.Storage.DB.Where("expression_id = ?", expressionId).First(&expression)

	return c.Status(200).JSON(expression)
}

func (h *Handler) GetExpressions(c *fiber.Ctx) error {
	expressions := []entities.Expression{}
	// h.Storage.DB.Find(&expressions)

	return c.Status(200).JSON(expressions)
}

func (h *Handler) GetValidExpressionToWork(c *fiber.Ctx) error {
	var response CreateExpressionWork

	expression := entities.Expression{}
	// h.Storage.DB.Where("state = ?", "valid").First(&expression)

	response.Expression = expression.Expression
	response.ExpressionId = expression.ExpressionId

	expression.State = "in_process"
	// h.Storage.DB.Where("expression_id = ?", expression.ExpressionId).Updates(expression)

	return c.Status(200).JSON(response)
}

func (h *Handler) UpdateExpressionInWork(c *fiber.Ctx) error {
	var request CreateExpressionWork
	expression := entities.Expression{}

	if err := c.BodyParser(&request); err != nil {
		return fmt.Errorf("body parser: %w", err)
	}
	// h.Storage.DB.Where("expression_id = ?", request.ExpressionId).First(&expression)
	expression.State = request.State
	expression.Result = request.Result
	expression.Message = request.Message

	// h.Storage.DB.Where("expression_id = ?", request.ExpressionId).Updates(&expression)

	return c.Status(200).JSON(expression)
}
