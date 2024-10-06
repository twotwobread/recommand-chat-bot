package rest

import (
	"fmt"

	"recommand-chat-bot/domain"

	"github.com/gofiber/fiber/v3"
)

func Store(c fiber.Ctx) error {
	m := new(domain.CreateMovieInput)
	if err := c.Bind().Body(m); err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Sprintf("Wrong movie data when bind req to domain: %s", err.Error()))
	}

	ctx := c.Context()
	uc := c.Locals("UseCase").(domain.MovieUsecase)
	id, err := uc.Store(ctx, m)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Sprintf("Wrong movie data to save in usecase: %s", err.Error()))
	}

	return c.Status(201).JSON(fiber.Map{"data": fiber.Map{"id": id}})
}
