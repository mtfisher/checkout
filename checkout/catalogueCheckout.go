package checkout

import (
	"strings"

	"github.com/mtfisher/checkout/pricing"
)

//CatalogueCheckout implements checkout with catalogue for lookup and cart
type CatalogueCheckout struct {
	catalogue map[string]pricing.PriceCalculator
	cart      map[string]int
}

//NewCatalogueCheckout returns new checkout catalouge instance
func NewCatalogueCheckout(catalogue map[string]pricing.PriceCalculator) CatalogueCheckout {
	checkout := new(CatalogueCheckout)
	checkout.catalogue = catalogue
	checkout.cart = make(map[string]int)

	return *checkout
}

//Scan scans a item in to cart
func (c *CatalogueCheckout) Scan(item string) {
	key := strings.ToUpper(item)

	if _, ok := c.catalogue[key]; !ok {
		return
	}

	if _, ok := c.cart[key]; ok {
		c.cart[key]++
	} else {
		c.cart[key] = 1
	}
}

//GetTotal returns the cart
func (c CatalogueCheckout) GetTotal() float64 {
	total := 0.0

	for item, qty := range c.cart {
		if _, ok := c.catalogue[item]; ok {
			priceResult := c.catalogue[item].CalculatePrice(qty)

			for _, result := range priceResult.TotalPrices {
				total += result.Total
			}
		}
	}

	return total
}
