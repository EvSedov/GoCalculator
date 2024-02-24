package handlers

import (
	"encoding/json"
	"errors"
	// "os"
	// "github.com/evsedov/GoCalculator/web/utils"
	"github.com/gofiber/fiber/v2"
)

type (
	Expression struct {
		ExpressionId string `json:"expression_id"`
		Expression   string `json:"expression"`
		Result       string `json:"result"`
		State        string `json:"state"`
		Message      string `json:"message"`
	}

	MainHandler struct{}
)

func (mh *MainHandler) GetMain(c *fiber.Ctx) error {
	
	var expressions []Expression
	URL := "http://orchestrator:8081/expressions"
	agent := fiber.Get(URL)
	agent.ContentType("application/json")
	_, body, errs := agent.Bytes()
	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	err := json.Unmarshal(body, &expressions)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	return c.Render("index", fiber.Map{
		"Title":       "Distributed calculator",
		"Expressions": expressions,
	})
}
