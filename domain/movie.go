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

func (g Genre) IsValid() bool {
	switch g {
	case Action, Adventure, Animation, Comedy, Crime, Documentary, Drama, Family, Fantasy,
		History, Horror, Music, Mystery, Romance, ScienceFiction, TVMovie, Thriller, War, Western:
		return true
	}
	return false
}

type Movie struct {
	Title       string     `json:"title" validate:"required,validateEmptyStr,maxlen=10"`
	Genre       Genre      `json:"genre" validate:"required,validateGenre"`
	Director    string     `json:"director" validate:"required,validateEmptyStr,maxlen=10"`
	Actors      []string   `json:"actors" validate:"required"`
	Description string     `json:"description" validate:"required,validateEmptyStr,maxlen=10"`
	ReleaseDate CustomTime `json:"release_date"`
}

type MovieDetailOutput struct {
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
	Store(ctx context.Context, m *Movie) (int64, error)
	GetRandom(ctx context.Context) (MovieDetailOutput, error)
}

type MovieRepository interface {
	Store(ctx context.Context, m *Movie) (int64, error)
	GetByID(ctx context.Context, id int64) (MovieDetailOutput, error)
	GetAll(ctx context.Context) ([]MovieDetailOutput, error)
	GetCount(ctx context.Context) (int, error)
}

/*
	context.Context가 컨트롤러에서 생성되어 모든 레이어로 전파되는 이유.
	1. 먼저 go에서 context.Context의 역할을 알아보자.
		- 작업 취소 (Cancellation): 장기 실행 작업을 안전하게 중단할 수 있게 해줍니다.
			예) HTTP 요청 처리 중 클라이언트 연결이 끊어졌을 때 관련 작업을 모두 취소할 수 있습니다.

		- 데드라인 및 타임아웃 (Deadline and Timeout): 작업에 시간 제한을 설정할 수 있습니다.
			예) API 요청이나 데이터베이스 쿼리에 타임아웃을 적용할 수 있습니다.

		- 요청 범위 값 전달 (Request-scoped values): 요청별로 고유한 값(예: 사용자 ID, 요청 ID)을 안전하게 전달할 수 있습니다.
			예) 로깅이나 추적에 필요한 정보를 컨텍스트를 통해 전파할 수 있습니다.

		- 고루틴 간 동기화: 여러 고루틴에 걸친 작업을 조정하는 데 사용됩니다.
			예) 한 작업이 완료되면 관련된 모든 고루틴을 종료할 수 있습니다.

		- API 경계 통일: 다양한 패키지와 API에서 일관된 인터페이스를 제공합니다.
			예) 데이터베이스, HTTP 클라이언트, 서버 등 다양한 라이브러리에서 동일한 컨텍스트 인터페이스를 사용합니다.

		- 주요 사용 사례:
			- 웹 서버
				- 각 HTTP 요청에 대해 컨텍스트를 생성하여 요청의 생명주기를 관리합니다.
				- 클라이언트 연결 종료 시 관련 작업을 모두 취소할 수 있습니다.

			- 데이터베이스 작업
				- 쿼리 실행 시 타임아웃을 설정하거나, 필요시 작업을 취소할 수 있습니다.

			- 외부 API 호출
				- API 요청에 타임아웃을 설정하고, 필요 시 요청을 취소할 수 있습니다.

			- 로깅 및 모니터링
				- 요청 ID나 사용자 정보를 컨텍스트에 포함시켜 로그에 일관되게 기록할 수 있습니다.

			- 마이크로서비스
				- 서비스 간 요청 추적이나 인증 정보 전달에 사용됩니다.

		- context.Context를 사용할 때 주의할 점
			1. 컨텍스트는 요청의 수명주기와 관련된 정보만 포함해야 합니다.
			2. 함수의 첫 번째 매개변수로 전달하는 것이 관례입니다.
			3. 컨텍스트를 키로 사용하여 값을 저장하지 말아야 합니다.
			4. 항상 context.Background() 또는 context.TODO()로 시작하여 컨텍스트 체인을 만듭니다.

		- 간단히 말하자면 context.Context는 Go 프로그램의 동시성 제어와 요청 범위 데이터 관리에 중요한 역할을 수행합니다.
	2. 이 애플리케이션에서 context.Context를 전파하는 이유는 ent ORM을 이용 시 context.Context를 이용해서 데이터를 저장하고
		또 추후에 요청마다의 필요한 정보가 있을 때 각 레이어에서 이를 이용할 수 있습니다.
*/
