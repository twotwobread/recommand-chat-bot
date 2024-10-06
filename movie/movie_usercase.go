package movie

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"recommand-chat-bot/domain"
)

type movieUsecase struct {
	repo domain.MovieRepository
}

func NewMovieUsecase(repo domain.MovieRepository) domain.MovieUsecase {
	return &movieUsecase{repo: repo}
}

func (u *movieUsecase) Store(ctx context.Context, m *domain.Movie) (int64, error) {
	got, err := u.repo.Store(ctx, m)
	if err != nil {
		return -1, fmt.Errorf("failed to store movie - %v, input: %v", err, m)
	}

	return got, nil
}

func (u movieUsecase) GetRandom(ctx context.Context) (domain.MovieDetailOutput, error) {
	cnt, err := u.repo.GetCount(ctx)
	if err != nil {
		return domain.MovieDetailOutput{}, fmt.Errorf("failed to get movie count - %v", err)
	}
	if cnt <= 0 {
		return domain.MovieDetailOutput{}, fmt.Errorf("failed to get movie count - movie count is zero")
	}

	for {
		seed := time.Now().UnixNano()
		r := rand.New(rand.NewSource(seed))
		rNum := int64(r.Intn(cnt) + 1)

		m, err := u.repo.GetByID(ctx, rNum)
		if err == nil {
			return m, nil
		}
	}
}
