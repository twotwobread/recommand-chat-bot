package repository

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"recommand-chat-bot/domain"
	"recommand-chat-bot/internal/db"
	"recommand-chat-bot/internal/ent"
)

type MovieRepoSuite struct {
	suite.Suite
	client *ent.Client
	ctx    context.Context
}

func (suite *MovieRepoSuite) SetupSuite() {
	client, err := db.InitInMemDB()
	suite.Require().NoError(err)
	suite.client = client
	suite.ctx = context.Background()
}

func (suite *MovieRepoSuite) TearDownSuite() {
	suite.client.Close()
}

var releaseDate = domain.CustomTime{Time: time.Date(2023, 12, 15, 0, 0, 0, 0, time.UTC)}

func (suite *MovieRepoSuite) TestMovieInMemRepositoryStore() {
	tests := []struct {
		name    string
		m       *domain.CreateMovieInput
		wantErr bool
	}{
		{
			name: "Success store movie",
			m: &domain.CreateMovieInput{
				Title:       "title",
				Genre:       domain.Action,
				Director:    "director",
				Actors:      []string{"Kim", "Lee"},
				Description: "description",
				ReleaseDate: releaseDate,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			r := NewMovieRepository(suite.client)
			mId, err := r.Store(suite.ctx, tt.m)
			assert.NoError(suite.T(), err)
			assert.NotEqual(suite.T(), nil, mId)
		})
	}
}

func (suite *MovieRepoSuite) TestMovieInMemRepositoryGetByID() {
	tests := []struct {
		name    string
		m       *domain.CreateMovieInput
		wantErr bool
	}{
		{
			name: "Success store movie",
			m: &domain.CreateMovieInput{
				Title:       "title",
				Genre:       domain.Action,
				Director:    "director",
				Actors:      []string{"Kim", "Lee"},
				Description: "description",
				ReleaseDate: releaseDate,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			r := NewMovieRepository(suite.client)
			mId, err := r.Store(suite.ctx, tt.m)
			assert.NoError(suite.T(), err)

			movie, err := r.GetByID(suite.ctx, mId)
			assert.NoError(suite.T(), err)

			assert.Equal(suite.T(), tt.m.Title, movie.Title)
			assert.Equal(suite.T(), tt.m.Genre, movie.Genre)
			assert.Equal(suite.T(), tt.m.Director, movie.Director)
			assert.Equal(suite.T(), tt.m.Actors, movie.Actors)
			assert.Equal(suite.T(), tt.m.Description, movie.Description)
			assert.Equal(suite.T(), tt.m.ReleaseDate, movie.ReleaseDate)
		})
	}
}

func TestMovieRepositorySuite(t *testing.T) {
	suite.Run(t, new(MovieRepoSuite))
}
