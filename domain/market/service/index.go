package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/hafiihzafarhana/Cokro-Binance/domain/market"
	"github.com/hafiihzafarhana/Cokro-Binance/domain/market/dto"
	"github.com/hafiihzafarhana/Cokro-Binance/entities"
)

type MarketService struct{}

func NewMarketService() market.MarketServiceInterface {
	return &MarketService{}
}

func (s *MarketService) GetBinanceOrderBook(data *dto.GetBinanceOrderBookReq) (*entities.MarketOrderBookEntity, error) {
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
	var orderBook entities.MarketOrderBookEntity
	if err := json.NewDecoder(resp.Body).Decode(&orderBook); err != nil {
		return nil, fmt.Errorf("failed to decode Binance response: %v", err)
	}
	return &orderBook, nil
}

func (s *MarketService) GetBinanceRecentTradeList(data *dto.GetBinanceRecentTradeListReq) ([]*entities.MarketRecentTradeListEntity, error) {
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
	var trades []*entities.MarketRecentTradeListEntity
	if err := json.NewDecoder(resp.Body).Decode(&trades); err != nil {
		return nil, fmt.Errorf("failed to decode Binance response: %v", err)
	}
	for _, t := range trades {
		t.DateTime = time.UnixMilli(t.Time).Local().Format("2006-01-02 15:04:05")
	}
	return trades, nil
}

func (s *MarketService) GetBinanceOldTradeLookup(data *dto.GetBinanceOldTradeLookupReq) ([]*entities.MarketOldTradeLookupEntity, error) {
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
	var trades []*entities.MarketOldTradeLookupEntity
	if err := json.NewDecoder(resp.Body).Decode(&trades); err != nil {
		return nil, fmt.Errorf("failed to decode Binance response: %v", err)
	}
	for _, t := range trades {
		t.DateTime = time.UnixMilli(t.Time).Local().Format("2006-01-02 15:04:05")
	}
	return trades, nil
}

func (s *MarketService) GetAgregateTradeList(data *dto.GetBinanceAgregateTradeListReq) ([]*entities.MarketAgregateTradeListEntity, error) {
	baseURL := "https://api.binance.com/api/v3/aggTrades"
	params := url.Values{}
	params.Add("symbol", data.Symbol)
	if data.Limit > 0 {
		params.Add("limit", fmt.Sprintf("%d", data.Limit))
	}
	if data.FromId > 0 {
		params.Add("fromId", fmt.Sprintf("%d", data.FromId))
	}
	if !data.StartTime.IsZero() {
		params.Add("startTime", fmt.Sprintf("%d", data.StartTime.UnixMilli()))
	}
	if !data.EndTime.IsZero() {
		params.Add("endTime", fmt.Sprintf("%d", data.EndTime.UnixMilli()))
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
	var trades []*entities.MarketAgregateTradeListEntity
	if err := json.NewDecoder(resp.Body).Decode(&trades); err != nil {
		return nil, fmt.Errorf("failed to decode Binance response: %v", err)
	}
	for _, t := range trades {
		t.DateTime = time.UnixMilli(t.Timestamp).Local().Format("2006-01-02 15:04:05")
	}
	return trades, nil
}

func (s *MarketService) GetCandleStickData(data *dto.GetBinanceCandleStickDataReq) ([]*entities.MarketCandleStickDataEntity, error) {
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
	if !data.StartTime.IsZero() {
		params.Add("startTime", fmt.Sprintf("%d", data.StartTime.UnixMilli()))
	}
	if !data.EndTime.IsZero() {
		params.Add("endTime", fmt.Sprintf("%d", data.EndTime.UnixMilli()))
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
	candles := make([]*entities.MarketCandleStickDataEntity, 0, len(rawData))
	for _, item := range rawData {
		if len(item) < 11 {
			continue
		}
		openTime, _ := item[0].(float64)
		openPrice, _ := item[1].(string)
		highPrice, _ := item[2].(string)
		lowPrice, _ := item[3].(string)
		closePrice, _ := item[4].(string)
		volume, _ := item[5].(string)
		closeTime, _ := item[6].(float64)
		quoteAssetVolume, _ := item[7].(string)
		numberOfTrades, _ := item[8].(float64)
		takerBuyBaseAssetVolume, _ := item[9].(string)
		takerBuyQuoteAssetVolume, _ := item[10].(string)

		candle := &entities.MarketCandleStickDataEntity{
			OpenTime:               int64(openTime),
			OpenTimeStr:            time.UnixMilli(int64(openTime)).Local().Format("2006-01-02 15:04:05"),
			OpenPrice:              openPrice,
			HighPrice:              highPrice,
			LowPrice:               lowPrice,
			ClosePrice:             closePrice,
			Volume:                 volume,
			CloseTime:              int64(closeTime),
			CloseTimeStr:           time.UnixMilli(int64(closeTime)).Local().Format("2006-01-02 15:04:05"),
			QuoteAssetVolume:       quoteAssetVolume,
			NumberOfTrades:         int(numberOfTrades),
			TakerBuyBaseAssetVolume: takerBuyBaseAssetVolume,
			TakerBuyQuoteAssetVolume: takerBuyQuoteAssetVolume,
			TypeKline: data.TypeKline,
		}
		candles = append(candles, candle)
	}
	return candles, nil
}