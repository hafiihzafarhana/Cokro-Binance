package usecase

import (
	"context"

	"github.com/hafiihzafarhana/Cokro-Binance/domain/spot/general/entity"
)

type MarketUsecaseInterface interface {
	GetServerTime(ctx context.Context) (*entity.GeneralServerTimeEntity, error)
}