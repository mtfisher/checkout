package pricing_test

import (
	"testing"

	"github.com/mtfisher/checkout/pricing"
	"github.com/stretchr/testify/assert"
)

//TestXForY_CalculatePrice tests XForY.CalculatePrice
func TestXForY_CalculatePrice(t *testing.T) {
	s := pricing.NewStock(50)
	x := pricing.NewXForY(s, 3, 130)

	result := x.CalculatePrice(7)
	assert.Exactly(t, 260.0, result.TotalPrices[0].Total)
	assert.Exactly(t, 6, result.TotalPrices[0].ItemQty)
	assert.Exactly(t, 50.0, result.TotalPrices[1].Total)
	assert.Exactly(t, 1, result.TotalPrices[1].ItemQty)
	assert.Exactly(t, 50.0, result.UnitPrice)

	secondResult := x.CalculatePrice(6)
	assert.Exactly(t, 260.0, secondResult.TotalPrices[0].Total)
	assert.Exactly(t, 6, secondResult.TotalPrices[0].ItemQty)
	assert.Exactly(t, 1, len(secondResult.TotalPrices))

	thirdResult := x.CalculatePrice(2)
	assert.Exactly(t, 100.0, thirdResult.TotalPrices[0].Total)
	assert.Exactly(t, 2, thirdResult.TotalPrices[0].ItemQty)
	assert.Exactly(t, 1, len(thirdResult.TotalPrices))
}
