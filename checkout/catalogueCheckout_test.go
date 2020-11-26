package checkout_test

import (
	"testing"

	"github.com/mtfisher/checkout/checkout"
	"github.com/mtfisher/checkout/pricing"
	"github.com/stretchr/testify/assert"
)

type mockPricing struct {
	Price float64
}

//CalculatePrice returns mock price
func (m mockPricing) CalculatePrice(itemQty int) pricing.PriceResult {
	return pricing.PriceResult{
		TotalPrices: []pricing.LineItem{
			pricing.LineItem{ItemQty: itemQty, Total: m.mockPricing},
		},
	}
}

//TestCatalogueCheckout_NewCatalogueCheckout tests create method
func TestCatalogueCheckout_NewCatalogueCheckout(t *testing.T) {
	item1 := mockPricing{Price: 10}

	checkout := checkout.NewCatalogueCheckout(map[string]pricing.PriceCalculator{"A": item1})
	assert.NotEmpty(t, checkout)
}

//TestCatalogueCheckout_GetTotal tests CatalogueCheckout.GetTotal
func TestCatalogueCheckout_GetTotal(t *testing.T) {
	item1 := mockPricing{Price: 10}
	item2 := mockPricing{Price: 20}

	checkout := checkout.NewCatalogueCheckout(map[string]pricing.PriceCalculator{"A": item1, "B": item2})

	checkout.Scan("a")
	checkout.Scan("B")

	assert.Exactly(t, 30.0, checkout.GetTotal)
}
