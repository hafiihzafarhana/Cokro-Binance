package market

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hafiihzafarhana/Cokro-Binance/domain/market/dto"
	"github.com/hafiihzafarhana/Cokro-Binance/entities"
)

type MarketServiceInterface interface {
	GetBinanceOrderBook(data *dto.GetBinanceOrderBookReq) (*entities.MarketOrderBookEntity, error)
	GetBinanceRecentTradeList(data *dto.GetBinanceRecentTradeListReq) ([]*entities.MarketRecentTradeListEntity, error)
	GetBinanceOldTradeLookup(data *dto.GetBinanceOldTradeLookupReq) ([]*entities.MarketOldTradeLookupEntity, error)
	GetAgregateTradeList(data *dto.GetBinanceAgregateTradeListReq) ([]*entities.MarketAgregateTradeListEntity, error)
	GetCandleStickData(data *dto.GetBinanceCandleStickDataReq) ([]*entities.MarketCandleStickDataEntity, error)
}

type MarketControllerInterface interface {
	CheckOrderBook(c *fiber.Ctx) error
	CheckRecentTradeList(c *fiber.Ctx) error
	CheckOldTradeLookup(c *fiber.Ctx) error
	CheckAgregateTradeList(c *fiber.Ctx) error
	CheckCandleStickData(c *fiber.Ctx) error
}