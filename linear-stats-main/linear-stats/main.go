package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("File not provided")
		return
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Cannot open file")
		return
	}
	defer f.Close()

	var y []float64
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		n, _ := strconv.ParseFloat(sc.Text(), 64)
		y = append(y, n)
	}

	if len(y) == 0 {
		fmt.Println("Empty file")
		return
	}

	n := float64(len(y))

	var x []float64
	for i := 0; i < int(n); i++ {
		x = append(x, float64(i))
	}

	var sumX, sumY, sumXY, sumX2, sumY2 float64
	for i := 0; i < int(n); i++ {
		sumX += x[i]
		sumY += y[i]
		sumXY += x[i] * y[i]
		sumX2 += x[i] * x[i]
		sumY2 += y[i] * y[i]
	}

	m := (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)
	b := (sumY - m*sumX) / n

	r := (n*sumXY - sumX*sumY) /
		(math.Sqrt((n*sumX2 - sumX*sumX) * (n*sumY2 - sumY*sumY)))

	fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n", m, b)
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", r)
}
