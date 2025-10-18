package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/hafiihzafarhana/Cokro-Binance/domain/market"
	"github.com/hafiihzafarhana/Cokro-Binance/domain/market/dto"
	"github.com/hafiihzafarhana/Cokro-Binance/domain/market/entity"
	"github.com/hafiihzafarhana/Cokro-Binance/helper/convert"
)

type MarketService struct{}

func NewMarketService() market.MarketServiceInterface {
	return &MarketService{}
}

func (s *MarketService) GetBinanceOrderBook(data *dto.GetBinanceOrderBookReq) (*entity.MarketOrderBookEntity, error) {
	baseURL := "https://api.binance.com/api/v3/depth"
	params := url.Values{}
	params.Add("symbol", data.Symbol)
	if data.Limit > 0 {
		params.Add("limit", fmt.Sprintf("%d", data.Limit))
	}
	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
	resp, err := http.Get(fullURL)
	if err != nil {
		return nil, fmt.Errorf("failed to request Binance API: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Binance API returned status: %d", resp.StatusCode)
	}
	var orderBook entity.MarketOrderBookEntity
	if err := json.NewDecoder(resp.Body).Decode(&orderBook); err != nil {
		return nil, fmt.Errorf("failed to decode Binance response: %v", err)
	}
	return &orderBook, nil
}

func (s *MarketService) GetBinanceRecentTradeList(data *dto.GetBinanceRecentTradeListReq) ([]*entity.MarketRecentTradeListEntity, error) {
	baseURL := "https://api.binance.com/api/v3/trades"
	params := url.Values{}
	params.Add("symbol", data.Symbol)
	if data.Limit > 0 {
		params.Add("limit", fmt.Sprintf("%d", data.Limit))
	}
	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
	resp, err := http.Get(fullURL)
	if err != nil {
		return nil, fmt.Errorf("failed to request Binance API: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Binance API returned status: %d", resp.StatusCode)
	}
	var trades []*entity.MarketRecentTradeListEntity
	if err := json.NewDecoder(resp.Body).Decode(&trades); err != nil {
		return nil, fmt.Errorf("failed to decode Binance response: %v", err)
	}
	for _, t := range trades {
		t.DateTime = convert.UnixToDateTimeString(t.Time, "")
	}
	return trades, nil
}

func (s *MarketService) GetBinanceOldTradeLookup(data *dto.GetBinanceOldTradeLookupReq) ([]*entity.MarketOldTradeLookupEntity, error) {
	baseURL := "https://api.binance.com/api/v3/historicalTrades"
	params := url.Values{}
	params.Add("symbol", data.Symbol)
	if data.Limit > 0 {
		params.Add("limit", fmt.Sprintf("%d", data.Limit))
	}
	params.Add("fromId", fmt.Sprintf("%d", data.FromId))
	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
	resp, err := http.Get(fullURL)
	if err != nil {
		return nil, fmt.Errorf("failed to request Binance API: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Binance API returned status: %d", resp.StatusCode)
	}
	var trades []*entity.MarketOldTradeLookupEntity
	if err := json.NewDecoder(resp.Body).Decode(&trades); err != nil {
		return nil, fmt.Errorf("failed to decode Binance response: %v", err)
	}
	for _, t := range trades {
		t.DateTime = convert.UnixToDateTimeString(t.Time, "UTC")
	}
	return trades, nil
}

func (s *MarketService) GetAgregateTradeList(data *dto.GetBinanceAgregateTradeListReq) ([]*entity.MarketAgregateTradeListEntity, error) {
	baseURL := "https://api.binance.com/api/v3/aggTrades"
	params := url.Values{}
	params.Add("symbol", data.Symbol)
	if data.Limit > 0 {
		params.Add("limit", fmt.Sprintf("%d", data.Limit))
	}
	if data.FromId > 0 {
		params.Add("fromId", fmt.Sprintf("%d", data.FromId))
	}
	loc, _ := time.LoadLocation("Asia/Jakarta")
	if data.StartTime != "" {
		localStart, err := time.ParseInLocation("2006-01-02T15:04:05", data.StartTime, loc)
		if err == nil {
			params.Add("startTime", fmt.Sprintf("%d", localStart.UTC().UnixMilli()))
		}
	}
	if data.EndTime != "" {
		localEnd, err := time.ParseInLocation("2006-01-02T15:04:05", data.EndTime, loc)
		if err == nil {
			params.Add("endTime", fmt.Sprintf("%d", localEnd.UTC().UnixMilli()))
		}
	}
	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
	resp, err := http.Get(fullURL)
	if err != nil {
		return nil, fmt.Errorf("failed to request Binance API: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Binance API returned status: %d", resp.StatusCode)
	}
	var trades []*entity.MarketAgregateTradeListEntity
	if err := json.NewDecoder(resp.Body).Decode(&trades); err != nil {
		return nil, fmt.Errorf("failed to decode Binance response: %v", err)
	}
	for _, t := range trades {
		t.DateTime = convert.UnixToDateTimeString(t.Timestamp, "")
	}
	return trades, nil
}

func (s *MarketService) GetCandleStickData(data *dto.GetBinanceCandleStickDataReq) ([]*entity.MarketCandleStickDataEntity, error) {
	baseURL := "https://api.binance.com/api/v3/klines"
	if data.TypeKline == "uiklines" {
		baseURL = "https://api.binance.com/api/v3/uiKlines"
	}
	params := url.Values{}
	params.Add("symbol", data.Symbol)
	params.Add("interval", data.Interval)
	if data.Limit > 0 {
		params.Add("limit", fmt.Sprintf("%d", data.Limit))
	}
	loc, _ := time.LoadLocation("Asia/Jakarta")
	if data.StartTime != "" {
		localStart, err := time.ParseInLocation("2006-01-02T15:04:05", data.StartTime, loc)
		if err == nil {
			params.Add("startTime", fmt.Sprintf("%d", localStart.UTC().UnixMilli()))
		}
	}
	if data.EndTime != "" {
		localEnd, err := time.ParseInLocation("2006-01-02T15:04:05", data.EndTime, loc)
		if err == nil {
			params.Add("endTime", fmt.Sprintf("%d", localEnd.UTC().UnixMilli()))
		}
	}
	if data.TimeZone != "" {
		params.Add("timeZone", data.TimeZone)
	}
	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
	resp, err := http.Get(fullURL)
	if err != nil {
		return nil, fmt.Errorf("failed to request Binance API: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Binance API returned status: %d", resp.StatusCode)
	}
	var rawData [][]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&rawData); err != nil {
		return nil, fmt.Errorf("failed to decode Binance response: %v", err)
	}
	candles := entity.MapMarketCandleStickData(rawData, data.TypeKline)
	return candles, nil
}

func (s *MarketService) GetCurrentAveragePrice(symbol string) (*entity.MarketCurrentAveragePriceEntity, error) {
	baseURL := "https://api.binance.com/api/v3/avgPrice"
	params := url.Values{}
	params.Add("symbol", symbol)
	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
	resp, err := http.Get(fullURL)
	if err != nil {
		return nil, fmt.Errorf("failed to request Binance API: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Binance API returned status: %d", resp.StatusCode)
	}
	var rawData map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&rawData); err != nil {
		return nil, fmt.Errorf("failed to decode Binance response: %v", err)
	}
	currentAvgPrice := entity.MapMarketCurrentAveragePrice(rawData)
	return currentAvgPrice, nil
}