# Algo Trading Backtesting System

A Go-based algorithmic trading backtesting system.

## Project Structure

```
/algo
├── api/                    # API endpoints
├── cmd/                    # Command line applications
│   └── backtestd/         # Main backtesting daemon
├── configs/                # Configuration files
├── internal/               # Internal packages
│   ├── config/            # Configuration management
│   ├── data/              # Data access layer
│   ├── engine/            # Backtesting engine
│   ├── indicator/         # Technical indicators
│   ├── logx/              # Logging utilities
│   ├── optimizer/         # Strategy optimization
│   ├── results/           # Results processing
│   └── strategy/          # Trading strategies
├── migrations/             # Database migrations
└── scripts/                # Utility scripts
```

## Getting Started

1. Install Go 1.21+
2. Clone the repository
3. Run `go mod tidy`
4. Run `go run cmd/backtestd/main.go`

## Configuration

Configuration files are located in the `configs/` directory.

## Implementation Progress

### ✅ Step 3 - Configuration Files
- Created `configs/default.yaml` with storage, engine, and workers configuration
- Created `configs/local.yaml` with local development overrides

### ✅ Step 4.1 - Engine Skeleton
- Implemented Engine struct with Run() method
- Added Candle type definition
- Basic candle iteration and logging functionality

### ✅ Step 4.3 - SimBroker
- Created OrderIntent and Fill structs
- Implemented SimBroker with ExecuteOrder method
- Added fees and slippage calculation
- Portfolio integration for fill processing

### ✅ Step 4.4 - Worker Isolation
- Each Engine instance has its own Portfolio and SimBroker
- NewEngine constructor ensures proper isolation
- Prevents race conditions in concurrent backtests

### ✅ Step 4.5 - Concurrency Setup
- Implemented concurrent candle feed using chan Candle
- Worker pool for parallel processing
- Each worker runs Engine.Run() independently
- Isolated Portfolio and SimBroker instances per worker

### ✅ Step 5.1 - RSI Indicator
- Implemented RSI function with standard formula
- Input: slice of close prices, Output: RSI values
- Gain/Loss calculation, RS = AverageGain/AverageLoss
- RSI = 100 - (100 / (1 + RS))

### ✅ Step 5.2 - Strategy Interface
- Created Strategy interface with OnCandle(c Candle) Signal method
- Defined Signal struct with Action ("BUY", "SELL", "HOLD") and Qty
- Strategy can be plugged into Engine

### ✅ Step 5.3 - RSI Reversion Strategy
- Implemented RSIReversion strategy struct
- Parameters: RSI period, buy threshold, sell threshold
- OnCandle method generates trading signals based on RSI values

### ✅ Step 5.4 - Engine Integration
- Integrated strategy into Engine.Run() method
- Each candle sent to Strategy.OnCandle()
- Signal processed and sent to SimBroker
- Broker executes orders and updates Portfolio
- Thread-safe for multi-worker setup

### ✅ Step 6.1 - CandleRepo Struct
- Created CandleRepo struct with Path field
- Implemented Stream method for CSV file reading
- Parses CSV lines into Candle structs
- Sends candles to channel and closes when done

### ✅ Step 6.2 - MarketData Feed Interface
- Created MarketDataFeed interface with Stream method
- CandleRepo implements MarketDataFeed interface
- Standardized contract for streaming market data

### ✅ Step 6.3 - Multi-thread Safe Streaming
- Enhanced CandleRepo.Stream with chunk processing
- 100-candle chunks to avoid blocking
- Graceful error handling with logging
- Channel closure guaranteed with defer
- Multi-thread safety improvements
