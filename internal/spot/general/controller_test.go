package general

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/hafiihzafarhana/Cokro-Binance/domain/spot/general/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockGeneralUsecase struct {
	mock.Mock
}

func (m *mockGeneralUsecase) GetServerTime(ctx context.Context) (*entity.GeneralServerTimeEntity, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.GeneralServerTimeEntity), args.Error(1)
}

func TestNewGeneralController(t *testing.T) {
	usecase := new(mockGeneralUsecase)
	ctrl := NewGeneralController(usecase)
	if ctrl == nil {
		t.Fatal("expected non-nil controller")
	}
}

func TestCheckServerTime_Success(t *testing.T) {
	app := fiber.New()
	ctx := context.Background()
	expectedData := &entity.GeneralServerTimeEntity{
		ServerTime: 1499827319559,
		ServerTimeStr: "2006-01-02 15:04:05",
	}

	mockUsecase := new(mockGeneralUsecase)
	mockUsecase.On("GetServerTime", ctx).Return(expectedData, nil)

	ctrl := NewGeneralController(mockUsecase)
	app.Get("/server-time", ctrl.CheckServerTime)

	req := httptest.NewRequest(http.MethodGet, "/server-time", nil)
	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}

func TestCheckServerTime_Error(t *testing.T) {
	app := fiber.New()
	ctx := context.Background()

	mockUsecase := new(mockGeneralUsecase)
	mockUsecase.On("GetServerTime", ctx).Return(nil, errors.New("binance API error"))

	ctrl := NewGeneralController(mockUsecase)
	app.Get("/server-time", ctrl.CheckServerTime)

	req := httptest.NewRequest(http.MethodGet, "/server-time", nil)
	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	mockUsecase.AssertExpectations(t)
}