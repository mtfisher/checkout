package pricing_test

import (
	"testing"

	"github.com/mtfisher/checkout/pricing"
	"github.com/stretchr/testify/assert"
)

func TestStock_CalculatePrice(t *testing.T) {
	s := pricing.Stock{UnitPrice: 20}
	result := s.CalculatePrice(2)
	assert.Exactly(t, 40.0, result.TotalPrices[0].Total)
	assert.Exactly(t, 2, result.TotalPrices[0].ItemQty)
	assert.Exactly(t, 20.0, result.UnitPrice)
}
