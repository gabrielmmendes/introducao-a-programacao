package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	var fahrenheit = make([]float32, x)

	for i := 0; i < x; i++ {
		fmt.Scan(&fahrenheit[i])
	}

	for i := 0; i < x; i++ {
		celsius := (5 * (fahrenheit[i] - 32)) / 9
		fmt.Printf("%.2f FAHRENHEIT EQUIVALE A %.2f CELSIUS \n", fahrenheit[i], celsius)
	}
}
