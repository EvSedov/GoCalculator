package middleware

import (
	"errors"
	"fmt"

	"github.com/evsedov/GoCalculator/orchestrator/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthenticateMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Отсутствует заголовок 'Authorization'",
		})
	}
	if len(authHeader) < 7 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Токен не валидный, т.к. длина токена менее 7 символов",
		})
	}

	tokenString := authHeader[7:] // Удаляем префикс "Bearer "

	// Проверяем валидность и подпись токена
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			str := fmt.Sprintf("unexpected signing method: %v", token.Header["alg"])

			return nil, errors.New(str)
		}

		return []byte(utils.Secret), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Ошибка при проверке токена: " + err.Error(),
		})
	}

	// Получаем данные пользователя из токена и сохраняем их в контексте запроса
	claims := token.Claims.(jwt.MapClaims)
	email := claims["email"].(string)
	c.Locals("email", email)

	return c.Next()
}
