package strategy

import "algo/internal/engine"

// RSIReversion implements a mean reversion strategy based on RSI
type RSIReversion struct {
	period        int
	buyThreshold  float64
	sellThreshold float64
}

// NewRSIReversion creates a new RSI Reversion strategy
func NewRSIReversion(period int, buyThreshold, sellThreshold float64) *RSIReversion {
	return &RSIReversion{
		period:        period,
		buyThreshold:  buyThreshold,
		sellThreshold: sellThreshold,
	}
}

// OnCandle processes a new candle and generates trading signals
func (r *RSIReversion) OnCandle(c engine.Candle) engine.Signal {
	// This is a simplified implementation
	// In a real scenario, you would need to maintain historical data
	// and calculate RSI based on the last 'period' candles

	// For now, return HOLD signal
	// TODO: Implement actual RSI calculation and signal generation
	return engine.Signal{
		Action: "HOLD",
		Qty:    0,
	}
}
