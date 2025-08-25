package indicator

// RSI calculates the Relative Strength Index
// Input: slice of close prices
// Output: slice of RSI values (same length, first 'period' elements can be 0)
func RSI(values []float64, period int) []float64 {
	if len(values) < period+1 {
		return make([]float64, len(values))
	}

	result := make([]float64, len(values))

	// Calculate price changes
	changes := make([]float64, len(values)-1)
	for i := 1; i < len(values); i++ {
		changes[i-1] = values[i] - values[i-1]
	}

	// Calculate initial average gain and loss
	var avgGain, avgLoss float64
	for i := 0; i < period; i++ {
		if changes[i] > 0 {
			avgGain += changes[i]
		} else {
			avgLoss += -changes[i]
		}
	}

	avgGain /= float64(period)
	avgLoss /= float64(period)

	// Calculate RSI for the first valid period
	if avgLoss == 0 {
		result[period] = 100
	} else {
		rs := avgGain / avgLoss
		result[period] = 100 - (100 / (1 + rs))
	}

	// Calculate RSI for remaining values using smoothing
	for i := period + 1; i < len(values); i++ {
		var gain, loss float64
		if changes[i-1] > 0 {
			gain = changes[i-1]
		} else {
			loss = -changes[i-1]
		}

		avgGain = (avgGain*float64(period-1) + gain) / float64(period)
		avgLoss = (avgLoss*float64(period-1) + loss) / float64(period)

		if avgLoss == 0 {
			result[i] = 100
		} else {
			rs := avgGain / avgLoss
			result[i] = 100 - (100 / (1 + rs))
		}
	}

	return result
}
