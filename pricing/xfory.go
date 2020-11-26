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

	if itemQty == 0 {
		//return stock if someone tries to feed in 0
		return initialResult
	}

	groupNumbers := itemQty / x.GroupQty

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
