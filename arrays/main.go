package main

import "fmt"

func main() {
	var myArray [3]int
	myArray[0] = 10
	myArray[1] = 20
	myArray[2] = 30

	//last position
	fmt.Println(len(myArray) - 1)
	fmt.Println(myArray[1])

	for i, v := range myArray {
		fmt.Printf("O Valor do indice é %d  e o valor é %d\n", i, v)
	}
}
