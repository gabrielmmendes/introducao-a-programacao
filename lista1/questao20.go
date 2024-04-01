package main

import "fmt"

func main() {
	var (
		horas    = 0
		minutos  = 0
		segundos = 0
	)

	fmt.Scan(&horas)
	fmt.Scan(&minutos)
	fmt.Scan(&segundos)

	fmt.Printf("O TEMPO EM SEGUNDOS E = %d \n", (horas*60*60)+(minutos*60)+segundos)
}
