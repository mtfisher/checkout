package pricing

//XForY implements price calculator for group pricing
type XForY struct {
	StockPrice PriceCalculator
	GroupQty   int
	GroupPrice float64
}

//CalculatePrice calculates the price for given qty
func (x XForY) CalculatePrice(itemQty int) PriceResult {
	initialResult := x.StockPrice.CalculatePrice(itemQty)

	groupNumbers := itemQty / x.GroupQty
	if groupNumbers < 1 {
		return initialResult
	}

	//now work out single qty
	remaining := itemQty % x.GroupQty
	LineItems := []LineItem{
		LineItem{ItemQty: groupNumbers * x.GroupQty, Total: float64(groupNumbers) * x.GroupPrice},
	}

	if remaining > 0 {
		stockPrice := x.StockPrice.CalculatePrice(remaining)
		LineItems = append(LineItems, stockPrice.TotalPrices[0])
	}

	return PriceResult{
		UnitPrice:   initialResult.UnitPrice,
		TotalPrices: LineItems,
	}
}
