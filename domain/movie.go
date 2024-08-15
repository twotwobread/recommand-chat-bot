package domain

import (
	"context"
	"encoding/json"
	"time"
)

type Movie struct {
	ID          int64      `json:"id"`
	Title       string     `json:"title"`
	Director    string     `json:"director"`
	Actors      []string   `json:"actors"`
	Description string     `json:"description"`
	ReleaseDate CustomTime `json:"release_date"`
	UpdatedAt   time.Time  `json:"updated_at"`
	CreatedAt   time.Time  `json:"created_at"`
}

type CustomTime struct {
	time.Time
}

func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(ct.Time.Format("2006-01-02"))
}

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	s := string(b)
	s = s[1 : len(s)-1]
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	ct.Time = t
	return nil
}

type MovieUsecase interface {
	GetByID(ctx context.Context, id int64) (Movie, error)
	Store(ctx context.Context, m *Movie) error
	GetAll(ctx context.Context) ([]Movie, error)
}

type MovieRepository interface{}
