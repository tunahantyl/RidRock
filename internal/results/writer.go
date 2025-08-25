package results

// Writer interface defines the contract for writing backtest results
type Writer interface {
	Write(result BacktestResult) error
}

// WriterJSON implements Writer interface for JSON output (optional for future)
type WriterJSON struct {
	Path string
}

// Write writes BacktestResult to JSON file (placeholder for future implementation)
func (w *WriterJSON) Write(result BacktestResult) error {
	// TODO: Implement JSON writing functionality
	return nil
}
