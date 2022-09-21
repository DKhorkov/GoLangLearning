//package firstGoProject
package main

import (
	"fmt"
	"os"
	"strings"
)

var cont bool = true

func main() {
	for cont {
		var number1, number2, result float64
		var operator string

		fmt.Print("Введите первое число: ")
		fmt.Scanln(&number1)
		fmt.Print("Введите второе число: ")
		fmt.Scanln(&number2)

		fmt.Print("Введите операцию, которую необходимо совершить с числами (+, -, *, /): ")
		fmt.Scanln(&operator)

		result = getResult(number1, number2, operator)
		fmt.Printf("%.2f %s %.2f = %.2f \n", number1, operator, number2, result)

		answer := checkCont()
		for answer {
			number1 = result
			fmt.Print("Введите второе число: ")
			fmt.Scanln(&number2)

			fmt.Print("Введите операцию, которую необходимо совершить с числами (+, -, *, /): ")
			fmt.Scanln(&operator)

			result = getResult(number1, number2, operator)
			fmt.Printf("%.2f %s %.2f = %.2f \n", number1, operator, number2, result)

			answer = checkCont()
		}
	}
}

func checkCont() bool {
	var input string
	fmt.Print("Желаете продолжить? (да/нет): ")
	fmt.Scanln(&input)
	for input != "да" {
		if input != "нет" {
			fmt.Print("Вы ввели что-то непонятное. Попробуйте снова: ")
			fmt.Scanln(&input)
		} else {
			break
		}
	}
	input = strings.ToLower(input)
	switch input {
	case "да":
		cont = true
		answer := contWithResult()
		return answer
	case "нет":
		os.Exit(3)
	}
	return false
}

func contWithResult() bool {
	var input string
	var answer bool
	fmt.Print("Желаете продолжить с полученным результатом? (да/нет): ")
	fmt.Scanln(&input)
	for input != "да" {
		if input != "нет" {
			fmt.Print("Вы ввели что-то непонятное. Попробуйте снова: ")
			fmt.Scanln(&input)
		} else {
			break
		}
	}
	switch input {
	case "да":
		answer = true
	case "нет":
		answer = false
	}
	return answer
}

func getResult(number1 float64, number2 float64, operator string) float64 {
	var result float64 = 0
	switch operator {
	case "+":
		result := number1 + number2
		return result
	case "-":
		result := number1 - number2
		return result
	case "*":
		result := number1 * number2
		return result
	case "/":
		for number2 == 0 {
			//panic("Нельзя делить на ноль!"
			fmt.Println("Нельзя делить на ноль! Введите другое второе число: ")
			fmt.Scanln(&number2)
		}
		result := number1 / number2
		return result
	default:
		fmt.Println("Некорректный ввод оператора! ")
	}
	return result
}
