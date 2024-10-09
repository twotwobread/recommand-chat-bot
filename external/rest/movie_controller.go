package rest

import (
	"fmt"
	"strconv"

	"recommand-chat-bot/domain"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v3"
)

func Store(c fiber.Ctx) error {
	m := new(domain.Movie)
	if err := c.Bind().Body(m); err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Sprintf("Wrong movie data when bind req to domain: %s", err.Error()))
	}

	validate := c.Locals("MovieValidator").(*validator.Validate)
	if err := validate.Struct(m); err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Sprintf("Wrong movie validate results: %s", err.Error()))
	}

	ctx := c.Context()
	uc := c.Locals("MUC").(domain.MovieUsecase)
	id, err := uc.Store(ctx, m)
	if err != nil {
		return fiber.NewError(
			fiber.ErrInternalServerError.Code,
			fmt.Sprintf("Wrong movie data to save in usecase: %s", err.Error()))
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"data": fiber.Map{"id": id}})
}

func GetByID(c fiber.Ctx) error {
	ctx := c.Context()
	mr := c.Locals("MRepo").(domain.MovieRepository)

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Sprintf("Wrong movie id param (can't convert %s to int64): %s", c.Params("id"), err.Error()))
	}

	m, err := mr.GetByID(ctx, int64(id))
	if err != nil {
		return fiber.NewError(
			fiber.ErrInternalServerError.Code,
			fmt.Sprintf("Wrong view movie data in usecase: %s", err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": m})
}

func GetAll(c fiber.Ctx) error {
	ctx := c.Context()
	mr := c.Locals("MRepo").(domain.MovieRepository)

	movies, err := mr.GetAll(ctx)
	if err != nil {
		return fiber.NewError(
			fiber.ErrInternalServerError.Code,
			fmt.Sprintf("Wrong view movie list data in usecase: %s", err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": movies})
}

func GetRandom(c fiber.Ctx) error {
	ctx := c.Context()
	uc := c.Locals("MUC").(domain.MovieUsecase)

	movie, err := uc.GetRandom(ctx)
	if err != nil {
		return fiber.NewError(
			fiber.ErrInternalServerError.Code,
			fmt.Sprintf("Wrong view random movie in usecase: %s", err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": movie})
}
