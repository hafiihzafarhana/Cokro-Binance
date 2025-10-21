package usecase

import (
	"github.com/hafiihzafarhana/Cokro-Binance/domain/spot/general/entity"
	"github.com/hafiihzafarhana/Cokro-Binance/internal/spot/general/dto"
)

func ToGeneralServerTimeRes(e *entity.GeneralServerTimeEntity) *dto.GeneralServerTimeRes {
	return &dto.GeneralServerTimeRes{
		ServerTime: e.ServerTime,
	}
}
