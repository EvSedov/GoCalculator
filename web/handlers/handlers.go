package handlers

import (
	// "encoding/json"
	// "errors"
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

func (mh *MainHandler) GetHome(c *fiber.Ctx) error {

	// var expressions []Expression
	// URL := "http://orchestrator:8081/expressions"
	// agent := fiber.Get(URL)
	// agent.ContentType("application/json")
	// _, body, errs := agent.Bytes()
	// if len(errs) > 0 {
	// 	return errors.Join(errs...)
	// }

	// err := json.Unmarshal(body, &expressions)
	// if err != nil {
	// 	return c.Status(400).SendString(err.Error())
	// }

	return c.Render("home", fiber.Map{
		"Title": "Distributed calculator",
		// "Expressions": expressions,x
	})
}

func (mh *MainHandler) GetLogin(c *fiber.Ctx) error {

	return c.Render("login", fiber.Map{
		"Title": "Distributed calculator",
	})
}

func (mh *MainHandler) GetRegister(c *fiber.Ctx) error {

	return c.Render("register", fiber.Map{
		"Title": "Distributed calculator",
	})
}
