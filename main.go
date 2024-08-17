package main

import (
	"log"
	"os"

	"recommand-chat-bot/internal/db"
	"recommand-chat-bot/internal/rest"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {
	app := setupApp()
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
		panic(err)
	}

	profile := os.Getenv("PROFILE")
	if profile == "prod" {
		db.InitPostgreDB()
	} else {
		db.InitInMemDB()
	}

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
