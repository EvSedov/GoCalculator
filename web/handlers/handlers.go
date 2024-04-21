package handlers

import (
	// "encoding/json"
	// "errors"
	// "os"
	// "github.com/evsedov/GoCalculator/web/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
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

var store = session.New()

func (mh *MainHandler) GetHome(c *fiber.Ctx) error {

	sess, err := store.Get(c)
	if err != nil {
		return c.Redirect("/login")
	}

	auth := sess.Get("auth")
	if auth == nil {
		return c.Redirect("/login")
	}

	return c.Render("home", fiber.Map{
		"Title": "Distributed calculator",
		// "Expressions": expressions,
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
