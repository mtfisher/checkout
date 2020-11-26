package pricing

//Stock implements price calculator for stock items
type Stock struct {
	unitPrice float64
}

//NewStock returns stock price calculator
func NewStock(unitPrice float64) Stock {
	return Stock{
		unitPrice: unitPrice,
	}
}

//CalculatePrice calculates the price for given qty
func (s Stock) CalculatePrice(itemQty int) PriceResult {
	return PriceResult{
		UnitPrice: s.unitPrice,
		TotalPrices: []LineItem{
			LineItem{ItemQty: itemQty, Total: float64(itemQty) * s.unitPrice},
		},
	}
}
