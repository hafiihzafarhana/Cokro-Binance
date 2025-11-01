package mapper

import (
	"github.com/hafiihzafarhana/Cokro-Binance/domain/spot/general/entity"
	"github.com/hafiihzafarhana/Cokro-Binance/internal/spot/general/model"
	"github.com/hafiihzafarhana/Cokro-Binance/shared/utils/convert"
)

func ToServerTimeEntity(e model.GetServerTimeModel) *entity.GeneralServerTimeEntity {
	return &entity.GeneralServerTimeEntity{
		ServerTime:    e.ServerTime,
		ServerTimeStr: convert.UnixToDateTimeString(e.ServerTime, ""),
	}
}