package general

import (
	"context"

	"github.com/hafiihzafarhana/Cokro-Binance/domain/spot/general/entity"
)

type GeneralService interface {
	GetServerTime(ctx context.Context) (*entity.GeneralServerTimeEntity, error)
}