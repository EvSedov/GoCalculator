package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/evsedov/GoCalculator/calculator/client"
	"github.com/gofiber/fiber/v2"
)

type (
	CalculatorHendler struct {
	}

	CreateRegistrationRespons struct {
		Expression string `json:"expression"`
	}
)

func (calc *CalculatorHendler) Start(ctx *fiber.Ctx) error {
	client := client.NewClient()

	expression, err := client.GetExpressionToWork(ctx)
	if err != nil {
		return ctx.Status(400).SendString("Error from GetExpressionToWork: " + err.Error())
	}

	fmt.Println("GET on /task response and calc expression =>", expression)

	jsonData, err := json.Marshal(expression)
	if err != nil {
		return ctx.Status(500).SendString("error: parse json")
	}

	resp, err := http.Post("http://orchestrator:8081/task", "application/json", bytes.NewReader(jsonData))
	if err != nil {
		return ctx.Status(500).SendString("error: send post response")
	}

	defer resp.Body.Close()
	var expressionFromServer []byte
	responseBytes, err := io.ReadAll(resp.Body)

	fmt.Println("responseBytes", string(responseBytes))

	if err != nil {
		return ctx.Status(500).SendString("error: read body to bytes")
	}

	err = json.Unmarshal(responseBytes, &expressionFromServer)
	if err != nil {
		return ctx.Status(500).SendString("выражение посчитано и возвращено в оркестратор для обновления базы данных")
	}

	return ctx.Status(200).JSON(expressionFromServer)
}
