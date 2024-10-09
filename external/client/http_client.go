package client

import (
	"fmt"
	"sync"
	"time"

	"recommand-chat-bot/domain"

	"github.com/valyala/fasthttp"
)

var (
	instance domain.HttpClient
	once     sync.Once
)

type httpClient struct {
	Client *fasthttp.Client
}

func (c httpClient) Get(url string, headers map[string]string, timeout time.Duration) (int, []byte, error) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.SetTimeout(timeout)
	req.Header.SetMethod(fasthttp.MethodGet)
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp := fasthttp.AcquireResponse()
	if err := c.Client.Do(req, resp); err != nil {
		return fasthttp.StatusInternalServerError, "", err
	}

	return resp.StatusCode(), resp.Body(), nil
}

func (c httpClient) Post(url string, payload any, headers map[string]string, timeout time.Duration) (int, []byte, error) {
	return fasthttp.StatusNotImplemented, []byte{}, fmt.Errorf("Yet Implemented")
}

func NewHttpClient() domain.HttpClient {
	once.Do(func() {
		instance = &httpClient{
			Client: &fasthttp.Client{},
		}
	})
	return instance
}
