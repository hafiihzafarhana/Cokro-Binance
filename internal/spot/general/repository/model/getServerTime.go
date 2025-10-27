package model

import (
	"github.com/hafiihzafarhana/Cokro-Binance/domain/spot/general/entity"
	"github.com/hafiihzafarhana/Cokro-Binance/shared/utils/convert"
)

type GetServerTimeModel struct {
	ServerTime int64  `json:"serverTime"`
}

func (m *GetServerTimeModel) ToServerTimeEntity() *entity.GeneralServerTimeEntity {
	return &entity.GeneralServerTimeEntity{
		ServerTime:    m.ServerTime,
		ServerTimeStr: convert.UnixToDateTimeString(m.ServerTime, ""),
	}
}