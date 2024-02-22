package client

import (
	"encoding/json"
	"errors"

	// "errors"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/evsedov/GoCalculator/calculator/utils"
	"github.com/gofiber/fiber/v2"
)

type (
	ExpressionToWork struct {
		ExpressionId string `json:"expression_id"`
		Expression   string `json:"expression"`
		Result       string `json:"result"`
		State        string `json:"state"`
		Message      string `json:"message"`
	}
)

const (
	// BaseURL = "http://172.18.0.1:81/task"
	BaseURL = "http://172.24.113.173:81/task"
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

func (c *Client) GetExpressionToWork(ctx *fiber.Ctx) (ExpressionToWork, error) {

	var expression ExpressionToWork
	var body []byte
	var errs []error

	request := fiber.Get(BaseURL)
	request.ContentType("application/json")
	_, body, errs = request.Bytes()

	if len(errs) > 0 {
		return ExpressionToWork{}, errors.Join(errs...)
	}

	// читаем тело ответа
	err := json.Unmarshal(body, &expression)
	if err != nil {
		return ExpressionToWork{}, err
	}

	result, _ := utils.CalcExpression(expression.Expression)
	if result == "" {
		return ExpressionToWork{}, errors.New("not valid expressions on server, result is empty string")
	}

	expression.Result = result
	expression.State = "ok"
	expression.Message = "the expression is calculated"

	return expression, nil
}
