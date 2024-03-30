package main

import (
	"fmt"
	"os"
)

func main() {
	var nota float32
	var conceito string

	fmt.Scan(&nota)

	if nota < 0 || nota > 10 {
		fmt.Printf("NOTA INVALIDA")
		os.Exit(0)
	}

	if nota >= 0 && nota < 6 {
		conceito = "D"
	} else if nota >= 6 && nota < 7.5 {
		conceito = "C"
	} else if nota >= 7.5 && nota < 9 {
		conceito = "B"
	} else {
		conceito = "A"
	}

	fmt.Printf("NOTA = %.1f CONCEITO = %s", nota, conceito)
}
