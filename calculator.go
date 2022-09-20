//package firstGoProject
package main

import (
	"fmt"
	"os"
	"strings"
)

var cont bool = true

func checkCont() {
	var input string
	fmt.Print("Желаете продолжить? (да/нет): ")
	fmt.Scanln(&input)
	for input != "да" {
		if input != "нет" {
			fmt.Print("Вы ввели что-то непонятное. Попробуйте снова!")
			fmt.Scanln(&input)
		} else {
			break
		}
	}
	input = strings.ToLower(input)
	switch input {
	case "да":
		cont = true
	case "нет":
		cont = false
	}
}

func main() {
	for cont {
		var number1, number2 float64
		var operator string

		fmt.Print("Введите первое число: ")
		fmt.Scanln(&number1)
		fmt.Print("Введите второе число: ")
		fmt.Scanln(&number2)

		fmt.Print("Введите операцию, которую необходимо совершить с числами (+, -, *, /): ")
		fmt.Scanln(&operator)

		switch operator {
		case "+":
			fmt.Printf("%.2f %s %.2f = %.2f \n", number1, operator, number2, number1+number2)
		case "-":
			fmt.Printf("%.2f %s %.2f = %.2f \n", number1, operator, number2, number1-number2)
		case "*":
			fmt.Printf("%.2f %s %.2f = %.2f \n", number1, operator, number2, number1*number2)
		case "/":
			for number2 == 0 {
				//panic("Нельзя делить на ноль!"
				fmt.Print("Нельзя делить на ноль! Введите другое второе число: ")
				fmt.Scanln(&number2)
			}
			fmt.Printf("%.2f %s %.2f = %.2f \n", number1, operator, number2, number1/number2)
		default:
			fmt.Println("Некорректный ввод оператора!")
		}
		checkCont()
	}
	os.Exit(3)

}
