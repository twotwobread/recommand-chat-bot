package repository

import (
	"context"

	"recommand-chat-bot/domain"
	"recommand-chat-bot/internal/ent"
	"recommand-chat-bot/internal/ent/movie"
)

type movieRepository struct {
	client *ent.Client
}

func NewMovieRepository(client *ent.Client) *movieRepository {
	return &movieRepository{client: client}
}

func (r *movieRepository) Store(ctx context.Context, m *domain.CreateMovieInput) (int64, error) {
	movie, err := r.client.Movie.
		Create().
		SetTitle(m.Title).
		SetGenre(string(m.Genre)).
		SetDirector(m.Director).
		SetActors(m.Actors).
		SetDescription(m.Description).
		SetReleaseDate(m.ReleaseDate.Time).
		Save(ctx)
	if err != nil {
		return -1, err
	}

	return movie.ID, nil
}

func (r *movieRepository) GetByID(ctx context.Context, id int64) (*domain.Movie, error) {
	m, err := r.client.Movie.
		Query().
		Where(movie.ID(id)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	newMovie := domain.Movie{
		ID:          m.ID,
		Title:       m.Title,
		Genre:       domain.Genre(m.Genre),
		Director:    m.Director,
		Actors:      m.Actors,
		Description: m.Description,
		ReleaseDate: domain.CustomTime{Time: m.ReleaseDate},
		UpdatedAt:   m.UpdatedAt,
		CreatedAt:   m.CreatedAt,
	}
	return &newMovie, nil
}
