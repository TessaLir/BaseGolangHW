package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var operation string
	GetOperation(&operation)

	var numbers []int
	GetNumbersForCalc(numbers)

	var operationResult float64
	CalculateNumbers(operation, numbers, &operationResult)

	fmt.Println(operationResult)
}

func GetOperation(operation *string) {
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
	}
}

func GetNumbersForCalc(numbers []int) {
	fmt.Print("Введите числа разделенные запятой, с которыми необходимо выполнить поерацию: ")

	reader := bufio.NewReader(os.Stdin)
	numbersLine, _ := reader.ReadString('\n')

	for _, line := range strings.Split(numbersLine, ",") {
		number, _ := strconv.Atoi(strings.TrimSpace(line))
		numbers = append(numbers, number)
	}
}

func CalculateNumbers(operation string, numbers []int, operationResult *float64) {
	var result float64
	switch operation {
	case "AVG":
		var sum float64
		CalculateNumbers("SUM", numbers, &sum)
		result = float64(sum) / float64(len(numbers))
	case "SUM":
		for _, number := range numbers {
			result += float64(number)
		}
	case "MED":
		sort.Ints(numbers)
		numbersLength := len(numbers)
		if numbersLength%2 != 0 {
			result = float64(numbers[numbersLength/2])
		} else {
			result = float64(numbers[numbersLength/2-1]+numbers[numbersLength/2]) / 2.0
		}
	}
}
