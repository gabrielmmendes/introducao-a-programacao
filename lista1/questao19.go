package main

import (
	"fmt"
	"os"
)

func main() {
	var n int
	var soma float64

	fmt.Scan(&n)

	if n <= 1 {
		fmt.Printf("Numero invalido!")
		os.Exit(0)
	}

	for i := 0; i < n; i++ {
		soma += 1 / (float64(i) + 1)
	}

	fmt.Printf("%.6f", soma)
}
