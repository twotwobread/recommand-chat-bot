package movie

import (
	"context"
	"fmt"

	"recommand-chat-bot/domain"
)

type movieUsecase struct {
	repo domain.MovieRepository
}

func NewMovieUsecase(repo domain.MovieRepository) domain.MovieUsecase {
	return &movieUsecase{repo: repo}
}

func (u *movieUsecase) Store(ctx context.Context, m *domain.CreateMovieInput) (int64, error) {
	got, err := u.repo.Store(ctx, m)
	if err != nil {
		return -1, fmt.Errorf("failed to store movie - %v, input: %v", err, m)
	}

	return got, nil
}

func (u *movieUsecase) GetByID(ctx context.Context, id int64) (domain.Movie, error) {
	got, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return domain.Movie{}, fmt.Errorf("failed to store movie - %v, input: %v", err, id)
	}

	return got, nil
}

func (u *movieUsecase) GetAll(ctx context.Context) ([]domain.Movie, error) {
	got, err := u.repo.GetAll(ctx)
	if err != nil {
		return []domain.Movie{}, fmt.Errorf("failed to get movies - %v", err)
	}

	return got, nil
}
