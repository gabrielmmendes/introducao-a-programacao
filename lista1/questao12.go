package main

import "fmt"

func main() {
	var horas int

	fmt.Scan(&horas)

	var taxa10 = horas / 3
	var taxa5 = horas % 3
	var valorDoAluguel = float32((taxa10 * 10) + (taxa5 * 5))

	fmt.Printf("O VALOR A PAGAR E = %.2f \n", valorDoAluguel)
}
