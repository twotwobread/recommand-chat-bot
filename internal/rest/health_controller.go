package rest

import "github.com/gofiber/fiber/v3"

func CheckHealth(c fiber.Ctx) error {
	return c.SendString("Check Health")
}
