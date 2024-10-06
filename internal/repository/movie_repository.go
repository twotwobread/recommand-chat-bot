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

func NewMovieRepository(client *ent.Client) domain.MovieRepository {
	return &movieRepository{client: client}
}

func (r *movieRepository) Store(ctx context.Context, m *domain.Movie) (int64, error) {
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
		return 0, err
	}

	return movie.ID, nil
}

func (r *movieRepository) GetByID(ctx context.Context, id int64) (domain.MovieDetailOutput, error) {
	m, err := r.client.Movie.
		Query().
		Where(movie.ID(id)).
		Only(ctx)
	if err != nil {
		return domain.MovieDetailOutput{}, err
	}

	newMovie := domain.MovieDetailOutput{
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
	return newMovie, nil
}

func (r *movieRepository) GetAll(ctx context.Context) ([]domain.MovieDetailOutput, error) {
	gots, err := r.client.Movie.Query().All(ctx)
	if err != nil {
		return []domain.MovieDetailOutput{}, err
	}

	movies := []domain.MovieDetailOutput{}
	for _, m := range gots {
		m := domain.MovieDetailOutput{
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
		movies = append(movies, m)
	}
	return movies, nil
}
