package main

import (
	"fmt"
	"os"
)

func main() {
	var conta int
	var consumo float32
	var tipoConsumidor string
	_, err := fmt.Scan(&conta, &consumo, &tipoConsumidor)
	if err != nil {
		return
	}

	var valorDaConta float32

	switch tipoConsumidor {
	case "R":
		{
			valorDaConta = 5 + (0.05 * consumo)
		}
	case "C":
		{
			if consumo <= 80 {
				valorDaConta = 500
			} else {
				valorDaConta = 500 + ((consumo - 80) * 0.25)
			}
		}
	case "I":
		{
			if consumo <= 100 {
				valorDaConta = 800
			} else {
				valorDaConta = 800 + ((consumo - 100) * 0.04)
			}
		}
	default:
		fmt.Printf("TIPO INVÁLIDO")
		os.Exit(0)
	}

	fmt.Printf("CONTA = %d \n", conta)
	fmt.Printf("VALOR DA CONTA = %.2f \n", valorDaConta)
}
