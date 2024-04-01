package main

import (
	"fmt"
	"os"
)

func main() {
	var numeroPar int
	var numeroDeImpressoes int

	fmt.Scan(&numeroPar, &numeroDeImpressoes)

	if numeroPar%2 != 0 {
		fmt.Printf("O PRIMEIRO NUMERO NAO E PAR")
		os.Exit(0)
	} else {
		for i := 0; i < numeroDeImpressoes; i++ {
			fmt.Printf("%d ", numeroPar)
			numeroPar += 2
		}
	}
}
