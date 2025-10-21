package general

import (
	"context"

	"github.com/hafiihzafarhana/Cokro-Binance/domain/spot/general/entity"
)

type GeneralRepository interface {
    GetServerTime(ctx context.Context) (*entity.GeneralServerTimeEntity, error)
}