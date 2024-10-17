package batch

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"recommand-chat-bot/domain"

	"github.com/valyala/fasthttp"
)

const baseURL = "https://api.themoviedb.org/3/discover/movie"

var tmdbToken = os.Getenv("TMDB_TOKEN")

type TMDBMovie struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	ReleaseDate string `json:"release_date"`
}

type TMDBResponse struct {
	Page         int         `json:"page"`
	TotalResults int         `json:"total_results"`
	TotalPages   int         `json:"total_pages"`
	Results      []TMDBMovie `json:"results"`
}

type tmdbBatchUsecase struct {
	Client domain.HttpClient
}

func (t tmdbBatchUsecase) Process() {
	endDate := time.Now().Format("2006-01-02")
	startDate := time.Now().AddDate(-3, 0, 0).Format("2006-01-02")

	allMovies := t.fetchAllMovies(startDate, endDate)
	fmt.Printf("Total movies fetched: %d\n", len(allMovies))
}

func (t tmdbBatchUsecase) fetchAllMovies(startDate, endDate string) []TMDBMovie {
	var allMovies []TMDBMovie
	page := 1

	for {
		response := t.fetchPage(page, startDate, endDate)
		allMovies = append(allMovies, response.Results...)

		if page >= response.TotalPages {
			break
		}

		page += 1
	}

	return allMovies
}

func (t tmdbBatchUsecase) fetchPage(page int, startDate, endDate string) TMDBResponse {
	params := url.Values{}
	params.Set("sort_by", "release_date.desc")
	params.Set("include_adult", "true")
	params.Set("include_video", "false")
	params.Set("page", fmt.Sprintf("%d", page))
	params.Set("primary_release_date.gte", startDate)
	params.Set("primary_release_date.lte", endDate)

	url := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	statusCode, body, err := t.Client.Get(
		url,
		map[string]string{
			"accept":        "application/json",
			"Authorization": fmt.Sprintf("Bearer %s", tmdbToken),
		},
		time.Duration(30),
	)
	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}

	if statusCode != fasthttp.StatusOK {
		log.Fatalf("Unexpected status code: %d", statusCode)
	}

	var response TMDBResponse
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	return response
}

func NewTmdbBatchUsecase(client domain.HttpClient) domain.BatchUsecase {
	return &tmdbBatchUsecase{Client: client}
}
