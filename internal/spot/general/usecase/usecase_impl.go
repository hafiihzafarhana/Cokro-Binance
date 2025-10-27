package usecase

import (
	"context"

	"github.com/hafiihzafarhana/Cokro-Binance/domain/spot/general"
	"github.com/hafiihzafarhana/Cokro-Binance/domain/spot/general/entity"
)

type GeneralUsecaseImpl struct{
	repo general.GeneralRepository
}

func NewGeneralUsecase(repo general.GeneralRepository) MarketUsecaseInterface {
    return &GeneralUsecaseImpl{repo: repo}
}

func (s *GeneralUsecaseImpl) GetServerTime(ctx context.Context) (*entity.GeneralServerTimeEntity, error) {
	result, err := s.repo.GetServerTime(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}