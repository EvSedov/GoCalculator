package client

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/evsedov/GoCalculator/calculator/utils"
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

	Response struct {
		Expression Expression `json:"expression"`
		Message    string     `json:"message"`
	}
)

const (
	BaseURL = "http://orchestrator:8081/api/v1/calculate"
)

type Client struct {
	BaseURL    string
	Delay      int
	HTTPClient *http.Client
}

func NewClient() *Client {
	delay, _ := strconv.Atoi(os.Getenv("delay_calculator"))
	return &Client{
		BaseURL: BaseURL,
		Delay:   delay,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *Client) GetExpression(ctx *fiber.Ctx) Response {

	var getResponse Response
	var body []byte
	var errs []error

	request := fiber.Get(BaseURL)
	request.ContentType("application/json")
	_, body, errs = request.Bytes()
	if len(errs) > 0 {
		return Response{
			Message: errors.Join(errs...).Error(),
		}
	}

	// читаем тело ответа
	err := json.Unmarshal(body, &getResponse)
	if err != nil {
		return Response{
			Message: err.Error(),
		}
	}

	result, _ := utils.CalcExpression(getResponse.Expression.Expression)
	if result == "" {
		return Response{
			Message: "not valid expressions on server, result is empty string",
		}
	}

	getResponse.Expression.Result = result
	getResponse.Expression.State = "ok"
	getResponse.Expression.Message = "the expression is calculated"
	getResponse.Message = "calculated"

	return getResponse
}
