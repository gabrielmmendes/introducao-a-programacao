package main

import "fmt"

func main() {
	var valorInicial int
	var razao int
	var n int
	var soma int

	fmt.Scan(&valorInicial, &razao, &n)

	for i := 0; i < n; i++ {
		soma += valorInicial + (i * razao)
	}

	fmt.Printf("%d", soma)
}
