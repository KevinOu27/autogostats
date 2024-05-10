package main

import (
	"fmt"

	"gonum.org/v1/gonum/stat"
)

func main() {
	x := []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	y := []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68}

	alpha, beta := stat.LinearRegression(x, y, nil, false)
	fmt.Printf("y = %f + %fx\n", alpha, beta)
}
