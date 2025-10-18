package market

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hafiihzafarhana/Cokro-Binance/domain/market/dto"
	"github.com/hafiihzafarhana/Cokro-Binance/domain/market/entity"
)

type MarketServiceInterface interface {
	GetBinanceOrderBook(data *dto.GetBinanceOrderBookReq) (*entity.MarketOrderBookEntity, error)
	GetBinanceRecentTradeList(data *dto.GetBinanceRecentTradeListReq) ([]*entity.MarketRecentTradeListEntity, error)
	GetBinanceOldTradeLookup(data *dto.GetBinanceOldTradeLookupReq) ([]*entity.MarketOldTradeLookupEntity, error)
	GetAgregateTradeList(data *dto.GetBinanceAgregateTradeListReq) ([]*entity.MarketAgregateTradeListEntity, error)
	GetCandleStickData(data *dto.GetBinanceCandleStickDataReq) ([]*entity.MarketCandleStickDataEntity, error)
}

type MarketControllerInterface interface {
	CheckOrderBook(c *fiber.Ctx) error
	CheckRecentTradeList(c *fiber.Ctx) error
	CheckOldTradeLookup(c *fiber.Ctx) error
	CheckAgregateTradeList(c *fiber.Ctx) error
	CheckCandleStickData(c *fiber.Ctx) error
}