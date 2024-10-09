package client

import (
	"fmt"
	"sync"
	"time"

	"github.com/valyala/fasthttp"
)

var (
	instance HttpClient
	once     sync.Once
)

type HttpClient interface {
	Get(url string, headers map[string]string, timeout time.Duration) (int, string, error)
	Post(url string, payload any, headers map[string]string, timeout time.Duration) (int, string, error)
}

type httpClient struct {
	Client *fasthttp.Client
}

func (c httpClient) Get(url string, headers map[string]string, timeout time.Duration) (int, string, error) {
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

	return resp.StatusCode(), string(resp.Body()), nil
}

func (c httpClient) Post(url string, payload any, headers map[string]string, timeout time.Duration) (int, string, error) {
	return fasthttp.StatusNotImplemented, "", fmt.Errorf("Yet Implemented")
}

func NewHttpClient() HttpClient {
	once.Do(func() {
		instance = &httpClient{
			Client: &fasthttp.Client{},
		}
	})
	return instance
}
