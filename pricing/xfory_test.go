package pricing_test

import (
	"testing"

	"github.com/mtfisher/checkout/pricing"
	"github.com/stretchr/testify/assert"
)

//TestXForY_CalculatePrice tests XForY.CalculatePrice
func TestXForY_CalculatePrice(t *testing.T) {
	s := pricing.Stock{UnitPrice: 50}
	x := pricing.XForY{StockPrice: s, GroupQty: 3, GroupPrice: 130}

	result := x.CalculatePrice(7)
	assert.Exactly(t, 260.0, result.TotalPrices[0].Total)
	assert.Exactly(t, 6, result.TotalPrices[0].ItemQty)
	assert.Exactly(t, 50.0, result.TotalPrices[1].Total)
	assert.Exactly(t, 1, result.TotalPrices[1].ItemQty)
	assert.Exactly(t, 50.0, result.UnitPrice)
}
