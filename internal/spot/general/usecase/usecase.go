package usecase

import (
	"context"

	"github.com/hafiihzafarhana/Cokro-Binance/internal/spot/general/dto"
)

type MarketUseCaseInterface interface {
	GetServerTime(ctx context.Context) (*dto.GeneralServerTimeRes, error)
}