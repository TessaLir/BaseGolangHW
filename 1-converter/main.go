package main

import (
	"errors"
	"fmt"
	"strconv"
)

const (
	USDtoEUR = 0.85
	USDtoRUB = 79.78
	EURtoUSD = 1.16
	EURtoRUB = (USDtoRUB / USDtoEUR) * EURtoUSD
	RUBtoEUR = 0.010671
	RUBtoUSD = 0.012497
)

var firstCurrency, secondCurrency string
var cash int
var err error

func main() {

	for {
		firstCurrency, err = GetUserCurrencyData()
		if isValidData(err) {
			break
		}
	}

	for {
		cash, err = GetUserCashData()
		if isValidData(err) {
			break
		}
	}

	for {
		secondCurrency, err = GetUserCurrencyData()
		if isValidData(err) {
			break
		}
	}

	result := calculateIMT(cash, firstCurrency, secondCurrency)
	fmt.Printf("%.2f", result)

}

func GetUserCurrencyData() (data string, err error) {
	printLineForEnterCurrency()

	fmt.Scan(&data)

	if firstCurrency != "" && data == firstCurrency {
		err = errors.New("NOT_VALID")
	}

	if data != "USD" && data != "EUR" && data != "RUB" {
		err = errors.New("NOT_VALID")
	}

	return
}

func GetUserCashData() (data int, err error) {
	fmt.Print("Введите количество денег: ")

	var input string
	fmt.Scan(&input)

	if data, err = strconv.Atoi(input); err != nil {
		err = errors.New("NOT_VALID")
	}

	return
}

func isValidData(err error) bool {
	if err != nil {
		fmt.Println("Вы ввели не корректное значение, повторите пожалуйста попытку ввода.")
		return false
	}
	return true
}

func printLineForEnterCurrency() {
	var currencyLine string = "EUR, USD, RUB"

	switch firstCurrency {
	case "EUR":
		currencyLine = "USD, RUB"
	case "USD":
		currencyLine = "EUR, RUB"
	case "RUB":
		currencyLine = "EUR, USD"
	}

	fmt.Printf("Введите желаемую валюту (%s): ", currencyLine)
}

func calculateIMT(coin int, currencySource string, currencyTarget string) float64 {
	var result float64

	switch {
	case currencySource == "EUR" && currencyTarget == "RUB":
		result = float64(coin) * EURtoRUB
	case currencySource == "EUR" && currencyTarget == "USD":
		result = float64(coin) * EURtoUSD
	case currencySource == "USD" && currencyTarget == "RUB":
		result = float64(coin) * USDtoRUB
	case currencySource == "USD" && currencyTarget == "EUR":
		result = float64(coin) * USDtoEUR
	case currencySource == "RUB" && currencyTarget == "USD":
		result = float64(coin) * RUBtoUSD
	case currencySource == "RUB" && currencyTarget == "EUR":
		result = float64(coin) * RUBtoEUR
	}

	return result
}
