package main

import (
	"fmt"
	"math"
)

func main() {
	var alturaPiramide float64
	var arestaHexagono float64

	fmt.Scan(&alturaPiramide, &arestaHexagono)

	var areaB = (3 * arestaHexagono * arestaHexagono * math.Sqrt(3)) / 2
	var volume = areaB * alturaPiramide / 3

	fmt.Printf("O VOLUME DA PIRAMIDE E = %.2f METROS CUBICOS", volume)
}
