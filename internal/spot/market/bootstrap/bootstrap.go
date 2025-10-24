package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hafiihzafarhana/Cokro-Binance/internal/spot/market"
	"github.com/hafiihzafarhana/Cokro-Binance/internal/spot/market/repository"
	"github.com/hafiihzafarhana/Cokro-Binance/internal/spot/market/usecase"
	"github.com/hafiihzafarhana/Cokro-Binance/shared/httpclient/binancespot"
)

func SetupSpotMarketModule(app *fiber.App) {
	httpClient := binancespot.NewBinanceSpotHttpClient()
	repo := repository.NewMarketRepository(httpClient)
	usecase := usecase.NewMarketUsecase(repo)
	ctrl := market.NewMarketController(usecase)
	market.RegisterMarketRoutes(app, ctrl)
}
