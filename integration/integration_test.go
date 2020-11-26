package integration_test

import (
	"testing"

	"github.com/mtfisher/checkout/checkout"
	"github.com/mtfisher/checkout/pricing"
	"github.com/stretchr/testify/assert"
)

//TestCatalogueCheckout_Integration main integration test
func TestCatalogueCheckout_Integration(t *testing.T) {
	integrationCatalogue := map[string]pricing.PriceCalculator{
		"A": pricing.NewXForY(pricing.NewStock(50), 3, 130),
		"B": pricing.NewXForY(pricing.NewStock(30), 2, 45),
		"C": pricing.NewStock(20),
		"D": pricing.NewStock(15),
	}

	checkout := checkout.NewCatalogueCheckout(integrationCatalogue)

	for i := 1; i <= 4; i++ {
		checkout.Scan("A")
	}

	for i := 1; i <= 3; i++ {
		checkout.Scan("B")
	}

	for i := 1; i <= 2; i++ {
		checkout.Scan("C")
	}

	checkout.Scan("D")

	//so i shoudl get 180(4 of a)+75(3 of b)+40(2 of c)+15(1 of d) which is 310 according to hand calculation
	assert.Exactly(t, 310.0, checkout.GetTotal())
}
