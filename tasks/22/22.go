package main

import (
	"fmt"
	"math/big"
)

// для работы с большими числами можно использовать пакет big
func main() {
	aBig := big.Int{}
	bBig := big.Int{}
	aBig.SetString("1267650600228229401496703205376", 31)
	bBig.SetString("5267650600229222401296703205376", 31)

	fmt.Printf(
		"sum (a + b) is %v\n"+
			"multiply (a * b) is %v\n"+
			"divide (b / a) is %v\n"+
			"sub (b - a) is %v",
		new(big.Int).Add(&aBig, &bBig),
		new(big.Int).Mul(&aBig, &bBig),
		new(big.Int).Div(&bBig, &aBig),
		new(big.Int).Sub(&aBig, &bBig),
	)

}
