package main

import "fmt"

func main() {
	var a float32
	var b float32
	var c float32
	var d float32

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Scan(&d)

	fmt.Printf("O VALOR DO DETERMINANTE E = %.2f \n", (a*d)-(b*c))
}
