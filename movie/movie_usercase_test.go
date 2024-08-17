package movie

import (
	"context"
	"testing"
	"time"

	"recommand-chat-bot/domain"
	"recommand-chat-bot/domain/mocks"
	"recommand-chat-bot/test/assertion"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MovieUsercaseSuite struct {
	suite.Suite
	m        domain.CreateMovieInput
	ctx      context.Context
	expected int64
}

func (suite *MovieUsercaseSuite) SetupSuite() {
	suite.ctx = context.Background()
	suite.m = domain.CreateMovieInput{
		Title:       "title",
		Genre:       domain.Action,
		Director:    "director",
		Actors:      []string{"Kim", "Lee"},
		Description: "description",
		ReleaseDate: releaseDate,
	}
	suite.expected = int64(1)
}

var releaseDate = domain.CustomTime{Time: time.Date(1998, 12, 3, 0, 0, 0, 0, time.UTC)}

func (suite *MovieUsercaseSuite) TestMovieUsecase_Store() {
	tests := []struct {
		name                string
		mockMovieRepository func() domain.MovieRepository
	}{
		{
			name: "Success movie usecase store",
			mockMovieRepository: func() domain.MovieRepository {
				movieRepo := &mocks.MockMovieRepository{}
				movieRepo.On("Store", suite.ctx, &suite.m).Return(suite.expected, nil).Once()
				return movieRepo
			},
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			movieUsecase := NewMovieUsecase(tt.mockMovieRepository())
			id, err := movieUsecase.Store(suite.ctx, &suite.m)
			assert.NoError(suite.T(), err)
			assert.Equal(suite.T(), suite.expected, id)
		})
	}
}

func (suite *MovieUsercaseSuite) TestMovieUsecase_GetByID() {
	tests := []struct {
		name          string
		mockMovieRepo func() domain.MovieRepository
	}{
		{
			name: "Success movie usecase GetByID",
			mockMovieRepo: func() domain.MovieRepository {
				mockMovieRepo := &mocks.MockMovieRepository{}
				mockMovieRepo.On("GetByID", suite.ctx, suite.expected).Return(domain.Movie{
					ID:          suite.expected,
					Title:       suite.m.Title,
					Genre:       suite.m.Genre,
					Director:    suite.m.Director,
					Actors:      suite.m.Actors,
					Description: suite.m.Description,
					ReleaseDate: suite.m.ReleaseDate,
					CreatedAt:   time.Now().UTC(),
					UpdatedAt:   time.Now().UTC(),
				}, nil).Once()
				return mockMovieRepo
			},
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			u := NewMovieUsecase(tt.mockMovieRepo())
			got, err := u.GetByID(suite.ctx, suite.expected)
			assert.NoError(suite.T(), err)
			expectedMovie := domain.Movie{
				Title:       suite.m.Title,
				Genre:       suite.m.Genre,
				Director:    suite.m.Director,
				Actors:      suite.m.Actors,
				Description: suite.m.Description,
				ReleaseDate: suite.m.ReleaseDate,
			}
			assertion.AssertMovieInputFields(suite.T(), expectedMovie, got)
		})
	}
}

func TestMovieUsercaseSuite(t *testing.T) {
	suite.Run(t, new(MovieUsercaseSuite))
}
