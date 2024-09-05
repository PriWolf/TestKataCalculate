package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {

	var num1, num2, symbol, end = getInput()
	checkBaseInput(num1, symbol, num2, end)

	var num1Int, num2Int, types = convertTypeNumber(num1, num2)
	//types = true - Arabian | false - Rome

	var sum = arithmetic(num1Int, num2Int, symbol)

	answer(sum, types)

}

func getInput() (string, string, string, string) {
	//получение данных от пользователя
	fmt.Println(
		"Это приложение калькулятор\n" +
			"Введите математическое выражение по примеру: '1 + 2'\n" +
			"Вы можете использовать римские(I,V,X) и арабские(1,5,10)\n" +
			"Числа могут быть от 1 до 10 включительно\n" +
			"Типы поддерживаемых операций: '+ - / *'\n")

	var num1, num2, symbol, end string
	fmt.Scanln(&num1, &symbol, &num2, &end)

	return num1, num2, symbol, end
}

func checkBaseInput(num1, symbol, num2, end string) {
	//Проверка первичных введенных данных под условия
	if num2 == "" || end != "" {
		panic("неправильно введины данные")
		return
	}
}

func convertTypeNumber(num1, num2 string) (int, int, bool) {
	var num1Int, num2Int int

	var num1Rome, num2Rome bool
	num1Rome, _ = regexp.MatchString("^(X|IX|IV|V?I{0,3})$", num1)
	num2Rome, _ = regexp.MatchString("^(X|IX|IV|V?I{0,3})$", num2)

	//fmt.Println(num1Rome, "1 Условия 2", num2Rome) //проверка условий

	if num1Rome && num2Rome {
		//Римская империя
		num1Int = convertRomeToInt(num1)
		num2Int = convertRomeToInt(num2)

		//fmt.Println("Число1: ", num1Int, "Число2: ", num2Int) //проверка выводимых

		if checkNumValue(num1Int, num2Int) {
			return num1Int, num2Int, false
		}
	}

	num1Int, err1 := strconv.Atoi(num1)
	num2Int, err2 := strconv.Atoi(num2)
	if err1 == nil && err2 == nil {
		//Арабская империя
		if checkNumValue(num1Int, num2Int) {
			return num1Int, num2Int, true
		}
	}

	panic("Введенные числа не подходят под критерии работы приложения")
}

func convertRomeToInt(num string) int {
	//Перевод из римских чисел в арабские
	var sum, lastNum int = 0, 0
	for i := 0; i < len(num); i++ {
		var nowNum int = 0
		switch string(num[i]) {
		case "X":
			nowNum = 10
			sum += nowNum
			break
		case "V":
			nowNum = 5
			sum += nowNum
			break
		case "I":
			nowNum = 1
			sum += nowNum
			break
		}

		if lastNum < nowNum {
			sum = sum - lastNum - lastNum
			lastNum = nowNum
		}

		//fmt.Println("Последнее число: ", lastNum, " | число сейчас: ", nowNum, "| сумма: ", sum) // проверка числа
	}
	return sum
}

func checkNumValue(num1, num2 int) bool {
	if num1 <= 10 && num2 <= 10 && num1 > 0 && num2 > 0 {
		return true
	}
	panic("Числа не подходит под условие")
	return false
}

func arithmetic(num1, num2 int, symbol string) int {
	//математические действия
	switch symbol {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		return num1 / num2
	default:
		panic("введина несуществующая операция")
	}
}

func answer(sum int, types bool) {
	var answer string
	//if Arabian
	if types {
		answer = strconv.Itoa(sum)
	} else {
		//else Rome

		answer = convertRomeToString(sum)
	}

	fmt.Println("Ответ: ", answer)
}

func convertRomeToString(sum int) string {
	if sum <= 0 {
		panic("Число ниже или равно нулю")
	}

	var number string
	for sum > 0 {
		if sum >= 100 {
			sum -= 100
			number += "C"
		} else if sum >= 90 {
			sum -= 90
			number += "XC"
		} else if sum >= 50 {
			sum -= 50
			number += "L"
		} else if sum >= 40 {
			sum -= 40
			number += "XL"
		} else if sum >= 10 {
			sum -= 10
			number += "X"
		} else if sum >= 9 {
			sum -= 9
			number += "IX"
		} else if sum >= 5 {
			sum -= 5
			number += "V"
		} else if sum >= 4 {
			sum -= 4
			number += "IV"
		} else if sum >= 1 {
			sum -= 1
			number += "I"
		}
	}
	return number
}
