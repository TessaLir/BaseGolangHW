package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	operation := GetOperation()
	numbers := GetNumbersForCalc()
	operationResult := CalculateNumbers(operation, numbers)
	fmt.Println(operationResult)
}

func GetOperation() string {
	for {
		fmt.Print("Введите желаемую операкию для калькулятора (AVG, SUM или MED): ")
		var operation string
		fmt.Scanln(&operation)
		switch operation {
		case "AVG":
		case "SUM":
		case "MED":
		default:
			fmt.Println("Вы ввели не корректную команду.")
			continue
		}
		return operation
	}
}

func GetNumbersForCalc() []int {
	var numbers []int
	fmt.Print("Введите числа разделенные запятой, с которыми необходимо выполнить поерацию: ")

	// var numbersLine string
	// fmt.Scanln(&numbersLine)

	reader := bufio.NewReader(os.Stdin)
	numbersLine, _ := reader.ReadString('\n')

	for _, line := range strings.Split(numbersLine, ",") {
		number, _ := strconv.Atoi(strings.TrimSpace(line))
		numbers = append(numbers, number)
	}
	return numbers
}

func CalculateNumbers(operation string, numbers []int) float64 {
	var result float64
	switch operation {
	case "AVG":
		sum := CalculateNumbers("SUM", numbers)
		result = float64(sum) / float64(len(numbers))
	case "SUM":
		for _, number := range numbers {
			result += float64(number)
		}
	case "MED":
		numbersLength := len(numbers)
		if numbersLength%2 != 0 {
			result = float64(numbers[numbersLength/2+1])
		} else {
			result = float64((numbers[numbersLength/2-1] + numbers[numbersLength/2]) / 2)
		}
	}
	return result
}
