package repository

import (
	"context"
	"reflect"
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

func (suite *MovieRepoSuite) assertMovieInputFields(expected, actual *domain.Movie) {
	// Reflection-based field comparison
	vExpected := reflect.ValueOf(expected).Elem()
	vActual := reflect.ValueOf(actual).Elem()

	// Iterate over the struct fields and assert equality
	for i := 0; i < vExpected.NumField(); i++ {
		fieldName := vExpected.Type().Field(i).Name

		// Skip auto-generated fields like ID, CreatedAt, UpdatedAt
		if fieldName == "ID" || fieldName == "CreatedAt" || fieldName == "UpdatedAt" {
			continue
		}

		expectedValue := vExpected.Field(i).Interface()
		actualValue := vActual.Field(i).Interface()
		suite.Equalf(expectedValue, actualValue, "Field %s should be equal", fieldName)
	}
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
			assert.NotZero(suite.T(), mId)
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

			assert.NotZero(suite.T(), movie.ID, "ID should be set by the database")
			assert.False(suite.T(), movie.CreatedAt.IsZero(), "CreatedAt should be set by the database")
			assert.False(suite.T(), movie.UpdatedAt.IsZero(), "UpdatedAt should be set by the database")

			expectedMovie := &domain.Movie{
				Title:       tt.m.Title,
				Genre:       tt.m.Genre,
				Director:    tt.m.Director,
				Actors:      tt.m.Actors,
				Description: tt.m.Description,
				ReleaseDate: tt.m.ReleaseDate,
			}

			// Compare user-provided fields only
			suite.assertMovieInputFields(expectedMovie, movie)
		})
	}
}

func TestMovieRepositorySuite(t *testing.T) {
	suite.Run(t, new(MovieRepoSuite))
}