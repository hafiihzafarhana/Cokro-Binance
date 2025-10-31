package usecase

import (
	"context"

	"github.com/hafiihzafarhana/Cokro-Binance/domain/spot/market"
	"github.com/hafiihzafarhana/Cokro-Binance/domain/spot/market/entity"
)

type MarketUseCaseInterface interface {
	GetBinanceOrderBook(ctx context.Context, data *market.GenericSymbolLimitParams) (*entity.MarketOrderBookEntity, error)
	GetBinanceRecentTradeList(ctx context.Context, data *market.GenericSymbolLimitParams) ([]*entity.MarketRecentTradeListEntity, error)
	GetBinanceOldTradeLookup(ctx context.Context, data *market.GetOldTradeLookupParams) ([]*entity.MarketOldTradeLookupEntity, error)
	GetBinanceAgregateTradeList(ctx context.Context, data *market.GetAgregateTradeListParams) ([]*entity.MarketAgregateTradeListEntity, error)
	GetBinanceCandleStickData(ctx context.Context, data *market.GetCandleStickDataParams) ([]*entity.MarketCandleStickDataEntity, error)
	GetBinanceCurrentAveragePrice(ctx context.Context, symbol string) (*entity.MarketCurrentAveragePriceEntity, error)
	GetBinancePriceChange24hr(ctx context.Context, data *market.GetPriceChange24hrParams) ([]*entity.MarketTickerPrice24hrEntity, error)
}