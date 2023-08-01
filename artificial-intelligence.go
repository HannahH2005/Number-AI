package main

import (
	"fmt"
)

func linearRegression(x, y []float64) (float64, float64) {
	n := float64(len(x))

	sumX, sumY, sumXY, sumXSquare := 0.0, 0.0, 0.0, 0.0
	for i := 0; i < len(x); i++ {
		sumX += x[i]
		sumY += y[i]
		sumXY += x[i] * y[i]
		sumXSquare += x[i] * x[i]
	}

	m := (n*sumXY - sumX*sumY) / (n*sumXSquare - sumX*sumX)
	b := (sumY - m*sumX) / n

	return m, b
}

func predictNextNumber(x float64, m float64, b float64) float64 {
	return m*x + b
}

func main() {
	x := []float64{1, 2, 3, 4, 5}
	y := []float64{2, 4, 6, 8, 10}

	m, b := linearRegression(x, y)

	nextNumber := predictNextNumber(float64(len(x)+1), m, b)

	fmt.Printf("Die nächste Zahl in der Folge ist: %.2f\n", nextNumber)
}
