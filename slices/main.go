package main

import "fmt"

func main() {
	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	//Exibe o tamanho, capacidade e valores gerais
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	// Não pega nenhuma posição
	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])

	//Pega apenas as 4 primeiras posições
	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4])

	//Ignora as duas primeiras posições
	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])

	s = append(s, 110) // dobra a capacidade do slice
	//Pega as duas primeiras posições
	fmt.Printf("len=%d cap=%d %v\n", len(s[:2]), cap(s[:2]), s[:2])
}
