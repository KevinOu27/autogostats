package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

// RegressionCoefficients holds the slope and intercept for linear regression
type RegressionCoefficients struct {
	Slope     float64
	Intercept float64
}

// LinearRegression calculates the slope and intercept for linear regression
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

// ReadCSV reads and parses a CSV file in the specified format
func ReadCSV(filename string) ([]float64, []float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	// Skip header
	_, err = reader.Read()
	if err != nil {
		return nil, nil, err
	}

	var x, y []float64
	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, nil, err
		}

		// Parse X and Y values
		xVal, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			return nil, nil, err
		}
		yVal, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			return nil, nil, err
		}
		x = append(x, xVal)
		y = append(y, yVal)
	}

	return x, y, nil
}

func main() {
	// Define the filenames for the datasets
	filenames := []string{"anscombes.csv"}

	// Iterate over each dataset
	for _, filename := range filenames {
		// Read and parse the CSV file
		x, y, err := ReadCSV(filename)
		if err != nil {
			fmt.Printf("Error reading %s: %v\n", filename, err)
			continue
		}

		// Calculate linear regression coefficients
		slope, intercept := LinearRegression(x, y)

		// Print the coefficients for the current dataset
		fmt.Printf("Dataset: %s\n", filename)
		fmt.Printf("Slope: %.2f, Intercept: %.2f\n", slope, intercept)
		fmt.Println()
	}
}
