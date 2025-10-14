package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hafiihzafarhana/Cokro-Binance/domain/general"
)

type GeneralService struct{}

func NewGeneralService() general.GeneralServiceInterface {
	return &GeneralService{}
}

func (s *GeneralService) GetBinanceServerTime() (map[string]interface{}, error) {
	resp, err := http.Get("https://api.binance.com/api/v3/time")
	if err != nil {
		return nil, fmt.Errorf("failed to request Binance API: %v", err)
	}
	
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode Binance response: %v", err)
	}

	return result, nil
}
