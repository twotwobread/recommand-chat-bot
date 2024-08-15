package rest

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

func TestCheckHealth(t *testing.T) {
	type args struct {
		c fiber.Ctx
	}
	tests := []struct {
		name         string
		wantErr      bool
		expectedCode int
		expectedBody string
	}{
		{
			name:         "Success health check",
			wantErr:      false,
			expectedCode: fiber.StatusOK,
			expectedBody: "Check Health",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			app.Get("/health", CheckHealth)

			req := httptest.NewRequest("GET", "/health", nil)
			resp, err := app.Test(req)
			body := make([]byte, 1024)
			n, _ := resp.Body.Read(body)
			bodyStr := string(body[:n])

			assert.Equal(t, tt.wantErr, (err != nil))
			assert.Equal(t, tt.expectedCode, resp.StatusCode)
			assert.Equal(t, tt.expectedBody, bodyStr)
		})
	}
}
