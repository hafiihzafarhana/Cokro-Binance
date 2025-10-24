package usecase

import (
	"context"

	"github.com/hafiihzafarhana/Cokro-Binance/domain/spot/market/entity"
	"github.com/hafiihzafarhana/Cokro-Binance/internal/spot/market/dto"
)

type MarketUseCaseInterface interface {
	GetBinanceOrderBook(ctx context.Context, data *dto.GenericSymbolLimitReq) (*entity.MarketOrderBookEntity, error)
	GetBinanceRecentTradeList(ctx context.Context, data *dto.GenericSymbolLimitReq) ([]*entity.MarketRecentTradeListEntity, error)
	GetBinanceOldTradeLookup(ctx context.Context, data *dto.GetOldTradeLookupReq) ([]*entity.MarketOldTradeLookupEntity, error)
	GetBinanceAgregateTradeList(ctx context.Context, data *dto.GetAgregateTradeListReq) ([]*entity.MarketAgregateTradeListEntity, error)
	GetBinanceCandleStickData(ctx context.Context, data *dto.GetCandleStickDataReq) ([]*entity.MarketCandleStickDataEntity, error)
	GetBinanceCurrentAveragePrice(ctx context.Context, symbol string) (*entity.MarketCurrentAveragePriceEntity, error)
}