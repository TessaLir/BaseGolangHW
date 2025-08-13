package main

import "fmt"

func main() {
	const (
		USDtoEUR = 0.85
		USDtoRUB = 79.78
		EURtoUSD = 1.16
		EURtoRUB = (USDtoRUB / USDtoEUR) * EURtoUSD
	)

	data := GetUserData()
	fmt.Println(data)

	fmt.Print(EURtoRUB)
}

func GetUserData() (data string) {
	fmt.Scan(&data)
	return
}

func calculateIMT(coin int, currencySource string, currencyTarget string) {

}
