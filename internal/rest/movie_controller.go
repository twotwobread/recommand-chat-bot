package rest

import (
	"fmt"
	"strconv"

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

	return c.Status(201).JSON(fiber.Map{"data": id})
}

func GetByID(c fiber.Ctx) error {
	ctx := c.Context()
	uc := c.Locals("UseCase").(domain.MovieUsecase)

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Sprintf("Wrong movie id param (can't convert %s to int64): %s", c.Params("id"), err.Error()))
	}

	m, err := uc.GetByID(ctx, int64(id))
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Sprintf("Wrong view movie data in usecase: %s", err.Error()))
	}

	return c.Status(201).JSON(fiber.Map{"data": m})
}

func GetAll(c fiber.Ctx) error {
	ctx := c.Context()
	uc := c.Locals("UseCase").(domain.MovieUsecase)

	movies, err := uc.GetAll(ctx)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Sprintf("Wrong view movie list data in usecase: %s", err.Error()))
	}
	return c.Status(201).JSON(fiber.Map{"data": movies})
}
