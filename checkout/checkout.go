package checkout

//Checkout is for checkout implementations
type Checkout interface {
	Scan(item string)
	GetTotal() float64
}
