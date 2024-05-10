//go:generate gotests -w -all auto_regression_test.go

package main

import (
	"math"
	"testing"

	"gonum.org/v1/gonum/stat"
)

func TestLinearRegression(t *testing.T) {
	x := []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	y := []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68}

	alpha, beta := stat.LinearRegression(x, y, nil, false)
	expectedAlpha, expectedBeta := 3.00, 0.500
	tolerance := 0.001 // Define a tolerance level

	if math.Abs(alpha-expectedAlpha) > tolerance || math.Abs(beta-expectedBeta) > tolerance {
		t.Errorf("Expected alpha=%f, beta=%f; Got alpha=%f, beta=%f", expectedAlpha, expectedBeta, alpha, beta)
	}
}

func BenchmarkLinearRegression(b *testing.B) {
	x := []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	y := []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68}
	for i := 0; i < b.N; i++ {
		stat.LinearRegression(x, y, nil, false)
	}
}

func TestBenchmarkLinearRegression(t *testing.T) {
	type args struct {
		b *testing.B
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BenchmarkLinearRegression(tt.args.b)
		})
	}
}
