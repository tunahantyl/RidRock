package strategy

import "algo/internal/engine"

// Strategy interface is now defined in engine package
// This file provides type aliases for convenience

// Strategy represents the trading strategy interface
type Strategy = engine.Strategy

// Signal represents a trading signal
type Signal = engine.Signal
