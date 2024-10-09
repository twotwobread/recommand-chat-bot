package domain

import "time"

type HttpClient interface {
	Get(url string, headers map[string]string, timeout time.Duration) (int, []byte, error)
	Post(url string, payload any, headers map[string]string, timeout time.Duration) (int, []byte, error)
}

type BatchUsecase interface {
	Process()
}
