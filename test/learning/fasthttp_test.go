package learning

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/valyala/fasthttp"
)

var TMDB_TOKEN string

func setup() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	TMDB_TOKEN = os.Getenv("TMDB_TOKEN")
	fmt.Printf("Token: %s", TMDB_TOKEN)
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func TestGet(t *testing.T) {
	url := "https://api.themoviedb.org/3/genre/movie/list"

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetMethod(fasthttp.MethodGet)
	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", TMDB_TOKEN))

	resp := fasthttp.AcquireResponse()
	client := &fasthttp.Client{}
	client.Do(req, resp)

	bodyBytes := resp.Body()
	println(string(bodyBytes))
}
