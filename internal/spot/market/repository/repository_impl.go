package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/hafiihzafarhana/Cokro-Binance/domain/spot/market"
	"github.com/hafiihzafarhana/Cokro-Binance/domain/spot/market/entity"
	"github.com/hafiihzafarhana/Cokro-Binance/internal/spot/market/mapper"
	"github.com/hafiihzafarhana/Cokro-Binance/internal/spot/market/model"
	binance "github.com/hafiihzafarhana/Cokro-Binance/shared/httpclient/binancespot"
	arraydata "github.com/hafiihzafarhana/Cokro-Binance/shared/utils/arrayData"
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

	var model model.GetOrderBookModel
	if err := json.Unmarshal(resp, &model); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return mapper.ToMarketOrderBookEntity(model), nil
}

func (r *MarketRepositoryImpl) GetRecentTradeList(ctx context.Context, params *market.GenericSymbolLimitParams) ([]*entity.MarketRecentTradeListEntity, error) {
	// Build query parameters
	paramsQ := url.Values{}
	paramsQ.Add("symbol", params.Symbol)
	if params.Limit > 0 {
		paramsQ.Add("limit", fmt.Sprintf("%d", params.Limit))
	}

	// Compose endpoint
	endpoint := fmt.Sprintf("%s?%s", "/trades", paramsQ.Encode())

	// Request to API
	resp, err := r.client.Get(ctx, endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %w", err)
	}

	// Decode JSON to model
	var models []model.GetRecentTradeListModel
	if err := json.Unmarshal(resp, &models); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	var results []*entity.MarketRecentTradeListEntity
	for _, m := range models {
		results = append(results, mapper.ToMarketRecentTradeListEntity(m))
	}

	return results, nil
}

func (r *MarketRepositoryImpl) GetOldTradeLookup(ctx context.Context, params *market.GetOldTradeLookupParams) ([]*entity.MarketOldTradeLookupEntity, error) {
	paramsQ := url.Values{}
	paramsQ.Add("fromId", fmt.Sprintf("%d", params.FromId))
	paramsQ.Add("symbol", params.Symbol)
	if params.Limit > 0 {
		paramsQ.Add("limit", fmt.Sprintf("%d", params.Limit))
	}

	endpoint := fmt.Sprintf("%s?%s", "/historicalTrades", paramsQ.Encode())

	resp, err := r.client.Get(ctx, endpoint)
	if err != nil {
		return nil, err
	}

	var models []model.GetOldTradeLookupModel
	if err := json.Unmarshal(resp, &models); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	var results []*entity.MarketOldTradeLookupEntity
	for _, m := range models {
		results = append(results, mapper.ToMarketOldTradeLookupEntity(m))
	}

	return results, nil
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

	var models []model.GetAgregateTradeListModel
	if err := json.Unmarshal(resp, &models); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	var results []*entity.MarketAgregateTradeListEntity
	for _, m := range models {
		results = append(results, mapper.ToMarketAgregateTradeListEntity(m))
	}

	return results, nil
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

	var rawData [][]interface{}
	if err := json.Unmarshal(resp, &rawData); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	var results []*entity.MarketCandleStickDataEntity
	for _, item := range rawData {
		if len(item) < 11 {
			continue
		}
		results = append(results, mapper.ToMarketCandleStickDataEntity(item, endpointType))
	}

	return results, nil
}

func (r *MarketRepositoryImpl) GetCurrentAveragePrice(ctx context.Context, symbol string) (*entity.MarketCurrentAveragePriceEntity, error) {
	paramsQ := url.Values{}
	paramsQ.Add("symbol", symbol)
	endpoint := fmt.Sprintf("%s?%s", "/avgPrice", paramsQ.Encode())

	resp, err := r.client.Get(ctx, endpoint)
	if err != nil {
		return nil, err
	}

	var models model.GetCurrentAveragePriceModel
	if err := json.Unmarshal(resp, &models); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return mapper.ToMarketCurrentAveragePrice(models), nil
}

func (r *MarketRepositoryImpl) GetTickerPrice24hr(ctx context.Context, params *market.GetPriceChange24hrParams) ([]*entity.MarketTickerPrice24hrEntity, error) {
    splitSymbol := arraydata.SplitAndTrim(params.Symbol, ",")
    if len(splitSymbol) == 0 {
        return nil, fmt.Errorf("no symbol provided")
    }

    paramsQ := url.Values{}
    if len(splitSymbol) > 1 {
        jsonSymbols, _ := json.Marshal(splitSymbol)
        paramsQ.Add("symbols", string(jsonSymbols))
    } else {
        paramsQ.Add("symbol", splitSymbol[0])
    }

    if params.Type != "" {
        paramsQ.Add("type", params.Type)
    }
    if params.SymbolStatus != "" {
        paramsQ.Add("symbolStatus", params.SymbolStatus)
    }

    endpoint := fmt.Sprintf("/ticker/24hr?%s", paramsQ.Encode())

    resp, err := r.client.Get(ctx, endpoint)
    if err != nil {
        return nil, err
    }

    var results []*entity.MarketTickerPrice24hrEntity

	if len(splitSymbol) > 1 {
		var rawList []model.GetTickerPrice24hrModel
		if err := json.Unmarshal(resp, &rawList); err != nil {
			return nil, fmt.Errorf("failed to parse multi response: %w", err)
		}
		for _, raw := range rawList {
			results = append(results, mapper.ToMarketTickerPrice24hrEntity(raw))
		}
	} else {
		var raw model.GetTickerPrice24hrModel
		if err := json.Unmarshal(resp, &raw); err != nil {
			return nil, fmt.Errorf("failed to parse single response: %w", err)
		}
		results = append(results, mapper.ToMarketTickerPrice24hrEntity(raw))
	}

	return results, nil
}

