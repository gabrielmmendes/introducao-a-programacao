package main

import "fmt"

func main() {
	var salarioMinimo float32

	_, err := fmt.Scan(&salarioMinimo)
	if err != nil {
		return
	}

	var cemKW = salarioMinimo * 0.7
	var umKW = cemKW / 100
	var qtdKWGasta float32

	_, err = fmt.Scan(&qtdKWGasta)
	if err != nil {
		return
	}

	var custoDoConsumo = umKW * qtdKWGasta
	var custoComDesconto = umKW * qtdKWGasta * 0.9

	fmt.Printf("Custo por kW: R$ %.2f \n", umKW)
	fmt.Printf("Custo do consumo: R$ %.2f \n", custoDoConsumo)
	fmt.Printf("Custo com desconto: R$ %.2f \n", custoComDesconto)
}
