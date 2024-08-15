package main

import (
	"log"

	"recommand-chat-bot/internal/rest"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()
	restMapping(app)
	log.Fatal(app.Listen(":5003"))
}

func restMapping(a *fiber.App) {
	v1 := a.Group("/v1")
	v1.Get("/health", rest.CheckHealth)
}
