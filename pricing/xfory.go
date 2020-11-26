package pricing

//XForY implements price calculator for group pricing
type XForY struct {
	stockPrice PriceCalculator
	groupQty   int
	groupPrice float64
}

//NewXForY returns xfory price calculator
func NewXForY(stockPrice PriceCalculator, groupQty int, groupPrice float64) XForY {
	return XForY{
		stockPrice: stockPrice,
		groupQty:   groupQty,
		groupPrice: groupPrice,
	}
}

//CalculatePrice calculates the price for given qty
func (x XForY) CalculatePrice(itemQty int) PriceResult {
	initialResult := x.stockPrice.CalculatePrice(itemQty)

	groupNumbers := itemQty / x.groupQty
	if groupNumbers < 1 {
		return initialResult
	}

	//now work out single qty
	remaining := itemQty % x.groupQty
	LineItems := []LineItem{
		LineItem{ItemQty: groupNumbers * x.groupQty, Total: float64(groupNumbers) * x.groupPrice},
	}

	if remaining > 0 {
		stockPrice := x.stockPrice.CalculatePrice(remaining)
		LineItems = append(LineItems, stockPrice.TotalPrices[0])
	}

	return PriceResult{
		UnitPrice:   initialResult.UnitPrice,
		TotalPrices: LineItems,
	}
}
