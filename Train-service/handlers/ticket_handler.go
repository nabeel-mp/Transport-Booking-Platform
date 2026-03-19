package handlers

import "github.com/gofiber/fiber/v3"

func GetTicket() fiber.Handler {
	return func(c fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"message": "get ticket - coming in phase 4"})
	}
}

func VerifyTicket() fiber.Handler {
	return func(c fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"message": "verify ticket - coming in phase 4"})
	}
}
