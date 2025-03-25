package response

import "github.com/gofiber/fiber/v2"

func SuccessPayload(msg any, data any) fiber.Map {
	return fiber.Map{
		"data":    data,
		"message": msg,
		"success": true,
	}
}

func FailedPayload(msg any, data any) fiber.Map {
	return fiber.Map{
		"data":    data,
		"message": msg,
		"success": false,
	}
}
