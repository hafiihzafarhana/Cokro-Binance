package general

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

type mockGeneralController struct{}

func (m *mockGeneralController) CheckServerTime(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).SendString("server-time ok")
}

func TestRegisterMarketRoutes(t *testing.T) {
	app := fiber.New()
	handler := &mockGeneralController{}

	RegisterGeneralRoutes(app, handler)

	tests := []struct {
		route string
		want  string
	}{
		{"/general/server-time", "server-time ok"},
	}

	for _, tt := range tests {
		req := httptest.NewRequest(http.MethodGet, tt.route, nil)
		resp, _ := app.Test(req, -1)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		body := make([]byte, resp.ContentLength)
		resp.Body.Read(body)
		assert.Contains(t, string(body), tt.want)
		resp.Body.Close()
	}
}