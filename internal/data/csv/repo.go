package csv

import (
	"algo/internal/engine"
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"time"
)

// CandleRepo handles CSV data reading for candles
// Implements data.MarketDataFeed interface
type CandleRepo struct {
	Path string
}

// Stream reads CSV file and sends candles to the provided channel
// Implements data.MarketDataFeed.Stream method
func (r *CandleRepo) Stream(c chan<- engine.Candle) error {
	defer close(c) // Ensure channel is closed after all candles are sent

	file, err := os.Open(r.Path)
	if err != nil {
		return err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)

	// Skip header row
	_, err = csvReader.Read()
	if err != nil {
		return err
	}

	// Read CSV records
	records, err := csvReader.ReadAll()
	if err != nil {
		return err
	}

	// Process candles in chunks to avoid blocking
	chunkSize := 100
	for i := 0; i < len(records); i += chunkSize {
		end := i + chunkSize
		if end > len(records) {
			end = len(records)
		}

		// Process chunk
		for j := i; j < end; j++ {
			record := records[j]

			if len(record) < 2 {
				log.Printf("Skipping invalid record at line %d: insufficient columns", j+2)
				continue // Skip invalid records
			}

			// Parse timestamp (assuming first column is timestamp)
			timestamp, err := time.Parse("2006-01-02 15:04:05", record[0])
			if err != nil {
				log.Printf("Skipping record at line %d: invalid timestamp '%s': %v", j+2, record[0], err)
				continue // Skip records with invalid timestamp
			}

			// Parse close price (assuming second column is close price)
			closePrice, err := strconv.ParseFloat(record[1], 64)
			if err != nil {
				log.Printf("Skipping record at line %d: invalid price '%s': %v", j+2, record[1], err)
				continue // Skip records with invalid price
			}

			candle := engine.Candle{
				Time:  timestamp,
				Close: closePrice,
			}

			c <- candle
		}

		// Small delay between chunks to prevent blocking
		if end < len(records) {
			time.Sleep(1 * time.Millisecond)
		}
	}

	return nil
}
