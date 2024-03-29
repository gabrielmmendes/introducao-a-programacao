package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n1, n2, n3 int

	_, err := fmt.Scan(&n1, &n2, &n3)
	if err != nil {
		return
	}

	if n1 > 9 || n2 > 9 || n3 > 9 {
		fmt.Print("DIGITO INVALIDO")
		os.Exit(0)
	}

	n, _ := strconv.Atoi(strconv.Itoa(n1) + strconv.Itoa(n2) + strconv.Itoa(n3))
	quadrado := n * n
	fmt.Printf("%d, %d", n, quadrado)
}
