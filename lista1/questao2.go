package main

import "fmt"

func main() {
	var publicoTotal float32

	var porcentagemPopular float32
	var porcentagemGeral float32
	var porcentagemArquibancada float32
	var porcentagemCadeiras float32

	var valorPopular float32 = 1
	var valorGeral float32 = 5
	var valorArquibancada float32 = 10
	var valorCadeiras float32 = 20

	var numeroDeJogos int

	_, err := fmt.Scan(&numeroDeJogos)
	if err != nil {
		return
	}

	var rendaTotalDosJogos = make([]float32, numeroDeJogos)

	for i := 0; i < numeroDeJogos; i++ {
		_, err := fmt.Scan(&publicoTotal, &porcentagemPopular, &porcentagemGeral, &porcentagemArquibancada, &porcentagemCadeiras)
		if err != nil {
			return
		}
		rendaTotalDosJogos[i] =
			(publicoTotal * (porcentagemPopular / 100) * valorPopular) +
				(publicoTotal * (porcentagemGeral / 100) * valorGeral) +
				(publicoTotal * (porcentagemArquibancada / 100) * valorArquibancada) +
				(publicoTotal * (porcentagemCadeiras / 100) * valorCadeiras)
	}

	for i := 0; i < numeroDeJogos; i++ {
		fmt.Printf("A RENDA DO JOGO N. %d E = %.2f\n", i+1, rendaTotalDosJogos[i])
	}

}
