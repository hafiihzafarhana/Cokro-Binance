package market

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hafiihzafarhana/Cokro-Binance/domain/market/dto"
	"github.com/hafiihzafarhana/Cokro-Binance/entities"
)

type MarketServiceInterface interface {
	GetBinanceOrderBook(data *dto.GetBinanceOrderBookReq) (*entities.MarketOrderBookEntity, error)
	GetBinanceRecentTradeList(data *dto.GetBinanceRecentTradeListReq) ([]*entities.MarketRecentTradeListEntity, error)
}

type MarketControllerInterface interface {
	CheckOrderBook(c *fiber.Ctx) error
	CheckRecentTradeList(c *fiber.Ctx) error
}