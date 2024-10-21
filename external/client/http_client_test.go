package client

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"recommand-chat-bot/batch"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

var tmdbToken string

func setup() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	tmdbToken = os.Getenv("TMDB_TOKEN")
	fmt.Printf("Token: %s", tmdbToken)
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func Test_httpClient_Get(t *testing.T) {
	type args struct {
		url     string
		headers map[string]string
		timeout time.Duration
	}
	type result struct {
		page    int
		results int
		pages   int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantRes result
	}{
		{
			name: "Http client get test",
			args: args{
				url: "https://api.themoviedb.org/3/discover/movie?page=1&primary_release_date.gte=2024-09-10&primary_release_date.lte=2024-10-10",
				headers: map[string]string{
					"accept":        "application/json",
					"Authorization": fmt.Sprintf("Bearer %s", tmdbToken),
				},
				timeout: time.Duration(30 * time.Second),
			},
			want: fasthttp.StatusOK,
			wantRes: result{
				page:    1,
				results: 3619,
				pages:   181,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewHttpClient()
			got, got1, err := c.Get(tt.args.url, tt.args.headers, tt.args.timeout)

			assert.NoError(t, err)
			assert.Nil(t, err, "err not nil: %v", err)
			assert.Equal(t, tt.want, got, "status code not ok: %v", got)
			assert.NotNil(t, got, "response is nil")

			var response batch.TMDBResponse
			err = json.Unmarshal(got1, &response)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantRes.page, response.Page, "page not equal: %v", response)
			assert.Equal(t, tt.wantRes.results, response.TotalResults, "results count not equal: %v", response)
			assert.Equal(t, tt.wantRes.pages, response.TotalPages, "pages count not equal: %v", response)
		})
	}
}
