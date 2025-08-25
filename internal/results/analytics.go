package results

import "math"

// Stats holds computed statistics for a backtest
type Stats struct {
	TotalPnL      float64
	MaxDrawdown   float64
	SharpeRatio   float64
	WinRate       float64
	TotalTrades   int
	WinningTrades int
}

// ComputeStats computes all statistics for the backtest result
func (r *BacktestResult) ComputeStats() Stats {
	stats := Stats{}

	// Total PnL
	stats.TotalPnL = r.computeTotalPnL()

	// Max Drawdown
	stats.MaxDrawdown = r.computeMaxDrawdown()

	// Sharpe Ratio
	stats.SharpeRatio = r.computeSharpeRatio()

	// Win Rate
	stats.WinRate = r.computeWinRate()

	// Trade counts
	stats.TotalTrades = len(r.Trades)
	stats.WinningTrades = r.countWinningTrades()

	return stats
}

// computeTotalPnL calculates total profit/loss
func (r *BacktestResult) computeTotalPnL() float64 {
	total := 0.0
	for _, trade := range r.Trades {
		total += trade.PnL
	}
	return total
}

// computeMaxDrawdown calculates maximum drawdown from equity curve
func (r *BacktestResult) computeMaxDrawdown() float64 {
	if len(r.EquityCurve) < 2 {
		return 0.0
	}

	maxDrawdown := 0.0
	peak := r.EquityCurve[0]

	for _, equity := range r.EquityCurve {
		if equity > peak {
			peak = equity
		}
		drawdown := (peak - equity) / peak
		if drawdown > maxDrawdown {
			maxDrawdown = drawdown
		}
	}

	return maxDrawdown
}

// computeSharpeRatio calculates Sharpe ratio (simplified version)
func (r *BacktestResult) computeSharpeRatio() float64 {
	if len(r.EquityCurve) < 2 {
		return 0.0
	}

	// Calculate returns
	returns := make([]float64, len(r.EquityCurve)-1)
	for i := 1; i < len(r.EquityCurve); i++ {
		if r.EquityCurve[i-1] != 0 {
			returns[i-1] = (r.EquityCurve[i] - r.EquityCurve[i-1]) / r.EquityCurve[i-1]
		}
	}

	// Calculate mean return
	meanReturn := 0.0
	for _, ret := range returns {
		meanReturn += ret
	}
	meanReturn /= float64(len(returns))

	// Calculate standard deviation
	variance := 0.0
	for _, ret := range returns {
		variance += math.Pow(ret-meanReturn, 2)
	}
	variance /= float64(len(returns))
	stdDev := math.Sqrt(variance)

	// Sharpe ratio (assuming risk-free rate = 0)
	if stdDev == 0 {
		return 0.0
	}
	return meanReturn / stdDev
}

// computeWinRate calculates percentage of winning trades
func (r *BacktestResult) computeWinRate() float64 {
	if len(r.Trades) == 0 {
		return 0.0
	}

	winningTrades := r.countWinningTrades()
	return float64(winningTrades) / float64(len(r.Trades))
}

// countWinningTrades counts trades with positive PnL
func (r *BacktestResult) countWinningTrades() int {
	count := 0
	for _, trade := range r.Trades {
		if trade.PnL > 0 {
			count++
		}
	}
	return count
}
