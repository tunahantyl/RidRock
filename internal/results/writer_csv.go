package results

import (
	"encoding/csv"
	"os"
	"strconv"
	"sync"
)

// WriterCSV implements Writer interface for CSV output with thread-safety
type WriterCSV struct {
	Path string
	mu   sync.Mutex // Mutex for thread-safety
}

// Write writes BacktestResult to CSV file with thread-safety
func (w *WriterCSV) Write(result BacktestResult) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	file, err := os.Create(w.Path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write trades header
	tradesHeader := []string{"Time", "Action", "Price", "Quantity", "PnL"}
	if err := writer.Write(tradesHeader); err != nil {
		return err
	}

	// Write trades
	for _, trade := range result.Trades {
		record := []string{
			trade.Time.Format("2006-01-02 15:04:05"),
			trade.Action,
			strconv.FormatFloat(trade.Price, 'f', 8, 64),
			strconv.FormatFloat(trade.Qty, 'f', 8, 64),
			strconv.FormatFloat(trade.PnL, 'f', 8, 64),
		}
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	// Add empty line between trades and equity
	if err := writer.Write([]string{}); err != nil {
		return err
	}

	// Write equity curve header
	equityHeader := []string{"Index", "Equity"}
	if err := writer.Write(equityHeader); err != nil {
		return err
	}

	// Write equity curve
	for i, equity := range result.EquityCurve {
		record := []string{
			strconv.Itoa(i),
			strconv.FormatFloat(equity, 'f', 8, 64),
		}
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}
