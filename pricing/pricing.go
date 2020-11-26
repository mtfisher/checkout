package pricing

//LineItem stores item quantity and whats the total for that quantity
type LineItem struct {
	ItemQty int
	Total   float64
}

//PriceResult stores data on what the pricing caulculators
type PriceResult struct {
	Name        string
	UnitPrice   float64
	TotalPrices []LineItem
}

//PriceCalculator is interface for price calulators
type PriceCalculator interface {
	CalculatePrice(itemQty int) PriceResult
}
