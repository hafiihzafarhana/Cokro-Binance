package response

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func setupTestApp() *fiber.App {
	app := fiber.New()
	return app
}

func TestSendStatusOkResponse(t *testing.T) {
	app := setupTestApp()
	app.Get("/", func(c *fiber.Ctx) error {
		return SendStatusOkResponse(c, "success")
	})

	req := httptest.NewRequest("GET", "/", nil)
	resp, _ := app.Test(req, -1)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var body GeneralMessage
	json.NewDecoder(resp.Body).Decode(&body)
	assert.Equal(t, "success", body.Message)
}

func TestStatusCreatedResponse(t *testing.T) {
	app := setupTestApp()
	app.Post("/", func(c *fiber.Ctx) error {
		return SendStatusCreatedResponse(c, "created")
	})

	req := httptest.NewRequest("POST", "/", nil)
	resp, _ := app.Test(req, -1)

	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)

	var body GeneralMessage
	json.NewDecoder(resp.Body).Decode(&body)
	assert.Equal(t, "created", body.Message)
}

func TestSendStatusCreatedWithDataResponse(t *testing.T) {
	app := setupTestApp()
	app.Get("/", func(c *fiber.Ctx) error {
		data := map[string]string{"id": "123"}
		return SendStatusCreatedWithDataResponse(c, "created with data", data)
	})

	req := httptest.NewRequest("GET", "/", nil)
	resp, _ := app.Test(req, -1)

	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)

	var body GeneralMessageWithData
	json.NewDecoder(resp.Body).Decode(&body)
	assert.Equal(t, "created with data", body.Message)
	assert.Equal(t, map[string]interface{}{"id": "123"}, body.Data)
}

func TestSendStatusOkWithDataResponse(t *testing.T) {
	app := setupTestApp()
	app.Get("/", func(c *fiber.Ctx) error {
		data := map[string]string{"name": "btc"}
		return SendStatusOkWithDataResponse(c, "ok", data)
	})

	req := httptest.NewRequest("GET", "/", nil)
	resp, _ := app.Test(req, -1)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var body GeneralMessageWithData
	json.NewDecoder(resp.Body).Decode(&body)
	assert.Equal(t, "ok", body.Message)
	assert.Equal(t, map[string]interface{}{"name": "btc"}, body.Data)
}

func TestSendStatusBadRequest(t *testing.T) {
	app := setupTestApp()
	app.Get("/", func(c *fiber.Ctx) error {
		return SendStatusBadRequest(c, "bad request")
	})

	req := httptest.NewRequest("GET", "/", nil)
	resp, _ := app.Test(req, -1)

	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

	var body GeneralMessage
	json.NewDecoder(resp.Body).Decode(&body)
	assert.Equal(t, "bad request", body.Message)
}

func TestSendStatusNotFound(t *testing.T) {
	app := setupTestApp()
	app.Get("/", func(c *fiber.Ctx) error {
		return SendStatusNotFound(c, "not found")
	})

	req := httptest.NewRequest("GET", "/", nil)
	resp, _ := app.Test(req, -1)

	assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)

	var body GeneralMessage
	json.NewDecoder(resp.Body).Decode(&body)
	assert.Equal(t, "not found", body.Message)
}

func TestSendStatusInternalServerError(t *testing.T) {
	app := setupTestApp()
	app.Get("/", func(c *fiber.Ctx) error {
		return SendStatusInternalServerError(c, "server error")
	})

	req := httptest.NewRequest("GET", "/", nil)
	resp, _ := app.Test(req, -1)

	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

	var body GeneralMessage
	json.NewDecoder(resp.Body).Decode(&body)
	assert.Equal(t, "server error", body.Message)
}

func TestSendStatusUnauthorized(t *testing.T) {
	app := setupTestApp()
	app.Get("/", func(c *fiber.Ctx) error {
		return SendStatusUnauthorized(c, "unauthorized")
	})

	req := httptest.NewRequest("GET", "/", nil)
	resp, _ := app.Test(req, -1)

	assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)

	var body GeneralMessage
	json.NewDecoder(resp.Body).Decode(&body)
	assert.Equal(t, "unauthorized", body.Message)
}

func TestSendStatusForbidden(t *testing.T) {
	app := setupTestApp()
	app.Get("/", func(c *fiber.Ctx) error {
		return SendStatusForbidden(c, "forbidden")
	})

	req := httptest.NewRequest("GET", "/", nil)
	resp, _ := app.Test(req, -1)

	assert.Equal(t, fiber.StatusForbidden, resp.StatusCode)

	var body GeneralMessage
	json.NewDecoder(resp.Body).Decode(&body)
	assert.Equal(t, "forbidden", body.Message)
}