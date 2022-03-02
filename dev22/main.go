//Разработать программу,
//которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.
package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(222222222)
	b := big.NewInt(111111111)

	fmt.Println(new(big.Int).Add(a, b))
	fmt.Println(new(big.Int).Sub(a, b))
	fmt.Println(new(big.Int).Mul(a, b))
	fmt.Println(new(big.Int).Div(a, b))
}
