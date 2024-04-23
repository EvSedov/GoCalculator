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

	PostRespons struct {
		Expression client.Expression `json:"expression"`
		Message    string            `json:"message"`
	}
)

func (calc *CalculatorHendler) Start(ctx *fiber.Ctx) error {
	client := client.NewClient()

	GETResponse := client.GetExpression(ctx)
	fmt.Println("GET on /calculate response and calc expression =>", GETResponse)
	if GETResponse.Message != "calculated" {
		return ctx.Status(400).SendString("Error from GetExpression: " + GETResponse.Message)
	}

	data, err := json.Marshal(GETResponse.Expression)
	if err != nil {
		return ctx.Status(500).JSON("error: parse json")
	}

	POSTResponse, err := http.Post("http://orchestrator:8081/api/v1/calculate", "application/json", bytes.NewReader(data))
	if err != nil {
		return ctx.Status(500).SendString("error: send post response")
	}

	defer POSTResponse.Body.Close()
	var expression []byte
	body, err := io.ReadAll(POSTResponse.Body)
	if err != nil {
		return ctx.Status(500).SendString("error: read body to bytes")
	}

	err = json.Unmarshal(body, &expression)
	if err != nil {
		return ctx.Status(500).SendString("выражение посчитано и возвращено в оркестратор для обновления базы данных")
	}

	return ctx.Status(200).JSON(POSTResponse)
}
