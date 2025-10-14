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