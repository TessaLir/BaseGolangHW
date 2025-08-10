package main

import "fmt"

func main() {
	const (
		USDtoEUR = 0.85
		USDtoRUB = 79.78
		EURtoUSD = 1.16
		EURtoRUB = (USDtoRUB / USDtoEUR) * EURtoUSD
	)

	fmt.Print(EURtoRUB)
}
