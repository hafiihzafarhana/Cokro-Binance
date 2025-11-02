package mapper

import (
	"testing"

	"github.com/hafiihzafarhana/Cokro-Binance/internal/spot/general/model"
	"github.com/stretchr/testify/assert"
)

func TestMapper(t *testing.T) {
	data := model.GetServerTimeModel {
		ServerTime: 1700000000000,
	}

	resp := ToServerTimeEntity(data)

	assert.Equal(t, resp.ServerTime, int64(1700000000000))
	assert.Equal(t, resp.ServerTimeStr, "2023-11-15 05:13:20")
}