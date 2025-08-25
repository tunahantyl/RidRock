package results

import "time"

// BacktestResult holds the results of a backtest
type BacktestResult struct {
	EquityCurve []float64
	PnL         []float64
	Trades      []Trade
}

// Trade represents a single trade execution
type Trade struct {
	Time   time.Time
	Action string
	Price  float64
	Qty    float64
	PnL    float64
}

// AddTrade adds a new trade to the results
func (r *BacktestResult) AddTrade(trade Trade) {
	r.Trades = append(r.Trades, trade)
}

// UpdateEquity adds a new equity value to the curve
func (r *BacktestResult) UpdateEquity(value float64) {
	r.EquityCurve = append(r.EquityCurve, value)
}
