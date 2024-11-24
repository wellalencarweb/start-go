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

	//Declaring array with values
	myArray2 := [3]int{50, 60, 70}
	fmt.Println(myArray2)

	//Declaring multidimensional array
	myArray3 := [3][2]int{{10, 20}, {30, 40}, {50, 60}}
	fmt.Println(myArray3)
}
