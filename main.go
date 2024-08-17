package main

import (
	"log"
	"os"

	"recommand-chat-bot/internal/db"
	"recommand-chat-bot/internal/ent"
	"recommand-chat-bot/internal/rest"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {
	app := setupApp()
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	profile := os.Getenv("PROFILE")
	client, err := getDBClient(profile)
	if err != nil {
		log.Fatalf("Error setup db: %v", err)
	}

	restMapping(app)
	log.Fatal(app.Listen(":5003"))
}

func setupApp() *fiber.App {
	app := fiber.New()
	return app
}

func getDBClient(profile string) (*ent.Client, error) {
	if profile == "prod" {
		return db.InitPostgreDB()
	}
	return db.InitInMemDB()
}

func restMapping(a *fiber.App) {
	v1 := a.Group("/v1")
	v1.Get("/health", rest.CheckHealth)
}
