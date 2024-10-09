package domain

import "time"

type HttpClient interface {
	Get(url string, headers map[string]string, timeout time.Duration) (int, string, error)
	Post(url string, payload any, headers map[string]string, timeout time.Duration) (int, string, error)
}
