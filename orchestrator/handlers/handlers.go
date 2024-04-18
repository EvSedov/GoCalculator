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

	CreateExpressionWork struct {
		ExpressionId string `json:"expression_id"`
		Expression   string `json:"expression"`
		Result       string `json:"result"`
		State        string `json:"state"`
		Message      string `json:"message"`
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

func (h *OrchestratorHandler) GetValidExpressionToWork(c *fiber.Ctx) error {
	var response CreateExpressionWork

	expression := models.Expression{}
	database.DB.Db.Where("state = ?", "valid").First(&expression)

	response.Expression = expression.Expression
	response.ExpressionId = expression.ExpressionId

	expression.State = "in_process"
	database.DB.Db.Where("expression_id = ?", expression.ExpressionId).Updates(expression)

	return c.Status(200).JSON(response)
}

func (h *OrchestratorHandler) UpdateExpressionInWork(c *fiber.Ctx) error {
	var request CreateExpressionWork
	expression := models.Expression{}

	if err := c.BodyParser(&request); err != nil {
		return fmt.Errorf("body parser: %w", err)
	}
	database.DB.Db.Where("expression_id = ?", request.ExpressionId).First(&expression)
	expression.State = request.State
	expression.Result = request.Result
	expression.Message = request.Message

	database.DB.Db.Where("expression_id = ?", request.ExpressionId).Updates(&expression)

	return c.Status(200).JSON(expression)
}
