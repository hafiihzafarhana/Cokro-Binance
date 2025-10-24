package market

import (
	"context"

	"github.com/hafiihzafarhana/Cokro-Binance/domain/spot/market/entity"
)

type MarketRepository interface {
	GetOrderBook(ctx context.Context, params *GenericSymbolLimitParams) (*entity.MarketOrderBookEntity, error)
	GetRecentTradeList(ctx context.Context, params *GenericSymbolLimitParams) ([]*entity.MarketRecentTradeListEntity, error)
	GetOldTradeLookup(ctx context.Context, params *GetOldTradeLookupParams) ([]*entity.MarketOldTradeLookupEntity, error)
	GetAgregateTradeList(ctx context.Context, params *GetAgregateTradeListParams) ([]*entity.MarketAgregateTradeListEntity, error)
	GetCandleStickData(ctx context.Context, params *GetCandleStickDataParams) ([]*entity.MarketCandleStickDataEntity, error)
	GetCurrentAveragePrice(ctx context.Context, symbol string) (*entity.MarketCurrentAveragePriceEntity, error)
}