package main

import (
	"math/rand"
	"testing"
	"time"
)

// LinearRegression calculates the slope and intercept for linear regression
func LinearRegression(x, y []float64) (slope, intercept float64) {
	n := len(x)
	if n != len(y) || n == 0 {
		return 0, 0
	}

	var sumX, sumY, sumXY, sumXSquare float64

	for i := 0; i < n; i++ {
		sumX += x[i]
		sumY += y[i]
		sumXY += x[i] * y[i]
		sumXSquare += x[i] * x[i]
	}

	meanX := sumX / float64(n)
	meanY := sumY / float64(n)

	slope = (float64(n)*sumXY - sumX*sumY) / (float64(n)*sumXSquare - sumX*sumX)
	intercept = meanY - slope*meanX

	return slope, intercept
}

// BenchmarkLinearRegression benchmarks the LinearRegression function
func BenchmarkLinearRegression(b *testing.B) {
	// Generate random test data
	rand.Seed(time.Now().UnixNano())
	x := make([]float64, b.N)
	y := make([]float64, b.N)
	for i := 0; i < b.N; i++ {
		x[i] = rand.Float64()
		y[i] = rand.Float64()
	}

	// Run the benchmark
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LinearRegression(x, y)
	}
}
