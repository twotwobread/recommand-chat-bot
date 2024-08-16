package domain

import (
	"context"
	"encoding/json"
	"time"
)

type Genre string

const (
	Action         Genre = "액션"
	Adventure      Genre = "모험"
	Animation      Genre = "애니메이션"
	Comedy         Genre = "코미디"
	Crime          Genre = "범죄"
	Documentary    Genre = "다큐멘터리"
	Drama          Genre = "드라마"
	Family         Genre = "가족"
	Fantasy        Genre = "판타지"
	History        Genre = "역사"
	Horror         Genre = "공포"
	Music          Genre = "음악"
	Mystery        Genre = "미스테리"
	Romance        Genre = "로맨스"
	ScienceFiction Genre = "SF"
	TVMovie        Genre = "TV영화"
	Thriller       Genre = "스릴러"
	War            Genre = "전쟁"
	Western        Genre = "서부"
)

type Movie struct {
	ID          int64      `json:"id"`
	Title       string     `json:"title"`
	Genre       Genre      `json:"genre"`
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
