package checkout

import (
	"github.com/mtfisher/checkout/pricing"
)

//CatalogueCheckout implements checkout with catalogue for lookup and cart
type CatalogueCheckout struct {
	catalogue [string]PriceCalculator
	cart      [string]int
}

//NewCatalogueCheckout returns new checkout catalouge instance
func NewCatalogueCheckout(catalogue [string]pricing.PriceCalculator) CatalogueCheckout {
	checkout := new(CatalogueCheckout)
	checkout.catalogue = catalogue

	return checkout
}

//Scan scans a item in to cart
func (c *CatalogueCheckout) Scan(item string) {

}

//GetTotal returns the cart
func (c CatalogueCheckout) GetTotal(): float64 {
	//TODO: implement once tests are writen
	return 0.0
}
