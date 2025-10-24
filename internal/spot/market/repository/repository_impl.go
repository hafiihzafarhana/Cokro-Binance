package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/hafiihzafarhana/Cokro-Binance/domain/spot/market"
	"github.com/hafiihzafarhana/Cokro-Binance/domain/spot/market/entity"
	binance "github.com/hafiihzafarhana/Cokro-Binance/shared/httpclient/binancespot"
)

type MarketRepositoryImpl struct {
	client *binance.HttpClient
}

func NewMarketRepository(client *binance.HttpClient) market.MarketRepository {
	return &MarketRepositoryImpl{client: client}
}

func (r *MarketRepositoryImpl) GetOrderBook(ctx context.Context, params *market.GenericSymbolLimitParams) (*entity.MarketOrderBookEntity, error) {
	paramsQ := url.Values{}
    paramsQ.Add("symbol", params.Symbol)
    if params.Limit > 0 {
        paramsQ.Add("limit", fmt.Sprintf("%d", params.Limit))
    }

	endpoint := fmt.Sprintf("%s?%s", "/depth", paramsQ.Encode())

	resp, err := r.client.Get(ctx, endpoint)
	if err != nil {
		return nil, err
	}

	var result entity.MarketOrderBookEntity
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}

func (r *MarketRepositoryImpl) GetRecentTradeList(ctx context.Context, params *market.GenericSymbolLimitParams) ([]*entity.MarketRecentTradeListEntity, error) {
	paramsQ := url.Values{}
	paramsQ.Add("symbol", params.Symbol)
	if params.Limit > 0 {
		paramsQ.Add("limit", fmt.Sprintf("%d", params.Limit))
	}

	endpoint := fmt.Sprintf("%s?%s", "/trades", paramsQ.Encode())

	resp, err := r.client.Get(ctx, endpoint)
	if err != nil {
		return nil, err
	}

	var result []*entity.MarketRecentTradeListEntity
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}
	
	return result, nil
}

func (r *MarketRepositoryImpl) GetOldTradeLookup(ctx context.Context, params *market.GetOldTradeLookupParams) ([]*entity.MarketOldTradeLookupEntity, error) {
	paramsQ := url.Values{}
	paramsQ.Add("fromId", fmt.Sprintf("%d", params.FromId))
	paramsQ.Add("symbol", params.Symbol)
	if params.Limit > 0 {
		paramsQ.Add("limit", fmt.Sprintf("%d", params.Limit))
	}

	endpoint := fmt.Sprintf("%s?%s", "/oldTradeLookup", paramsQ.Encode())

	resp, err := r.client.Get(ctx, endpoint)
	if err != nil {
		return nil, err
	}

	var result []*entity.MarketOldTradeLookupEntity
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return result, nil
}

func (r *MarketRepositoryImpl) GetAgregateTradeList(ctx context.Context, params *market.GetAgregateTradeListParams) ([]*entity.MarketAgregateTradeListEntity, error) {
	paramsQ := url.Values{}
	paramsQ.Add("symbol", params.Symbol)
	if params.Limit > 0 {
		paramsQ.Add("limit", fmt.Sprintf("%d", params.Limit))
	}
	if params.FromId > 0 {
		paramsQ.Add("fromId", fmt.Sprintf("%d", params.FromId))
	}
	if params.StartTime != "" {
		paramsQ.Add("startTime", params.StartTime)
	}
	if params.EndTime != "" {
		paramsQ.Add("endTime", params.EndTime)
	}

	endpoint := fmt.Sprintf("%s?%s", "/aggTrades", paramsQ.Encode())

	resp, err := r.client.Get(ctx, endpoint)
	if err != nil {
		return nil, err
	}

	var result []*entity.MarketAgregateTradeListEntity
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return result, nil
}

func (r *MarketRepositoryImpl) GetCandleStickData(ctx context.Context, params *market.GetCandleStickDataParams) ([]*entity.MarketCandleStickDataEntity, error) {
	paramsQ := url.Values{}
	paramsQ.Add("symbol", params.Symbol)
	paramsQ.Add("interval", params.Interval)
	if params.StartTime != "" {
		paramsQ.Add("startTime", params.StartTime)
	}
	if params.EndTime != "" {
		paramsQ.Add("endTime", params.EndTime)
	}
	if params.Limit > 0 {
		paramsQ.Add("limit", fmt.Sprintf("%d", params.Limit))
	}
	if params.TimeZone != "" {
		paramsQ.Add("timeZone", params.TimeZone)
	}
	endpointType := params.TypeKline
	if endpointType == "uiklines" {
		endpointType = "uiKlines"
	}

	endpoint := fmt.Sprintf("/%s?%s", endpointType, paramsQ.Encode())

	resp, err := r.client.Get(ctx, endpoint)
	if err != nil {
		return nil, err
	}

	var result []*entity.MarketCandleStickDataEntity
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return result, nil
}

func (r *MarketRepositoryImpl) GetCurrentAveragePrice(ctx context.Context, symbol string) (*entity.MarketCurrentAveragePriceEntity, error) {
	return nil, nil
}

