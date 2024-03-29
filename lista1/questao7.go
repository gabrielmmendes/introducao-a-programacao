package main

import "fmt"

func main() {
	var fahrenheit float32
	var qtdChuvaPolegadas float32

	fmt.Scan(&fahrenheit)
	fmt.Scan(&qtdChuvaPolegadas)

	var celsius = ((5 * fahrenheit) - 160) / 9
	var qtdChuvaMilimetro = qtdChuvaPolegadas * 25.4

	fmt.Printf("O VALOR EM CELSIUS = %.2f \n", celsius)
	fmt.Printf("A QUANTIDADE DE CHUVA E = %.2f \n", qtdChuvaMilimetro)
}
