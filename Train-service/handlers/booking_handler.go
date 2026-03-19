package handlers

import "github.com/gofiber/fiber/v3"

func BookTrain() fiber.Handler {
	return func(c fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"message": "book train - coming in phase 4"})
	}
}

func GetBooking() fiber.Handler {
	return func(c fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"message": "get booking - coming in phase 4"})
	}
}

func CancelBooking() fiber.Handler {
	return func(c fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"message": "cancel booking - coming in phase 4"})
	}
}
