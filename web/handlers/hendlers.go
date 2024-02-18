package handlers

import (
	"github.com/evsedov/GoCalculator/web/storage"
	"github.com/gofiber/fiber/v2"
)

type (
	MainHandler struct {
		Storage storage.WebExpressionGetter
	}
)

func (mh *MainHandler) GetMain(c *fiber.Ctx) error {

	var exp storage.Expression
	c.BodyParser(exp)
	mh.Storage.SetExpressions(exp)
	expressions := mh.Storage.GetExpressions()

	return c.Render("index", fiber.Map{
		"Title":       "Distributed calculator",
		"Expressions": expressions,
	})
}

func (mh *MainHandler) GetExpressions(c *fiber.Ctx) error {
	var expressions []storage.Expression

	agent := fiber.Get("http://127.0.0.1:81/expressions")
	statusCode, _, errs := agent.Bytes()
	if len(errs) > 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errs": errs,
		})
	}
	err := c.BodyParser(expressions)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"err": err,
		})
	}

	return c.Status(statusCode).JSON(expressions)
}
