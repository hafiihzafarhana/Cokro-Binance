package entity

import "github.com/hafiihzafarhana/Cokro-Binance/helper/convert"

type MarketCurrentAveragePriceEntity struct {
	IntervalInMinute  int  `json:"intervalInMinute"`
	Price 		  string `json:"price"`
	CloseTime 	  int64  `json:"closeTime"`
	CloseTimeStr string `json:"closeDateTimeStr"`
}

func MapMarketCurrentAveragePrice(rawData interface{}) *MarketCurrentAveragePriceEntity {
	dataMap, ok := rawData.(map[string]interface{})
	if !ok {
		return nil
	}

	interval := int(dataMap["mins"].(float64))
	price := dataMap["price"].(string)
	closeTime := int64(dataMap["closeTime"].(float64))

	return &MarketCurrentAveragePriceEntity{
		IntervalInMinute: interval,
		Price:            price,
		CloseTime:        closeTime,
		CloseTimeStr:     convert.UnixToDateTimeString(closeTime, ""),
	}
}