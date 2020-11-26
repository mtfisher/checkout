package pricing

//Stock implements price calculator for stock items
type Stock struct {
	UnitPrice float64
}

func (s Stock) CalculatePrice(itemQty int) PriceResult {
	return PriceResult{
		UnitPrice: s.UnitPrice,
		TotalPrices: []LineItem{
			LineItem{ItemQty: itemQty, Total: float64(itemQty) * s.UnitPrice},
		},
	}
}
