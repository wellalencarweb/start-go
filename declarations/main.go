package main

/*importation packeges*/
import "fmt"

/*create types example*/
type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Tom"
	e float64 = 35.11
	f ID      = 1
)

func main() {
	a := "First atribuition"
	println(a)

	a = "Second atribuition"
	println(a)

	println(b, c, d, e)

	//print type creted
	println(f)

	//print use packege fmt type and value
	fmt.Printf("O Tipo de E é: %T, O Valor de E é:%v\n", e, e)
}
