package engine

type OrderIntent struct {
	Symbol   string
	Side     string
	Quantity float64
	Price    float64
}

type Fill struct {
	Symbol   string
	Side     string
	Quantity float64
	Price    float64
	Fees     float64
	Slippage float64
}

type Portfolio struct{}

func (p *Portfolio) ProcessFill(fill Fill) {
	// Portfolio will process the fill
}

type SimBroker struct {
	portfolio *Portfolio
	fees      float64
	slippage  float64
}

func (b *SimBroker) ExecuteOrder(order OrderIntent) Fill {
	// Apply fees and slippage
	actualPrice := order.Price
	if order.Side == "buy" {
		actualPrice += order.Price * b.slippage
	} else {
		actualPrice -= order.Price * b.slippage
	}

	fees := order.Price * order.Quantity * b.fees

	fill := Fill{
		Symbol:   order.Symbol,
		Side:     order.Side,
		Quantity: order.Quantity,
		Price:    actualPrice,
		Fees:     fees,
		Slippage: b.slippage,
	}

	// Send Fill to Portfolio
	b.portfolio.ProcessFill(fill)

	return fill
}
