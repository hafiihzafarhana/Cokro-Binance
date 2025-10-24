package usecase

import (
	"context"

	"github.com/hafiihzafarhana/Cokro-Binance/internal/spot/general/dto"
)

type MarketUsecaseInterface interface {
	GetServerTime(ctx context.Context) (*dto.GeneralServerTimeRes, error)
}