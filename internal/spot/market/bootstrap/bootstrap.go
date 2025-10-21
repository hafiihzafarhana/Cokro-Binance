package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hafiihzafarhana/Cokro-Binance/internal/spot/general"
	"github.com/hafiihzafarhana/Cokro-Binance/internal/spot/general/repository"
	"github.com/hafiihzafarhana/Cokro-Binance/internal/spot/general/usecase"
	"github.com/hafiihzafarhana/Cokro-Binance/shared/httpclient/binancespot"
)

func SetupSpotGeneralModule(app *fiber.App) {
	httpClient := binancespot.NewBinanceSpotHttpClient()
	repo := repository.NewGeneralRepository(httpClient)
	usecase := usecase.NewGeneralService(repo)
	ctrl := general.NewGeneralController(usecase)
	general.RegisterGeneralRoutes(app, ctrl)
}
