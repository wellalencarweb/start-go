package main

import "fmt"

func main() {
	//declarando maps
	sal := make(map[string]int)
	valores := map[string]int{}

	sal["teste"] = 100
	valores["test"] = 200

	salarios := map[string]int{"Tom": 20000, "Well": 15000, "Wellington": 30000}

	fmt.Println(salarios)

	fmt.Println(salarios["Well"])

	delete(salarios, "Tom")

	fmt.Println(salarios)

	salarios["WellTom"] = 35000

	fmt.Println(salarios)

	valuesDistinct := map[string]interface{}{"intValue": 10, "stringValue": "Hello", "floatValue": 3.14}

	fmt.Println(valuesDistinct)
	fmt.Println(valuesDistinct["stringValue"])

	type Valor struct {
		IntVal    int
		StringVal string
		FloatVal  float64
	}

	valWithStruct := map[string]Valor{
		"val1": {IntVal: 10},
		"val2": {StringVal: "Hello"},
		"val3": {FloatVal: 3.14},
	}

	// Acessar os valores diretamente pela struct
	fmt.Println("Int value:", valWithStruct["val1"].IntVal)
	fmt.Println("String value:", valWithStruct["val2"].StringVal)
	fmt.Println("Float value:", valWithStruct["val3"].FloatVal)

}
