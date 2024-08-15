package main

import (
	"log"

	"recommand-chat-bot/internal/rest"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := setupApp()
	restMapping(app)
	log.Fatal(app.Listen(":5003"))
}

func setupApp() *fiber.App {
	app := fiber.New()
	return app
}

func restMapping(a *fiber.App) {
	v1 := a.Group("/v1")
	v1.Get("/health", rest.CheckHealth)
}
