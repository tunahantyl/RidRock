package data

import "algo/internal/engine"

// MarketDataFeed interface defines the contract for streaming market data
type MarketDataFeed interface {
	Stream(c chan<- engine.Candle) error
}
