package main

import (
	"log"
	"os"

	"recommand-chat-bot/domain"
	"recommand-chat-bot/internal/db"
	"recommand-chat-bot/internal/ent"
	"recommand-chat-bot/internal/repository"
	"recommand-chat-bot/internal/rest"
	"recommand-chat-bot/movie"

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

	restMapping(app, client)
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

func restMapping(a *fiber.App, client *ent.Client) {
	v1 := a.Group("/v1")
	uc := movie.NewMovieUsecase(repository.NewMovieRepository(client))
	v1.Use(DependencyMiddleware(&uc))

	v1.Get("/health", rest.CheckHealth)
	v1.Post("/movies", rest.Store)
}

func DependencyMiddleware(useCase *domain.MovieUsecase) fiber.Handler {
	return func(c fiber.Ctx) error {
		c.Locals("UseCase", useCase)
		return c.Next()
	}
}
