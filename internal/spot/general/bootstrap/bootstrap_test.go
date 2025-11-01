package bootstrap

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestSetupSpotGeneralModule(t *testing.T) {
	app := fiber.New()

	assert.NotPanics(t, func() {
		SetupSpotGeneralModule(app)
	})

	// call one endpoint
	req := httptest.NewRequest(http.MethodGet, "/general/server-time", nil)
	resp, err := app.Test(req, -1)

	assert.NoError(t, err)
	assert.NotEqual(t, http.StatusNotFound, resp.StatusCode)
}