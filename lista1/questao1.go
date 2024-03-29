package main

import "fmt"

func main() {

	var nota1 float32
	var nota2 float32
	var nota3 float32
	var media float32

	_, err := fmt.Scan(&nota1, &nota2, &nota3)
	if err != nil {
		return
	}

	media = (nota1 + nota2 + nota3) / 3

	fmt.Printf("MEDIA = %.2f \n", media)

	if media >= 6 {
		fmt.Printf("APROVADO \n")
	} else {
		fmt.Printf("REPROVADO \n")
	}
}
