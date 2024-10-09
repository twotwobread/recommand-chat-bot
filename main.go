package main

import (
	"log"
	"os"

	vld "recommand-chat-bot/domain/validator"
	"recommand-chat-bot/external/db"
	"recommand-chat-bot/external/ent"
	"recommand-chat-bot/external/repository"
	"recommand-chat-bot/external/rest"
	"recommand-chat-bot/movie"

	"github.com/go-playground/validator"
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

	registerValidators(app)
	restMapping(app, client)
	log.Fatal(app.Listen(":5003"))
}

func setupApp() *fiber.App {
	// Setup your validator in the config
	app := fiber.New()
	return app
}

func getDBClient(profile string) (*ent.Client, error) {
	if profile == "prod" {
		return db.InitPostgreDB()
	}
	return db.InitInMemDB()
}

func registerValidators(app *fiber.App) {
	v := validator.New()
	if err := vld.RegisterMovieValidation(v); err != nil {
		log.Fatalf("Impossible register movie validators: %v", err)
	}

	app.Use(func(c fiber.Ctx) error {
		c.Locals("MovieValidator", v)
		return c.Next()
	})
}

func restMapping(app *fiber.App, client *ent.Client) {
	v1 := app.Group("/v1")
	mr := repository.NewMovieRepository(client)
	uc := movie.NewMovieUsecase(mr)
	v1.Use(func(c fiber.Ctx) error {
		c.Locals("MUC", uc)
		return c.Next()
	})
	v1.Use(func(c fiber.Ctx) error {
		c.Locals("MRepo", mr)
		return c.Next()
	})

	v1.Get("/health", rest.CheckHealth)

	v1.Post("/movies", rest.Store)
	v1.Get("/movies", rest.GetAll)
	v1.Get("/movies/random", rest.GetRandom)
	v1.Get("/movies/:id", rest.GetByID)
}
