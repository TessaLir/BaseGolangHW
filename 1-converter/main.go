package main

import (
	"errors"
	"fmt"
	"strconv"
)

var currenciesMap = make(map[string]float64)

var firstCurrency, secondCurrency string
var cash int
var err error

func main() {

	initCurency()

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

func initCurency() {
	currenciesMap["USDtoEUR"] = 0.85
	currenciesMap["USDtoRUB"] = 79.78
	currenciesMap["EURtoUSD"] = 1.16
	currenciesMap["EURtoRUB"] = (currenciesMap["USDtoRUB"] / currenciesMap["USDtoEUR"]) * currenciesMap["EURtoUSD"]
	currenciesMap["RUBtoEUR"] = 0.010671
	currenciesMap["RUBtoUSD"] = 0.012497
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
	var currency = fmt.Sprintf("%sto%s", currencySource, currencyTarget)
	return float64(coin) * currenciesMap[currency]
}
