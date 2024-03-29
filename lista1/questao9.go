package main

import "fmt"

func main() {
	var a float32
	var b float32
	var c float32

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	fmt.Printf("O VALOR DE DELTA E = %.2f \n", (b*b)-(4*a*c))
}
