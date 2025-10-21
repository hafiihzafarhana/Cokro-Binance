package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hafiihzafarhana/Cokro-Binance/domain/spot/general"
	"github.com/hafiihzafarhana/Cokro-Binance/domain/spot/general/entity"
	binance "github.com/hafiihzafarhana/Cokro-Binance/shared/httpclient/binancespot"
)

type GeneralRepositoryImpl struct {
	client *binance.HttpClient
}

func NewGeneralRepository(client *binance.HttpClient) general.GeneralRepository {
	return &GeneralRepositoryImpl{client: client}
}

func (r *GeneralRepositoryImpl) GetServerTime(ctx context.Context) (*entity.GeneralServerTimeEntity, error) {
	resp, err := r.client.Get(ctx, "/time")
	if err != nil {
		return nil, err
	}
	
	var result entity.GeneralServerTimeEntity
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}