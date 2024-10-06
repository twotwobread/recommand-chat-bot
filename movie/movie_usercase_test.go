package movie

import (
	"context"
	"testing"
	"time"

	"recommand-chat-bot/domain"
	"recommand-chat-bot/domain/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MovieUsercaseSuite struct {
	suite.Suite
	m        domain.Movie
	ctx      context.Context
	expected int64
}

func (suite *MovieUsercaseSuite) SetupSuite() {
	suite.ctx = context.Background()
	suite.m = domain.Movie{
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

func TestMovieUsercaseSuite(t *testing.T) {
	suite.Run(t, new(MovieUsercaseSuite))
}
