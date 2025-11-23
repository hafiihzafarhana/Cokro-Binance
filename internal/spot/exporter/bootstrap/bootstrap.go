package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hafiihzafarhana/Cokro-Binance/internal/spot/exporter"
	"github.com/hafiihzafarhana/Cokro-Binance/internal/spot/exporter/usecase"
	"github.com/hafiihzafarhana/Cokro-Binance/internal/spot/market/repository"
	"github.com/hafiihzafarhana/Cokro-Binance/shared/httpclient/binancespot"
)

func SetupSpotExporterModule(app *fiber.App) {
	httpClient := binancespot.NewBinanceSpotHttpClient()
	repo := repository.NewMarketRepository(httpClient)
	usecase := usecase.NewExporterUsecase(repo)
	ctrl := exporter.NewExporterController(usecase)
	exporter.RegisterExporterRoutes(app, ctrl)
}
