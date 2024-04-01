package main

import (
	"fmt"
	"os"
)

func main() {
	var n int

	fmt.Scan(&n)

	if n <= 5 || n >= 2000 {
		fmt.Printf("NUMERO INVALIDO")
		os.Exit(0)
	}

	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			fmt.Printf("%d ^ 2 = %d \n", i, i*i)
		}
	}
}
