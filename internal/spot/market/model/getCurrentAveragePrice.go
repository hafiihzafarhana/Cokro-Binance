package model

type GetCurrentAveragePriceModel struct {
	IntervalInMinute  int `json:"mins"`
	Price 		  string `json:"price"`
	CloseTime 	  int64 `json:"closeTime"`
}