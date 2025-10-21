package usecase

import (
	"context"

	"github.com/hafiihzafarhana/Cokro-Binance/domain/spot/general"
	"github.com/hafiihzafarhana/Cokro-Binance/internal/spot/general/dto"
)

type GeneralServiceImpl struct{
	repo general.GeneralRepository
}

func NewGeneralService(repo general.GeneralRepository) MarketUseCaseInterface {
    return &GeneralServiceImpl{repo: repo}
}

func (s *GeneralServiceImpl) GetServerTime(ctx context.Context) (*dto.GeneralServerTimeRes, error) {
	result, err := s.repo.GetServerTime(ctx)
	if err != nil {
		return nil, err
	}
	return ToGeneralServerTimeRes(result), nil
}