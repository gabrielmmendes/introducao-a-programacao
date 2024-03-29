package main

import "fmt"

func main() {
	const PI = 3.14159

	var custoAluminioM2 float32 = 100
	var raio float32
	var altura float32

	fmt.Scan(&raio)
	fmt.Scan(&altura)

	custoTotal := custoAluminioM2 * ((2 * (PI * raio * raio)) + (2 * PI * raio * altura))

	fmt.Printf("O VALOR DO CUSTO E = %.2f \n", custoTotal)
}
